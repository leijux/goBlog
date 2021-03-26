package models

import (
	"time"

	"goBlog/database/orm"

	json "github.com/json-iterator/go"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

var (
	ErrIDIsEmpty = errors.New("Id is empty")
)

//Blog 博客
type Blog struct {
	gorm.Model
	Title    string `gorm:"not null"            `
	Content  string `gorm:"type:text;not null"  `
	Email    string `gorm:"size:30;not null"    `
	Like     uint
	Likes    []Likes
	Comments []Comment
}

//blog id 和创建时间
type BlogApi struct {
	ID        uint      `json:"id"        form:"id"                               `
	Title     string    `json:"title"     form:"title"     binding:"required"      `
	Content   string    `json:"content"   form:"content"   binding:"required"      `
	Email     string    `json:"author"    form:"author"    binding:"required,email"`
	Like      uint      `json:"like"      form:"like"                             `
	CreatedAt time.Time `json:"createdAt" form:"createdAt"                        `
	UpdatedAt time.Time `json:"updatedAt" form:"updatedAt" `
}

func NewBlog() BlogApi {
	return BlogApi{}
}

func (blog Blog) ToBlogApi() BlogApi {
	return BlogApi{
		ID:        blog.ID,
		Title:     blog.Title,
		Content:   blog.Content,
		Email:     blog.Email,
		Like:      blog.Like,
		CreatedAt: blog.CreatedAt,
		UpdatedAt: blog.UpdatedAt,
	}
}

func (blog BlogApi) ToBlog() *Blog {
	return &Blog{
		Model: gorm.Model{
			ID: blog.ID,
		},
		Title:   blog.Title,
		Content: blog.Content,
		Email:   blog.Email,
		Like:    blog.Like,
	}
}

//AddBlog 添加博客
func (blog BlogApi) CreateBlog() (bool, error) {
	b := blog.ToBlog()
	return createBlog(b)
}

func createBlog(blog *Blog) (bool, error) {
	if result := orm.Db.Create(blog); result.Error != nil {
		return false, errors.Wrap(result.Error, "create blog err")
	}
	return true, nil
}

//DelBlog 删除博客
func (blog *BlogApi) DelBlog() error {
	if blog.ID != 0 {
		b := blog.ToBlog()
		return delBlog(b)
	}
	return ErrIDIsEmpty
}

func delBlog(b *Blog) error {
	return orm.Db.Delete(&b).Error
}

//UpBlog 更新博客
func (blog BlogApi) UpBlog() error {
	b := blog.ToBlog()
	if b.ID != 0 {
		err := orm.Db.Model(&b).Updates(Blog{
			Title:   b.Title,
			Content: b.Content,
		}).Error
		return err
	}
	return ErrIDIsEmpty
}

//GetBlogs 全部分页
func (blog BlogApi) GetBlogs(page int, pageSize int) ([]BlogApi, error) {
	var blogs []Blog
	if err := orm.Db.Scopes(Paginate(page, pageSize)).Find(&blogs).Error; err != nil {
		return nil, err
	}
	if blogs != nil {
		l := len(blogs)
		blogApis := make([]BlogApi, l)
		for i := 0; i < l; i++ {
			blogApis[i] = blogs[i].ToBlogApi()
		}
		return blogApis, nil
	}
	return nil, errors.New("blogs is empty")
}

//分页器
func Paginate(page int, pageSize int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		offset := (page - 1) * pageSize
		return db.Order("`created_at` desc").Offset(offset).Limit(pageSize)
	}
}

//AuthoeToBlogs 作者分页
func (blog BlogApi) AuthoeToBlogs(page int, pageSize int) ([]BlogApi, error) {
	return authoeToBlogs(blog.Email, page, pageSize)
}

func authoeToBlogs(email string, page int, pageSize int) ([]BlogApi, error) {
	var blogs []Blog
	if err := orm.Db.Scopes(Paginate(page, pageSize)).
		Where("email=?", email).
		Find(&blogs).Error; err != nil {
		return nil, err
	}
	if blogs != nil {
		l := len(blogs)
		blogApis := make([]BlogApi, l)
		for i := 0; i < l; i++ {
			blogApis[i] = blogs[i].ToBlogApi()
		}
		return blogApis, nil
	}
	return nil, errors.New("blogs is empty")
}

//Count 统计blog的数量
func Count() (count int64, err error) {
	err = orm.Db.Model(&Blog{}).Count(&count).Error
	return
}

//GetTop 得到前10个blog
func GetTop() ([]BlogApi, error) {
	return getTop()
}

func getTop() ([]BlogApi, error) {
	blogs := make([]Blog, 0, 10)

	err := orm.Db.Select("id", "title", "content", "email", "like").
		Order("`like` desc").
		Limit(10).
		Find(&blogs).Error

	if err != nil {
		return nil, err
	}
	l := len(blogs)
	blogApis := make([]BlogApi, l)
	for i := 0; i < l; i++ {
		blogApis[i] = blogs[i].ToBlogApi()
	}
	return blogApis, err
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
