package apis

import (
	"goBlog/models"
	"strconv"

	"goBlog/log"
	"goBlog/middleware"
	"goBlog/src/common"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

//AddBlogAPI 添加博客
func AddBlogAPI(c *gin.Context) {
	b := models.NewBlog()
	err := c.Bind(&b)
	if err != nil {
		const msg string = "should bind err"
		log.Errorf("%+v", err)
		common.Rmsg(c, false, msg)
		return
	}
	u, _ := c.Get(middleware.GetIdentityKey()) //得到用户信息
	if v, ok := u.(models.UserApi); ok {
		bol, err := addBlogAPI(v, b)
		if err != nil {
			log.Errorf("%+v", err)
			common.Rmsg(c, bol, err.Error())
			return
		}
		const msg = "success !"
		common.Rmsg(c, bol, msg)
		return
	}
}

func addBlogAPI(v models.UserApi, b models.BlogApi) (bool, error) {
	if v.Email == b.Email {
		bol, err := b.CreateBlog()
		if err != nil {
			const msg string = "add blog err"
			return bol, errors.Wrap(err, msg)
		}
		return bol, nil
	} else {
		const msg string = "add blog err"
		return false, errors.New(msg)
	}
}

//GetBlogsAPI 有作者就返回作者的分页  没有就返回全部的分页
func GetBlogsAPI(c *gin.Context) {
	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil {
		const msg = "page err"
		log.Error(msg,
			zap.Int("page", page),
		)
		common.Rmsg(c, false, msg)
		return
	}
	if page <= 0 {
		page = 1
	}
	pageSize, err := strconv.Atoi(c.DefaultQuery("page_size", "5"))
	if err != nil {
		const msg = "pageSize err"
		log.Error(msg,
			zap.Int("page_size", pageSize),
		)
		common.Rmsg(c, false, msg)
		return
	}
	switch {
	case pageSize > 50:
		pageSize = 50
	case pageSize <= 0:
		pageSize = 5
	}

	authoe := c.Query("authoe")
	if authoe != "" {
		bs, err := getAuthoeBlogs(authoe, page, pageSize)
		if err != nil {
			log.Errorf("%+v", err)
			common.Rmsg(c, false, err.Error())
			return
		}
		common.Rmsg(c, true, "", bs)
		return
	}
	bs, err := getBlogs(page, pageSize)
	if err != nil {
		log.Errorf("%+v", err)
		common.Rmsg(c, false, err.Error())
		return
	}
	common.Rmsg(c, true, "", bs)
}

func getBlogs(page, pageSize int) ([]models.BlogApi, error) {
	b := models.NewBlog()
	bs, err := b.GetBlogs(page, pageSize)
	return bs, err
}

func getAuthoeBlogs(email string, page, pageSize int) ([]models.BlogApi, error) {
	b := models.NewBlog()
	b.Email = email
	bs, err := b.AuthoeToBlogs(page, pageSize)
	return bs, err
}

//UpGlog 更新文章
func UpGlog() {

}

//BlogSizeAPI 得到文章数量
func BlogSizeAPI(c *gin.Context) {
	i, err := models.Count()
	if err != nil {
		log.Errorf("%+v", err)
		common.Rmsg(c, false, "")
		return
	}
	common.Rmsg(c, true, "", i)
}

//GetTopAPI 得到to排名
func GetTopAPI(c *gin.Context) {
	bs, err := models.GetTop()
	if err != nil {
		log.Errorf("%+v", err)
		common.Rmsg(c, false, "")
		return
	}
	common.Rmsg(c, true, "", bs)
}
