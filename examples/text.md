****** API resources ******
group: rbac.authorization.k8s.io, kind: RoleBinding, resource: rolebindings, versions: map[v1:[]]
group: infrastructure.cluster.x-k8s.io, kind: DockerMachine, resource: dockermachines, versions: map[v1alpha4:[catalog.repo-catalog.infra:workload-cluster.capi-docker.ws1] v1beta1:[catalog.repo-catalog.infra:workload-cluster.capi-docker.ws1]]
group: rbac.authorization.k8s.io, kind: ClusterRole, resource: clusterroles, versions: map[v1:[]]
group: bootstrap.cluster.x-k8s.io, kind: KubeadmConfig, resource: kubeadmconfigs, versions: map[v1alpha4:[catalog.repo-catalog.infra:workload-cluster.capi-bootstrap.ws1] v1beta1:[catalog.repo-catalog.infra:workload-cluster.capi-bootstrap.ws1]]
group: cluster.x-k8s.io, kind: Cluster, resource: clusters, versions: map[v1alpha4:[catalog.repo-catalog.infra:workload-cluster.capi-core.ws1] v1beta1:[catalog.repo-catalog.infra:workload-cluster.capi-core.ws1]]
group: infrastructure.cluster.x-k8s.io, kind: DockerCluster, resource: dockerclusters, versions: map[v1alpha4:[catalog.repo-catalog.infra:workload-cluster.capi-docker.ws1] v1beta1:[catalog.repo-catalog.infra:workload-cluster.capi-docker.ws1]]
group: , kind: Service, resource: services, versions: map[v1:[]]
group: , kind: Namespace, resource: namespaces, versions: map[v1:[]]
group: apps, kind: ReplicaSet, resource: replicasets, versions: map[v1:[]]
group: rbac.authorization.k8s.io, kind: ClusterRoleBinding, resource: clusterrolebindings, versions: map[v1:[]]
group: config.pkg.kform.dev, kind: Repository, resource: repositories, versions: map[v1alpha1:[]]
group: addons.cluster.x-k8s.io, kind: ClusterResourceSet, resource: clusterresourcesets, versions: map[v1alpha4:[catalog.repo-catalog.infra:workload-cluster.capi-core.ws1] v1beta1:[catalog.repo-catalog.infra:workload-cluster.capi-core.ws1]]
group: cluster.x-k8s.io, kind: Machine, resource: machines, versions: map[v1alpha4:[catalog.repo-catalog.infra:workload-cluster.capi-core.ws1] v1beta1:[catalog.repo-catalog.infra:workload-cluster.capi-core.ws1]]
group: cluster.x-k8s.io, kind: MachineDeployment, resource: machinedeployments, versions: map[v1alpha4:[catalog.repo-catalog.infra:workload-cluster.capi-core.ws1] v1beta1:[catalog.repo-catalog.infra:workload-cluster.capi-core.ws1]]
group: , kind: ConfigMap, resource: configmaps, versions: map[v1:[]]
group: apps, kind: DaemonSet, resource: daemonsets, versions: map[v1:[]]
group: batch, kind: Job, resource: jobs, versions: map[v1:[]]
group: storage.k8s.io, kind: CSIDriver, resource: csidrivers, versions: map[v1:[]]
group: infrastructure.cluster.x-k8s.io, kind: DockerMachineTemplate, resource: dockermachinetemplates, versions: map[v1alpha4:[catalog.repo-catalog.infra:workload-cluster.capi-docker.ws1] v1beta1:[catalog.repo-catalog.infra:workload-cluster.capi-docker.ws1]]
group: policy, kind: PodDisruptionBudget, resource: poddisruptionbudgets, versions: map[v1:[]]
group: discovery.k8s.io, kind: EndpointSlice, resource: endpointslices, versions: map[v1:[]]
group: inv.sdcio.dev, kind: TargetSyncProfile, resource: targetsyncprofiles, versions: map[v1alpha1:[catalog.repo-catalog.config:sw.config-server.ws1]]
group: , kind: ServiceAccount, resource: serviceaccounts, versions: map[v1:[]]
group: , kind: ResourceQuota, resource: resourcequotas, versions: map[v1:[]]
group: events.k8s.io, kind: Event, resource: events, versions: map[v1:[]]
group: autoscaling, kind: HorizontalPodAutoscaler, resource: horizontalpodautoscalers, versions: map[v2:[]]
group: bootstrap.cluster.x-k8s.io, kind: KubeadmConfigTemplate, resource: kubeadmconfigtemplates, versions: map[v1alpha4:[catalog.repo-catalog.infra:workload-cluster.capi-bootstrap.ws1] v1beta1:[catalog.repo-catalog.infra:workload-cluster.capi-bootstrap.ws1]]
group: config.sdcio.dev, kind: apiService, resource: apiService, versions: map[v1alpha1:[catalog.repo-catalog.config:sw.config-server.ws1]]
group: addons.cluster.x-k8s.io, kind: ClusterResourceSetBinding, resource: clusterresourcesetbindings, versions: map[v1alpha4:[catalog.repo-catalog.infra:workload-cluster.capi-core.ws1] v1beta1:[catalog.repo-catalog.infra:workload-cluster.capi-core.ws1]]
group: cluster.x-k8s.io, kind: ClusterClass, resource: clusterclasses, versions: map[v1alpha4:[catalog.repo-catalog.infra:workload-cluster.capi-core.ws1] v1beta1:[catalog.repo-catalog.infra:workload-cluster.capi-core.ws1]]
group: , kind: PodTemplate, resource: podtemplates, versions: map[v1:[]]
group: flowcontrol.apiserver.k8s.io, kind: PriorityLevelConfiguration, resource: prioritylevelconfigurations, versions: map[v1beta3:[]]
group: inv.sdcio.dev, kind: DiscoveryRule, resource: discoveryrules, versions: map[v1alpha1:[catalog.repo-catalog.config:sw.config-server.ws1]]
group: inv.sdcio.dev, kind: Target, resource: targets, versions: map[v1alpha1:[catalog.repo-catalog.config:sw.config-server.ws1]]
group: inv.sdcio.dev, kind: TargetConnectionProfile, resource: targetconnectionprofiles, versions: map[v1alpha1:[catalog.repo-catalog.config:sw.config-server.ws1]]
group: , kind: ComponentStatus, resource: componentstatuses, versions: map[v1:[]]
group: batch, kind: CronJob, resource: cronjobs, versions: map[v1:[]]
group: pkg.kform.dev, kind: PackageRevision, resource: packagerevisions, versions: map[v1alpha1:[]]
group: pkg.kform.dev, kind: PackageRevisionResources, resource: packagerevisionresources, versions: map[v1alpha1:[]]
group: inv.sdcio.dev, kind: Schema, resource: schemas, versions: map[v1alpha1:[catalog.repo-catalog.config:sw.config-server.ws1]]
group: cluster.x-k8s.io, kind: MachineHealthCheck, resource: machinehealthchecks, versions: map[v1alpha4:[catalog.repo-catalog.infra:workload-cluster.capi-core.ws1] v1beta1:[catalog.repo-catalog.infra:workload-cluster.capi-core.ws1]]
group: , kind: Binding, resource: bindings, versions: map[v1:[]]
group: , kind: ReplicationController, resource: replicationcontrollers, versions: map[v1:[]]
group: apiregistration.k8s.io, kind: APIService, resource: apiservices, versions: map[v1:[]]
group: apiextensions.k8s.io, kind: CustomResourceDefinition, resource: customresourcedefinitions, versions: map[v1:[]]
group: , kind: PersistentVolume, resource: persistentvolumes, versions: map[v1:[]]
group: , kind: Secret, resource: secrets, versions: map[v1:[]]
group: runtime.cluster.x-k8s.io, kind: ExtensionConfig, resource: extensionconfigs, versions: map[v1alpha1:[catalog.repo-catalog.infra:workload-cluster.capi-core.ws1]]
group: apps, kind: Deployment, resource: deployments, versions: map[v1:[]]
group: networking.k8s.io, kind: IngressClass, resource: ingressclasses, versions: map[v1:[]]
group: networking.k8s.io, kind: NetworkPolicy, resource: networkpolicies, versions: map[v1:[]]
group: cluster.x-k8s.io, kind: MachineSet, resource: machinesets, versions: map[v1alpha4:[catalog.repo-catalog.infra:workload-cluster.capi-core.ws1] v1beta1:[catalog.repo-catalog.infra:workload-cluster.capi-core.ws1]]
group: authorization.k8s.io, kind: SelfSubjectAccessReview, resource: selfsubjectaccessreviews, versions: map[v1:[]]
group: infrastructure.cluster.x-k8s.io, kind: DockerClusterTemplate, resource: dockerclustertemplates, versions: map[v1alpha4:[catalog.repo-catalog.infra:workload-cluster.capi-docker.ws1] v1beta1:[catalog.repo-catalog.infra:workload-cluster.capi-docker.ws1]]
group: admissionregistration.k8s.io, kind: ValidatingWebhookConfiguration, resource: validatingwebhookconfigurations, versions: map[v1:[]]
group: ipam.cluster.x-k8s.io, kind: IPAddress, resource: ipaddresses, versions: map[v1alpha1:[catalog.repo-catalog.infra:workload-cluster.capi-core.ws1] v1beta1:[catalog.repo-catalog.infra:workload-cluster.capi-core.ws1]]
group: infrastructure.cluster.x-k8s.io, kind: DockerMachinePool, resource: dockermachinepools, versions: map[v1alpha4:[catalog.repo-catalog.infra:workload-cluster.capi-docker.ws1] v1beta1:[catalog.repo-catalog.infra:workload-cluster.capi-docker.ws1]]
group: infrastructure.cluster.x-k8s.io, kind: DockerMachinePoolTemplate, resource: dockermachinepooltemplates, versions: map[v1beta1:[catalog.repo-catalog.infra:workload-cluster.capi-docker.ws1]]
group: , kind: Pod, resource: pods, versions: map[v1:[]]
group: , kind: PersistentVolumeClaim, resource: persistentvolumeclaims, versions: map[v1:[]]
group: storage.k8s.io, kind: CSINode, resource: csinodes, versions: map[v1:[]]
group: storage.k8s.io, kind: CSIStorageCapacity, resource: csistoragecapacities, versions: map[v1:[]]
group: scheduling.k8s.io, kind: PriorityClass, resource: priorityclasses, versions: map[v1:[]]
group: coordination.k8s.io, kind: Lease, resource: leases, versions: map[v1:[]]
group: ipam.cluster.x-k8s.io, kind: IPAddressClaim, resource: ipaddressclaims, versions: map[v1alpha1:[catalog.repo-catalog.infra:workload-cluster.capi-core.ws1] v1beta1:[catalog.repo-catalog.infra:workload-cluster.capi-core.ws1]]
group: , kind: Node, resource: nodes, versions: map[v1:[]]
group: , kind: Event, resource: events, versions: map[v1:[]]
group: authorization.k8s.io, kind: SubjectAccessReview, resource: subjectaccessreviews, versions: map[v1:[]]
group: networking.k8s.io, kind: Ingress, resource: ingresses, versions: map[v1:[]]
group: controlplane.cluster.x-k8s.io, kind: KubeadmControlPlane, resource: kubeadmcontrolplanes, versions: map[v1alpha4:[catalog.repo-catalog.infra:workload-cluster.capi-controlplane.ws1] v1beta1:[catalog.repo-catalog.infra:workload-cluster.capi-controlplane.ws1]]
group: , kind: LimitRange, resource: limitranges, versions: map[v1:[]]
group: apps, kind: ControllerRevision, resource: controllerrevisions, versions: map[v1:[]]
group: authorization.k8s.io, kind: SelfSubjectRulesReview, resource: selfsubjectrulesreviews, versions: map[v1:[]]
group: flowcontrol.apiserver.k8s.io, kind: FlowSchema, resource: flowschemas, versions: map[v1beta3:[]]
group: storage.k8s.io, kind: StorageClass, resource: storageclasses, versions: map[v1:[]]
group: admissionregistration.k8s.io, kind: MutatingWebhookConfiguration, resource: mutatingwebhookconfigurations, versions: map[v1:[]]
group: node.k8s.io, kind: RuntimeClass, resource: runtimeclasses, versions: map[v1:[]]
group: config.pkg.kform.dev, kind: PackageVariant, resource: packagevariants, versions: map[v1alpha1:[]]
group: , kind: Endpoints, resource: endpoints, versions: map[v1:[]]
group: certificates.k8s.io, kind: CertificateSigningRequest, resource: certificatesigningrequests, versions: map[v1:[]]
group: rbac.authorization.k8s.io, kind: Role, resource: roles, versions: map[v1:[]]
group: storage.k8s.io, kind: VolumeAttachment, resource: volumeattachments, versions: map[v1:[]]
group: cluster.x-k8s.io, kind: MachinePool, resource: machinepools, versions: map[v1alpha4:[catalog.repo-catalog.infra:workload-cluster.capi-core.ws1] v1beta1:[catalog.repo-catalog.infra:workload-cluster.capi-core.ws1]]
group: apps, kind: StatefulSet, resource: statefulsets, versions: map[v1:[]]
group: authentication.k8s.io, kind: TokenReview, resource: tokenreviews, versions: map[v1:[]]
group: authorization.k8s.io, kind: LocalSubjectAccessReview, resource: localsubjectaccessreviews, versions: map[v1:[]]
group: controlplane.cluster.x-k8s.io, kind: KubeadmControlPlaneTemplate, resource: kubeadmcontrolplanetemplates, versions: map[v1alpha4:[catalog.repo-catalog.infra:workload-cluster.capi-controlplane.ws1] v1beta1:[catalog.repo-catalog.infra:workload-cluster.capi-controlplane.ws1]]
****************************
******** GR MAPPING ********
 persistentvolumes PersistentVolume
