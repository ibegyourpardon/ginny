/***********************************************************************************************
 ***                                          GINNY                                          ***
 ***********************************************************************************************/

package routers

import (
	"github.com/gin-gonic/gin"
	apps "ginny/applications"
)

func InitRouter(router *gin.Engine) {

	//router := gin.Default()
	//router.GET("/ping", apps.Ping)
	wechat := router.Group("/wechat") // 使用微信作为一个示例
	{
		wechat.GET("/ping", apps.Ping)
	}
	//return router
}
