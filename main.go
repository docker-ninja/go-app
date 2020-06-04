package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"log"
	"net/http"
	"os"
	"strconv"
)


func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello\n")
}

func ping(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "pong\n")
}

func vote(w http.ResponseWriter, req *http.Request) {

	client, ctx := newConnection(req)
	currentVotes := getVotes(client, ctx)
	currentVotes++
	votes := increaseVote(currentVotes, client, ctx)

	fmt.Fprintf(w, fmt.Sprintf("total votes  %d ", votes))
}

func increaseVote(votes int, client *redis.Client, ctx context.Context) int {
	err := client.Set(ctx, "votes", votes, 0).Err()
	if err != nil {
		panic(err)
	}
	return votes
}

func getVotes(client *redis.Client, ctx context.Context) int {
	value, err := client.Get(ctx, "votes").Result()

	if err != nil {
		if err.Error() == "redis: nil" {
			increaseVote(0, client, ctx)
			return 0
		} else {
			panic(err)
		}
	}

	votes, err := strconv.Atoi(value)
	if err != nil {
		panic(err)
	}

	return votes
}

func newConnection(req *http.Request) (*redis.Client, context.Context) {
	redisHost := os.Getenv("REDIS_HOST")
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:6379", redisHost),
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	ctx := req.Context()
	_, err := client.Ping(ctx).Result()
	if err != nil {
		panic(err)
	}
	return client, ctx
}

func main() {
	addr := ":8080"
	http.HandleFunc("/", hello)
	http.HandleFunc("/ping", ping)
	http.HandleFunc("/vote", vote)
	log.Println("listen on", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
