package middleware

import (
	"binghambai.com/lowCodePlatform-mobile/app/common"
	"github.com/gin-gonic/gin"
	"log"
)

func CatchError() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if r := recover(); !common.IsNil(r) {
				switch t := interface{}(r).(type) {
				case nil:
					break
				case *common.Response:
					log.Printf("panic: %v\n", t.Msg)
					common.Error(c, 500, t.Msg)
				default:
					log.Printf("panic: internal error; %v\n", t)
					common.Error(c, 500, "服务器内部错误")
				}
				c.Abort()
			}
		}()

		c.Next()
	}
}
