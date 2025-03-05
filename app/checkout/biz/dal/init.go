package dal

import (
	"github.com/MosesHe/gomall/app/checkout/biz/dal/mysql"
	"github.com/MosesHe/gomall/app/checkout/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
