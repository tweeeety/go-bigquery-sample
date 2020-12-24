package main

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"time"

	"cloud.google.com/go/bigquery"
)

type Env struct {
	PROJECT string
	DATASET string
	TABLE   string
}

var env = Env{}
var rs1Letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func init() {
	rand.Seed(time.Now().UnixNano())
	env.PROJECT = os.Getenv("PROJECT")
	env.DATASET = os.Getenv("DATASET")
	env.TABLE = os.Getenv("TABLE")
}

type Item struct {
	Id        string    `bigquery:"id"`
	Data      string    `bigquery:"data"`
	Timestamp time.Time `bigquery:"timestamp"`
}

func RandString(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = rs1Letters[rand.Intn(len(rs1Letters))]
	}
	return string(b)
}

func Put(randStr string, putData string) error {

	ctx := context.Background()
	client, err := bigquery.NewClient(ctx, env.PROJECT)
	if err != nil {
		return err
	}
	defer client.Close()

	u := client.Dataset(env.DATASET).Table(env.TABLE).Uploader()
	now := time.Now()
	items := []*Item{
		{Id: randStr, Data: putData, Timestamp: now},
	}

	err = u.Put(ctx, items)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	var randStr = RandString(16)
	var putData = "{huga:123, fuga:234, piyo{test: 34567}}"

	fmt.Println(randStr)
	fmt.Printf("env: %+v\n", env)

	err := Put(randStr, putData)
	fmt.Printf("err: %+v", err)
}
