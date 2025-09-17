package main

import (
	"fmt"
	"os"
)

func main() {
	cache := NewWriteBackCache()

	var cmd string

	for {
		fmt.Print("Commands: write, read, flush, show, exit\n >>>")
		_, err := fmt.Scanln(&cmd)
		if err != nil {
			fmt.Println("error:", err)
			os.Exit(1)
		}

		switch cmd {
		case "write":
			var key, value string
			fmt.Print("Key: ")
			fmt.Scanln(&key)
			fmt.Print("Value: ")
			fmt.Scanln(&value)
			cache.Write(key, value)
			fmt.Printf("Written key=%s, value=%s\n", key, value)

		case "read":
			var key string
			fmt.Print("Key: ")
			fmt.Scanln(&key)
			if val, ok := cache.Read(key); ok {
				fmt.Printf("Read key=%s, value=%s\n", key, val)
			} else {
				fmt.Println("Key not found")
			}

		case "flush":
			cache.Flush()
			fmt.Println("Flushed dirty data to storage")

		case "show":
			fmt.Println("Cache:", cache.cache)
			fmt.Println("Dirty:", cache.dirty)
			fmt.Println("Storage:", cache.persistentStorage)

		case "exit":
			fmt.Println("Bye!")
			return

		default:
			fmt.Println("Unknown command:", cmd)
		}
	}
}
