package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Food Aggregator")
	fmt.Println("Endpoint Hit: homePage")
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/buyItem", getbuyItem)
	http.HandleFunc("/buyItemQty", getbuyItemQty)
	http.HandleFunc("/buyItemQtyPrice", getbuyItemQtyPrice)
	http.HandleFunc("/summary", summaryData)
	log.Fatal(http.ListenAndServe(":3000", nil))
}

func main() {
	handleRequests()
}

func getbuyItem(w http.ResponseWriter, r *http.Request) {
	time.Sleep(2100 * time.Millisecond)
	var product string
	fmt.Println(w, "Buy Item:")
	fmt.Scanln(&product)
	if product == "Apple" || product == "banana" {
		fmt.Fprintf(w, "The following products are: \n Apple \n Apple \n banana")
	} else if product == "Carrot" || product == "okra" || product == "Onion" {
		fmt.Fprintf(w, "The following products are: \n Carrot \n okra \n Onion")
	} else if product == "wheat" || product == "barley" || product == "rye" {
		fmt.Fprintf(w, "The following products are: \n wheat \n barley \n rye")
	} else {
		fmt.Fprintf(w, "Not Found")
	}
}

func getbuyItemQty(w http.ResponseWriter, r *http.Request) {
	time.Sleep(2100 * time.Millisecond)
	var product string
	fmt.Println(w, "Buy Item:")
	fmt.Scanln(&product)
	if product == "Apple" || product == "banana" {
		fmt.Fprintf(w, "The following products are: \n Product \t Quantity \n")
		fmt.Fprintf(w, " Apple \t 30 \n Apple \t 28 \n banana \t 21 \n")
	} else if product == "Carrot" || product == "okra" || product == "Onion" {
		fmt.Fprintf(w, "The following products are: \n Product \t Quantity \n")
		fmt.Fprintf(w, " Carrot \t 13 \n okra \t 15 \n Onion \t 20 \n")
	} else if product == "wheat" || product == "barley" || product == "rye" {
		fmt.Fprintf(w, "The following products are: \n Product \t Quantity \n")
		fmt.Fprintf(w, "wheat \t 22 \n barley \t 26 \n rye \t 14 \n")
	} else {
		fmt.Fprintf(w, "Not Found")
	}
}

func getbuyItemQtyPrice(w http.ResponseWriter, r *http.Request) {
	time.Sleep(2100 * time.Millisecond)
	var product string
	fmt.Println(w, "Enter a product Name:")
	fmt.Scanln(&product)
	if product == "Apple" || product == "banana" {
		response, err := http.Get("https://run.mocky.io/v3/c51441de-5c1a-4dc2-a44e-aab4f619926b")
		if err != nil {
			fmt.Fprintf(w, "Not Found! %s\n", err)
		} else {
			data, _ := ioutil.ReadAll(response.Body)
			fmt.Fprintf(w, string(data))
		}
	} else if product == "Carrot" || product == "okra" || product == "Onion" {
		response, err := http.Get("https://run.mocky.io/v3/4ec58fbc-e9e5-4ace-9ff0-4e893ef9663c")
		if err != nil {
			fmt.Fprintf(w, "Not Found! %s\n", err)
		} else {
			data, _ := ioutil.ReadAll(response.Body)
			fmt.Fprintf(w, string(data))
		}
	} else if product == "wheat" || product == "barley" || product == "rye" {
		response, err := http.Get("https://run.mocky.io/v3/e6c77e5c-aec9-403f-821b-e14114220148")
		if err != nil {
			fmt.Fprintf(w, "Not Found! %s\n", err)
		} else {
			data, _ := ioutil.ReadAll(response.Body)
			fmt.Fprintf(w, string(data))
		}
	} else {
		fmt.Fprintf(w, "Not FOUND")
	}
}

func summaryData(w http.ResponseWriter, r *http.Request) {
	time.Sleep(2100 * time.Millisecond)
	response, err := http.Get("https://run.mocky.io/v3/c51441de-5c1a-4dc2-a44e-aab4f619926b")
	if err != nil {
		fmt.Printf("The HTTP req failed with error %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		fmt.Println(string(data))
	}

	jsonData := map[string]string{"id": "24-583-0264", "name": "Apple", "quantity": "30", "price": "$62.02"}
	jsonValue, _ := json.Marshal(jsonData) //converts map into byte data
	request, _ := http.NewRequest("POST", "https://httpbin.org/post", bytes.NewBuffer(jsonValue)) //sends data to the server
	request.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	response, err = client.Do(request) //Error Handling
	if err != nil {
		fmt.Printf("The HTTP req failed with error %s\n", err)
	} else {
		fmt.Printf("Enpoint Posted successfully\n")
		data, _ := ioutil.ReadAll(response.Body)
		fmt.Fprintln(w, string(data))
	}
}
