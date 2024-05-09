package main

import (
	"LootData4/GET"
	"LootData4/POST"
	ws "LootData4/websocket"
	"fmt"
	"log"
	"net/http"
)

func main() {

	//POSTS
	http.HandleFunc("/", POST.Loot_Model_Handler)
	http.HandleFunc("/resources", POST.Resource_Model_Handler)
	http.HandleFunc("/returns", POST.Timer_Model_Handler)

	//GET
	http.HandleFunc("/overview", GET.World_Status_Model_Handler)
	http.HandleFunc("/report", GET.Loot_Model_Handler)

	//WS
	http.HandleFunc("/ws", ws.Websocket_Handler)

	// Start the server
	fmt.Println("Starting server at port 8080")
	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err)
	}
}
