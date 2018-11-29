package cnto

//
//import (
//	"net/http"
//
//	"github.com/gin-gonic/gin"
//	"github.com/hal-ms/driver/building/model"
//	"github.com/hal-ms/driver/building/service"
//)
//
//func Scene(c *gin.Context) {
//	scene := c.Param("scene")
//	var d model.Scene
//
//	err := c.BindJSON(&d)
//	if err != nil {
//		c.JSON(http.StatusBadRequest, "NG")
//		return
//	}
//
//	service.Scene.Write(scene, d)
//	c.JSON(http.StatusOK, "OK")
//}
