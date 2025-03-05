package dal

import (
	"github.com/MosesHe/gomall/app/payment/biz/dal/mysql"
	"github.com/MosesHe/gomall/app/payment/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
