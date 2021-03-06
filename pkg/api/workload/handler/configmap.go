package handler

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	v1 "k8s.io/api/core/v1"
)

// Get ConfigMaps
func (w *WorkloadsAPI) GetConfigMaps(g *gin.Context) {
	namespace := g.Param("namespace")
	name := g.Param("name")
	item, err := w.configMaps.Get(namespace, name)
	if err != nil {
		toInternalServerError(g, "", err)
		return
	}
	g.JSON(http.StatusOK, item)
}

// List ConfigMaps
func (w *WorkloadsAPI) ListConfigMaps(g *gin.Context) {
	list, err := w.configMaps.List("", "", 0, 0, nil)
	if err != nil {
		toInternalServerError(g, "", err)
		return
	}
	configMapList := &v1.ConfigMapList{}
	marshalData, err := json.Marshal(list)
	if err != nil {
		toInternalServerError(g, "", err)
		return
	}
	err = json.Unmarshal(marshalData, configMapList)
	if err != nil {
		toInternalServerError(g, "", err)
		return
	}
	g.JSON(http.StatusOK, configMapList)
}
