package main

import (
	"fmt"
	"log"
	"net/rpc"
)

type Item struct {
	Title string
	Body  string
}

func main() {
	var reply Item
	var db []Item

	client, err := rpc.DialHTTP("tcp", "localhost:4040")

	if err != nil {
		log.Fatal("Connection error: ", err)
	}

	a := Item{"First", "item 1"}
	b := Item{"Second", "item 2"}
	c := Item{"Third", "item 3"}

	client.Call("API.AddItem", a, &reply)
	client.Call("API.AddItem", b, &reply)
	client.Call("API.AddItem", c, &reply)
	client.Call("API.GetDB", "", &db)

	fmt.Println("db: ", db)
	client.Call("API.EditItem", Item{"Second", "a new item"}, &reply)
	client.Call("API.DeleteItem", "first", &reply)
	client.Call("GetDB", "", &db)
	fmt.Println("db: ", db)

}
