package dal

import (
	"github.com/MosesHe/gomall/app/cart/biz/dal/mysql"
	"github.com/MosesHe/gomall/app/cart/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
