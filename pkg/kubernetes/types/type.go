package client

import "k8s.io/apimachinery/pkg/runtime/schema"

type ResourceName string

// nuwa resource
const (
	Water                    ResourceName = "water"
	Deployment               ResourceName = "deployment"
	Stone                    ResourceName = "stone"
	StatefulSet              ResourceName = "statefulset"
	StatefulSet1             ResourceName = "statefulset1"
	DaemonSet                ResourceName = "daemonsets"
	Injector                 ResourceName = "injector"
	Pod                      ResourceName = "pod"
	Job                      ResourceName = "jobs"
	CronJobs                 ResourceName = "cronjobs"
	ReplicaSet               ResourceName = "replicasets"
	Event                    ResourceName = "events"
	Node                     ResourceName = "nodes"
	ConfigMaps               ResourceName = "configmaps"
	Secrets                  ResourceName = "secrets"
	ResourceQuota            ResourceName = "resourcequotas"
	Service                  ResourceName = "services"
	Ingress                  ResourceName = "ingresses"
	NetworkPolicy            ResourceName = "networkpolicies"
	HorizontalPodAutoscaler  ResourceName = "horizontalpodautoscalers"
	CustomResourceDefinition ResourceName = "customresourcedefinitions"
	PersistentVolume         ResourceName = "persistentvolumes"
	PersistentVolumeClaims   ResourceName = "persistentvolumeclaims"
	StorageClass             ResourceName = "storageclasses"
	ServiceAccount           ResourceName = "serviceaccounts"
	Role                     ResourceName = "roles"
	ClusterRole              ResourceName = "clusterroles"
	RoleBinding              ResourceName = "rolebindings"
	Namespace                ResourceName = "namesapces"
	FormRender               ResourceName = "formrenders"
	ClusterRoleBinding       ResourceName = "clusterrolebindings"
	WorkloadsTemplate        ResourceName = "workloads"
	Endpoint                 ResourceName = "endpoints"
	Page                     ResourceName = "pages"
	Form                     ResourceName = "forms"
	Field                    ResourceName = "fields"
	BaseDepartment           ResourceName = "basedepartments"
	BasePermission           ResourceName = "basepermissions"
	BaseRole                 ResourceName = "baseroles"
	BaseUser                 ResourceName = "baseusers"
	PodSecurityPolicie      ResourceName = "podsecuritypolicies"
	BaseRoleUser             ResourceName = "baseroleusers"

	//tekton
	Pipeline         ResourceName = "pipelines"
	PipelineRun      ResourceName = "pipelineruns"
	Task             ResourceName = "tasks"
	TaskRun          ResourceName = "taskruns"
	PipelineResource ResourceName = "pipelineresources"
)

// GroupVersionResources describe resource collection
var GroupVersionResources = map[ResourceName]schema.GroupVersionResource{
	Water:                    {Group: "nuwa.nip.io", Version: "v1", Resource: "waters"},
	Deployment:               {Group: "apps", Version: "v1", Resource: "deployments"},
	Stone:                    {Group: "nuwa.nip.io", Version: "v1", Resource: "stones"},
	StatefulSet1:             {Group: "nuwa.nip.io", Version: "v1", Resource: "statefulsets"},
	StatefulSet:              {Group: "apps", Version: "v1", Resource: "statefulsets"},
	DaemonSet:                {Group: "apps", Version: "v1", Resource: "daemonsets"},
	Injector:                 {Group: "nuwa.nip.io", Version: "v1", Resource: "injectors"},
	Pod:                      {Group: "", Version: "v1", Resource: "pods"},
	Node:                     {Group: "", Version: "v1", Resource: "nodes"},
	Event:                    {Group: "", Version: "v1", Resource: "events"},
	Job:                      {Group: "batch", Version: "v1", Resource: "jobs"},
	CronJobs:                 {Group: "batch", Version: "v1beta1", Resource: "cronjobs"},
	ReplicaSet:               {Group: "apps", Version: "v1", Resource: "replicasets"},
	ConfigMaps:               {Group: "", Version: "v1", Resource: "configmaps"},
	Secrets:                  {Group: "", Version: "v1", Resource: "secrets"},
	ResourceQuota:            {Group: "", Version: "v1", Resource: "resourcequotas"},
	Service:                  {Group: "", Version: "v1", Resource: "services"},
	Namespace:                {Group: "", Version: "v1", Resource: "namespaces"},
	Ingress:                  {Group: "extensions", Version: "v1beta1", Resource: "ingresses"},
	NetworkPolicy:            {Group: "networking.k8s.io", Version: "v1", Resource: "networkpolicies"},
	HorizontalPodAutoscaler:  {Group: "autoscaling", Version: "v2beta1", Resource: "horizontalpodautoscalers"},
	CustomResourceDefinition: {Group: "apiextensions.k8s.io", Version: "v1beta1", Resource: "customresourcedefinitions"},
	PersistentVolume:         {Group: "", Version: "v1", Resource: "persistentvolumes"},
	PersistentVolumeClaims:   {Group: "", Version: "v1", Resource: "persistentvolumeclaims"},
	StorageClass:             {Group: "storage.k8s.io", Version: "v1", Resource: "storageclasses"},
	ServiceAccount:           {Group: "", Version: "v1", Resource: "serviceaccounts"},
	ClusterRole:              {Group: "rbac.authorization.k8s.io", Version: "v1", Resource: "clusterroles"},
	Role:                     {Group: "rbac.authorization.k8s.io", Version: "v1", Resource: "roles"},
	ClusterRoleBinding:       {Group: "rbac.authorization.k8s.io", Version: "v1", Resource: "clusterrolebindings"},
	RoleBinding:              {Group: "rbac.authorization.k8s.io", Version: "v1", Resource: "rolebindings"},
	Endpoint:                 {Group: "", Version: "v1", Resource: "endpoints"},
	WorkloadsTemplate:        {Group: "fuxi.nip.io", Version: "v1", Resource: "workloads"},
	Page:                     {Group: "fuxi.nip.io", Version: "v1", Resource: "pages"},
	Form:                     {Group: "fuxi.nip.io", Version: "v1", Resource: "forms"},
	Field:                    {Group: "fuxi.nip.io", Version: "v1", Resource: "fields"},
	BaseDepartment:           {Group: "fuxi.nip.io", Version: "v1", Resource: "basedepartments"},
	BasePermission:           {Group: "fuxi.nip.io", Version: "v1", Resource: "basepermissions"},
	BaseRole:                 {Group: "fuxi.nip.io", Version: "v1", Resource: "baseroles"},
	BaseUser:                 {Group: "fuxi.nip.io", Version: "v1", Resource: "baseusers"},
	BaseRoleUser:             {Group: "fuxi.nip.io", Version: "v1", Resource: "baseroleusers"},

	// tekton.dev resource view
	Pipeline:         {Group: "tekton.dev", Version: "v1alpha1", Resource: "pipelines"},
	PipelineRun:      {Group: "tekton.dev", Version: "v1alpha1", Resource: "pipelineruns"},
	Task:             {Group: "tekton.dev", Version: "v1alpha1", Resource: "tasks"},
	TaskRun:          {Group: "tekton.dev", Version: "v1alpha1", Resource: "taskruns"},
	PipelineResource: {Group: "tekton.dev", Version: "v1alpha1", Resource: "pipelineresources"},

	PodSecurityPolicie: {Group: "policy", Version: "v1beta1", Resource: "podsecuritypolicies"},
}

