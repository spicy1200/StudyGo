package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/pkg/errors"

	_ "github.com/go-sql-driver/mysql"
)

type blog_article struct {
	title string
	count string
}
type mysqlConfig struct {
	username  string
	password  string
	localhost string
	port      int
	charset   string
}

func connectMysql() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:123456@tcp(localhost:3306)/blog?charset=utf8")
	if err != nil {
		errors.New(fmt.Sprintf("sql.Db connect %v", err))
	}
	return db, err
}
func getArticle(search string, db *sql.DB) string {
	var (
		title string
		id    int
	)
	err := db.QueryRow("SELECT title,id FROM blog_article WHERE title = ?", search).Scan(&title, &id)

	var result string
	switch {
	case err == sql.ErrNoRows:
		// errors.Wrap(err, " has no article")
		result = fmt.Sprintf("%s has no article.", title)
	case err != nil:
		// errors.Wrap(err, fmt.Sprintf("query row error."))
		log.Fatal(fmt.Sprintf("query row error %v", err))
	default:
		result = fmt.Sprintf("article is %s count %d.", title, id)
	}
	return result
}

func main() {

	/**
		1. 我们在数据库操作的时候，比如 dao 层中当遇到一个 sql.ErrNoRows 的时候，是否应该 Wrap 这个 error，抛给上层。为什么，应该怎么做请写出代码？
		答: 不能
		理由：
			分析: 选择使用 wrap error 是只有applications 可以选择使用应用的策略。具有最高可用性的包只能返回根错误值。
		sql.ErrNoRows 返回的错误信息是属于更错误信息值，因为sql.ErrNoRows 错误信息的产生是在Dao层也不输与Applications 应用层，
		所以在Dao层出现sql.ErrNoRows 应该选择直接返回错误信息

	**/
	db, err := connectMysql()
	if err != nil {
		log.Fatal("connect mysql error message", err)
	}
	// result := getArticle("文章标题", db)
	result := getArticle("c", db)
	fmt.Println("get article ", result)
}
