// Code generated by mdatagen. DO NOT EDIT.

package metadata

import "go.opentelemetry.io/collector/confmap"

// MetricConfig provides common config for a particular metric.
type MetricConfig struct {
	Enabled bool `mapstructure:"enabled"`

	enabledSetByUser bool
}

func (ms *MetricConfig) Unmarshal(parser *confmap.Conf) error {
	if parser == nil {
		return nil
	}
	err := parser.Unmarshal(ms, confmap.WithErrorUnused())
	if err != nil {
		return err
	}
	ms.enabledSetByUser = parser.IsSet("enabled")
	return nil
}

// MetricsConfig provides config for k8s_cluster metrics.
type MetricsConfig struct {
	K8sContainerCPULimit                MetricConfig `mapstructure:"k8s.container.cpu_limit"`
	K8sContainerCPURequest              MetricConfig `mapstructure:"k8s.container.cpu_request"`
	K8sContainerEphemeralstorageLimit   MetricConfig `mapstructure:"k8s.container.ephemeralstorage_limit"`
	K8sContainerEphemeralstorageRequest MetricConfig `mapstructure:"k8s.container.ephemeralstorage_request"`
	K8sContainerMemoryLimit             MetricConfig `mapstructure:"k8s.container.memory_limit"`
	K8sContainerMemoryRequest           MetricConfig `mapstructure:"k8s.container.memory_request"`
	K8sContainerReady                   MetricConfig `mapstructure:"k8s.container.ready"`
	K8sContainerRestarts                MetricConfig `mapstructure:"k8s.container.restarts"`
	K8sContainerStorageLimit            MetricConfig `mapstructure:"k8s.container.storage_limit"`
	K8sContainerStorageRequest          MetricConfig `mapstructure:"k8s.container.storage_request"`
	K8sCronjobActiveJobs                MetricConfig `mapstructure:"k8s.cronjob.active_jobs"`
	K8sDaemonsetCurrentScheduledNodes   MetricConfig `mapstructure:"k8s.daemonset.current_scheduled_nodes"`
	K8sDaemonsetDesiredScheduledNodes   MetricConfig `mapstructure:"k8s.daemonset.desired_scheduled_nodes"`
	K8sDaemonsetMisscheduledNodes       MetricConfig `mapstructure:"k8s.daemonset.misscheduled_nodes"`
	K8sDaemonsetReadyNodes              MetricConfig `mapstructure:"k8s.daemonset.ready_nodes"`
	K8sDeploymentAvailable              MetricConfig `mapstructure:"k8s.deployment.available"`
	K8sDeploymentDesired                MetricConfig `mapstructure:"k8s.deployment.desired"`
	K8sHpaCurrentReplicas               MetricConfig `mapstructure:"k8s.hpa.current_replicas"`
	K8sHpaDesiredReplicas               MetricConfig `mapstructure:"k8s.hpa.desired_replicas"`
	K8sHpaMaxReplicas                   MetricConfig `mapstructure:"k8s.hpa.max_replicas"`
	K8sHpaMinReplicas                   MetricConfig `mapstructure:"k8s.hpa.min_replicas"`
	K8sJobActivePods                    MetricConfig `mapstructure:"k8s.job.active_pods"`
	K8sJobDesiredSuccessfulPods         MetricConfig `mapstructure:"k8s.job.desired_successful_pods"`
	K8sJobFailedPods                    MetricConfig `mapstructure:"k8s.job.failed_pods"`
	K8sJobMaxParallelPods               MetricConfig `mapstructure:"k8s.job.max_parallel_pods"`
	K8sJobSuccessfulPods                MetricConfig `mapstructure:"k8s.job.successful_pods"`
	K8sNamespacePhase                   MetricConfig `mapstructure:"k8s.namespace.phase"`
	K8sPodPhase                         MetricConfig `mapstructure:"k8s.pod.phase"`
	K8sPodStatusReason                  MetricConfig `mapstructure:"k8s.pod.status_reason"`
	K8sReplicasetAvailable              MetricConfig `mapstructure:"k8s.replicaset.available"`
	K8sReplicasetDesired                MetricConfig `mapstructure:"k8s.replicaset.desired"`
	K8sReplicationControllerAvailable   MetricConfig `mapstructure:"k8s.replication_controller.available"`
	K8sReplicationControllerDesired     MetricConfig `mapstructure:"k8s.replication_controller.desired"`
	K8sResourceQuotaHardLimit           MetricConfig `mapstructure:"k8s.resource_quota.hard_limit"`
	K8sResourceQuotaUsed                MetricConfig `mapstructure:"k8s.resource_quota.used"`
	K8sStatefulsetCurrentPods           MetricConfig `mapstructure:"k8s.statefulset.current_pods"`
	K8sStatefulsetDesiredPods           MetricConfig `mapstructure:"k8s.statefulset.desired_pods"`
	K8sStatefulsetReadyPods             MetricConfig `mapstructure:"k8s.statefulset.ready_pods"`
	K8sStatefulsetUpdatedPods           MetricConfig `mapstructure:"k8s.statefulset.updated_pods"`
	OpenshiftAppliedclusterquotaLimit   MetricConfig `mapstructure:"openshift.appliedclusterquota.limit"`
	OpenshiftAppliedclusterquotaUsed    MetricConfig `mapstructure:"openshift.appliedclusterquota.used"`
	OpenshiftClusterquotaLimit          MetricConfig `mapstructure:"openshift.clusterquota.limit"`
	OpenshiftClusterquotaUsed           MetricConfig `mapstructure:"openshift.clusterquota.used"`
}

