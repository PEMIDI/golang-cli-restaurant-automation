package main

import "fmt"

func main() {
	fmt.Println("Welcome to the CLI Restaurant Reservation System!")

	type Food struct {
		Name  string
		Price float64
	}

	// init the foods
	pizza := Food{Name: "Pizza", Price: 99.99}
	burger := Food{Name: "burger", Price: 110.0}
	taco := Food{Name: "Taco", Price: 100.0}

	// show menu
	fmt.Println("- 1 - ", pizza.Name, " - ", pizza.Price)
	fmt.Println("- 2 - ", burger.Name, " - ", burger.Price)
	fmt.Println("- 3 - ", taco.Name, " - ", taco.Price)

	//get user input

	// init user menu
	var userMenu []Food

	//check user input

	for {

		var userInput int
		println("Choose a number to select a Food:")
		fmt.Scanln(&userInput)

		switch userInput {
		case 1:
			userMenu = append(userMenu, pizza)
		case 2:
			userMenu = append(userMenu, burger)
		case 3:
			userMenu = append(userMenu, taco)
		case 4:
			break

		default:
			print("Invalid input!")
		}

		fmt.Println(userMenu)
	}

}

// TODO: Load menu data (hardcoded or from a file).

// TODO: Display the menu with Food name, type, and price.
// TODO: Allow the user to select Food items from the menu.
// TODO: Calculate the total price based on selected items.
// TODO: Confirm the order and display a summary.

// TODO: Allow users to make a reservation by entering name, number of guests, and time.
// TODO: Store order history in a JSON file or SQLite database.
// TODO: Implement an "Admin Mode" where an admin can add/remove menu items.

// TODO: Implement a clean exit mechanism for the CLI.
