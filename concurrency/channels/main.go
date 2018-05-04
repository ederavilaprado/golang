package main

import (
	"flag"
	"fmt"
	"time"
)

// var pick string

var pick = flag.String("c", "concat", "choose one: concat, time, pingpong")

// flag.StringVar()

func main() {
	flag.Parse()

	switch *pick {
	case "concat":
		concatExample()
		break
	case "time":
		timeTickingExample()
		break
	case "pingpong":
		timeTickingExample()
		break
	default:
		panic("Not a valid option")
	}
}

// Concat example, simple one
func concatExample() {
	c := make(chan string)

	go concat("check_a", 1*time.Second, c)
	go concat("check_b", 5*time.Second, c)

	fmt.Printf("=> %+v\n", "code here")

	ra := <-c
	fmt.Printf("=> %+v\n", "get the first")
	fmt.Printf("=> %s\n", ra)

	rb := <-c
	fmt.Printf("=> %+v\n", "get the second")
	fmt.Printf("=> %s\n", rb)
}

func concat(inp string, wait time.Duration, c chan string) {
	fmt.Printf("=> inside concat with inp: %s\n", inp)
	time.Sleep(wait)
	r := "_concat_" + inp
	c <- r
}

// Time ticking example...
func timeTickingExample() {
	c := time.Tick(500 * time.Millisecond)
	for now := range c {
		fmt.Printf("-> %s\n", now)
	}
}

// Ping Pong example...

type ball struct{ hits int }

func pingPongExample() {
	table := make(chan *ball)
	go player("ping", table)
	go player("pong", table)

	table <- new(ball) // game on; toss the ball
	time.Sleep(2 * time.Second)
	<-table // game over; grab the ball
}

func player(name string, table chan *ball) {
	for {
		ball := <-table
		ball.hits++
		fmt.Println(name, ball.hits)
		time.Sleep(100 * time.Millisecond)
		table <- ball
	}
}
