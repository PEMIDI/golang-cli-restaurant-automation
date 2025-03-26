package main

import "fmt"

type Food struct {
	Name  string
	Price float64
}

const (
	pizzaPrice  = 99.99
	burgerPrice = 110.0
	tacoPrice   = 100.0
)

func generateFood() []Food {
	return []Food{
		{Name: "pizza", Price: pizzaPrice},
		{Name: "burger", Price: burgerPrice},
		{Name: "Taco", Price: tacoPrice},
	}
}

func main() {
	fmt.Println("Welcome to the CLI Restaurant Reservation System!")

	foods := generateFood()
	for i, food := range foods {
		println(i+1, food.Name)
	}

	var userMenu []Food

	for {
		var userInput string

		println("Choose a number to select a Food: enter 4 to confirm order ")
		fmt.Scanln(&userInput)

		if userInput == "exit" {
			break
		}

		switch userInput {
		case "1":
			userMenu = append(userMenu, foods[0])
		case "2":
			userMenu = append(userMenu, foods[1])
		case "3":
			userMenu = append(userMenu, foods[2])
		case "4":
			totalPrice := 0.0
			for _, food := range userMenu {
				totalPrice += food.Price
			}
			fmt.Printf("price is $%.2f", totalPrice)
			return

		case "exit":
			fmt.Println("goodbye")

		default:
			print("Invalid input!")
		}

		fmt.Println(userMenu)
	}
}
