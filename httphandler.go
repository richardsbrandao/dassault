package main

import (
	"fmt"
	"log"

	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
	"gopkg.in/pg.v4"
)

var db *pg.DB

func init() {
	db = connect()
}

func connect() *pg.DB {
	return pg.Connect(&pg.Options{
		Addr: "localhost:5432",
		User: "postgres",
		Password: "postgres",
		Database: "dassault",
		SSL:false,
		PoolSize: 13000,
	})
}

// Index is the index handler
func CreateOrder(ctx *fasthttp.RequestCtx, ps fasthttprouter.Params) {


	customer := &Customer{
		Name:"Sandro",
		Email:"mileno@gmail.com",
	}

	err := db.Create(customer)
	fmt.Fprint(ctx, err)


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
	router.GET("/order", CreateOrder)
	router.PUT("/order", ChangeOrder)
	router.GET("/order/:id", GetOrder)
	router.GET("/multi/:name/:word", MultiParams)

	log.Fatal(fasthttp.ListenAndServe(":8080", router.Handler))
}

type Customer struct {
	Id    int64
	Name  string
	Email string
}
