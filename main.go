package main

import (
	"bookingapp-techworld/helper"
	"fmt"
	"sync"
	"time"
)

type datastruct struct {
	username string
	city     string
	email    string
	exp      uint
} // we can have mixed data types
// these are like light weight classes

var bookingwithexp = make([]datastruct, 0)

// create variables loacl as much as possible
var applicationName string = "Go conference"

const totalTickets int = 50

var remainTickets uint = 50
var bookings = make([]map[string]string, 0)
var userName string
var userTickets uint
var city string
var email string
var exp uint

var wg = sync.WaitGroup{}

func main() {

	// uint is for positive numbers , so we can not assign a negative value later.
	greet()
	//var bookings = []string{}

	userData, userDatawithexp := userData()

	var isValidName bool = checkvalidname(userName)
	// we dont need to pass the values here since we have declared userName at the package level
	// but for refrence i have keep it like this, see the greet function above we did not pass the values
	// isValidName := checkvalidname(userName)
	if !isValidName {
		fmt.Println("Enter the name having greater than 3 chars")

	}

	fmt.Println("How many tickets you need:")
	fmt.Scan(&userTickets)

	isValidTickets := checkvaliTicket(userTickets, remainTickets)

	if !isValidTickets {
		fmt.Printf("we have only %v remaining tickets\n", remainTickets)

	}

	fmt.Printf("%v booked %v tickets\n", userName, userTickets)
	remainTickets = remainTickets - userTickets
	bookings = append(bookings, userData)
	firstName := getFirstnames()
	bookingwithexp = append(bookingwithexp, userDatawithexp)
	exp := getExp()

	fmt.Println("Remaining tickets are:", remainTickets)

	fmt.Println("Details of the teams:", bookings)
	fmt.Println("Details of the teams with exp:", bookingwithexp)

	fmt.Println("Names of Teams attending to event and exp:", firstName, exp)

	wg.Add(1)
	go sendticket(userName)

	noTicketsRemaining := checkticketavail(remainTickets)
	// we dont need to pass the values to the function since remainTickets declare at
	// package level

	// if remainTickets == 0 {
	// 	fmt.Println("Sold Out")
	// 	break
	// }

	// below two lines just for reference
	name, number := helper.ReturnTwoValues()
	fmt.Println(name, number)

	if noTicketsRemaining {
		fmt.Println("Sold Out")

	}

	wg.Wait()

}

func greet() {

	fmt.Println("Welcome to ", applicationName)
	fmt.Printf("we have total of %v tickets , available tickets are %v\n", totalTickets, remainTickets)

}

func userData() (map[string]string, datastruct) {

	var userData = make(map[string]string)

	fmt.Println("Enter your team name:")
	fmt.Scan(&userName)
	fmt.Println("Enter your city:")
	fmt.Scan(&city)
	fmt.Println("Enter your mail id")
	fmt.Scan(&email)
	fmt.Println("Enter your exp:")
	fmt.Scan(&exp)

	var userDatawithexp = datastruct{

		username: userName,
		city:     city,
		email:    email,
		exp:      exp,
	}

	userData["name"] = userName
	userData["city"] = city
	userData["email"] = email

	return userData, userDatawithexp

}

func checkvalidname(userName string) bool {

	isValidName := len(userName) > 3

	return isValidName

}

func checkvaliTicket(userTickets uint, remainTickets uint) bool {
	isValidTickets := userTickets <= remainTickets
	return isValidTickets
}

func checkticketavail(remainTickets uint) bool {
	noTicketsRemaining := remainTickets == 0

	return noTicketsRemaining
}

func getFirstnames() []string {
	firstName := []string{}
	for _, booking := range bookings {
		firstName = append(firstName, booking["name"])

	}

	return firstName
}

func getExp() []int {
	exp := []int{}
	for _, team := range bookingwithexp {
		exp = append(exp, int(team.exp))
	}

	return exp
}

func sendticket(userName string) {
	time.Sleep(30 * time.Second)
	fmt.Println("########################")
	fmt.Println("Tickets sent to", userName)
	fmt.Println("########################")
	wg.Done()

}
