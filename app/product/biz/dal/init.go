package dal

import (
	"github.com/MosesHe/gomall/app/product/biz/dal/mysql"
	"github.com/MosesHe/gomall/app/product/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
