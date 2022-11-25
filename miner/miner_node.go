package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"fmt"
	"github.com/gin-gonic/gin"
)

type transaction struct {
	//outline the basic transaction structure
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

func main() {
	//http request handlers
		//post a random transaction json every 1 second (send post request to transacts)
		//get a block json every 10 seconds and calculate merkelhash (get request to block)
		//post computed merkle hash to server
}






// import (
// 	"bytes"
// 	"encoding/json"
// 	"log"
// 	"net/http"
// 	"fmt"
// 	//"io/ioutil"
// 	//"github.com/gin-gonic/gin"
// )

 

// var dum = album{
// 	ID: "4", 
// 	Title: "new title", 
// 	Artist: "new artist", 
// 	Price: 24.99}

// func main() {

// 	//code for getting a block of transactions
// 	// resp, err := http.Get("http://localhost:8080/albums")

//     // if err != nil {
//     //     log.Fatal(err)
//     // }

//     // defer resp.Body.Close()

//     // body, err := ioutil.ReadAll(resp.Body)

//     // if err != nil {
//     //     log.Fatal(err)
//     // }

//     // fmt.Println(string(body))


// 	//values := map[string]string{ID: "4", Title: "new title", Artist: "new artist", Price: "24.99"}
// 	json_data, err := json.Marshal(dum)
// 	resp, err := http.Post("http://localhost:8080/albums", "application/json", bytes.NewBuffer(json_data))
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println(resp)
// 	var res album

//     json.NewDecoder(resp.Body).Decode(&res)

//     fmt.Println(res)
