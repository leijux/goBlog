package blog

import (
	"encoding/json"
	"goBlog/database"
	"goBlog/log"
	"goBlog/models"

	"github.com/jinzhu/gorm"
)

//Blog 博客
type Blog struct {
	gorm.Model
	Title   string `gorm:"not null"              db:"title"  `
	Content string `gorm:"type:text;not null"    db:"content"`
	Email   string `gorm:"not null;unique"       db:"author" `
	Likes   int    `                             db:"likes"`       
}

//blog id 和创建时间
type BlogApi struct{
	Title   string `json:"title"     form:"title"    binding:"required"`
	Content string `json:"content"   form:"content"  binding:"required"`
	Email   string `json:"author"    form:"author"   binding:"required"`
	Likes   int    `json:"likes"     form:"likes"`
}

var _ models.IModels = &Blog{}

//AddBlog 添加博客
func (blog *Blog) AddBlog() (id int64, err error) {
	res, err := database.Db.NamedExec("insert into blog(title,content,author,likes,created,updated)  values(:title,:content,:author,:likes,:created,:updated)", blog)
	if err != nil {
		log.Logger.Errorln(err)
		return
	}
	id, err = res.LastInsertId()
	if err != nil {
		log.Logger.Errorln(err)
		return
	}
	return
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
		log.Logger.Errorln(err)
		return
	}
	return
}

//AuthoeToBlogs 作者分页
func (blog *Blog) AuthoeToBlogs(page int, perpPage int) (blogs []Blog, err error) {
	n, m := (page-1)*perpPage, perpPage
	err = database.Db.Select(&blogs, "select * from blog where author=? order by created desc limit ?,?  ", blog.Email, n, m)
	if err != nil {
		log.Logger.Errorln(err)
		return
	}
	return
}

//Count 统计blog的数量
func (blog *Blog) Count() (count int, err error) {
	rows, err := database.Db.Query("select count(*) from blog")
	if err != nil {
		log.Logger.Errorln(err)
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
		log.Logger.Errorln(err)
		return
	}
	return
}

//ToJSON ...
func (blog *Blog) ToJSON() string {
	j, err := json.Marshal(blog)
	if err != nil {
		log.Logger.Debugln(err)
		return ""
	}
	return string(j)
}

//FromJSON ...
func (blog *Blog) FromJSON(data string) {
	err := json.Unmarshal([]byte(data), blog)
	if err != nil {
		log.Logger.Debugln(err)
		return
	}
	return
}