func DefaultMetricsConfig() MetricsConfig {
	return MetricsConfig{
		K8sContainerCPULimit: MetricConfig{
			Enabled: true,
		},
		K8sContainerCPURequest: MetricConfig{
			Enabled: true,
		},
		K8sContainerEphemeralstorageLimit: MetricConfig{
			Enabled: true,
		},
		K8sContainerEphemeralstorageRequest: MetricConfig{
			Enabled: true,
		},
		K8sContainerMemoryLimit: MetricConfig{
			Enabled: true,
		},
		K8sContainerMemoryRequest: MetricConfig{
			Enabled: true,
		},
		K8sContainerReady: MetricConfig{
			Enabled: true,
		},
		K8sContainerRestarts: MetricConfig{
			Enabled: true,
		},
		K8sContainerStorageLimit: MetricConfig{
			Enabled: true,
		},
		K8sContainerStorageRequest: MetricConfig{
			Enabled: true,
		},
		K8sCronjobActiveJobs: MetricConfig{
			Enabled: true,
		},
		K8sDaemonsetCurrentScheduledNodes: MetricConfig{
			Enabled: true,
		},
		K8sDaemonsetDesiredScheduledNodes: MetricConfig{
			Enabled: true,
		},
		K8sDaemonsetMisscheduledNodes: MetricConfig{
			Enabled: true,
		},
		K8sDaemonsetReadyNodes: MetricConfig{
			Enabled: true,
		},
		K8sDeploymentAvailable: MetricConfig{
			Enabled: true,
		},
		K8sDeploymentDesired: MetricConfig{
			Enabled: true,
		},
		K8sHpaCurrentReplicas: MetricConfig{
			Enabled: true,
		},
		K8sHpaDesiredReplicas: MetricConfig{
			Enabled: true,
		},
		K8sHpaMaxReplicas: MetricConfig{
			Enabled: true,
		},
		K8sHpaMinReplicas: MetricConfig{
			Enabled: true,
		},
		K8sJobActivePods: MetricConfig{
			Enabled: true,
		},
		K8sJobDesiredSuccessfulPods: MetricConfig{
			Enabled: true,
		},
		K8sJobFailedPods: MetricConfig{
			Enabled: true,
		},
		K8sJobMaxParallelPods: MetricConfig{
			Enabled: true,
		},
		K8sJobSuccessfulPods: MetricConfig{
			Enabled: true,
		},
		K8sNamespacePhase: MetricConfig{
			Enabled: true,
		},
		K8sPodPhase: MetricConfig{
			Enabled: true,
		},
		K8sPodStatusReason: MetricConfig{
			Enabled: false,
		},
		K8sReplicasetAvailable: MetricConfig{
			Enabled: true,
		},
		K8sReplicasetDesired: MetricConfig{
			Enabled: true,
		},
		K8sReplicationControllerAvailable: MetricConfig{
			Enabled: true,
		},
		K8sReplicationControllerDesired: MetricConfig{
			Enabled: true,
		},
		K8sResourceQuotaHardLimit: MetricConfig{
			Enabled: true,
		},
		K8sResourceQuotaUsed: MetricConfig{
			Enabled: true,
		},
		K8sStatefulsetCurrentPods: MetricConfig{
			Enabled: true,
		},
		K8sStatefulsetDesiredPods: MetricConfig{
			Enabled: true,
		},
		K8sStatefulsetReadyPods: MetricConfig{
			Enabled: true,
		},
		K8sStatefulsetUpdatedPods: MetricConfig{
			Enabled: true,
		},
		OpenshiftAppliedclusterquotaLimit: MetricConfig{
			Enabled: true,
		},
		OpenshiftAppliedclusterquotaUsed: MetricConfig{
			Enabled: true,
		},
		OpenshiftClusterquotaLimit: MetricConfig{
			Enabled: true,
		},
		OpenshiftClusterquotaUsed: MetricConfig{
			Enabled: true,
		},
	}
}

