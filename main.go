package main

import (
	"fmt"
)

func main() {

	var applicationName string = "Go conference"
	const totalTickets int = 50
	var remainTickets uint = 50
	// uint is for positive numbers , so we can not assign a negative value later.
	greet(applicationName, totalTickets, remainTickets)
	//var bookings = []string{}
	var bookings []string
	var userName string
	var userTickets uint

	for {

		fmt.Println("Enter your team name:")
		fmt.Scan(&userName)

		var isValidName bool = checkvalidname(userName)
		// isValidName := checkvalidname(userName)
		if !isValidName {
			fmt.Println("Enter the name having greater than 3 chars")
			continue
		}

		fmt.Println("How many tickets you need:")
		fmt.Scan(&userTickets)

		isValidTickets := checkvaliTicket(userTickets, remainTickets)

		if !isValidTickets {
			fmt.Printf("we have only %v remaining tickets\n", remainTickets)
			continue
		}

		fmt.Printf("%v booked %v tickets\n", userName, userTickets)
		remainTickets = remainTickets - userTickets
		bookings = append(bookings, userName)

		fmt.Println("Remaining tickets are:", remainTickets)

		fmt.Println("Teams attending to event:", bookings)

		noTicketsRemaining := checkticketavail(remainTickets)

		// if remainTickets == 0 {
		// 	fmt.Println("Sold Out")
		// 	break
		// }

		if noTicketsRemaining {
			fmt.Println("Sold Out")
			break
		}

	}

}

func greet(applicationName string, totalTickets int, remainTickets uint) {

	fmt.Println("Welcome to ", applicationName)
	fmt.Printf("we have total of %v tickets , available tickets are %v\n", totalTickets, remainTickets)

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
