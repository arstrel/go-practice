package main

import (
	"fmt"
	"strings"
)

func main() {
	conferenceName := "Go Conference"
	const conferenceTickets = 50
	var remainingTickets uint = 50

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	// var bookings = [50]string{"Nama", "Nuckole", "Jams"}
	var bookings []string

	// infinite loop
	for {
		var firstName string
		var lastName string
		var email string
		var userTickets uint
		// ask the user for their name
		fmt.Println("Enter your first name:")
		fmt.Scan(&firstName)

		fmt.Println("Enter your last name:")
		fmt.Scan(&lastName)

		fmt.Println("Enter your email:")
		fmt.Scan(&email)

		fmt.Println("Enter how many tickets you want:")
		fmt.Scan(&userTickets)

		isValidName := len(firstName) >= 2 && len(lastName) >= 2
		isValidEmail := strings.Contains(email, "@")
		isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets

		if isValidName && isValidEmail && isValidTicketNumber {
			remainingTickets = remainingTickets - userTickets
			bookings = append(bookings, firstName+" "+lastName)

			fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
			fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

			firstNames := []string{}

			for _, fullName := range bookings {
				name := strings.Fields(fullName)[0]
				firstNames = append(firstNames, name)
			}

			fmt.Printf("These are all our bookinds %v\n", firstNames)

			if remainingTickets == 0 {
				fmt.Printf("All tickets are sold out!\n")
				break
			}
		} else {
			fmt.Println("Your input data is invalid. Try again")
			continue
		}

	}

}
