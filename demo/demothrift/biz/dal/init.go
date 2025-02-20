package dal

import (
	"github.com/MosesHe/gomall/demo/demothrift/biz/dal/mysql"
	"github.com/MosesHe/gomall/demo/demothrift/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
