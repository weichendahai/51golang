package models

import (
	//"database/sql"
	_ "errors"
	_ "fmt"
	//_ "github.com/mattn/go-sqlite3"
	"time"
)

type article struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

// Return a list of all the articles
func GetAllArticles() []article {
	//查询数据

	db := InitDB()
	defer db.Close()
	//rows, err := db.Query("SELECT id, title, content, created_at FROM articles")
	rows, _ := db.Query("SELECT id, title, content, created_at FROM articles")
	var articleList = []article{}
	for rows.Next() {
		var id int
		var title string
		var content string
		var created time.Time
		//err := rows.Scan(&id, &title, &content, &created)
		_ = rows.Scan(&id, &title, &content, &created)
		a := article{ID: id, Title: title, Content: content}
		// Add the article to the list of articles
		articleList = append(articleList, a)
	}
	db.Close()
	return articleList
}

// Fetch an article based on the ID supplied
func GetArticleByID(id int) (*article, error) {

	//查询数据
	var title string
	var content string
	var created time.Time

	db := InitDB()
	_ = db.QueryRow("SELECT id, title, content, created_at FROM articles where id=?", id).Scan(&id, &title, &content, &created)
	a := article{ID: id, Title: title, Content: content}
	db.Close()
	return &a, nil
}

// Create a new article with the title and content provided
func CreateNewArticle(title, content string) (*article, error) {
	//插入数据

	db := InitDB()
	//stmt, err := db.Prepare("INSERT INTO articles(title, content, created_at) values(?,?,?)")
	stmt, _ := db.Prepare("INSERT INTO articles(title, content, created_at) values(?,?,?)")
	res, _ := stmt.Exec(title, content, time.Now())
	id, _ := res.LastInsertId()
	db.Close()

	a := article{ID: int(id), Title: title, Content: content}
	return &a, nil
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
