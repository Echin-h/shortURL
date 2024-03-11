package function

import (
	"ShortUrl/internal/global/database"
	"ShortUrl/internal/global/errors"
	"ShortUrl/internal/global/log"
	"ShortUrl/internal/model"
	"ShortUrl/tools"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"time"
)

/* TODO: 短链接转化成长链接并且转到长链接
1. 从数据库中查询出长链接
2. 如果过期时间不为空，那么就需要判断是否过期,如果过期了，那么就需要删除
3. 如果salt值不为空，那么就需要解密
4. 如果salt值为空，实现跳转的操作
5. 还有什么情况需要处理？
6. 如果短链接不存在，那么就需要返回错误信息
7. 如果短链接存在，那么就需要增加访问次数
*/

func Longer(c *gin.Context) {
	shortUrl := c.Param("short")
	uri := strings.TrimPrefix(shortUrl, "/")
	// 从数据库中查询出长链接
	u := &model.ShortUrl{}
	err2 := database.Url.Model(&model.ShortUrl{}).Where("short = ?", uri).First(u).Error
	if err2 != nil {
		errors.Fail(c, errors.DB_ERROR.WithOrigin(err2))
		return
	}
	today := time.Now().UnixNano()
	if today > u.Expire.UnixNano() {
		// 如果过期时间不为空，那么就需要判断是否过期,如果过期了，那么就需要删除
		database.Url.Model(&model.ShortUrl{}).Where("short = ?", uri).Delete(u)
		errors.Fail(c, errors.REQUEST_ERROR.WithMessage("short url is expired"))
		return
	}
	slice := make([]byte, 0)
	if u.Salt != "" {
		index := tools.GenKmp(u.Long, u.Salt)
		if index == -1 {
			log.SugarLogger.Warnln("salt is not exist")
			return
		} else if index == 0 {
			log.SugarLogger.Warnln("longUrl should not be empty")
			return
		} else {
			slice = append(slice, u.Long[:index]...)
		}
	} else {
		slice = append(slice, u.Long...)
	}
	c.Redirect(http.StatusFound, string(slice))
	errors.Success(c, "longUrl:", string(slice))

}
