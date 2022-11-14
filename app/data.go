package app

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

// Reading data from a JSON file with translation into a slice consisting of Car structures
func readDataJSON(id string) []Car {
	// Reading from a JSON file and logging in case of an error
	file, err := ioutil.ReadFile("./data.json")
	if err != nil {
		log.Fatal(err)
	}

	// Declaring a slice from Car structures
	var data []Car

	// Writing from JSON format (which was a string) to the data slice and logging in case of an error
	if err := json.Unmarshal([]byte(file), &data); err != nil {
		log.Fatal(err)
	}

	// id = "all" when all data is displayed to the user
	if id == "all" {

		// return data
		return data
	}

	// Declaring item variable when id belongs to a specific car
	var item Car

	// Selecting a value for item by id
	for i := 0; i < len(data); i++ {
		if data[i].Id == id {
			item = data[i]
		}
	}

	// Returning data as a slice, because function return type - slice from car structures
	return []Car{item}
}

// Writing data to a JSON file
func writeDataJSON(item Car) error {
	// Reading all data from source data file
	data := readDataJSON("all")

	// Adding a new cars to a data slice
	data = append(data, item)

	// Converting a slice from car structures to a set of bytes, returning an error if there is one
	dataBytes, err := json.Marshal(data)
	if err != nil {
		return err
	}

	//Write data to source file data.json with permission 0644, return error if any
	if err := ioutil.WriteFile("./data.json", dataBytes, 0644); err != nil {
		return err
	}

	//Returning an error if there is one
	return err
}

// Updating data in one Car
func updateDataJSON(item Car) error {
	// Reading all data from source data file
	data := readDataJSON("all")

	// Finding old data about an car by id and writing new data from item
	for i := 0; i < len(data); i++ {
		if data[i].Id == item.Id {
			data[i] = item
		}
	}

	// Converting a slice from Car structures to a set of bytes, returning an error if there is one
	dataBytes, err := json.Marshal(data)
	if err != nil {
		return err
	}

	//Write data to source file data.json with permission 0644, return error if any
	if err := ioutil.WriteFile("./data.json", dataBytes, 0644); err != nil {
		return err
	}

	//Returning an error if there is one
	return err
}

// Removing an car by id from the resulting slice
func removeCarFromSlice(s []Car, id string) []Car {
	i := 0

	// Поиск численного индекса из слайса по id
	for index, item := range s {
		if item.Id == id {
			i = index
		}
	}
	// Returning a new slice without the element with the passed id
	return append(s[:i], s[i+1:]...)
}
