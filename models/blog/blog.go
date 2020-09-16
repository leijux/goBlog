package blog

import (
	"encoding/json"
	"goBlog/database"
	"goBlog/database/orm"
	"goBlog/log"
	"goBlog/models"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

//Blog 博客
type Blog struct {
	gorm.Model
	Title   string `gorm:"not null"              `
	Content string `gorm:"type:text;not null"    `
	Email   string `gorm:"not null"       `
	Likes   int    `                             `
}

//blog id 和创建时间
type BlogApi struct {
	Title   string `json:"title"     form:"title"    binding:"required"`
	Content string `json:"content"   form:"content"  binding:"required"`
	Email   string `json:"author"    form:"author"   binding:"required"`
	Likes   int    `json:"likes"     form:"likes"`
}

var _ models.IModels = &Blog{}

func NewBolg() BlogApi {
	return BlogApi{}
}
func (blog Blog) ToBlogApi() BlogApi {
	return BlogApi{
		Title:   blog.Title,
		Content: blog.Content,
		Email:   blog.Email,
		Likes:   blog.Likes,
	}
}

func (blog BlogApi) ToBlog() *Blog {
	return &Blog{
		Title:   blog.Title,
		Content: blog.Content,
		Email:   blog.Email,
		Likes:   blog.Likes,
	}
}

//AddBlog 添加博客
func (blog *BlogApi) CreateBlog() (bool, error) {
	// res, err := database.Db.NamedExec("insert into blog(title,content,author,likes,created,updated)  values(:title,:content,:author,:likes,:created,:updated)", blog)
	// if err != nil {
	// 	//log.Logger.Errorln(err)
	// 	return
	// }
	// id, err = res.LastInsertId()
	// if err != nil {
	// 	//log.Logger.Errorln(err)
	// 	return
	// }
	b := blog.ToBlog()
	return createBlog(b)
}

func createBlog(blog *Blog) (bool, error) {
	log.Infoln(blog)
	if result := orm.Create(blog); result.Error != nil {
		return false, errors.Wrap(result.Error, "create blog err")
	}
	return true, nil
}

//DelBlog 删除博客
func (blog *Blog) DelBlog() {

}

//UpBlog 更新博客
func (blog *Blog) UpBlog() {

}

//GetBlogs 全部分页
func (blog *Blog) GetBlogs(page int, perpPage int) (blogs []Blog, err error) {
	n, m := (page-1)*perpPage, perpPage //page 当前页数  perpage 每页第几个
	err = database.Db.Select(&blogs, "select * from blog order by created desc limit ?,?  ", n, m)
	if err != nil {
		//log.Logger.Errorln(err)
		return
	}
	return
}

//AuthoeToBlogs 作者分页
func (blog *Blog) AuthoeToBlogs(page int, perpPage int) (blogs []Blog, err error) {
	n, m := (page-1)*perpPage, perpPage
	err = database.Db.Select(&blogs, "select * from blog where author=? order by created desc limit ?,?  ", blog.Email, n, m)
	if err != nil {
		//log.Logger.Errorln(err)
		return
	}
	return
}

//Count 统计blog的数量
func (blog *Blog) Count() (count int, err error) {
	rows, err := database.Db.Query("select count(*) from blog")
	if err != nil {
		//log.Logger.Errorln(err)
		return
	}
	if rows.Next() {
		err = rows.Scan(&count)
	}
	return
}

//GetTop 得到blog的数量
func (blog *Blog) GetTop() (blogs []Blog, err error) {
	err = database.Db.Select(&blogs, "select * from blog order by likes desc limit 10")
	if err != nil {
		//log.Logger.Errorln(err)
		return
	}
	return
}

//ToJSON ...
func (blog *Blog) ToJSON() string {
	j, err := json.Marshal(blog)
	if err != nil {
		//log.Logger.Debugln(err)
		return ""
	}
	return string(j)
}

//FromJSON ...
func (blog *Blog) FromJSON(data string) {
	err := json.Unmarshal([]byte(data), blog)
	if err != nil {
		//log.Logger.Debugln(err)
		return
	}
	return
}