apiregistration.k8s.io apiservices APIService
apps replicasets ReplicaSet
authorization.k8s.io selfsubjectaccessreviews SelfSubjectAccessReview
scheduling.k8s.io priorityclasses PriorityClass
inv.sdcio.dev discoveryrules DiscoveryRule
config.sdcio.dev apiService apiService
infrastructure.cluster.x-k8s.io dockermachines DockerMachine
 persistentvolumeclaims PersistentVolumeClaim
 configmaps ConfigMap
batch cronjobs CronJob
rbac.authorization.k8s.io roles Role
storage.k8s.io volumeattachments VolumeAttachment
node.k8s.io runtimeclasses RuntimeClass
bootstrap.cluster.x-k8s.io kubeadmconfigs KubeadmConfig
inv.sdcio.dev targetconnectionprofiles TargetConnectionProfile
 events Event
apps statefulsets StatefulSet
config.pkg.kform.dev packagevariants PackageVariant
pkg.kform.dev packagerevisions PackageRevision
cluster.x-k8s.io machinedeployments MachineDeployment
 nodes Node
 resourcequotas ResourceQuota
apps daemonsets DaemonSet
apps controllerrevisions ControllerRevision
storage.k8s.io csinodes CSINode
storage.k8s.io csidrivers CSIDriver
apiextensions.k8s.io customresourcedefinitions CustomResourceDefinition
 endpoints Endpoints
