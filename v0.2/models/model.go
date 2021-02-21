package models

import (
	"fmt"
	"time"

	"github.com/astaxie/beego/orm"
	"github.com/go-redis/redis"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/ziutek/mymysql/godrv"
)

type Blog struct {
	Id      int `pk`
	Title   string
	Content string
	Created time.Time
}

func init() {
	orm.RegisterDataBase("default", "mysql", "root:@tcp(127.0.0.1:3306)/myblog?charset=utf8")
	// 需要在init中注册定义的model
	orm.RegisterModel(new(Blog))

	orm.RegisterDriver("mysql", orm.DRMySQL)

	// create table
	orm.RunSyncdb("default", false, true)
}

// 声明一个全局的redisdb变量
var redisdb *redis.Client

func GetAll() (blogs []Blog) {
	o := orm.NewOrm()
	_, err := o.QueryTable("blog").All(&blogs)
	if err != nil {
		fmt.Println("find all:", err)
	}
	return
}

func GetBlog(id int) (blog Blog) {
	blog.Id = id
	o := orm.NewOrm()
	err := o.Read(&blog)
	if err != nil {
		fmt.Println("read error:", err)
	}
	return
}

func SaveBlog(blog Blog) (bg Blog) {
	blogTmp := Blog{Id: blog.Id}
	o := orm.NewOrm()
	err := o.Read(&blogTmp)
	if err != nil {
		fmt.Println("read error:", err)
		_, err := o.Insert(&blog)
		if err != nil {
			fmt.Println("insert error:", err)
		}
	} else {
		_, err := o.Update(&blog)
		if err != nil {
			fmt.Println("update error:", err)
		}
	}

	return bg
}

func DelBlog(blog Blog) {
	blogTmp := new(Blog)
	blogTmp.Id = blog.Id
	o := orm.NewOrm()
	err := o.Read(&blog)
	if err != nil {
		fmt.Println("read error:", err)
		return
	}
	_, err = o.Delete(&blog)
	if err != nil {
		fmt.Println("del failed:", err)
	}
	return
}
