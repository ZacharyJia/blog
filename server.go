package main

import (
	"net/http"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"

	"time"
	"strconv"
	"math"
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
	Content string `json:"content" gorm:"type: text;"`
	//ContentBrief string `json:"content" gorm:"-"`
	Tags []Tag `gorm:"many2many:blog_tags" json:"tags"`
}

type BlogTag struct {
	BlogId int
	TagId int
}

func (blog *Blog) fmt(brief bool)  {
	blog.DateStr = blog.Date.Format("2006-01-02")
	content := []rune(blog.Content)
	if brief && len(blog.Content) > 51 {
		blog.Content = string(content[:50])
	}
}

var BLOG_PER_PAGE = 1

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

	e.GET("/blog/list/:page", func(c echo.Context) error {
		blogs := []Blog{}
		page, err:= strconv.Atoi(c.Param("page"))
		if err != nil {
			page = 1
		}
		var count int
		db.Table("blogs").Count(&count).Order("Date").Offset((page - 1) * BLOG_PER_PAGE).Limit(BLOG_PER_PAGE).Find(&blogs)

		for i := 0; i < len(blogs); i++ {
			db.Model(&blogs[i]).Related(&(blogs[i].Tags), "Tags")
			blogs[i].fmt(true)
		}
		res := make(map[string]interface{})
		res["data"] = blogs
		res["cur_page"] = page
		res["total_page"] = math.Ceil(float64(count) / float64(BLOG_PER_PAGE))
		return c.JSON(http.StatusOK, res)
	})

	e.GET("/blog/:id", func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.String(http.StatusNotFound, "this blog is not found")
		}
		blog := Blog{}
		db.Where("id = ?", id).Find(&blog)
		db.Model(&blog).Related(&blog.Tags, "Tags")
		blog.fmt(false)
		return c.JSON(http.StatusOK, blog)
	})
	e.Logger.Fatal(e.Start(":8000"))
}