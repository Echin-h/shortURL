package model

import (
	ulid "github.com/oklog/ulid/v2"
	"gorm.io/gorm"
	"time"
)

type ShortUrl struct {
	Base
	Count  int       `json:"count" gorm:"comment:访问次数;default:0"`
	Long   string    `json:"long" gorm:"not null;comment:原始链接"`
	Short  string    `json:"short" gorm:"type:varchar(255);unique;not null;comment:短链接"`
	Salt   string    `json:"salt" gorm:"type:varchar(255);comment:解决hash冲突问题的盐值"`
	Expire time.Time `json:"expire" gorm:"comment:过期时间"`
}

func (u *ShortUrl) BeforeCreate(*gorm.DB) error {
	u.ID = ulid.Make().String()
	return nil
}

/* 步骤
1. 使用MurMurHash算法生成短链接
2. 使用十进制转六十二进制的方法生成短链接
3. 使用数据库存储短链接，短链接为唯一索引
4. salt值解决hash冲突问题(怎么生成这个salt值还没想好)
5. 过期时间设置
6.

*/
