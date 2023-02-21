package main

import (
	"fmt"
)

func main() {

	var applicationName string = "Go conference"
	const totalTickets int = 50
	var remainTickets uint = 50
	// uint is for positive numbers , so we can not assign a negative value later.
	fmt.Println("Welcome to ", applicationName)
	fmt.Printf("we have total of %v tickets , available tickets are %v\n", totalTickets, remainTickets)

	//var bookings = []string{}
	var bookings []string

	var userName string
	var userTickets uint
	for {

		fmt.Println("Enter your name:")
		fmt.Scan(&userName)
		// isValidName := len(userName) > 3
		// if !isValidName {
		// 	fmt.Println("Enter the name having greater than 3 chars")
		// 	continue
		// }
		fmt.Println("How many tickets you need:")
		fmt.Scan(&userTickets)

		if userTickets < 0 {
			fmt.Println(" enter positive number")
		}

		isValidTickets := userTickets < remainTickets

		if !isValidTickets {
			fmt.Printf("we have only %v remaining tickets\n", remainTickets)
			continue
		}

		fmt.Printf("%v booked %v tickets\n", userName, userTickets)
		remainTickets = remainTickets - userTickets
		bookings = append(bookings, userName)

		fmt.Println("Remaining tickets are:", remainTickets)

		fmt.Println("booking are:", bookings)

		noTicketsRemaining := remainTickets == 0

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
