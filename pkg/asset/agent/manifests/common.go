package manifests

import (
	"github.com/openshift/installer/pkg/asset/agent"
	"github.com/openshift/installer/pkg/asset/installconfig"
	"github.com/openshift/installer/pkg/asset/releaseimage"
	"github.com/openshift/installer/pkg/types"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/utils/pointer"
)

var (
	// TestSSHKey provides a ssh key for unit tests
	TestSSHKey = `|
	ssh-rsa AAAAB3NzaC1y1LJe3zew1ghc= root@localhost.localdomain`
	// TestSecret provides a ssh key for unit tests
	TestSecret = `'{"auths":{"cloud.openshift.com":{"auth":"b3BlUTA=","email":"test@redhat.com"}}}`
)

// GetValidAgentPullSecret returns a valid agent pull secret
func GetValidAgentPullSecret() *AgentPullSecret {
	return &AgentPullSecret{
		Config: &corev1.Secret{
			TypeMeta: metav1.TypeMeta{
				Kind:       "Secret",
				APIVersion: "v1",
			},
			ObjectMeta: metav1.ObjectMeta{
				Name:      "pull-secret",
				Namespace: "cluster-0",
			},
			StringData: map[string]string{
				".dockerconfigjson": TestSecret,
			},
		},
	}
}

// GetValidOptionalInstallConfig returns a valid optional install config
func GetValidOptionalInstallConfig() *agent.OptionalInstallConfig {
	return &agent.OptionalInstallConfig{
		InstallConfig: installconfig.InstallConfig{
			Config: &types.InstallConfig{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "ocp-edge-cluster-0",
					Namespace: "cluster-0",
				},
				BaseDomain: "testing.com",
				PullSecret: TestSecret,
				SSHKey:     TestSSHKey,
				ControlPlane: &types.MachinePool{
					Name:     "master",
					Replicas: pointer.Int64Ptr(3),
					Platform: types.MachinePoolPlatform{},
				},
				Compute: []types.MachinePool{
					{
						Name:     "worker-machine-pool-1",
						Replicas: pointer.Int64Ptr(2),
					},
					{
						Name:     "worker-machine-pool-2",
						Replicas: pointer.Int64Ptr(3),
					},
				},
			},
		},
		Supplied: true,
	}
}

// GetValidReleaseimage returns a valid release image
func GetValidReleaseimage() *releaseimage.Image {
	return &releaseimage.Image{
		PullSpec: releaseimage.GetDefaultReleaseImageOriginal(),
	}
}
