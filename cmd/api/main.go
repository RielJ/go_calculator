package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/rielj/go_calculator/api"
)

const PROMPT = ">> "

func printHelp(store api.Store) {
	fmt.Println()
	fmt.Println("Here are the items available for purchase:")

	store.PrintItems()

	fmt.Println("Please enter the name of the item you would like to purchase.")
	fmt.Println("Type 'total' to see the total of your purchases.")
	fmt.Println("Type 'member' to see the total of your purchases with membership.")
	fmt.Println("Type 'exit' to exit the program.")
	fmt.Println()
}

func main() {
	store := api.Store{}
	store.Init()

	fmt.Println("Welcome to the store!")
	printHelp(store)

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println()
		fmt.Print(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}
		line := scanner.Text()
		fmt.Println()

		if line == "help" {
			printHelp(store)
			continue
		}
		if line == "total" {
			fmt.Printf("Total: $%.2f\n", store.CalculateTotal())
			continue
		}
		if line == "member" {
			fmt.Printf("Total: $%.2f\n", store.CalculateTotalCustomerMember())
			continue
		}
		if line == "exit" {
			break
		}

		if item, ok := store.GetItemByName(line); ok {
			store.AddOrder(item.Set)
			fmt.Printf("%s Order added!\n", item.Name)
		} else {
			fmt.Println()
			fmt.Println("Item not found in store. Please try again.")
			fmt.Println()

			fmt.Println("Here are the items available for purchase:")
			store.PrintItems()
		}

	}
}
