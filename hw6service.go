package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type CurrentBestBid struct {
	Amount int    `json:"amount"`
	Bidder string `json:"bidder"`
}

type Item struct {
	Name           string           `json:"name"`
	Description    string           `json:"description"`
	MinimumBid     int              `json:"minimumBid"`
	CurrentBestBid CurrentBestBid   `json:"currentBestBid"`
	Allbids        []CurrentBestBid `json:"allbids"`
}

var itemsArray []Item

// This handler adds item to the itemsArray

func writeResponse(w http.ResponseWriter, data interface{}) {
	resp := struct {
		Message string      `json:"message"`
		Data    interface{} `json:"data"`
	}{
		Message: "success",
		Data:    data,
	}

	err := json.NewEncoder(w).Encode(&resp)

	if err != nil {
		log.Println(err)
		return
	}

}

func AddHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	description := r.URL.Query().Get("description")
	minimum := r.URL.Query().Get("minimum")
	bidAmount, err := strconv.Atoi(minimum)
	if err != nil {
		writeResponse(w, "Bid Amount should Be Integer")
		return
	}
	current := r.URL.Query().Get("current")
	currentAmount, err := strconv.Atoi(current)
	if err != nil {
		writeResponse(w, "current Amount should Be Integer")
		return
	}

	bidder := r.URL.Query().Get("bidder")

	for i := 0; i < len(itemsArray); i++ {
		if itemsArray[i].Name == name {
			writeResponse(w, "Item Already Exists")

			return

		}

	}

	// If item doesn't exists we can add item

	cb := CurrentBestBid{
		currentAmount,
		bidder,
	}
	allbids := []CurrentBestBid{
		cb,
	}

	item := Item{
		Name:           name,
		Description:    description,
		MinimumBid:     bidAmount,
		CurrentBestBid: cb,
		Allbids:        allbids,
	}

	itemsArray = append(itemsArray, item)
	writeResponse(w, "Item Added SuccessFully")

}

// This handler Places bid on particular item searching on the basis of its name

func BidHandler(w http.ResponseWriter, r *http.Request) {

	name := r.URL.Query().Get("name")

	bidder := r.URL.Query().Get("bidder")
	minimum := r.URL.Query().Get("amt")

	bidAmount, err := strconv.Atoi(minimum)
	if err != nil {
		writeResponse(w, "Bid Amount should Be Integer")
		return
	}

	for i := 0; i < len(itemsArray); i++ {
		if itemsArray[i].Name == name {

			if bidAmount >= itemsArray[i].MinimumBid {

				// Add the bid to the all bids array
				cb := CurrentBestBid{
					bidAmount,
					bidder,
				}
				itemsArray[i].Allbids = append(itemsArray[i].Allbids, cb)

				// checking the  current best bid if it has lesser amount then will allot this bid to current best bid

				if bidAmount > itemsArray[i].CurrentBestBid.Amount {
					itemsArray[i].CurrentBestBid.Amount = bidAmount
					itemsArray[i].CurrentBestBid.Bidder = bidder
				}
				writeResponse(w, "bidding Made SuccessFully")
				return

			} else {
				writeResponse(w, fmt.Sprintf("Minimum bid of %d is needed", itemsArray[i].MinimumBid))
				return

			}

		}

	}
	writeResponse(w, "bidding item not found")

}

func HealthHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("arrived")
	w.Header().Set("Content-Type", "application/json")

	writeResponse(w, "Server Running")
}

// This handler will look up for the item from the slice

func LookupHandler(w http.ResponseWriter, r *http.Request) {
	// Returning the bid by its name

	name := r.URL.Query().Get("name")

	for i := 0; i < len(itemsArray); i++ {
		if itemsArray[i].Name == name {

			writeResponse(w, itemsArray[i])
			return

		}

	}
	writeResponse(w, "item Not found")
}

func main() {
	http.HandleFunc("/", HealthHandler)
	http.HandleFunc("/add", AddHandler)
	http.HandleFunc("/bid", BidHandler)
	http.HandleFunc("/lookup", LookupHandler)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalln(err)
	}
}
