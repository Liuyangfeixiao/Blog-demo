package controller

import (
	"blog/dao"
	"blog/model"
	"fmt"
	"html/template"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/russross/blackfriday/v2"
)

func Index(c *gin.Context) {
	c.HTML(200, "index.html", nil)
}

func ListUser(c *gin.Context) {
	c.HTML(200, "userlist.html", nil)
}

func Register(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	user := model.User{
		Username: username,
		Password: password,
	}

	dao.Mgr.AddUser(&user)

	c.Redirect(301, "/")
}

func GoRegister(c *gin.Context) {
	c.HTML(200, "register.html", nil)
}

func GoLogin(c *gin.Context) {
	c.HTML(200, "login.html", nil)
}
func Login(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	fmt.Println(username)
	u := dao.Mgr.Login(username)

	if u.Username == "" {
		c.HTML(200, "login.html", "用户名不存在！")
		fmt.Println("用户名不存在")
	} else {
		if u.Password != password {
			fmt.Println("密码错误")
			c.HTML(200, "login.html", "密码错误")
		} else {
			fmt.Println("登录成功")
			c.Redirect(301, "/")
		}

	}
}

// Blog Operation
func GetPostIndex(c *gin.Context) {
	// username := c.PostForm("username")
	posts := dao.Mgr.GetAllPost()
	c.HTML(200, "postindex.html", posts)
}

func AddPost(c *gin.Context) {
	title := c.PostForm("title")
	tag := c.PostForm(("tag"))
	content := c.PostForm("content")

	post := model.Post{
		Title:   title,
		Tag:     tag,
		Content: content,
	}

	dao.Mgr.AddPost(&post)

	c.Redirect(302, "/post_index")
}

func GoAddPost(c *gin.Context) {
	c.HTML(200, "post.html", nil)
}

func PostDetails(c *gin.Context) {
	s := c.Query("pid")

	pid, _ := strconv.Atoi(s)

	p := dao.Mgr.GetPost(pid)

	content := blackfriday.Run([]byte(p.Content))

	c.HTML(200, "postDetails.html", gin.H{
		"Title":   p.Title,
		"Content": template.HTML(content),
	})
}
