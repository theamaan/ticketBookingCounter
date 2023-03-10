package main

import (
	"fmt"
	//"strconv"
	"strings"
	"time"
)

// package level variables:-defined at the top outside all the function
var conferenceName = "Go Conference"

const conferenceTickets int = 50 //can't change the value
var remainingTickets uint = 50

// var bookings = make([]map[string]string,0)//creating empty list of maps
var bookings = make([]UserData, 0)

type UserData struct { //struct can be like a lightweight class which doesn't support inheritence
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

func main() {
	//conferenceName := "Go Conference"
	//const conferenceTickets int = 50 //can't change the value
	//var remainingTickets uint = 50
	//Array
	//var bookings [50]string //strings with 50 elements mazimum
	//bookings[0] = "Amaan"
	//bookings[1] = "AmaanUllah"
	//to create a slice we are gonna use array without a size definition
	//var bookings []string
	//Alternative way to create a slice is :-
	//var bookings = []string{}
	//bookings := []string{}

	//greetUsers(conferenceName, conferenceTickets, remainingTickets) //calling functions explicitly
	//commented the above lines because we don't need to pass the arguments when everything is defined as package level variables
	greetUsers()
	/*
		fmt.Printf("ConferenceTickets is %T, remainingTickets is %T,conferenceName is %T\n", conferenceTickets, remainingTickets, conferenceName) //basically %T tells the type of (datatype) used int if there is int value uint if the datatype is uint or string if the datatype stored is String
		fmt.Println("")
	*/

	// ask the user for their name
	for remainingTickets > 0 && len(bookings) < 50 /*size of bookings list is less than 50*/ { //this will keep asking new booking after one booking is done
		firstName, lastName, email, userTickets := greetUserInput()
		/*
			fmt.Println("Enter Your FirstName:")
			fmt.Scan(&firstName) //fmt stands for Formatted functions: can do Print,Collect user input and write into a file
			fmt.Println("Enter Your LastName:")
			fmt.Scan(&lastName)
			fmt.Println("Enter Your E-mail address please:")
			fmt.Scan(&email)
			//fmt.Println(remainingTickets)  //printing 50
			//fmt.Println(&remainingTickets) //this will print out memory adderess for the remaining tickets // Printing:- 0xc000018098
			fmt.Println("Enter the number of tickets you want:")
			fmt.Scan(&userTickets)
		*/
		//Check wheter user ticket is more than remaining tickets
		/*
			if userTickets > remainingTickets {
				fmt.Printf("Invallid Input. We only have %v tickets, so you can't book %v tickets\n", remainingTickets, userTickets)
				//break //tells the rest of the code to skip iteration and stop the loop
				continue //instead of break we will use continue because we want to give another chance to the user to book his tickets
				//Continue causes loop to skip the remainder of its body and immmediately retesting its condition
				//just like break it will it will skip the below codes and start from all over the for loop again
			}*/

		isValidName, isValidEmail, isValidTicketNumber := validateUserInput(firstName, lastName, email, userTickets)

		//user validation check(User has to put his first name and last name atleast more than 2)
		//isValidName := len(firstName) >= 2 && len(lastName) >= 2
		//user has to put @ in his mail address
		//isValidEmail := strings.Contains(email, "@")
		//user has to enter valid number of tickets should be positive and greater than 0
		//isValidTicketNumber := userTickets > 0 && userTickets < remainingTickets

		/*
					//city has to have either Singapore or London
					isValidCity := city =="Singapore"|| city=="London"
					//city cannot have Lucknow or Delhi
			        isInValidCity := city !="Lucknow" && city!="Delhi"
		*/

		if isValidName && isValidEmail && isValidTicketNumber {
			bookTicket(userTickets, firstName, lastName, email)
			go sendTickets(userTickets, firstName, lastName, email)
			//now in function := remainingTickets = remainingTickets - userTickets

			//bookings[0] = firstName + " " + lastName
			//in slice we are gonna basically add the next element and remove the index
			//now in function :=bookings = append(bookings, firstName+" "+lastName) //much nicer because we dont have to keep the track of all the indexes
			//fmt.Printf("The whole array: %v\n", bookings) //retreiving the slice value is actually the same as array
			/*
				fmt.Printf("The whole slice: %v\n", bookings)
				fmt.Printf("The first value: %v\n", bookings[0])
				fmt.Printf("Slice type %v\n", bookings)
				fmt.Printf("Slice length: %T\n", len(bookings)) //built-in function to keep the size of Array/Slice
			*/

			/* The output(ARRAY) for the line 40 to 44 is:-
			The whole array: [Amaan Ullah                                                 ] // the line here is for the rest of the array element
			The first value: Amaan Ullah
			Array type [Amaan Ullah                                                 ]
			Array length: int
			*/

			/*
				The whole slice: [Amaan Ullah]  //You may notice that the brackets actually wrap around....So as we add elements it automatically expands
				The first value: Amaan Ullah
				Slice type [Amaan Ullah]
				Slice length: int
			*/

			//now in function :=fmt.Printf("Thankyou %v %v for booking %v tickets. You will be receiving a conformation message on your email address at %v\n", firstName, lastName, userTickets, email)
			//now in function :=fmt.Printf("%v tickets are left for %v \n", remainingTickets, conferenceName)

			// making the function to call
			firstNames := getFirstNames()
			fmt.Printf("The first names of bookings are: %v\n", firstNames)

			/*
				firstNames := []string{}
				//loop through or iterate through our booking list, grabbing one element at a time and to iterate through a list we have a pretty simple syntax actually
				for _, booking := range bookings { //inplace of index we use a Blank Indentifier (underscore) we have to explicitly make it if we are not using any variable but have to put it neccessariliy
					var names = strings.Fields(booking) /*string.Fields:-Splits the string with white space as a seperator and the fields function comes from string package so this will take our full name string split it on empty space and gives us a slice of strings separated by space. For Example:- "Amaan Ullah" will be ["Amaan","Ullah"]

					firstNames = append(firstNames, names[0])
				}

				fmt.Printf("The first names of bookings are: %v\n", firstNames)
			*/

			if remainingTickets == 0 {
				//end the program
				fmt.Println("The Conference tickets is sold out.Please come back next year.")
				break
			}
		} else { //else has to be in the same line as '}'
			if !isValidEmail {
				fmt.Println(" Your first name or last name is too short")
			}
			if !isValidEmail {
				fmt.Println("Your email address is missing @ symbol")
			}
			if !isValidTicketNumber {
				fmt.Println(" The number of tickets you entered is Invalid")
			}
			fmt.Println("Your input data is invalid. \nPlease try again :)")
			//fmt.Printf("Invallid Input. We only have %v tickets, so you can't book %v tickets\n", remainingTickets, userTickets)
			//break //tells the rest of the code to skip iteration and stop the loop
			//continue(not needed in if-else) //instead of break we will use continue because we want to give another chance to the user to book his tickets
			//Continue causes loop to skip the remainder of its body and immmediately retesting its condition
			//just like break it will it will skip the below codes and start from all over the for loop again
		}

	}

}
func greetUsers() {
	fmt.Printf("Welcome to %v booking applications\n", conferenceName)
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Println("")
	fmt.Printf("We have a total of  %v tickets and %v tickets are available\n", conferenceTickets, remainingTickets)
	fmt.Println("")
	fmt.Println("Get your tickets here to attend")
	fmt.Println("")
}

func getFirstNames() []string { //bookings is the slice of the string ([]string,outside the paranthesis is the output parameter i.e. return type)

	firstNames := []string{}
	//loop through or iterate through our booking list, grabbing one element at a time and to iterate through a list we have a pretty simple syntax actually
	for _, booking := range bookings { //inplace of index we use a Blank Indentifier (underscore) we have to explicitly make it if we are not using any variable but have to put it neccessariliy

		//var names = strings.Fields(booking) /*string.Fields:-Splits the string with white space as a seperator and the fields function comes from string package so this will take our full name string split it on empty space and gives us a slice of strings separated by space. For Example:- "Amaan Ullah" will be ["Amaan","Ullah"] */

		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames

}

// function to validate input
func validateUserInput(firstName string, lastName string, email string, userTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := userTickets > 0 && userTickets < remainingTickets
	return isValidName, isValidEmail, isValidTicketNumber
}

func greetUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Println("Enter Your FirstName:")
	fmt.Scan(&firstName) //fmt stands for Formatted functions: can do Print,Collect user input and write into a file

	fmt.Println("Enter Your LastName:")
	fmt.Scan(&lastName)

	fmt.Println("Enter Your E-mail address please:")
	fmt.Scan(&email)

	//fmt.Println(remainingTickets)  //printing 50
	//fmt.Println(&remainingTickets) //this will print out memory adderess for the remaining tickets // Printing:- 0xc000018098

	fmt.Println("Enter the number of tickets you want:")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets

}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets

	//create a map for a user

	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}

	/*
		make(map[string]string)
		userData["firstName"] = firstName
		userData["lastName"] = lastName
		userData["email"] = email
		userData["numberOfTickets"] = strconv.FormatUint(uint64(userTickets), 10)
	*/

	//bookings = append(bookings, firstName+" "+lastName)
	bookings = append(bookings, userData) //instead of a string we are now adding a map to our list
	fmt.Printf("List of bookings is %v\n", bookings)
	fmt.Printf("Thankyou %v %v for booking %v tickets. You will be receiving a confirmation message on your email address at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets are left for %v \n", remainingTickets, conferenceName)
}

func sendTickets(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("###################")
	fmt.Printf("Sending ticket %v to email address %v", ticket, email)
	fmt.Println("###################")
}
