/***********************************************************************************************
 ***                                          GINNY                                          ***
 ***********************************************************************************************/

package applications

import (
	"github.com/gin-gonic/gin"
)

// use "Ping" showing my respect
func Ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "it woks",
	})
}