// ResourceAttributeConfig provides common config for a particular resource attribute.
type ResourceAttributeConfig struct {
	Enabled bool `mapstructure:"enabled"`

	enabledSetByUser bool
}

func (rac *ResourceAttributeConfig) Unmarshal(parser *confmap.Conf) error {
	if parser == nil {
		return nil
	}
	err := parser.Unmarshal(rac, confmap.WithErrorUnused())
	if err != nil {
		return err
	}
	rac.enabledSetByUser = parser.IsSet("enabled")
	return nil
}

// ResourceAttributesConfig provides config for k8s_cluster resource attributes.
type ResourceAttributesConfig struct {
	ContainerID                  ResourceAttributeConfig `mapstructure:"container.id"`
	ContainerImageName           ResourceAttributeConfig `mapstructure:"container.image.name"`
	ContainerImageTag            ResourceAttributeConfig `mapstructure:"container.image.tag"`
	K8sContainerName             ResourceAttributeConfig `mapstructure:"k8s.container.name"`
	K8sCronjobName               ResourceAttributeConfig `mapstructure:"k8s.cronjob.name"`
	K8sCronjobUID                ResourceAttributeConfig `mapstructure:"k8s.cronjob.uid"`
	K8sDaemonsetName             ResourceAttributeConfig `mapstructure:"k8s.daemonset.name"`
	K8sDaemonsetUID              ResourceAttributeConfig `mapstructure:"k8s.daemonset.uid"`
	K8sDeploymentName            ResourceAttributeConfig `mapstructure:"k8s.deployment.name"`
	K8sDeploymentUID             ResourceAttributeConfig `mapstructure:"k8s.deployment.uid"`
	K8sHpaName                   ResourceAttributeConfig `mapstructure:"k8s.hpa.name"`
	K8sHpaUID                    ResourceAttributeConfig `mapstructure:"k8s.hpa.uid"`
	K8sJobName                   ResourceAttributeConfig `mapstructure:"k8s.job.name"`
	K8sJobUID                    ResourceAttributeConfig `mapstructure:"k8s.job.uid"`
	K8sKubeletVersion            ResourceAttributeConfig `mapstructure:"k8s.kubelet.version"`
	K8sKubeproxyVersion          ResourceAttributeConfig `mapstructure:"k8s.kubeproxy.version"`
	K8sNamespaceName             ResourceAttributeConfig `mapstructure:"k8s.namespace.name"`
	K8sNamespaceUID              ResourceAttributeConfig `mapstructure:"k8s.namespace.uid"`
	K8sNodeName                  ResourceAttributeConfig `mapstructure:"k8s.node.name"`
	K8sNodeUID                   ResourceAttributeConfig `mapstructure:"k8s.node.uid"`
	K8sPodName                   ResourceAttributeConfig `mapstructure:"k8s.pod.name"`
	K8sPodUID                    ResourceAttributeConfig `mapstructure:"k8s.pod.uid"`
	K8sReplicasetName            ResourceAttributeConfig `mapstructure:"k8s.replicaset.name"`
	K8sReplicasetUID             ResourceAttributeConfig `mapstructure:"k8s.replicaset.uid"`
	K8sReplicationcontrollerName ResourceAttributeConfig `mapstructure:"k8s.replicationcontroller.name"`
	K8sReplicationcontrollerUID  ResourceAttributeConfig `mapstructure:"k8s.replicationcontroller.uid"`
	K8sResourcequotaName         ResourceAttributeConfig `mapstructure:"k8s.resourcequota.name"`
	K8sResourcequotaUID          ResourceAttributeConfig `mapstructure:"k8s.resourcequota.uid"`
	K8sStatefulsetName           ResourceAttributeConfig `mapstructure:"k8s.statefulset.name"`
	K8sStatefulsetUID            ResourceAttributeConfig `mapstructure:"k8s.statefulset.uid"`
	OpencensusResourcetype       ResourceAttributeConfig `mapstructure:"opencensus.resourcetype"`
	OpenshiftClusterquotaName    ResourceAttributeConfig `mapstructure:"openshift.clusterquota.name"`
	OpenshiftClusterquotaUID     ResourceAttributeConfig `mapstructure:"openshift.clusterquota.uid"`
}

