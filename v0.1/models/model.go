package models

import (
	"crypto/md5"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"time"

	"github.com/astaxie/beego/logs"
	"github.com/go-redis/redis"
	_ "github.com/ziutek/mymysql/godrv"
)

type Blog struct {
	Id      string
	Title   string
	Content string
	Created time.Time
}

// 声明一个全局的redisdb变量
var redisdb *redis.Client

// 初始化连接
func initClient() (err error) {
	redisdb = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		PoolSize: 128,
		DB:       0,
	})

	_, err = redisdb.Ping().Result()
	if err != nil {
		return err
	}
	return nil
}

func init() {
	err := initClient()
	if err != nil {
		panic(err)
	}
	logs.Info("王冉")
}

func GetAll() (blogs []Blog) {
	var cursor uint64
	keys, cursor, err := redisdb.Scan(cursor, "*", 100).Result()
	if err != nil {
		fmt.Println("scan keys failed err:", err)
		return
	}
	for _, key := range keys {
		value, err2 := redisdb.Get(key).Result()
		if err2 != nil {
			fmt.Println("get keys failed err:", err2)
			continue
		}
		if json.Valid([]byte(value)) {
			m := Blog{}

			err := json.Unmarshal([]byte(value), &m)
			if err != nil {

				fmt.Println("Umarshal failed:", err)
				continue
			}
			blogs = append(blogs, m)

		}
	}
	return
}

func GetBlog(id string) (blog Blog) {
	key := fmt.Sprintf("blog_%s", id)
	value, err := redisdb.Get(key).Result()
	if err != nil {
		fmt.Println("get failed:", err)
	}
	err2 := json.Unmarshal([]byte(value), &blog)
	if err2 != nil {
		fmt.Println("Umarshal failed:", err2)
		return
	}
	return
}

func SaveBlog(blog Blog) (bg Blog) {
	key := fmt.Sprintf("blog_%s", blog.Id)
	value, err := json.Marshal(blog)
	if err != nil {

		fmt.Println("Umarshal failed: %s", err)
		return
	}
	value2 := string(value)
	_, err2 := redisdb.Set(key, value2, -1).Result()
	if err2 != nil {
		fmt.Println("set failed:", err2)
	}
	return bg
}

func DelBlog(blog Blog) {
	key := fmt.Sprintf("blog_%s", blog.Id)
	_, err := redisdb.Del(key).Result()
	if err != nil {
		fmt.Println("del failed:", err)
	}
	return
}

//生成32位md5字串
func GetMd5String(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

//生成Guid字串
func UniqueId() string {
	b := make([]byte, 48)

	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return ""
	}
	return GetMd5String(base64.URLEncoding.EncodeToString(b))
}

func Counter() (num int) {
	var cursor uint64
	keys, cursor, err := redisdb.Scan(cursor, "*", 100).Result()
	if err != nil {
		fmt.Println("scan keys failed err:", err)
		return
	}
	num = len(keys)
	return
}
