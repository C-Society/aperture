package apply

import (
	"context"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
	appsv1 "k8s.io/api/apps/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/kubernetes/scheme"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
	"sigs.k8s.io/yaml"

	languagev1 "github.com/fluxninja/aperture/api/gen/proto/go/aperture/policy/language/v1"
	"github.com/fluxninja/aperture/operator/api"
	policyv1alpha1 "github.com/fluxninja/aperture/operator/api/policy/v1alpha1"
	"github.com/fluxninja/aperture/pkg/log"
)

var (
	file string
	dir  string
)

func init() {
	ApplyPolicyCmd.Flags().StringVar(&file, "file", "", "Path to Aperture Policy file")
	ApplyPolicyCmd.Flags().StringVar(&dir, "dir", "", "Path to directory containing Aperture Policy files")
}

// ApplyPolicyCmd is the command to apply a policy to the cluster.
var ApplyPolicyCmd = &cobra.Command{
	Use:           "policy",
	Short:         "Apply Aperture Policy to the cluster",
	Long:          `Use this command to apply the Aperture Policy to the cluster.`,
	SilenceErrors: true,
	Example: `aperturectl apply policy --file=policies/static-rate-limiting.yaml

aperturectl apply policy --dir=policies`,
	RunE: func(_ *cobra.Command, _ []string) error {
		if file != "" {
			return ApplyPolicy(file)
		} else if dir != "" {
			return ApplyPolicies(dir)
		} else {
			return errors.New("either --file or --dir must be provided")
		}
	},
}

// ApplyPolicies applies all policies in a directory to the cluster.
func ApplyPolicies(policyDir string) error {
	// walk the directory and apply all policies
	return filepath.Walk(policyDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		fileBase := info.Name()[:len(info.Name())-len(filepath.Ext(info.Name()))]
		if filepath.Ext(info.Name()) == ".yaml" && !strings.HasSuffix(fileBase, "-cr") {
			err = ApplyPolicy(path)
			if err != nil {
				return err
			}
		}
		return nil
	})
}

// ApplyPolicy applies a policy to the cluster.
func ApplyPolicy(policyFile string) error {
	policyFileBase := filepath.Base(policyFile)
	policyName := policyFileBase[:len(policyFileBase)-len(filepath.Ext(policyFileBase))]

	content, err := os.ReadFile(policyFile)
	if err != nil {
		return err
	}
	policy := &languagev1.Policy{}
	err = yaml.Unmarshal(content, policy)
	if err != nil {
		log.Warn().Msgf("Skipping apply for policy '%s' due to invalid spec.", policyFile)
		return nil
	}
	policyBytes, err := policy.MarshalJSON()
	if err != nil {
		return err
	}

	policyCR := &policyv1alpha1.Policy{}
	policyCR.Spec.Raw = policyBytes
	policyCR.Name = policyName
	return createAndApplyPolicy(policyCR)
}

func createAndApplyPolicy(policy *policyv1alpha1.Policy) error {
	err := api.AddToScheme(scheme.Scheme)
	if err != nil {
		return fmt.Errorf("failed to connect to Kubernetes: %w", err)
	}

	c, err := client.New(kubeRestConfig, client.Options{
		Scheme: scheme.Scheme,
	})
	if err != nil {
		return fmt.Errorf("failed to create Kubernetes client: %w", err)
	}

	deployment, err := getControllerDeployment()
	if err != nil {
		return err
	}

	policy.Namespace = deployment.GetNamespace()
	policy.Annotations = map[string]string{
		"fluxninja.com/validate": "true",
	}
	spec := policy.Spec
	_, err = controllerutil.CreateOrUpdate(context.Background(), c, policy, func() error {
		policy.Spec = spec
		return nil
	})
	if err != nil {
		return fmt.Errorf("failed to apply policy in Kubernetes: %w", err)
	}

	log.Info().Str("policy", policy.GetName()).Str("namespace", policy.GetNamespace()).Msg("Applied Policy successfully")
	return nil
}

func getControllerDeployment() (*appsv1.Deployment, error) {
	clientSet, err := kubernetes.NewForConfig(kubeRestConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to create new ClientSet: %w", err)
	}

	deployment, err := clientSet.AppsV1().Deployments("").List(context.Background(), metav1.ListOptions{
		FieldSelector: "metadata.name=aperture-controller",
	})
	if err != nil {
		if apierrors.IsNotFound(err) {
			return nil, fmt.Errorf(
				"no deployment with name 'aperture-controller' found on the Kubernetes cluster. The policy can be only applied in the namespace where the Aperture Controller is running")
		}
		return nil, fmt.Errorf("failed to fetch aperture-controller namespace in Kubernetes: %w", err)
	}

	if len(deployment.Items) != 1 {
		return nil, errors.New("no deployment with name 'aperture-controller' found on the Kubernetes cluster. The policy can be only applied in the namespace where the Aperture Controller is running")
	}

	return &deployment.Items[0], nil
}