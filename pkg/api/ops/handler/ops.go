package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yametech/fuxi/pkg/logging"
	"github.com/yametech/fuxi/pkg/service/ops"
)

type OpsController struct {
	Service ops.OpsService
}

func (o *OpsController) getUserName(c *gin.Context) string {
	userName := c.Request.Header.Get("x-auth-username")
	if len(userName) < 1 {
		c.AbortWithStatusJSON(http.StatusBadRequest, &gin.H{
			"msg":  "authorization not correct",
			"code": http.StatusBadRequest,
			"data": "",
		})
		return ""
	}
	return userName

}

//CheckParams check params if not null
func (o *OpsController) CheckParams(c *gin.Context) (bool, string, string) {
	namespace := c.Param("namespace")
	name := c.Param("name")
	if namespace == "" && name == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg":  "get task list error: namespace or name cannot be empty",
			"code": http.StatusBadRequest,
			"data": "",
		})
		return true, "", ""
	}
	return false, namespace, name
}

//ListRepos return current user department all repos
func (o *OpsController) ListRepos(c *gin.Context) {

	userName := o.getUserName(c)
	if userName == "" {
		return
	}

	namespace := c.Param("namespace")
	if namespace == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg":  "list repos error: namespace cannot be empty",
			"code": http.StatusBadRequest,
			"data": "",
		})
		return
	}

	repos, err := o.Service.ListRepos(userName, namespace)
	if err != nil {
		logging.Log.Error("ListRepos error:" + err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg":  "get repos failed",
			"code": http.StatusInternalServerError,
			"data": "",
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"msg":  "get repos success",
		"code": http.StatusOK,
		"data": repos,
	})
}

//ListBranchs return current user department  all branchs
func (o *OpsController) ListBranchs(c *gin.Context) {
	userName := o.getUserName(c)
	if userName == "" {
		return
	}

	namespace := c.Param("namespace")
	if namespace == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg":  "list branchs error: namespace cannot be empty",
			"code": http.StatusBadRequest,
			"data": "",
		})
		return
	}

	branchs, err := o.Service.ListBranchs(userName, namespace)

	if err != nil {
		logging.Log.Error("ListBranchs error:" + err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg":  "get branch error" + err.Error(),
			"code": http.StatusInternalServerError,
			"data": "",
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"msg":  "get branch success",
		"code": http.StatusOK,
		"data": branchs,
	})
}
