package main

import (
	"fmt"
	"github.com/yametech/fuxi/pkg/service/common"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/util/log"
	"github.com/micro/go-micro/web"
	"github.com/yametech/fuxi/pkg/api/workload/handler"
	"github.com/yametech/fuxi/pkg/preinstall"
	// swagger doc
	file "github.com/swaggo/files"
	swag "github.com/swaggo/gin-swagger"
	_ "github.com/yametech/fuxi/cmd/workload/docs"
)

// @title Gin swagger
// @version 1.0
// @description Gin swagger base
// @contact.name laik author
// @contact.url  github.com/yametech
// @contact.email laik.lj@me.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

const (
	name = "go.micro.api.workload"
	ver  = "latest"
)

func initNeed() (web.Service, *gin.Engine, *gin.RouterGroup, *handler.WorkloadsAPI) {
	service, apiInstallConfigure, err := preinstall.InitApi(50, name, ver, "")
	if err != nil {
		panic(err)
	}
	router := gin.Default()
	common.SharedK8sClient = &apiInstallConfigure.DefaultInstallConfigure
	handler.CreateSharedSessionManager(apiInstallConfigure.ClientV1, apiInstallConfigure.RestConfig)

	return service, router, router.Group("/workload"), handler.NewWorkladAPI()
}

var service, router, group, workloadsAPI = initNeed()

func WrapH(h http.Handler) gin.HandlerFunc {
	return func(c *gin.Context) {
		origin := c.Request.Header.Get("Origin")

		c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT, PATCH, DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Access-Control-Allow-Headers, Origin,Accept, X-Requested-With, Content-Type, Access-Control-Request-Method, Access-Control-Request-Headers,X-Access-Token,XKey,Authorization")

		h.ServeHTTP(c.Writer, c.Request)
	}
}

