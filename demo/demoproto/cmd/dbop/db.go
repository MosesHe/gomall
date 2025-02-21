package main

import (
	"github.com/MosesHe/gomall/demo/demoproto/biz/dal"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	dal.Init()
	// CURD
	// CREATE
	// mysql.DB.Create(&model.User{Email: "demo@example.com", Password: "roottoor"})

	// UPDATE
	// mysql.DB.Model(&model.User{}).Where("email = ?", "demo@example.com").Update("password", "toorroot")

	// READ
	// var row model.User
	// mysql.DB.Model(&model.User{}).Where("email = ?", "demo@example.com").First(&row)
	// fmt.Printf("row: %+v\n", row)

	// DELETE
	// mysql.DB.Where("email = ?", "demo@example.com").Delete(&model.User{})
	// mysql.DB.Unscoped().Where("email = ?", "demo@example.com").Delete(&model.User{})
}
