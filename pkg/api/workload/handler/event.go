package handler

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	corev1 "k8s.io/api/core/v1"
	"net/http"
	"strconv"
)

// Get Event
func (w *WorkloadsAPI) GetEvent(g *gin.Context) {
	namespace := g.Param("namespace")
	name := g.Param("name")
	item, err := w.event.Get(namespace, name)
	if err != nil {
		toInternalServerError(g, "", err)
		return
	}
	g.JSON(http.StatusOK, item)
}

// List Event
func (w *WorkloadsAPI) ListEvent(g *gin.Context) {
	limit := g.Param("limit")
	limitNum := int64(10000)
	var err error
	if limit != "" {
		limitNum, err = strconv.ParseInt(limit, 64, 10)
		if err != nil {
			toRequestParamsError(g, err)
			return
		}
	}
	list, err := w.event.List("", "", 0, limitNum, nil)
	if err != nil {
		toInternalServerError(g, "", err)
		return
	}
	eventList := &corev1.EventList{}
	marshalData, err := json.Marshal(list)
	if err != nil {
		toInternalServerError(g, "", err)
		return
	}
	err = json.Unmarshal(marshalData, eventList)
	if err != nil {
		toInternalServerError(g, "", err)
		return
	}

	g.JSON(http.StatusOK, eventList)
}