var (
	ResourceWater                    = getGvr(Water)
	ResourceJob                      = getGvr(Job)
	ResourceDeployment               = getGvr(Deployment)
	ResourceStone                    = getGvr(Stone)
	ResourceStatefulSet              = getGvr(StatefulSet)
	ResourceStatefulSet1             = getGvr(StatefulSet1)
	ResourceDaemonSet                = getGvr(DaemonSet)
	ResourceInjector                 = getGvr(Injector)
	ResourcePod                      = getGvr(Pod)
	ResourceCronJobs                 = getGvr(CronJobs)
	ResourceReplicaSet               = getGvr(ReplicaSet)
	ResourceEvent                    = getGvr(Event)
	ResourceNode                     = getGvr(Node)
	ResourceConfigMaps               = getGvr(ConfigMaps)
	ResourceSecrets                  = getGvr(Secrets)
	ResourceResourceQuota            = getGvr(ResourceQuota)
	ResourceService                  = getGvr(Service)
	ResourceIngress                  = getGvr(Ingress)
	ResourceNetworkPolicy            = getGvr(NetworkPolicy)
	ResourceHorizontalPodAutoscaler  = getGvr(HorizontalPodAutoscaler)
	ResourceCustomResourceDefinition = getGvr(CustomResourceDefinition)
	ResourcePersistentVolume         = getGvr(PersistentVolume)
	ResourcePersistentVolumeClaims   = getGvr(PersistentVolumeClaims)
	ResourceStorageClass             = getGvr(StorageClass)
	ResourceServiceAccount           = getGvr(ServiceAccount)
	ResourceRole                     = getGvr(Role)
	ResourceRoleBinding              = getGvr(RoleBinding)
	ResourceNamespace                = getGvr(Namespace)
	ResourceEndpoint                 = getGvr(Endpoint)
	ResourceClusterRoleBinding       = getGvr(ClusterRoleBinding)
	ResourceWorkloadsTemplate        = getGvr(WorkloadsTemplate)
	ResourceClusterRole              = getGvr(ClusterRole)
	ResourcePage                     = getGvr(Page)
	ResourceForm                     = getGvr(Form)
	ResourceField                    = getGvr(Field)
	ResourceBaseDepartment           = getGvr(BaseDepartment)
	ResourceBasePermission           = getGvr(BasePermission)
	ResourceBaseRole                 = getGvr(BaseRole)
	ResourceBaseUser                 = getGvr(BaseUser)
	ResourceBaseRoleUser             = getGvr(BaseRoleUser)
	ResourcePieline                  = getGvr(Pipeline)
	ResourcePipelineRun              = getGvr(PipelineRun)
	ResourceTask                     = getGvr(Task)
	ResourceTaskRun                  = getGvr(TaskRun)
	ResourcePipelineResource         = getGvr(PipelineResource)
	ResourcePodSecurityPolicie      = getGvr(PodSecurityPolicie)
)

func getGvr(rs ResourceName) schema.GroupVersionResource {
	gvr, exist := GroupVersionResources[rs]
	if !exist {
		panic("try to get an undefined resource")
	}
	return gvr
}
