package user

import (
	"gorm.io/gorm"
)

type Blog struct {
	gorm.Model
	Title       string
	Blog        string
	Upvotes     int32
	BlogComment BlogComment `gorm:"embedded`
}

type BlogCommnent struct {
	gorm.Model
	Desc    string
	upvotes int32
	BlogID  uint
	Blog    Blog `gorm:"foreignKey:BlogID"`
}

func CreateBlog() {
	db := database.DBConn
	blog := Blog{}

	db.Create()
}
