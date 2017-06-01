package main

import (
	"net/http"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"

	"time"
)

type Tag struct {
	Id int `json:"id"`
	Tag string `json:"name"`
}

type Blog struct {
	Id    int `json:"id"`
	Title string `json:"title"`
	Date  time.Time `json:"-"`
	DateStr string `json:"date" gorm:"-"`
	Content string `json:"-" gorm:"type: text;"`
	ContentBrief string `json:"content" gorm:"-"`
	Tags []Tag `gorm:"many2many:blog_tags" json:"tags"`
}

type BlogTag struct {
	BlogId int
	TagId int
}

func (blog *Blog) fmt()  {
	blog.DateStr = blog.Date.Format("2006-01-02")
	content := []rune(blog.Content)
	if len(blog.Content) > 51 {
		blog.ContentBrief = string(content[:50])
	} else {
		blog.ContentBrief = blog.Content
	}
}

func main() {

	db, err := gorm.Open("sqlite3", "data.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	db.AutoMigrate(&Blog{}, &Tag{})

	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:8000", "http://localhost:8080"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType},
	}))
	e.Static("/", "public/dist")
	e.GET("/hello", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, This is my blog project!")
	})

	e.GET("/blog/list", func(c echo.Context) error {
		blogs := []Blog{}
		db.Limit(5).Find(&blogs)

		for i := 0; i < len(blogs); i++ {
			db.Model(&blogs[i]).Related(&(blogs[i].Tags), "Tags")
			blogs[i].fmt()
		}
		//blog.fmt()
		//fmt.Println(blog.Tags)
		return c.JSON(http.StatusOK, blogs)
	})
	e.Logger.Fatal(e.Start(":8000"))
}