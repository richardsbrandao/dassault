package main

import (
	"fmt"
	"github.com/buaazp/fasthttprouter"
	"github.com/manveru/faker"
	"github.com/valyala/fasthttp"
	"gopkg.in/pg.v4"
	"log"
	"../db"
	"../structs"
)

var db_var *pg.DB

func init() {
	fmt.Printf("Starting Conections........\n")

	db_var = db.Connect()
}

func main() {
	fmt.Printf("Starting server........\n")

	router := fasthttprouter.New()
	router.POST("/order", CreateOrder)
	router.GET("/order", CreateOrder)
	router.PUT("/order", ChangeOrder)
	router.GET("/order/:id", GetOrder)
	router.GET("/multi/:name/:word", MultiParams)

	log.Fatal(fasthttp.ListenAndServe(":8080", router.Handler))
}

// Index is the index handler
func CreateOrder(ctx *fasthttp.RequestCtx, ps fasthttprouter.Params) {

	fmt.Printf("Call  CreateOrder........\n")

	fake, err := faker.New("en")
	if err != nil {
		panic(err)
	}

	customer := &structs.Customer{
		Name:  fake.Name(),
		Email: fake.Email(),
	}

	err2 := db_var.Create(customer)
	fmt.Fprint(ctx, err2)
}

func GetOrders(ctx *fasthttp.RequestCtx, ps fasthttprouter.Params) {
	fmt.Printf("Call  GetOrders........\n")

	fmt.Fprint(ctx, "Welcome!\n")
}

func ChangeOrder(ctx *fasthttp.RequestCtx, ps fasthttprouter.Params) {
	fmt.Printf("Call  ChangeOrder........\n")

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
