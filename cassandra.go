package main

import (
	"fmt"
	//"log"

	//"github.com/buaazp/fasthttprouter"
	//"github.com/valyala/fasthttp"
	"github.com/gocql/gocql"
)



func main() {
	cluster := gocql.NewCluster("172.17.0.2")
    cluster.Keyspace = "ironman_dev"
	cluster.Consistency = gocql.One

	session, _ := cluster.CreateSession()
    defer session.Close()

    /*if err := session.Query("INSERT INTO wallet_by_user_and_trip (user_id, trip_id, id_wallet, st_default_currency) VALUES (?, ?, ?, ?)",
        gocql.TimeUUID(), gocql.TimeUUID(), gocql.TimeUUID(), "BRL").Exec(); err != nil {
        fmt.Println(err)
    }*/

    var user_id gocql.UUID
    var trip_id gocql.UUID
    var id_wallet gocql.UUID
    var st_default_currency string

    /*if err := session.Query("SELECT * FROM wallet_by_user_and_trip WHERE user_id = ?",
        "5de312b2-546a-11e6-906e-74e6e2ce84aa").Scan(&user_id, &trip_id, &id_wallet, &st_default_currency); err != nil {
        fmt.Println(err)
    }

    fmt.Printf("FindOne Wallet: user_id=[%v] trip_id[%v] id_wallet[%v] st_default_currency[%v]", user_id, trip_id, id_wallet, st_default_currency)
	*/

	iter := session.Query("SELECT * FROM wallet_by_user_and_trip WHERE user_id = ?", "5de312b2-546a-11e6-906e-74e6e2ce84aa").Iter()
    
    for iter.Scan(&user_id, &trip_id, &id_wallet, &st_default_currency) {
        fmt.Printf("List Wallet: user_id=[%v] trip_id[%v] id_wallet[%v] st_default_currency[%v]", user_id, trip_id, id_wallet, st_default_currency)
    }
    if err := iter.Close(); err != nil {
        fmt.Println(err)
    }

}