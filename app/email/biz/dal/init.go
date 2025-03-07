package dal

import (
	"github.com/MosesHe/gomall/app/email/biz/dal/mysql"
	"github.com/MosesHe/gomall/app/email/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
