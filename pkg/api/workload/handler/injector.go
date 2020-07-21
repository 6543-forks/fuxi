package handler

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/yametech/fuxi/pkg/api/common"
	nuwav1 "github.com/yametech/nuwa/api/v1"
	"net/http"
)

// Get Injector
func (w *WorkloadsAPI) GetInjector(g *gin.Context) {
	namespace := g.Param("namespace")
	name := g.Param("name")

	item, err := w.statefulSet1.Get(namespace, name)
	if err != nil {
		common.ResourceNotFoundError(g, "", err)
		return
	}
	g.JSON(http.StatusOK, item)
}

// List StatefulSet
func (w *WorkloadsAPI) ListInjector(g *gin.Context) {
	list, err := resourceList(g, w.injector)
	if err != nil {
		common.ToInternalServerError(g, "", err)
		return
	}
	statefulSetList := &nuwav1.StatefulSetList{}
	marshalData, err := json.Marshal(list)
	if err != nil {
		common.ToInternalServerError(g, "", err)
		return
	}
	err = json.Unmarshal(marshalData, statefulSetList)
	if err != nil {
		common.ToInternalServerError(g, "", err)
		return
	}
	g.JSON(http.StatusOK, statefulSetList)
}
