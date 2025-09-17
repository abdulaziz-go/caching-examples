package main

import (
	"fmt"
)

func main() {
	store := NewEventualConsistentStore()

	for {
		fmt.Println("Choose: w, r")
		var choice string
		fmt.Scanln(&choice)

		switch choice {
		case "w":
			var key, val string
			fmt.Print("Enter key: ")
			fmt.Scanln(&key)
			fmt.Print("Enter value: ")
			fmt.Scanln(&val)
			store.Write(key, val)
			fmt.Println("Write successful (replicating in background...)")

		case "r":
			var node string
			fmt.Print("Enter node id (n1/n2/n3): ")
			fmt.Scanln(&node)
			data := store.Read(node)
			if data == nil {
				fmt.Println("Node not found")
			} else {
				fmt.Printf("Data at %s: %v\n", node, data)
			}
		default:
			fmt.Println("Invalid choice")
		}
	}
}
