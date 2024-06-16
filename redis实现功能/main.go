package main

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"log"
)

var ctx = context.Background()

func getClient() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	})
	//ping := rdb.Ping(ctx)
	//fmt.Println(ping.String()) //PONG
	return rdb
}

func setMathScore(studentName string, score int) {
	client := getClient()
	do := client.ZAdd(context.Background(), "math_score", redis.Z{Member: studentName, Score: float64(score)})
	if err := do.Err(); err != nil {
		log.Fatal(err)
	}
}

func getTheChartsDesc(start, end int64) {
	client := getClient()
	do := client.ZRevRangeWithScores(context.Background(), "math_score", start, end)
	result, err := do.Result()
	if err != nil {
		log.Fatal(err)
	}

	for _, z := range result {
		fmt.Printf("姓名:%s,分数:%d\n", z.Member, int(z.Score))
		//姓名:张三,分数:90
		//姓名:李四,分数:80
		//姓名:王五,分数:70
	}
}

func getTheChartsAsc(start, end int64) {
	client := getClient()
	do := client.ZRangeWithScores(context.Background(), "math_score", start, end)
	result, err := do.Result()
	if err != nil {
		log.Fatal(err)
	}

	for _, z := range result {
		fmt.Printf("姓名:%s,分数:%d\n", z.Member, int(z.Score))
		//姓名:江北,分数:40
		//姓名:江南,分数:50
		//姓名:初见,分数:60
	}
}

func total() {
	client := getClient()
	card := client.ZCard(context.Background(), "math_score")
	total, err := card.Uint64()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(total) //6
}

func remove(studentName string) {
	client := getClient()
	rem := client.ZRem(context.Background(), "math_score", studentName)
	if err := rem.Err(); err != nil {
		log.Fatal(err)
	}
}

func example() {
	setMathScore("张三", 90)
	setMathScore("李四", 80)
	setMathScore("王五", 70)
	setMathScore("初见", 60)
	setMathScore("江南", 50)
	setMathScore("江北", 40)
	total()
	getTheChartsDesc(0, 2)
	getTheChartsAsc(0, 2)
	remove("江北")
	total()
}

func main() {
	example()
}