func DefaultResourceAttributesConfig() ResourceAttributesConfig {
	return ResourceAttributesConfig{
		ContainerID: ResourceAttributeConfig{
			Enabled: true,
		},
		ContainerImageName: ResourceAttributeConfig{
			Enabled: true,
		},
		ContainerImageTag: ResourceAttributeConfig{
			Enabled: true,
		},
		K8sContainerName: ResourceAttributeConfig{
			Enabled: true,
		},
		K8sCronjobName: ResourceAttributeConfig{
			Enabled: true,
		},
		K8sCronjobUID: ResourceAttributeConfig{
			Enabled: true,
		},
		K8sDaemonsetName: ResourceAttributeConfig{
			Enabled: true,
		},
		K8sDaemonsetUID: ResourceAttributeConfig{
			Enabled: true,
		},
		K8sDeploymentName: ResourceAttributeConfig{
			Enabled: true,
		},
		K8sDeploymentUID: ResourceAttributeConfig{
			Enabled: true,
		},
		K8sHpaName: ResourceAttributeConfig{
			Enabled: true,
		},
		K8sHpaUID: ResourceAttributeConfig{
			Enabled: true,
		},
		K8sJobName: ResourceAttributeConfig{
			Enabled: true,
		},
		K8sJobUID: ResourceAttributeConfig{
			Enabled: true,
		},
		K8sKubeletVersion: ResourceAttributeConfig{
			Enabled: false,
		},
		K8sKubeproxyVersion: ResourceAttributeConfig{
			Enabled: false,
		},
		K8sNamespaceName: ResourceAttributeConfig{
			Enabled: true,
		},
		K8sNamespaceUID: ResourceAttributeConfig{
			Enabled: true,
		},
		K8sNodeName: ResourceAttributeConfig{
			Enabled: true,
		},
		K8sNodeUID: ResourceAttributeConfig{
			Enabled: true,
		},
		K8sPodName: ResourceAttributeConfig{
			Enabled: true,
		},
		K8sPodUID: ResourceAttributeConfig{
			Enabled: true,
		},
		K8sReplicasetName: ResourceAttributeConfig{
			Enabled: true,
		},
		K8sReplicasetUID: ResourceAttributeConfig{
			Enabled: true,
		},
		K8sReplicationcontrollerName: ResourceAttributeConfig{
			Enabled: true,
		},
		K8sReplicationcontrollerUID: ResourceAttributeConfig{
			Enabled: true,
		},
		K8sResourcequotaName: ResourceAttributeConfig{
			Enabled: true,
		},
		K8sResourcequotaUID: ResourceAttributeConfig{
			Enabled: true,
		},
		K8sStatefulsetName: ResourceAttributeConfig{
			Enabled: true,
		},
		K8sStatefulsetUID: ResourceAttributeConfig{
			Enabled: true,
		},
		OpencensusResourcetype: ResourceAttributeConfig{
			Enabled: true,
		},
		OpenshiftClusterquotaName: ResourceAttributeConfig{
			Enabled: true,
		},
		OpenshiftClusterquotaUID: ResourceAttributeConfig{
			Enabled: true,
		},
	}
}

// MetricsBuilderConfig is a configuration for k8s_cluster metrics builder.
type MetricsBuilderConfig struct {
	Metrics            MetricsConfig            `mapstructure:"metrics"`
	ResourceAttributes ResourceAttributesConfig `mapstructure:"resource_attributes"`
}

func DefaultMetricsBuilderConfig() MetricsBuilderConfig {
	return MetricsBuilderConfig{
		Metrics:            DefaultMetricsConfig(),
		ResourceAttributes: DefaultResourceAttributesConfig(),
	}
}
