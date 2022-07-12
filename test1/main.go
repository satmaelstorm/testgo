package main

import (
	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
	"log"
	"math/rand"
	"time"
)

func putToQueue(body []byte) {
	//тут по сети кладем какие-то байты в очередь
	time.Sleep(time.Duration(rand.Intn(100)) * time.Second)
	log.Println(body)
}

func main() {
	r := router.New()

	r.PUT("/queue", func(ctx *fasthttp.RequestCtx) {
		go func(body []byte) {
			putToQueue(body)
		}(ctx.Request.Body())
		ctx.Response.SetStatusCode(fasthttp.StatusAccepted)
	})

	err := fasthttp.ListenAndServe(":8088", r.Handler)
	log.Fatal(err)
}
