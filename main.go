package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"net/rpc"
)

type Item struct {
	Title string
	Body  string
}

type API int

var database []Item

func (a *API) GetDB(Title string, reply *[]Item) error {
	*reply = database
	fmt.Println("Retrieve DB")
	return nil
}

func (a *API) GetByName(Title string, reply *Item) error {
	var getItem Item
	for _, val := range database {
		if val.Title == Title {
			getItem = val
		}
	}
	fmt.Println("Retrieve item")
	*reply = getItem
	return nil
}

func (a *API) CreateItem(item Item, reply *Item) error {
	database = append(database, item)
	*reply = item
	fmt.Println("Create item")
	return nil
}

func (a *API) AddItem(item Item, reply *Item) error {
	database = append(database, item)
	*reply = item
	fmt.Println("Add item")
	return nil
}

func (a *API) EditItem(edit Item, reply *Item) error {
	var changed Item

	for idx, val := range database {
		if val.Title == edit.Title {
			database[idx] = Item{edit.Title, edit.Body}
			changed = database[idx]
		}
	}
	*reply = changed
	fmt.Println("Edit item")
	return nil
}

func (a *API) DeleteItem(item Item, reply *Item) error {
	var del Item

	for idx, val := range database {
		if val.Title == item.Title && val.Body == item.Body {
			database = append(database[:idx], database[idx+1:]...)
			del = item
			break
		}
	}
	fmt.Println("Delete item")
	*reply = del
	return nil
}

func main() {

	var api = new(API)
	err := rpc.Register(api)

	if err != nil {
		log.Fatal("error registering API", err)
	}

	rpc.HandleHTTP()

	listener, err := net.Listen("tcp", ":4040")

	if err != nil {
		log.Fatal("Listener error", err)
	}

	log.Printf("serving RPC on port %d", 4040)
	err = http.Serve(listener, nil)
	if err != nil {
		log.Fatal("error serving: ", err)
	}

	// fmt.Println("intial database: ", database)
	// a := Item{"first", "a test item"}
	// b := Item{"second", "a test item"}
	// c := Item{"third", "a test item"}

	// AddItem(a)
	// AddItem(b)
	// AddItem(c)
	// fmt.Println("second database: ", database)

	// DeleteItem(b)
	// fmt.Println("third database: ", database)
	// EditItem("third", Item{"butt", "stuff"})
	// fmt.Println("fourth database: ", database)

	// fmt.Println(GetByName("first"))

}
