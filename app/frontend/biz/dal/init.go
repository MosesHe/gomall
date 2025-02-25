package dal

import (
	"github.com/MosesHe/gomall/app/frontend/biz/dal/mysql"
	"github.com/MosesHe/gomall/app/frontend/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
