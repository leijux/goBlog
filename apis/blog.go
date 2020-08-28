package apis

import (
	"fmt"
	"strconv"

	"goBlog/log"
	"goBlog/middleware"
	"goBlog/models/blog"
	"goBlog/models/user"
	"goBlog/src/common"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

//AddBlogAPI 添加博客
func AddBlogAPI(c *gin.Context) {
	var b blog.Blog
	err := c.Bind(&b)
	if err != nil {
		const msg string = "should bind err"
		log.Logger.Errorln(err)
		common.Rmsg(c, false, msg)
		return
	}
	u, _ := c.Get(middleware.GetIdentityKey()) //得到用户信息
	if v, ok := u.(*user.UserApi); ok {
		bol, err := addBlogAPI(v, b)
		common.Rmsg(c, bol, err.Error())
		return
	}
}

func addBlogAPI(v *user.UserApi, b blog.Blog) (bool, error) {
	if v.Email == b.Email {
		//TODO  id没有处理
		_, err := b.AddBlog()
		if err != nil {
			const msg string = "add blog err"
			log.Logger.Errorln(err)
			//common.Rmsg(c, false, msg)
			return false, errors.Wrap(err, msg)
		}
		const msg string = "add blog id"
		// common.Rmsg(c, true, msg, id)
		return false, errors.New(msg)
	} else {
		const msg string = "add blog err"
		// common.Rmsg(c, false, msg)
		return false, errors.New(msg)
	}
}

//GetBlogsAPI 有作者就返回作者的分页  没有就返回全部的分页
func GetBlogsAPI(c *gin.Context) {
	page := c.DefaultQuery("page", "1")
	perpPage := c.DefaultQuery("per_page", "5")
	email := c.Query("email")
	p, err := strconv.Atoi(page)
	pp, err := strconv.Atoi(perpPage)

	if pp <= 0 || p <= 0 {
		err = errors.New("page/perpPage Less than 0")
	}
	if err != nil {
		msg := fmt.Sprintln("get blog err")
		log.Logger.Errorln(msg, err)
		common.Rmsg(c, false, msg)
		return
	}
	if email == "" {
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
		b.Email = email
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
func UpGlog() {

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
