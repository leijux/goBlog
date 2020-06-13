package apis

import (
	"fmt"
	"strconv"
	"time"

	"goBlog/log"
	"goBlog/middleware"
	"goBlog/models/blog"
	"goBlog/models/user"
	"goBlog/src/common"

	"github.com/gin-gonic/gin"
)

//AddBlogAPI 添加博客
func AddBlogAPI(c *gin.Context) {
	var b blog.Blog
	err := c.Bind(&b)
	if err != nil {
		msg := fmt.Sprintln("shoul bind err")
		log.Logger.Errorln(err)
		common.Rmsg(c, false, msg)
		return
	}
	t := time.Now()
	b.Created = t
	b.Updated = t
	u, _ := c.Get(middleware.AuthMiddleware.IdentityKey) //得到用户信息
	if v, ok := u.(*user.User); ok {
		if v.Emeil == b.Author {
			id, err := b.AddBlog()
			if err != nil {
				msg := fmt.Sprintln("add blog err")
				log.Logger.Errorln(err)
				common.Rmsg(c, false, msg)
				return
			}
			msg := fmt.Sprintf("add blog id")
			common.Rmsg(c, true, msg, id)
		} else {
			msg := fmt.Sprintf("add blog err")
			common.Rmsg(c, false, msg)
			return
		}
	}
}

//GetBlogsAPI 有作者就返回作者的分页  没有就返回全部的分页
func GetBlogsAPI(c *gin.Context) { 
	page := c.DefaultQuery("page", "1")
	perpPage := c.DefaultQuery("per_page", "5")
	author := c.Query("author")
	p, err := strconv.Atoi(page)
	pp, err := strconv.Atoi(perpPage)

	if err != nil {
		msg := fmt.Sprintln("get blog err")
		log.Logger.Errorln(msg, err)
		common.Rmsg(c, false, msg)
		return
	}
	if author == "" {
		var b blog.Blog
		//page=2&per_page=100
		bs, err := b.GetBlogs(p, pp)
		if err != nil {
			common.Rmsg(c, false, err.Error())
			return
		}
		common.Rmsg(c, true, "", bs)
	} else {
		b := new(blog.Blog)
		//page=2&per_page=100
		b.Author = author
		bs, err := b.AuthoeToBlogs(p, pp)
		if err != nil {
			common.Rmsg(c, false, err.Error())
			return
		}
		common.Rmsg(c, true, "", bs)
		return
	}
}
//UpGlog 更新文章
func UpGlog(){

}
//BlogSizeAPI 得到文章数量
func BlogSizeAPI(c *gin.Context) {
	var b blog.Blog
	a, err := b.Count()
	if err != nil {
		common.Rmsg(c, false, "")
		return
	}
	common.Rmsg(c, true, "", a)
}

//GetTopAPI 得到to排名
func GetTopAPI(c *gin.Context) {
	var b blog.Blog
	bs, err := b.GetTop()
	if err != nil {
		common.Rmsg(c, false, "")
		return
	}
	common.Rmsg(c, true, "", bs)
}
