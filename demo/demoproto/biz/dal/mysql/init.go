package mysql

import (
	"fmt"
	"os"

	"github.com/MosesHe/gomall/demo/demoproto/biz/model"
	"github.com/MosesHe/gomall/demo/demoproto/conf"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func Init() {
	dns := fmt.Sprintf(conf.GetConf().MySQL.DSN, os.Getenv("MYSQL_USER"),
		os.Getenv("MYSQL_PASSWORD"),
		os.Getenv("MYSQL_HOST"),
		os.Getenv("MYSQL_DATABASE"),
	)
	DB, err = gorm.Open(mysql.Open(dns),
		&gorm.Config{
			PrepareStmt:            true,
			SkipDefaultTransaction: true,
		},
	)
	if err != nil {
		panic(err)
	}

	type Version struct {
		Version string
	}

	var v Version

	err = DB.Raw("select version() as version").Scan(&v).Error
	if err != nil {
		panic(err)
	}
	DB.AutoMigrate(&model.User{})
	fmt.Println(v)
}
