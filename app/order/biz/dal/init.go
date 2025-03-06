package dal

import (
	"github.com/MosesHe/gomall/app/order/biz/dal/mysql"
	"github.com/MosesHe/gomall/app/order/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
