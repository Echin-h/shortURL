package tools

import (
	"math/rand"
	"strconv"
	"time"
)

// 将十进制数转化为六十二进制的数
const charset = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

// 除法+倒序  这里的代码需要修改，因为长度不符合规定
func Convert(num uint64) string {
	var result string
	for num > 0 {
		temp := num % 62
		result = string(charset[temp]) + result
		num = num / 62
	}
	return result
}

func GenSalt() string {
	rand.Seed(time.Now().UnixNano())
	random := rand.Intn(1000)
	return strconv.Itoa(random)
}