ipam.cluster.x-k8s.io ipaddressclaims IPAddressClaim
cluster.x-k8s.io machinepools MachinePool
infrastructure.cluster.x-k8s.io dockerclustertemplates DockerClusterTemplate
infrastructure.cluster.x-k8s.io dockermachinepooltemplates DockerMachinePoolTemplate
batch jobs Job
bootstrap.cluster.x-k8s.io kubeadmconfigtemplates KubeadmConfigTemplate
cluster.x-k8s.io clusters Cluster
addons.cluster.x-k8s.io clusterresourcesetbindings ClusterResourceSetBinding
authorization.k8s.io localsubjectaccessreviews LocalSubjectAccessReview
inv.sdcio.dev targets Target
cluster.x-k8s.io clusterclasses ClusterClass
infrastructure.cluster.x-k8s.io dockermachinetemplates DockerMachineTemplate
networking.k8s.io ingresses Ingress
coordination.k8s.io leases Lease
inv.sdcio.dev schemas Schema
controlplane.cluster.x-k8s.io kubeadmcontrolplanetemplates KubeadmControlPlaneTemplate
 replicationcontrollers ReplicationController
config.pkg.kform.dev repositories Repository
inv.sdcio.dev targetsyncprofiles TargetSyncProfile
 componentstatuses ComponentStatus
