package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// cli app for the service app

const URL = "http://localhost:8080"

func makeHttprequest(url string) string {
	resp, err := http.Get(fmt.Sprintf(url))
	if err != nil {
		log.Fatalln(err)
	}
	//We Read the response body on the line below.
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	//Convert the body to type string
	sb := string(body)

	return sb
}

func PlaceBid() string {

	var name string

	fmt.Println("Please enter the name of product you want to bid on") // Question to the user
	_, err := fmt.Scan(&name)
	if err != nil {
		log.Println(err)
	}

	var bidder string

	fmt.Println("Please enter the bidder name") // Question to the user
	_, err = fmt.Scan(&bidder)
	if err != nil {
		log.Println(err)
	}
	var amt int

	fmt.Println("Please enter the bid amount") // Question to the user
	_, err = fmt.Scan(&amt)
	if err != nil {
		log.Println(err)
	}

	return makeHttprequest(fmt.Sprintf("%s/bid?name=%s&bidder=%s&amt=%d", URL, name, bidder, amt))

}

func Lookup() string {

	var name string

	fmt.Println("Please enter the name of product you want to look") // Question to the user
	_, err := fmt.Scan(&name)
	if err != nil {
		log.Println(err)
	}

	return makeHttprequest(fmt.Sprintf("%s/lookup?name=%s", URL, name))

}

func Start() {

	var bidorlookup string
	var username string
	fmt.Println("Please enter your username") // Question to the user
	_, err := fmt.Scan(&username)
	if err != nil {
		log.Println(err)
	}
	fmt.Println("Bid Or Lookup (b/l)") // Question to the user

	_, err = fmt.Scan(&bidorlookup)
	if err != nil {
		log.Println(err)
	}

	switch bidorlookup {
	case "b":
		fmt.Println(PlaceBid())
		break
	case "l":
		fmt.Println(Lookup())

		break
	default:
		fmt.Println("Invalid Option Please Restart")
		break

	}
	Start()
}

func main() {
	Start()

}
