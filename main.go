package main 

import "fmt"
func main() {
	fmt.Printf(`Welcome to our my booking application
	`)
	const TotalTickets = 50

	fmt.Printf("Total tickets %v\n", TotalTickets)

	var userName string
	var userTickets int
	// ask user for their name

	userName = "Mykel"
	userTickets = 2

	fmt.Printf("User: %v booked %v tickets\n", userName, userTickets)

}