authorization.k8s.io subjectaccessreviews SubjectAccessReview
certificates.k8s.io certificatesigningrequests CertificateSigningRequest
rbac.authorization.k8s.io clusterroles ClusterRole
admissionregistration.k8s.io validatingwebhookconfigurations ValidatingWebhookConfiguration
cluster.x-k8s.io machines Machine
cluster.x-k8s.io machinesets MachineSet
networking.k8s.io networkpolicies NetworkPolicy
policy poddisruptionbudgets PodDisruptionBudget
rbac.authorization.k8s.io rolebindings RoleBinding
storage.k8s.io csistoragecapacities CSIStorageCapacity
infrastructure.cluster.x-k8s.io dockerclusters DockerCluster
authorization.k8s.io selfsubjectrulesreviews SelfSubjectRulesReview
 services Service
 pods Pod
storage.k8s.io storageclasses StorageClass
discovery.k8s.io endpointslices EndpointSlice
pkg.kform.dev packagerevisionresources PackageRevisionResources
infrastructure.cluster.x-k8s.io dockermachinepools DockerMachinePool
 podtemplates PodTemplate
 bindings Binding
 secrets Secret
 namespaces Namespace
 limitranges LimitRange
apps deployments Deployment
autoscaling horizontalpodautoscalers HorizontalPodAutoscaler
rbac.authorization.k8s.io clusterrolebindings ClusterRoleBinding
admissionregistration.k8s.io mutatingwebhookconfigurations MutatingWebhookConfiguration
events.k8s.io events Event
authentication.k8s.io tokenreviews TokenReview
networking.k8s.io ingressclasses IngressClass
flowcontrol.apiserver.k8s.io prioritylevelconfigurations PriorityLevelConfiguration
flowcontrol.apiserver.k8s.io flowschemas FlowSchema
cluster.x-k8s.io machinehealthchecks MachineHealthCheck
 serviceaccounts ServiceAccount
