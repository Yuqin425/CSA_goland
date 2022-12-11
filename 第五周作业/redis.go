package main

import (
	"fmt"
	"github.com/go-redis/redis"
	"log"
)

var redisDB *redis.Client

func initClient() (err error) {
	redisDB = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", //地址:端口
		Password: "",               //密码，没有就写空字符串""
		DB:       0,                //使用哪个数据库，0代表默认数据库
	})
	_, err = redisDB.Ping().Result()
	if err != nil {
		return err
	}
	return nil
}

func Zset(key string, val []redis.Z) error {
	_, err := redisDB.ZAdd(key, val...).Result()
	if err != nil {
		fmt.Printf("zadd failed, err:%v\n", err)
		return err
	}
	return nil
}

func main() {
	err := initClient()
	if err != nil {
		log.Fatalln("init redis is err:", err)
		return
	}
	key := "language_rank"
	languages := []redis.Z{
		redis.Z{Score: 80, Member: "go"},
		redis.Z{Score: 88, Member: "java"},
		redis.Z{Score: 70, Member: "c++"},
	}
	err = Zset(key, languages)
	if err != nil {
		log.Fatalln("zset redis is err:", err)
		return
	}
	op := redis.ZRangeBy{
		Min: "60",
		Max: "85",
	}
	ret, err := redisDB.ZRangeByScoreWithScores(key, op).Result()
	if err != nil {
		fmt.Printf("zrangebyscore failed, err:%v\n", err)
		return
	}
	for _, z := range ret {
		fmt.Println(z.Member, z.Score)
	}
}