func main() {
	// #api
	// #v1
	// Pod
	{
		serveHttp := WrapH(handler.CreateAttachHandler("/workload/shell/pod"))
		router.GET("/workload/shell/pod/*path", serveHttp)
		group.GET("/attach/namespace/:namespace/pod/:name/container/:container", PodAttach)
		group.GET("/api/v1/pods", PodList)
		group.GET("/api/v1/namespaces/:namespace/pods/:name", PodGet)
		group.GET("/api/v1/namespaces/:namespace/pods/:name/log", PodLog)
	}

	// Node
	{
		group.GET("/api/v1/nodes", NodeList)
		group.GET("/api/v1/nodes/:node", NodeGet)
	}

	// PersistentVolume
	{
		group.GET("/api/v1/persistentvolumes", PersistentVolumeList)
		group.GET("/api/v1/persistentvolumes/:name", PersistentVolumeGet)
		group.DELETE("/api/v1/persistentvolumes/:name", PersistentVolumeDelete)
	}

	// PersistentVolumeClaims
	{
		group.GET("/api/v1/persistentvolumeclaims", PersistentVolumeClaimsList)
		group.GET("/api/v1/namespaces/:namespace/persistentvolumeclaims/:name", PersistentVolumeClaimsGet)
	}

	// Event
	{
		group.GET("/api/v1/events", EventList)
		group.GET("/api/v1/namespace/:namespace/event/:name", EventGet)
	}

	// ResourceQuota
	{
		group.GET("/api/v1/resourcequotas", ResourceQuotaList)
		group.GET("/api/v1/namespaces/:namespace/resourcequotas/:name", ResourceQuotaGet)
		group.POST("/api/v1/namespaces/:namespace/resourcequotas", workloadsAPI.Apply)
	}

	// Service
	{
		group.GET("/api/v1/services", ServiceList)
		group.GET("/api/v1/namespaces/:namespace/services/:name", ServiceGet)
	}

	// Endpoint
	{
		group.GET("/api/v1/endpoints", EndpointList)
		group.GET("/api/v1/namespaces/:namespace/endpoints/:name", EndpointGet)
	}

	// ServiceAccount
	{
		group.GET("/api/v1/serviceaccounts", ServiceAccountList)
		group.GET("/api/v1/namespaces/:namespace/serviceaccounts/:name", ServiceAccountGet)
		//group.POST("/api/v1/namespaces/:namespace/serviceaccounts", workloadsAPI.Apply)
	}

	// ConfigMaps
	{
		group.GET("/api/v1/configmaps", ConfigMapsList)
		group.GET("/api/v1/namespaces/:namespace/configmaps/:name", ConfigMapsGet)
	}

	// Secret
	{
		group.GET("/api/v1/secrets", SecretList)
		group.GET("/api/v1/namespaces/:namespace/secrets/:name", SecretGet)
		group.POST("/api/v1/namespaces/:namespace/secrets", workloadsAPI.Apply)
	}

	// #apis
	// #apps v1
	// Deployment
	{
		group.GET("/apis/apps/v1/deployments", DeploymentList)
		group.GET("/apis/apps/v1/namespaces/:namespace/deployments/:name", DeploymentGet)
		// deployment scale
		group.GET("/apis/apps/v1/namespaces/:namespace/deployments/:name/scale", workloadsAPI.GetDeploymentScale)
		group.PUT("/apis/apps/v1/namespaces/:namespace/deployments/:name/scale", workloadsAPI.PutDeploymentScale)
	}

	// ReplicaSet
	{
		group.GET("/apis/apps/v1/replicasets", ReplicaSetList)
		group.GET("/apis/apps/v1/namespaces/:namespace/replicasets/:name", ReplicaSetGet)
	}

	// StatefulSet
	{
		group.GET("/apis/apps/v1/statefulsets", StatefulSetList)
		group.GET("/apis/apps/v1/namespaces/:namespace/statefulsets/:name", StatefulSetGet)
	}

	// StatefulSet1
	{
		group.GET("/apis/nuwa.nip.io/v1/statefulsets", StatefulSet1List)
		group.GET("/apis/nuwa.nip.io/v1/namespaces/:namespace/statefulsets/:name", StatefulSet1Get)
	}

	// DaemonSet
	{
		group.GET("/apis/apps/v1/daemonsets", DaemonSetList)
		group.GET("/apis/apps/v1/namespaces/:namespace/daemonsets/:name", DaemonSetGet)
	}

	// #batch
	// #v1beta1

	// CronJob
	{
		group.GET("/apis/batch/v1beta1/cronjobs", CronJobList)
		group.GET("/apis/batch/v1beta1/namespaces/:namespace/cronjobs/:name", CronJobGet)
	}

	// #v1

	// Job
	{
		group.GET("/apis/batch/v1/jobs", JobList) //TODO fix route
		group.GET("/apis/batch/v1/namespaces/:namespace/jobs/:name", JobGet)
	}

	// #extensions
	// #v1beta1

	// Ingress
	{
		group.GET("/apis/extensions/v1beta1/ingresses", IngressList)
		group.GET("/apis/extensions/v1beta1/namespaces/:namespace/ingresses/:name", IngressGet)
	}

	// #networking.k8s.io
	// #v1

	// NetworkPolicy
	{
		group.GET("/apis/networking.k8s.io/v1/networkpolicies", NetworkPolicyList)
		group.GET("/apis/networking.k8s.io/v1/namespaces/:namespace/networkpolicies/:name", NetworkPolicyGet)
	}

	// #storage.k8s.io
	// #v1
	{
		group.GET("/apis/storage.k8s.io/v1/storageclasses", StorageClassList)
		group.GET("/apis/storage.k8s.io/v1/storageclasses/:name", StorageClassGet)
	}

	// #autoscaling
	// #v2beta1
	// HorizontalPodAutoscaler
	{
		group.GET("/apis/autoscaling/v2beta1/horizontalpodautoscalers", HorizontalPodAutoscalerList)
		group.GET("/apis/autoscaling/v2beta1/namespaces/:namespace/horizontalpodautoscalers/:name", HorizontalPodAutoscalerGet)
	}

	// kubernetes RBAC
	// #rbac.authorization.k8s.io
	// v1
	// Roles
	{
		group.GET("/apis/rbac.authorization.k8s.io/v1/roles", RoleList)
		group.GET("/apis/rbac.authorization.k8s.io/v1/namespaces/:namespace/roles/:name", RoleGet)
	}
	// #rbac.authorization.k8s.io
	// #v1
	// Clusterroles
	{
		group.GET("/apis/rbac.authorization.k8s.io/v1/clusterroles", ClusterRoleList)
		group.GET("/apis/rbac.authorization.k8s.io/v1/namespaces/:namespace/clusterroles/:name", ClusterRoleGet)
		group.POST("/apis/rbac.authorization.k8s.io/v1/clusterroles", workloadsAPI.Apply)
	}

	// RoleBinding
	{
		group.GET("/apis/rbac.authorization.k8s.io/v1/rolebindings", RoleBindingList)
		group.GET("/apis/rbac.authorization.k8s.io/v1/namespaces/:namespace/rolebindings/:name", RoleBindingGet)
		group.POST("/apis/rbac.authorization.k8s.io/v1/namespaces/:namespace/rolebindings", workloadsAPI.Apply)
	}

	// ClusterRolebind
	{
		group.GET("/apis/rbac.authorization.k8s.io/v1/clusterrolebindings", ClusterRoleBindList)
		group.GET("/apis/rbac.authorization.k8s.io/v1/namespaces/:namespace/clusterrolebindings/:name ", ClusterRoleBindGet)
		group.POST("/apis/rbac.authorization.k8s.io/v1/namespaces/:namespace/clusterrolebindings", workloadsAPI.Apply)
	}

	// #policy
	// #v1beta1
	// #podsecuritypolicies
	{
		group.GET("/apis/policy/v1beta1/podsecuritypolicies", PodSecurityPolicieList)
		group.GET("/apis/policy/v1beta1/namespaces/:namespace/podsecuritypolicies/:name", PodSecurityPolicieGet)
		group.POST("/apis/policy/v1beta1/namespaces/:namespace/podsecuritypolicies", workloadsAPI.Apply)
	}

	// #apiextensions.k8s.io/v1beta1
	// #v1beta1
	// CustomResourceDefinition
	{
		group.GET("/apis/apiextensions.k8s.io/v1beta1/customresourcedefinitions", CustomResourceDefinitionList)

		ignores := []string{
			"fuxi.nip.io/v1/formrenders",
			"fuxi.nip.io/v1/workloads",
			"fuxi.nip.io/v1/fields",
			"fuxi.nip.io/v1/forms",
			"fuxi.nip.io/v1/pages",
			"fuxi.nip.io/v1/basedepartments",
			"fuxi.nip.io/v1/basepermissions",
			"fuxi.nip.io/v1/baseroles",
			"fuxi.nip.io/v1/baseusers",
			"nuwa.nip.io/v1/statefulsets",
			"tekton.dev/v1alpha1/pipelines",
			"tekton.dev/v1alpha1/pipelineruns",
			"tekton.dev/v1alpha1/pipelineresources",
			"tekton.dev/v1alpha1/tasks",
			"tekton.dev/v1alpha1/taskruns",
		}
		apiVersions, err := workloadsAPI.ListCustomResourceRouter(ignores)
		if err != nil {
			panic(err)
		}
		routerPath := "/apis/%s"
		for _, apiVersion := range apiVersions {
			relativePath := fmt.Sprintf(routerPath, apiVersion)
			group.GET(relativePath, GeneralCustomResourceDefinitionList)
		}
	}

	// Namespace
	{
		group.GET("/api/v1/namespaces", NamespaceList)
		group.GET("/api/v1/namespaces/:namespace", NamespaceGet)
		group.POST("/api/v1/namespaces", NamespaceCreate)
		group.DELETE("/api/v1/namespaces/:namespace", NamespaceDelete)

	}

	// post  workload/stack
	{
		// all resource apply
		group.POST("/stack", workloadsAPI.Apply)

		// other resource  api/apis resource
		group.DELETE("/api/v1/namespaces/:namespace/:resource/:name", workloadsAPI.Delete)
		group.DELETE("/apis/:group/:version/namespaces/:namespace/:resource/:name", workloadsAPI.Delete)
	}

	// fuxi.nip.io
	// FormRender
	{
		group.GET("/apis/fuxi.nip.io/v1/formrenders", FormRenderList)
		group.GET("/apis/fuxi.nip.io/v1/namespaces/:namespace/formrenders/:name", FormRenderGet)
	}

	// fuxi.nip.io
	// Workloads
	{
		group.GET("/apis/fuxi.nip.io/v1/workloads", WorkloadsTemplateList)
		group.GET("/apis/fuxi.nip.io/v1/namespaces/:namespace/workloads/:name", WorkloadsTemplateGet)
		group.POST("/apis/fuxi.nip.io/v1/workloads", WorkloadsTemplateCreate)
	}

	// Field
	{
		group.GET("/apis/fuxi.nip.io/v1/fields", FieldList)
		group.GET("/apis/fuxi.nip.io/v1/namespaces/:namespace/fields/:name", FieldGet)
		group.POST("/apis/fuxi.nip.io/v1/namespaces/:namespace/fields", FieldCreate)
	}

	// Form
	{
		group.GET("/apis/fuxi.nip.io/v1/forms", FormList)
		group.GET("/apis/fuxi.nip.io/v1/namespaces/:namespace/forms/:name", FormGet)
		group.POST("/apis/fuxi.nip.io/v1/namespaces/:namespace/forms", FormCreate)
	}

	// Page
	{
		group.GET("/apis/fuxi.nip.io/v1/pages", PageList)
		group.GET("/apis/fuxi.nip.io/v1/namespaces/:namespace/pages/:name", PageGet)
		group.POST("/apis/fuxi.nip.io/v1/namespaces/:namespace/pages", PageCreate)
	}

	// Metrics
	{
		group.POST("/metrics", workloadsAPI.Metrics)
		group.GET("/apis/metrics.k8s.io/v1beta1/nodes", workloadsAPI.NodeMetrics)

		group.GET("/apis/metrics.k8s.io/v1beta1/pods", workloadsAPI.PodMetricsList)
		// GET /workload/apis/metrics.k8s.io/v1beta1/namespaces/rook-ceph/pods
		group.GET("/apis/metrics.k8s.io/v1beta1/namespaces/:namespace/pods", workloadsAPI.PodMetrics)
	}

	{
		group.GET("/config", func(g *gin.Context) {
			g.JSON(http.StatusOK, gin.H{
				"lensVersion":       "1.0",
				"lensTheme":         "",
				"userName":          "admin",
				"token":             "",
				"allowedNamespaces": "[]",
				"isClusterAdmin":    true,
				"chartEnable":       true,
				"kubectlAccess":     true,
			})
		})
	}

	// tekton.dev
	// v1alpha1
	{
		// pipeline
		group.GET("/apis/tekton.dev/v1alpha1/pipelines", PipelineList)
		group.GET("/apis/tekton.dev/v1alpha1/namespaces/:namespace/pipelines/:name", PipelineGet)

		// pipelineRun
		group.GET("/apis/tekton.dev/v1alpha1/pipelineruns", PipelineRunList)
		group.GET("/apis/tekton.dev/v1alpha1/namespaces/:namespace/pipelineruns/:name", PipelineRunGet)

		// task
		group.GET("/apis/tekton.dev/v1alpha1/tasks", TaskList)
		group.GET("/apis/tekton.dev/v1alpha1/namespaces/:namespace/tasks/:name", TaskGet)

		// taskRun
		group.GET("/apis/tekton.dev/v1alpha1/taskruns", TaskRunList)
		group.GET("/apis/tekton.dev/v1alpha1/namespaces/:namespace/taskruns/:name", TaskRunGet)

		// pipelineResource
		group.GET("/apis/tekton.dev/v1alpha1/pipelineresources", PipelineResourceList)
		group.GET("/apis/tekton.dev/v1alpha1/namespaces/:namespace/pipelineresources/:name", PipelineResourceGet)
	}

	// watch the group resource
	{
		group.GET("/watch", WatchStream)
	}

	// Swag
	{
		/// Then, if you set envioment variable DEV_OPEN_SWAGGER to anything, /swagger/*any will respond 404, just like when route unspecified.
		/// Release production environment can be turned on
		group.GET("/swagger/*any", swag.DisablingWrapHandler(file.Handler, "DEV_OPEN_SWAGGER"))
	}

	service.Handle("/", router)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