controlplane.cluster.x-k8s.io kubeadmcontrolplanes KubeadmControlPlane
ipam.cluster.x-k8s.io ipaddresses IPAddress
addons.cluster.x-k8s.io clusterresourcesets ClusterResourceSet
runtime.cluster.x-k8s.io extensionconfigs ExtensionConfig
****************************
**** Package Catalog ****
package infra/workload-cluster/workload-cluster revision v2
  pkg gvk config.pkg.kform.dev/v1alpha1, Kind=PackageVariant, upstreams [repo-catalog:infra/workload-cluster/cluster-capi:*]
package infra/workload-cluster/workload-cluster revision v1
  pkg gvk config.pkg.kform.dev/v1alpha1, Kind=PackageVariant, upstreams [repo-catalog:infra/workload-cluster/cluster-capi:*]
package infra/workload-cluster/capi-docker revision v1
  gvk cert-manager.io/v1, Kind=Certificate, resolution error cannot find api reference: group: cert-manager.io, kind: Certificate
  gvk cert-manager.io/v1, Kind=Issuer, resolution error cannot find api reference: group: cert-manager.io, kind: Issuer
  resource gvk /v1, Kind=ServiceAccount
  resource gvk /v1, Kind=Service
  resource gvk /*, Kind=Secret
  resource gvk /v1, Kind=Namespace
  resource gvk /*, Kind=Event
  resource gvk /*, Kind=ConfigMap
  resource gvk admissionregistration.k8s.io/v1, Kind=MutatingWebhookConfiguration
  resource gvk admissionregistration.k8s.io/v1, Kind=ValidatingWebhookConfiguration
  resource gvk apiextensions.k8s.io/v1, Kind=CustomResourceDefinition
  resource gvk apps/v1, Kind=Deployment
  resource gvk authentication.k8s.io/*, Kind=TokenReview
  resource gvk authorization.k8s.io/*, Kind=SubjectAccessReview
  resource gvk coordination.k8s.io/*, Kind=Lease
  resource gvk rbac.authorization.k8s.io/v1, Kind=ClusterRole
  resource gvk rbac.authorization.k8s.io/v1, Kind=RoleBinding
  resource gvk rbac.authorization.k8s.io/v1, Kind=ClusterRoleBinding
  resource gvk rbac.authorization.k8s.io/v1, Kind=Role
  pkg gvk cluster.x-k8s.io/*, Kind=MachineSet, upstreams [repo-catalog:capi-core:*]
  pkg gvk cluster.x-k8s.io/*, Kind=MachinePool, upstreams [repo-catalog:capi-core:*]
  pkg gvk cluster.x-k8s.io/*, Kind=Cluster, upstreams [repo-catalog:capi-core:*]
  pkg gvk cluster.x-k8s.io/*, Kind=Machine, upstreams [repo-catalog:capi-core:*]
package infra/workload-cluster/capi-kind revision v2
  pkg gvk claim.pkg.kform.dev/v1alpha1, Kind=PackageDependency, upstreams [repo-catalog:capi-kind-templates:v1]
  pkg gvk cluster.x-k8s.io/v1beta1, Kind=Cluster, upstreams [repo-catalog:capi-core:*]
  pkg gvk cluster.x-k8s.io/v1beta1, Kind=ClusterClass, upstreams [repo-catalog:capi-core:*]
package infra/workload-cluster/capi-kind-nodepool revision v1
  pkg gvk cluster.x-k8s.io/v1beta1, Kind=Cluster, upstreams [repo-catalog:capi-core:*]
  pkg gvk infrastructure.cluster.x-k8s.io/v1beta1, Kind=DockerMachineTemplate, upstreams [repo-catalog:capi-docker:*]
  pkg gvk bootstrap.cluster.x-k8s.io/v1beta1, Kind=KubeadmConfigTemplate, upstreams [repo-catalog:capi-bootstrap:*]
  pkg gvk cluster.x-k8s.io/v1beta1, Kind=MachineDeployment, upstreams [repo-catalog:capi-core:*]
package infra/workload-cluster/capi-bootstrap revision v1
  gvk cert-manager.io/v1, Kind=Issuer, resolution error cannot find api reference: group: cert-manager.io, kind: Issuer
  gvk cert-manager.io/v1, Kind=Certificate, resolution error cannot find api reference: group: cert-manager.io, kind: Certificate
  resource gvk /*, Kind=Event
  resource gvk /*, Kind=ConfigMap
  resource gvk /*, Kind=Secret
  resource gvk /v1, Kind=Namespace
  resource gvk /v1, Kind=ServiceAccount
  resource gvk /v1, Kind=Service
  resource gvk admissionregistration.k8s.io/v1, Kind=ValidatingWebhookConfiguration
  resource gvk admissionregistration.k8s.io/v1, Kind=MutatingWebhookConfiguration
  resource gvk apiextensions.k8s.io/v1, Kind=CustomResourceDefinition
  resource gvk apps/v1, Kind=Deployment
  resource gvk authentication.k8s.io/*, Kind=TokenReview
  resource gvk authorization.k8s.io/*, Kind=SubjectAccessReview
  resource gvk coordination.k8s.io/*, Kind=Lease
  resource gvk rbac.authorization.k8s.io/v1, Kind=Role
  resource gvk rbac.authorization.k8s.io/v1, Kind=RoleBinding
  resource gvk rbac.authorization.k8s.io/v1, Kind=ClusterRoleBinding
  resource gvk rbac.authorization.k8s.io/v1, Kind=ClusterRole
  pkg gvk cluster.x-k8s.io/*, Kind=Cluster, upstreams [repo-catalog:capi-core:*]
  pkg gvk cluster.x-k8s.io/*, Kind=MachinePool, upstreams [repo-catalog:capi-core:*]
  pkg gvk cluster.x-k8s.io/*, Kind=Machine, upstreams [repo-catalog:capi-core:*]
  pkg gvk cluster.x-k8s.io/*, Kind=MachineSet, upstreams [repo-catalog:capi-core:*]
package infra/workload-cluster/capi-kind-templates revision v1
  pkg gvk infrastructure.cluster.x-k8s.io/v1beta1, Kind=DockerMachineTemplate, upstreams [repo-catalog:capi-docker:*]
  pkg gvk infrastructure.cluster.x-k8s.io/v1beta1, Kind=DockerClusterTemplate, upstreams [repo-catalog:capi-docker:*]
  pkg gvk cluster.x-k8s.io/v1beta1, Kind=ClusterClass, upstreams [repo-catalog:capi-core:*]
  pkg gvk controlplane.cluster.x-k8s.io/v1beta1, Kind=KubeadmControlPlaneTemplate, upstreams [repo-catalog:capi-controlplane:*]
  pkg gvk bootstrap.cluster.x-k8s.io/v1beta1, Kind=KubeadmConfigTemplate, upstreams [repo-catalog:capi-bootstrap:*]
package config/sw/config-server revision v1
  resource gvk /v1, Kind=Secret
  resource gvk /*, Kind=Secret
  resource gvk /v1, Kind=ServiceAccount
  resource gvk /v1, Kind=Namespace
  resource gvk /v1, Kind=PersistentVolumeClaim
  resource gvk /v1, Kind=Service
  resource gvk /*, Kind=Service
  resource gvk /v1, Kind=ConfigMap
  resource gvk /*, Kind=Namespace
  resource gvk /*, Kind=Pod
  resource gvk /*, Kind=ServiceAccount
  resource gvk admissionregistration.k8s.io/*, Kind=ValidatingWebhookConfiguration
  resource gvk admissionregistration.k8s.io/*, Kind=MutatingWebhookConfiguration
  resource gvk apiextensions.k8s.io/v1, Kind=CustomResourceDefinition
  resource gvk apiregistration.k8s.io/v1, Kind=APIService
  resource gvk apps/v1, Kind=Deployment
  resource gvk coordination.k8s.io/*, Kind=Lease
  resource gvk flowcontrol.apiserver.k8s.io/*, Kind=PriorityLevelConfiguration
  resource gvk flowcontrol.apiserver.k8s.io/*, Kind=FlowSchema
  resource gvk rbac.authorization.k8s.io/v1, Kind=ClusterRoleBinding
  resource gvk rbac.authorization.k8s.io/v1, Kind=ClusterRole
  resource gvk rbac.authorization.k8s.io/v1, Kind=Role
  resource gvk rbac.authorization.k8s.io/v1, Kind=RoleBinding
package infra/workload-cluster/capi-controlplane revision v1
  gvk cert-manager.io/v1, Kind=Certificate, resolution error cannot find api reference: group: cert-manager.io, kind: Certificate
  gvk cert-manager.io/v1, Kind=Issuer, resolution error cannot find api reference: group: cert-manager.io, kind: Issuer
  resource gvk /*, Kind=Event
  resource gvk /v1, Kind=Service
  resource gvk /v1, Kind=ServiceAccount
  resource gvk /v1, Kind=Namespace
  resource gvk /*, Kind=Secret
  resource gvk admissionregistration.k8s.io/v1, Kind=MutatingWebhookConfiguration
  resource gvk admissionregistration.k8s.io/v1, Kind=ValidatingWebhookConfiguration
  resource gvk apiextensions.k8s.io/v1, Kind=CustomResourceDefinition
  resource gvk apiextensions.k8s.io/*, Kind=CustomResourceDefinition
  resource gvk apps/v1, Kind=Deployment
  resource gvk authentication.k8s.io/*, Kind=TokenReview
  resource gvk authorization.k8s.io/*, Kind=SubjectAccessReview
  resource gvk coordination.k8s.io/*, Kind=Lease
  resource gvk rbac.authorization.k8s.io/v1, Kind=Role
  resource gvk rbac.authorization.k8s.io/v1, Kind=ClusterRole
  resource gvk rbac.authorization.k8s.io/v1, Kind=RoleBinding
  resource gvk rbac.authorization.k8s.io/v1, Kind=ClusterRoleBinding
  pkg gvk cluster.x-k8s.io/*, Kind=Cluster, upstreams [repo-catalog:capi-core:*]
  pkg gvk cluster.x-k8s.io/*, Kind=Machine, upstreams [repo-catalog:capi-core:*]
package infra/workload-cluster/capi-core revision v1
  gvk cert-manager.io/v1, Kind=Certificate, resolution error cannot find api reference: group: cert-manager.io, kind: Certificate
  gvk cert-manager.io/v1, Kind=Issuer, resolution error cannot find api reference: group: cert-manager.io, kind: Issuer
  resource gvk /*, Kind=ConfigMap
  resource gvk /*, Kind=Event
  resource gvk /v1, Kind=Namespace
  resource gvk /*, Kind=Node
  resource gvk /*, Kind=Secret
  resource gvk /v1, Kind=ServiceAccount
  resource gvk /v1, Kind=Service
  resource gvk /*, Kind=Namespace
  resource gvk admissionregistration.k8s.io/v1, Kind=ValidatingWebhookConfiguration
  resource gvk admissionregistration.k8s.io/v1, Kind=MutatingWebhookConfiguration
  resource gvk apiextensions.k8s.io/*, Kind=CustomResourceDefinition
  resource gvk apiextensions.k8s.io/v1, Kind=CustomResourceDefinition
  resource gvk apps/v1, Kind=Deployment
  resource gvk authentication.k8s.io/*, Kind=TokenReview
  resource gvk authorization.k8s.io/*, Kind=SubjectAccessReview
  resource gvk coordination.k8s.io/*, Kind=Lease
  resource gvk rbac.authorization.k8s.io/v1, Kind=RoleBinding
  resource gvk rbac.authorization.k8s.io/v1, Kind=ClusterRole
  resource gvk rbac.authorization.k8s.io/v1, Kind=ClusterRoleBinding
  resource gvk rbac.authorization.k8s.io/v1, Kind=Role
****************************