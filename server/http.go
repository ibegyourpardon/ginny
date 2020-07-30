/***********************************************************************************************
 ***                                          GINNY                                          ***
 ***********************************************************************************************/

package server

import (
	"ginny/configs"
	"ginny/routers"

	"github.com/gin-gonic/gin"
)

// 配置并启动服务
func Run(httpServer *gin.Engine) {
	// 服务配置
	serverConfig := configs.GetServerConfig()

	// gin 运行时 release debug test
	gin.SetMode(serverConfig["ENV"])

	httpServer = gin.Default()

	// 注册路由
	routers.InitRouter(httpServer)

	serverAddr := serverConfig["HOST"] + ":" + serverConfig["PORT"]

	// 启动服务
	err := httpServer.Run(serverAddr)

	if nil != err {
		panic("server run error: " + err.Error())
	}
}
