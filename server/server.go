package server

import (
	"flag"
	"log"
	"os"
	"redistr/common"
	"reflect"
	"strings"
	"sync"

	redis "github.com/go-redis/redis/v7"
	"github.com/tidwall/redcon"
)

var serverport = ":6379"
var mu sync.RWMutex
var primaryclient *redis.Client
var anabranchclient *redis.Client

// Server start server
func Server() {
	flag.Parse()
	if primaryclient == nil {
		log.Println("primaryclient is nill")
		os.Exit(1)
	}
	// var items = make(map[string][]byte)
	go log.Printf("started server at %s", serverport)
	defer primaryclient.Close()

	err := redcon.ListenAndServe(serverport, handle, accept, closed)

	if err != nil {
		log.Fatal(err)
	}

}

//SetPrimaryRedisClient 设置主rediscliet
func SetPrimaryRedisClient(opt redis.Options) {
	primaryclient = GetRedisClient(opt)
}

//SetAnabranchRedisClient 设置分支redisclient
func SetAnabranchRedisClient(opt redis.Options) {
	anabranchclient = GetRedisClient(opt)
}

//SetServerPort 设置server 端口
func SetServerPort(port string) {
	serverport = ":" + port
}

func handle(conn redcon.Conn, cmd redcon.Command) {
	var cmdstr []interface{}
	for i := 0; i < len(cmd.Args); i++ {
		cmdstr = append(cmdstr, string(cmd.Args[i]))
	}
	bak, err := primaryclient.Do(cmdstr...).Result()

	if anabranchclient != nil && common.UpdateCmd.Contains(strings.ToUpper(string(cmd.Args[0]))) {
		go anabranchclient.Do(cmdstr...)
	}

	if bak == nil {
		if err != nil {
			log.Println(err)
			conn.WriteError(err.Error())
		} else {
			conn.WriteString("nil")
		}
		return
	}

	if err != nil {
		log.Println(err)
		conn.WriteError(err.Error())
	} else {
		switch reflect.TypeOf(bak).String() {
		default:
			log.Println(reflect.TypeOf(bak))
			conn.WriteString("cannot find return type")
		case "[]interface {}":
			bakin := bak.([]interface{})
			mu.Lock()
			conn.WriteArray(len(bakin))
			for i := 0; i < len(bakin); i++ {

				log.Println(reflect.TypeOf(bakin[i]).String())
				if reflect.TypeOf(bakin[i]).String() == "int64" {
					conn.WriteInt64(bakin[i].(int64))
					continue
				}

				if reflect.TypeOf(bakin[i]).String() == "string" {
					conn.WriteBulkString(bakin[i].(string))
					continue
				}
				if reflect.TypeOf(bakin[i]).String() == "[]interface {}" {
					log.Println(bakin[i])
					conn.WriteNull()

				}
			}
			mu.Unlock()
		case "string":
			conn.WriteString(bak.(string))
		case "int64":
			conn.WriteInt64(bak.(int64))
		case "int":
			conn.WriteInt(bak.(int))

		}
	}
}

func accept(conn redcon.Conn) bool {
	// use this function to accept or deny the connection.
	// log.Printf("accept: %s", conn.RemoteAddr())
	return true
}

func closed(conn redcon.Conn, err error) {
	// this is called when the connection has been closed
	// log.Printf("closed: %s, err: %v", conn.RemoteAddr(), err)
}

//GetRedisClient 获取redis client
func GetRedisClient(opt redis.Options) *redis.Client {
	client := redis.NewClient(&opt)
	return client
}

//Writetoanabranch 分支写入
func Writetoanabranch(client *redis.Client, cmdstr ...interface{}) {
	_, err := client.Do(cmdstr).Result()
	if err != nil {
		log.Println(err)
	}
}
