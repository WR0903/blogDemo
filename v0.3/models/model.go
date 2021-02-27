package models

import (
	"fmt"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/ziutek/mymysql/godrv"
)

type Blog struct {
	Id      int `pk`
	Title   string
	Tags    string
	Short   string
	Content string
	Author  string
	Created string
}

//用户的数据结果
type User struct {
	Id         int
	Username   string
	Password   string
	Status     int // 0 正常状态， 1删除
	Createtime int64
}

func init() {
	orm.RegisterDataBase("default", "mysql", "root:@tcp(127.0.0.1:3306)/myblog?charset=utf8")
	// 需要在init中注册定义的model
	orm.RegisterModel(new(Blog), new(User))

	orm.RegisterDriver("mysql", orm.DRMySQL)

	// create table
	orm.RunSyncdb("default", true, true)
}

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
	return blog
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

func InserUser(user User) (int64, error) {
	userTmp := User{Id: user.Id}
	o := orm.NewOrm()
	err := o.Read(&userTmp)
	if err != nil {
		fmt.Println("read error:", err)
		_, err = o.Insert(&user)
		if err != nil {
			fmt.Println("insert error:", err)
		}
	} else {
		_, err := o.Update(&user)
		if err != nil {
			fmt.Println("update error:", err)
		}
	}

	return 1, err
}

func QueryUserWithParam(username, password string) error {
	userTmp := User{Username: username, Password: password}
	o := orm.NewOrm()
	err := o.Read(&userTmp)
	if err != nil {
		fmt.Println("QueryUserWithParam error:", err)
	}
	return err
}

func QueryUserWithUsername(username string) error {
	userTmp := User{Username: username}
	o := orm.NewOrm()
	err := o.Read(&userTmp)
	if err != nil {
		fmt.Println("QueryUserWithParam error:", err)
	}
	return err
}
