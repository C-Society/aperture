package podscaler

import (
	"context"
	"errors"
	"path"

	policylangv1 "github.com/fluxninja/aperture/api/gen/proto/go/aperture/policy/language/v1"
	policysyncv1 "github.com/fluxninja/aperture/api/gen/proto/go/aperture/policy/sync/v1"
	etcdclient "github.com/fluxninja/aperture/pkg/etcd/client"
	"github.com/fluxninja/aperture/pkg/policies/controlplane/iface"
	"github.com/fluxninja/aperture/pkg/policies/paths"
	clientv3 "go.etcd.io/etcd/client/v3"
	"go.uber.org/fx"
	"google.golang.org/protobuf/proto"
)

// Module returns the fx options for horizontal pod scaler in the main app.
func Module() fx.Option {
	return fx.Options(
		scaleReporterModule(),
	)
}

type podScalerConfigSync struct {
	policyReadAPI  iface.Policy
	podScalerProto *policylangv1.PodScaler
	etcdPath       string
	componentID    string
}

// NewPodScalerOptions creates fx options for HorizontalPodScaler and also returns the agent group name associated with it.
func NewPodScalerOptions(
	podScalerProto *policylangv1.PodScaler,
	componentStackID string,
	policyReadAPI iface.Policy,
) (fx.Option, string, error) {
	// Get Agent Group Name from HorizontalPodScaler.KubernetesObjectSelector.AgentGroup
	k8sObjectSelectorProto := podScalerProto.GetKubernetesObjectSelector()
	if k8sObjectSelectorProto == nil {
		return fx.Options(), "", errors.New("podScaler.Selector is nil")
	}
	agentGroup := k8sObjectSelectorProto.GetAgentGroup()
	etcdPath := path.Join(paths.PodScalerConfigPath,
		paths.AgentComponentKey(agentGroup, policyReadAPI.GetPolicyName(), componentStackID))
	configSync := &podScalerConfigSync{
		podScalerProto: podScalerProto,
		policyReadAPI:  policyReadAPI,
		etcdPath:       etcdPath,
		componentID:    componentStackID,
	}

	return fx.Options(
		fx.Invoke(
			configSync.doSync,
		),
	), agentGroup, nil
}

func (configSync *podScalerConfigSync) doSync(etcdClient *etcdclient.Client, lifecycle fx.Lifecycle) error {
	logger := configSync.policyReadAPI.GetStatusRegistry().GetLogger()
	// Add/remove file in lifecycle hooks in order to sync with etcd.
	lifecycle.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			wrapper := &policysyncv1.PodScalerWrapper{
				PodScaler: configSync.podScalerProto,
				CommonAttributes: &policysyncv1.CommonAttributes{
					PolicyName:  configSync.policyReadAPI.GetPolicyName(),
					PolicyHash:  configSync.policyReadAPI.GetPolicyHash(),
					ComponentId: configSync.componentID,
				},
			}
			dat, err := proto.Marshal(wrapper)
			if err != nil {
				logger.Error().Err(err).Msg("Failed to marshal flux meter config")
				return err
			}
			_, err = etcdClient.KV.Put(clientv3.WithRequireLeader(ctx),
				configSync.etcdPath, string(dat), clientv3.WithLease(etcdClient.LeaseID))
			if err != nil {
				logger.Error().Err(err).Msg("Failed to put flux meter config")
				return err
			}
			return nil
		},
		OnStop: func(ctx context.Context) error {
			_, err := etcdClient.KV.Delete(clientv3.WithRequireLeader(ctx), configSync.etcdPath)
			if err != nil {
				logger.Error().Err(err).Msg("Failed to delete flux meter config")
				return err
			}
			return nil
		},
	})

	return nil
}
