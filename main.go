package main

import (
	"LootData4/GET"
	"LootData4/POST"
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
	//need loot GET function still

	// Start the server
	fmt.Println("Starting server at port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
