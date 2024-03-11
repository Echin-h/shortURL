package function

import (
	"ShortUrl/internal/global/database"
	"ShortUrl/internal/global/errors"
	"ShortUrl/internal/global/log"
	"ShortUrl/internal/model"
	"ShortUrl/tools"
	"github.com/gin-gonic/gin"
	"github.com/spaolacci/murmur3"
	"time"
)

func Shorten(c *gin.Context) {
	longUrl := c.PostForm("longUrl")
	if longUrl == "" {
		errors.Fail(c, errors.REQUEST_ERROR.WithMessage("longUrl can't be empty"))
	}
	u := &model.ShortUrl{}
	var count int64 = 0 // 用于判断是否存在
	// long是保留关键字，需要加引号
	database.Url.Model(&model.ShortUrl{}).Where("`long` = ?", longUrl).First(&u).Count(&count)
	if count > 0 {
		log.SugarLogger.Info("short url is already exist")
		errors.Success(c, "short ult is already exist:", u.Short)
		return
	}
	num := murmur3.Sum64([]byte("longUrl"))
	ShortUrl := tools.Convert(num)
	u = &model.ShortUrl{
		Count:  0,
		Long:   longUrl,
		Short:  ShortUrl,
		Salt:   "",
		Expire: time.Now().Add(time.Hour * 24 * 30),
	}
	database.Url.Model(&model.ShortUrl{}).Where("short = ?", u.Short).Count(&count)
	var salt string
	for cnt := 0; count > 0 && cnt <= 3; cnt++ {
		if cnt == 3 {
			log.SugarLogger.Info("add salt is up to three times")
			errors.Fail(c, errors.SERVER_ERROR.WithMessage("add salt is up to three times"))
		}
		salt = tools.GenSalt()
		newLongUrl := longUrl
		newLongUrl += salt

		num = murmur3.Sum64([]byte(newLongUrl))
		ShortUrl = tools.Convert(num)

		u.Short = ShortUrl
		u.Salt = salt
		u.Long = newLongUrl
		database.Url.Model(&model.ShortUrl{}).Where("short = ?", u.Short).Count(&count)
	}
	err := database.Url.Model(&model.ShortUrl{}).Create(u).Error
	if err != nil {
		log.SugarLogger.Error("shorten error", err)
		errors.Fail(c, errors.DB_ERROR.WithOrigin(err))
	}
	log.SugarLogger.Info("shorten success")
	errors.Success(c, "shorten:", u.Short)
}
