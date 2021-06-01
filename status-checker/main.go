package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"https://google.com",
		"https://facebook.com",
		"https://instagram.com",
		"https://yahoo.com",
		"https://hotmail.com",
		"https://flipkart.com",
	}

	c := make(chan string) //a channel to save us!

	for _, link := range links {
		//the bad way/time taking way -> StatusCheck(link)
		go StatusCheck(link, c) //go -> create a new thread go routine!
	}
	//fmt.Println(<-c) //a blocking call

	//infinite loop for{....}

	for l := range c {
		go func(link string) { //function literal
			time.Sleep(time.Second) //pauses the go routine for 1 second.
			StatusCheck(link, c)
		}(l) // () this is used to call the function!
	}
}

func StatusCheck(link string, c chan string) {
	_, err := http.Get(link) //wait LOL

	if err != nil {
		fmt.Println(link, " site is Down!!!!!")
		c <- link
		return
	}

	fmt.Println(link, " working...")
	c <- link
}
