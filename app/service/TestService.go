package service

import (
	"binghambai.com/lowCodePlatform-mobile/app/conn"
	"binghambai.com/lowCodePlatform-mobile/app/model"
	"github.com/gin-gonic/gin"
)
import "binghambai.com/lowCodePlatform-mobile/app/common"

func TestService(c *gin.Context) {
	db := conn.Db
	var shortUrl model.ShortUrl
	db.First(&shortUrl, 838092237798912000)

	//返回hello world字符串
	common.Success(c, shortUrl)
}
