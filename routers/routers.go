package routers

import (
	"github.com/gin-gonic/gin"
	"objectStorageServer/api"
)

func InitRouter() (r *gin.Engine) {
	r = gin.Default()
	r.GET("/putObject", api.PutObjectHandler)
	r.GET("/getObject", api.GetObjectHandler)
	return
}
