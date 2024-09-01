package main

import (
	"fmt"
	"strings"
)

func main() {

	var conferenceName = " GO Conference"
	const conferenceTickets = 50
	var remainingTickets uint = 50
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have tickets %v available %v are avaliable \n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets now!")

	for {
		var firstName string
		var lastName string
		var userTickets uint
		var email string
		var booking []string

		fmt.Println("Please enter your first name")
		fmt.Scan(&firstName)
		fmt.Println("Please enter your last name")
		fmt.Scan(&lastName)
		fmt.Println("Please enter number of tickets you would like to buy")
		fmt.Scan(&userTickets)
		if userTickets > remainingTickets {
			fmt.Println("Sorry we only have ", remainingTickets, " tickets left")
			break
		}
		fmt.Println("Please enter your email")
		fmt.Scan(&email)
		remainingTickets = remainingTickets - userTickets
		booking = append(booking, firstName+" "+lastName)

		//fmt.Printf("Hello %v, How many tickets would you like to buy? \n", userName)
		fmt.Printf("%v like to buy %v tickets \n", firstName, userTickets)
		fmt.Printf("You will receive a confirmation email shortly on this: %v . Thank you for booking with us!\n", email)
		fmt.Println("Remaining tickets are ", remainingTickets)
		fmt.Println("Booking details are ", booking)
		//fmt.Printf("Arrey type %T\n", booking)

		firstNames := []string{}
		for _, b := range booking {
			names := strings.Fields(b)
			if len(names) > 0 {
				firstNames = append(firstNames, names[0])

			}
			fmt.Printf("First names in bookings: %v\n", firstNames)
			if remainingTickets == 0 {
				fmt.Println("Sorry we are sold out!")
				break
			}
			fmt.Println("------------------x----------------------\n")
		}
	}
}
