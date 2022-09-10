package main

import (
	"fmt"
	"strings"
)

func main() {
	conferenceName := "Go Conference"
	const conferenceTickets int = 50
	var remainingTickets int = 50
	var bookings = []string{}

	fmt.Printf("Welcome to %v booking application!\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v, are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	for {

		var firstName string
		var lastName string
		var email string
		var userTickets int
		//ask user for their name
		fmt.Println("Enter your first name?")
		fmt.Scan(&firstName)

		fmt.Println("Enter your last name?")
		fmt.Scan(&lastName)

		fmt.Println("Enter your email address?")
		fmt.Scan(&email)

		fmt.Println("Enter number of tickets?")
		fmt.Scan(&userTickets)

		if userTickets > remainingTickets {
			fmt.Printf("We only have %v tickets remaining, so you can't book %v tickets\n", remainingTickets, userTickets)
			break
		}

		remainingTickets = remainingTickets - userTickets
		// bookings[0] = firstName + " " + lastName
		bookings = append(bookings, firstName+" "+lastName)

		// fmt.Printf("the whole slice: %v\n", bookings)
		// fmt.Printf("the first value: %v\n", bookings[0])
		// fmt.Printf("Slice type: %T\n", bookings)
		// fmt.Printf("Slice length: %v\n", len(bookings))

		fmt.Printf("Thank you %v %v for booking %v tickets. You will recieve a confirmation email at %v\n", firstName, lastName, userTickets, email)
		fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
		// fmt.Printf("User %v booked %v tickets.\n", firstName, userTickets)
		firstNames := []string{}
		for _, booking := range bookings {
			var names = strings.Fields(booking)
			firstNames = append(firstNames, names[0])
		}
		fmt.Printf("The first names of the bookings %v\n", firstNames)
		// noTicketsRemaining := remainingTickets == 0
		if remainingTickets == 0 {
			fmt.Println("Our conference is booked! Come back next year")
			break

		}

	}
}
