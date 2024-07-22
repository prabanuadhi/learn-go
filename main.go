package main

import (
	"booking-app/helper"
	"fmt"
	"sync"
	"time"
)

var conferenceName string = "Go Conference"

const conferenceTickets = 50

var remainingTickets uint = 50

// var bookings = []string{}

// var bookings = make([]map[string]string, 0)

var bookings = make([]userData, 0)

type userData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

var wg = sync.WaitGroup{}

func main() {

	greetUser()

	// for {
	if remainingTickets == 0 {
		fmt.Printf("Sorry, there's no ticket remaining\n")
		// break
	}

	firstName, lastName, email, userTickets := getInput()
	isValidName, isValidEmail := helper.ValidateInput(firstName, lastName, email)

	if userTickets > remainingTickets {
		fmt.Printf("There are only %v remaining tickets, so you cannot order %v tickets\n\n", remainingTickets, userTickets)
	} else {
		if isValidEmail && isValidName {
			bookings = bookTicket(userTickets, firstName, lastName, email)
			wg.Add(1)
			go sendTicket(userTickets, firstName, lastName, email)

			fmt.Printf("First names in bookings : %v\n\n", printFirstNames())
		} else {
			if !isValidEmail {
				fmt.Println("Email format invalid!")
			}
			if !isValidName {
				fmt.Println("Name format invalid!")
				fmt.Println()
			}
		}

	}
	wg.Wait()
	// }
}

func greetUser() {
	fmt.Printf("Welcome to our %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your ticket here to attend")
}

func printFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		// fullname := strings.Fields(booking)
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}

func getInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint
	//ask user to enter name

	fmt.Print("Enter your first name : ")
	fmt.Scan(&firstName)

	fmt.Print("Enter your last name : ")
	fmt.Scan(&lastName)

	fmt.Print("Enter your email : ")
	fmt.Scan(&email)

	fmt.Print("Enter your number of tickets : ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint, firstName string, lastName string, email string) []userData {
	remainingTickets -= userTickets

	// var userData = make(map[string]string)
	// userData["firstName"] = firstName
	// userData["lastName"] = lastName
	// userData["email"] = email
	// userData["numberOfTickets"] = strconv.FormatUint(uint64(userTickets), 10)

	var userData = userData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}

	bookings = append(bookings, userData)

	fmt.Printf("Thank you %v %v for booking %v tickets, a confirmation email will send to %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

	return bookings
}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(10 * time.Second)
	fmt.Println("#############################")
	message := fmt.Sprintf("%v tickets for %v %v sent to %v\n", userTickets, firstName, lastName, email)
	fmt.Printf("Sending ticket :\n%v\n", message)
	fmt.Println("#############################")
	wg.Done()
}
