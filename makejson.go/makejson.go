package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	var name, address string

	fmt.Print("Enter name: ")
	fmt.Scanln(&name)

	fmt.Print("Enter address: ")
	fmt.Scanln(&address)

	data := map[string]string{
		"name":    name,
		"address": address,
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		os.Exit(1)
	}

	fmt.Println(string(jsonData))
}
