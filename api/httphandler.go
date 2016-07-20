package main

import (
	"fmt"
	"log"

	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
)

// Index is the index handler
func CreateOrder(ctx *fasthttp.RequestCtx, ps fasthttprouter.Params) {
	fmt.Fprint(ctx, "Welcome!\n")
}

func GetOrders(ctx *fasthttp.RequestCtx, ps fasthttprouter.Params) {
	fmt.Fprint(ctx, "Welcome!\n")
}

func ChangeOrder(ctx *fasthttp.RequestCtx, ps fasthttprouter.Params) {
	fmt.Fprint(ctx, "Welcome!\n")
}

func GetOrder(ctx *fasthttp.RequestCtx, ps fasthttprouter.Params) {
	fmt.Fprint(ctx, "Welcome!\n")
}

// Hello is the Hello handler
func Hello(ctx *fasthttp.RequestCtx, ps fasthttprouter.Params) {
	fmt.Fprintf(ctx, "hello, %s!\n", ps.ByName("name"))
}

// MultiParams is the multi params handler
func MultiParams(ctx *fasthttp.RequestCtx, ps fasthttprouter.Params) {
	fmt.Fprintf(ctx, "hi, %s, %s!\n", ps.ByName("name"), ps.ByName("word"))
}

func main() {
	router := fasthttprouter.New()
	router.POST("/order", CreateOrder)
	router.GET("/order", GetOrders)
	router.PUT("/order", ChangeOrder)
	router.GET("/order/:id", GetOrder)
	router.GET("/multi/:name/:word", MultiParams)

	log.Fatal(fasthttp.ListenAndServe(":8080", router.Handler))
}