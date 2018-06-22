package main

import (
	"context"
	"time"
	"log"
)



func someHandler() {
	ctx, cancel := context.WithCancel(context.Background())

	go doStuff(ctx)

	time.Sleep(10 * time.Second)

	cancel()
}

func timeoutHandler()  {
	ctx, cancel := context.WithTimeout(context.Background(),5 * time.Second)
		//context.WithDeadline(context.Background(),time.Now().Add(5*time.Second))

	//go doStuff(ctx)

	go doTimeoutStuff(ctx)
	time.Sleep(10 * time.Second)

	cancel()
}

func valueHandler()  {

	ctx :=context.WithValue(context.Background(),"id",1)
	ctx = context.WithValue(ctx,"name","lyh")
	go doValueStuff(ctx)
	time.Sleep(10*time.Second)
}

func doStuff(ctx context.Context) {
	for {
		time.Sleep(1 * time.Second)
		select {
		case <- ctx.Done():
			log.Printf("done\n")
			return
		default:
			log.Printf("work\n")
		}
	}
}

func doTimeoutStuff(ctx context.Context){
	for{
		time.Sleep(1 * time.Second)
		if deadline, ok := ctx.Deadline();ok{
			log.Printf("deadline set\n")
			if time.Now().After(deadline){
				log.Printf("%s\n",ctx.Err().Error())
			}
		}
		select {
		case <- ctx.Done():
			log.Printf("done\n")
			return
		default:
			log.Printf("work\n")
		}
	}
}

func doValueStuff(ctx context.Context)  {
	log.Printf("`%s` value:%v\n", "id",ctx.Value("id"))
	log.Printf("`%s` value:%v\n", "name",ctx.Value("name"))
}

func main() {
	//someHandler()
	//timeoutHandler()
	valueHandler()
	log.Printf("down\n")
}
