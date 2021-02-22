## 自己搭建的博客
算是对beego的学习

### v0.1
第一次接触beego，用的redis数据库，写的不是很通用，不过大体框架搭起来了。
* router,这里一般是定义url得匹配规则，然后将对应得pose或者get请求连接到一个结构体，该结构体包含了get或者pose得处理。 
``` 
beego.Router("/edit/:id", &controllers.EditController{})   
```
* controllers, 这里一般是定义一个结构体，里面包含beego.Controller,然后该结构体对pose和get等方法进行重写  
* views，这里是视图相关，可以用html或者tpl文件

### v0.2
在v0.1的基础上，将数据库改成mysql，并用orm进行数据库的操作，这里包括表的模型，自动创建表单，对数据库的基本的增删改查等操作。  
* orm的初始化
```   
orm.RegisterDataBase("default", "mysql", "root:@tcp(127.0.0.1:3306)/myblog?charset=utf8")
// 需要在init中注册定义的model
orm.RegisterModel(new(Blog))
orm.RegisterDriver("mysql", orm.DRMySQL)
// create table
orm.RunSyncdb("default", false, true)
```
* 增加数据  
```
blogTmp := Blog{Id: blog.Id}
o := orm.NewOrm()
_, err := o.Insert(&blog)
if err != nil {
	fmt.Println("insert error:", err)
}
```
* 删除数据   
```
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
```
* 修改数据
```
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
```
* 查找数据
```
1、查找单个：
blog.Id = id
o := orm.NewOrm()
err := o.Read(&blog)
if err != nil {
	fmt.Println("read error:", err)
}
2、查找多个：
o := orm.NewOrm()
_, err := o.QueryTable("blog").All(&blogs)
if err != nil {
	fmt.Println("find all:", err)
}
```
