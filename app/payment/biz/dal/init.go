package dal

import (
	"github.com/MosesHe/gomall/app/payment/biz/dal/mysql"
)

func Init() {
	// redis.Init()
	mysql.Init()
}
