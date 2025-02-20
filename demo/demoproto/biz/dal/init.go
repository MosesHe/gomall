package dal

import (
	"github.com/MosesHe/gomall/demo/demoproto/biz/dal/mysql"
	"github.com/MosesHe/gomall/demo/demoproto/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
