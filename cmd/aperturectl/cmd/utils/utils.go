package utils

import (
	"context"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v2"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"

	"github.com/Masterminds/semver/v3"
	"github.com/fluxninja/aperture/pkg/log"
	"github.com/fluxninja/aperture/pkg/policies/controlplane"
	"github.com/fluxninja/aperture/pkg/policies/controlplane/circuitfactory"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/config"
	"github.com/go-git/go-git/v5/storage/memory"
)

// GenerateDotFile generates a DOT file from the circuit.
func GenerateDotFile(circuit *circuitfactory.Circuit, dotFilePath string) error {
	d := circuitfactory.DOT(circuit.ToGraphView())
	f, err := os.Create(dotFilePath)
	if err != nil {
		log.Error().Err(err).Msg("error creating file")
		return err
	}
	defer f.Close()

	_, err = f.WriteString(d)
	if err != nil {
		log.Error().Err(err).Msg("error writing to file")
		return err
	}
	return nil
}

// GenerateMermaidFile generates a mermaid file from the circuit.
func GenerateMermaidFile(circuit *circuitfactory.Circuit, mermaidFile string) error {
	m := circuitfactory.Mermaid(circuit.ToGraphView())
	f, err := os.Create(mermaidFile)
	if err != nil {
		log.Error().Err(err).Msg("error creating file")
		return err
	}
	defer f.Close()

	_, err = f.WriteString(m)
	if err != nil {
		log.Error().Err(err).Msg("error writing to file")
		return err
	}
	return nil
}

// CompilePolicy compiles the policy and returns the circuit.
func CompilePolicy(path string) (*circuitfactory.Circuit, error) {
	yamlFile, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	ctx := context.Background()

	// FIXME This ValidateAndCompile function validates the policy as a whole –
	// circuit, but also the other resource classifiers, fluxmeters.  This
	// command is called "circuit-compiler" though, so it's bit... surprising.
	// If we compiled just a circuit, we could drop dependency on
	// `controlplane` package.
	circuit, valid, msg, err := controlplane.ValidateAndCompile(ctx, filepath.Base(path), yamlFile)
	if err != nil {
		return nil, err
	}
	if !valid {
		return nil, fmt.Errorf("invalid circuit: %s", msg)
	}
	return circuit, nil
}

// FetchPolicyFromCR extracts the spec key from a CR and saves it to a temp file.
func FetchPolicyFromCR(crPath string) (string, error) {
	// extract spec key from CR and save it a temp file
	// call compilePolicy with the temp file
	// delete the temp file
	crFile, err := os.ReadFile(crPath)
	if err != nil {
		log.Error().Err(err).Msg("failed to read CR file")
		return "", err
	}
	// unmarshal yaml to map struct and extract spec key
	var cr map[string]interface{}
	err = yaml.Unmarshal(crFile, &cr)
	if err != nil {
		log.Error().Err(err).Msg("failed to unmarshal CR file")
		return "", err
	}
	spec, ok := cr["spec"]
	if !ok {
		log.Error().Msg("failed to find spec key in CR file")
		return "", err
	}
	// marshal spec to yaml
	specYaml, err := yaml.Marshal(spec)
	if err != nil {
		log.Error().Err(err).Msg("failed to marshal spec key in CR file")
		return "", err
	}
	// get filename from path
	filename := filepath.Base(crPath)
	// create temp file
	tmpfile, err := os.CreateTemp("", filename)
	if err != nil {
		log.Error().Err(err).Msg("failed to create temp file")
		return "", err
	}
	// write spec to temp file
	_, err = tmpfile.Write(specYaml)
	if err != nil {
		log.Error().Err(err).Msg("failed to write to temp file")
		return "", err
	}
	// close temp file
	err = tmpfile.Close()
	if err != nil {
		log.Error().Err(err).Msg("failed to close temp file")
		return "", err
	}

	return tmpfile.Name(), nil
}

// GetKubeConfig prepares Kubernetes config to connect with the cluster using provided or default kube config file location.
func GetKubeConfig(kubeConfig string) (*rest.Config, error) {
	if kubeConfig == "" {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			return nil, err
		}
		kubeConfig = filepath.Join(homeDir, ".kube", "config")
		log.Info().Msgf("Using Kubernetes config '%s'", kubeConfig)
	}
	restConfig, err := clientcmd.BuildConfigFromFlags("", kubeConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to Kubernetes. Error: %s", err.Error())
	}
	kubeRestConfig := restConfig
	return kubeRestConfig, nil
}

func ResolveLatestVersion() (string, error) {
	remote := git.NewRemote(memory.NewStorage(), &config.RemoteConfig{
		Name: "origin",
		URLs: []string{apertureRepo},
	})

	refs, err := remote.List(&git.ListOptions{})
	if err != nil {
		return "", err
	}

	var latestRelease *semver.Version

	tagsRefPrefix := "refs/tags/v"

	for _, ref := range refs {
		reference := ref.Name().String()
		if ref.Name().IsTag() && strings.HasPrefix(reference, tagsRefPrefix) {
			version := strings.TrimPrefix(reference, tagsRefPrefix)

			release, err := semver.NewVersion(version)
			if err != nil {
				return "", err
			}

			if release.Prerelease() != "" {
				continue
			}

			if latestRelease == nil || release.GreaterThan(latestRelease) {
				latestRelease = release
			}
		}
	}

	if latestRelease == nil {
		return "", errors.New("unable to resolve release tags to find latest release")
	}
	return fmt.Sprintf("v%s", latestRelease.String()), nil
}