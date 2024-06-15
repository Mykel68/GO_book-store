package main 

import "fmt"
func main() {
	

	const conferenceName string = "Go Conference"
	const TotalTickets int = 50
	var RemainingTickets uint  = 50

	fmt.Printf("Welcome to %v booking application\n", conferenceName)

	fmt.Printf("Total tickets %v\n", TotalTickets)

	var bookings [50]string

	var firstName string
	var lastName string
	var email string
	var userTickets uint
	// ask user for their name
	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email: ")
	fmt.Scan(&email)

	fmt.Println("Enter number of tickets: ")
	fmt.Scan(&userTickets)

	RemainingTickets = RemainingTickets - userTickets

	bookings[0] = firstName + " " + lastName

	fmt.Printf("The whole array: %v\n", bookings)
	fmt.Printf("The first value: %v\n", bookings[0])
	fmt.Printf("Array type %T\n", bookings)
	fmt.Printf("Array length %v\n", len(bookings))

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)

	fmt.Printf("%v tickets remaining for %v\n", RemainingTickets, conferenceName)


}