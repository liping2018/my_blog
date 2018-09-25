//file:redis.go
//auth:liping
//date:2018-09-24
//desc:redis key-value 服务器

package utils

import (
	"fmt"

	"github.com/garyburd/redigo/redis"
)

var con redis.Conn

func initRedis() {
	var err error
	host := GetString("redis::host")
	port := GetString("redis::port")
	pwd := GetString("redis::password")
	con, err = redis.Dial("tcp", host+":"+port)
	if err != nil {
		fmt.Println("初始化redis失败", err)
		return
	}
	_, err = con.Do("AUTH", pwd)
	if err != nil {
		fmt.Println("验证redis失败")
		return
	}
}

//获取redis值
func SetRedisValue(key, value, outtime string) {
	initRedis()
	defer con.Close()
	_, err := con.Do("SET", key, value, "EX", outtime)
	if err != nil {
		fmt.Println("redis set failed:", err)
	}
}

//设置redis值
func GetRedisValue(key string) string {
	initRedis()
	defer con.Close()
	value, err := redis.String(con.Do("GET", key))
	if err != nil {
		fmt.Println("获取redis值失败:", err)
		return ""
	}
	return value
}

//移除redis值
func RemoveRedisValue(key string) {
	initRedis()
	defer con.Close()
	_, err := con.Do("DEL", key)
	if err != nil {
		fmt.Println("移除redis值失败:", err)
	}
}

//查询redis 值是否已经存在

func IsExistRedisValue(key string) bool {
	initRedis()
	defer con.Close()
	is_key_exit, err := redis.Bool(con.Do("EXISTS", key))
	if err != nil {
		fmt.Println("error:", err)
	}
	return is_key_exit
}
