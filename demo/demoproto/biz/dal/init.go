package dal

import (
	"github.com/MosesHe/gomall/demo/demoproto/biz/dal/mysql"
)

func Init() {
	// redis.Init()
	mysql.Init()
}
