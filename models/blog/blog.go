package blog

import (
	"task-system/database"
	"task-system/log"
	"time"
)

type Blog struct {
	ID      int       `db:"id"      json:"id"`
	Title   string    `db:"title"   json:"title"     form:"title"    binding:"required"`
	Content string    `db:"content" json:"content"   form:"content"  binding:"required"`
	Author  string    `db:"author"  json:"author"    form:"author"   binding:"required"`
	Likes   int       `db:"likes"   json:"likes"     form:"likes"`
	Created time.Time `db:"created" json:"created"`
	Updated time.Time `db:"updated" json:"updated"`
}

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

func (blog *Blog) DelBlog() {

}

func (blog *Blog) UpBlog() {

}

func (blog *Blog) GetBlogs(page int, perpPage int) (blogs []Blog, err error) {
	n, m := (page-1)*perpPage, perpPage
	err = database.Db.Select(&blogs, "select * from blog ORDER BY created DESC limit ?,?  ", n, m)
	if err != nil {
		log.Logger.Errorln(err)
		return
	}
	return
}
func (blog *Blog) AuthoeToBlogs(page int, perpPage int) (blogs []Blog, err error) {
	n, m := (page-1)*perpPage, perpPage
	err = database.Db.Select(&blogs, "select * from blog where author=? ORDER BY created DESC limit ?,?  ", blog.Author, n, m)
	if err != nil {
		log.Logger.Errorln(err)
		return
	}
	return
}
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
