package dao

import (
	"blog/model"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Manager interface {
	AddUser(user *model.User)
	Login(username string) model.User

	// blog operation
	AddPost(post *model.Post)
	GetAllPost() []model.Post
	GetPost(pid int) model.Post
}

type manager struct {
	db *gorm.DB
}

var Mgr Manager

func init() {
	dsn := "root:1a2b3c4d@tcp(127.0.0.1:3306)/blog?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	Mgr = &manager{db: db}
	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.Post{})
}

func (m *manager) AddUser(user *model.User) {
	m.db.Create(user)
}

func (m *manager) Login(username string) model.User {
	var user model.User
	m.db.Where("username = ?", username).First(&user)
	return user
}

// 博客操作
func (m *manager) AddPost(post *model.Post) {
	m.db.Create(post)
}

func (m *manager) GetAllPost() []model.Post {
	var posts = make([]model.Post, 10)
	m.db.Find(&posts)
	return posts
}

func (m *manager) GetPost(pid int) model.Post {
	var post model.Post
	m.db.First(&post, pid)
	return post
}
