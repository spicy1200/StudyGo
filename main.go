// package main

// import "fmt"

// func main() {
// 	/**
// 		main 函数是作为goroutine 执行
// 		操作系统 线程 和 goroutine 的关系？
// 			线程被操作系统调度
// 			goroutine 实际上是被Go 运行时runtime来调度的，最终这些goroutine 它都可以映射到
// 			到某一个线程上，执行单元对操作系统来说它不认识goroutine,它只认识操作系统线程
// 			go runtime 实际上就是负责把这些goroutine 调度到某一个Go runtime的一个逻辑处理器（P）
// 			P 像一个队列，把这些任务挂到队列上，然后相当于最终是在各个线程去队列里面捞到一个任务去执行
// 			通过go runtime 的特性使的我们可以调度数以万计的goroutine,以惊人的效率和能力执行并发运行

// 			runtime.GomaxProcs  设置逻辑处理器的个数

// 			go routine
// 			什么时间会结束
// 			有没有办法结束它

// 			log.Fatal 调用os.exit 不会调用 defers
// 				使用 log.Fatal 只能在 go init 或者 main 函数中

// 			errgroup
// 	**/
// 	fmt.Println("ceshi")
// }

package main

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"
	"strings"

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
func Dao(search string, db *sql.DB) error {
	var (
		title string
		id    int
	)
	err := db.QueryRow("SELECT title,id FROM blog_article WHERE title = ?", search).Scan(&title, &id)

	switch {
	case err == sql.ErrNoRows:
		// errors.Wrap(err, " has no article")
		return errors.Wrapf(err, "%d not found", errNoRows)
	case err != nil:
		// errors.Wrap(err, fmt.Sprintf("query row error."))
		return errors.Wrapf(err, "%d error", dbErr)
	default:
		fmt.Println("blog_article id %d", id)
	}
	return nil
}

var (
	errNoRows int = 404
	dbErr     int = 500
)

func main() {

	/**
		1. 我们在数据库操作的时候，比如 dao 层中当遇到一个 sql.ErrNoRows 的时候，是否应该 Wrap 这个 error，抛给上层。为什么，应该怎么做请写出代码？
		答: 可以
		理由：
			分析: 选择使用 wrap error 是只有applications 可以选择使用应用的策略。具有最高可用性的包只能返回根错误值。
		sql.ErrNoRows 返回的错误信息是属于根错误信息值，因为sql.ErrNoRows 错误信息的产生是在Dao层通过Wrap error 抛出错误 让Applications 应用层进行处理

	**/
	db, err := connectMysql()
	if err != nil {
		log.Fatal("connect mysql error message", err)
	}
	// result := Dao("文章标题", db)
	errRes := Dao("c", db)
	if strings.HasPrefix(errRes.Error(), strconv.Itoa(errNoRows)) {
		fmt.Println("not found")
	}
	if strings.HasPrefix(errRes.Error(), strconv.Itoa(dbErr)) {
		fmt.Println("查询数据出现问题  %v", errRes.Error())
	}
}
