package main

import (
	"log"
	"objectStorageServer/command"
	"objectStorageServer/routers"
)

func main() {
	//gin.SetMode(gin.ReleaseMode)
	r := routers.InitRouter()
	err := r.Run(":" + command.Port)
	if err != nil {
		log.Println(err)
	}
	return
}
