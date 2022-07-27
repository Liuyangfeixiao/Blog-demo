# Blog-demo
使用gin, gorm, mysql构建

具体项目结构如下图

```text
// gin中的处理函数
|- controller
// 涉及到数据库访问的内容
|- dao
// 路由，访问路径关联到controller
|- router
// 生成后台数据库的gin的结构体
|- model
// 保存静态资源
|- assets
// 模板文件
|- templates
```
