package app

import (
	"encoding/json"
	"io/ioutil"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// The function of rendering (drawing) the main page using the main template - "layout/main",
// empty fiber.Map{} is being passed due to c.Render requiring 3 parameters
func index(c *fiber.Ctx) error {
	return c.Render("index", fiber.Map{}, "layout/main")
}

// Function for rendering the car creation page using the main template - "layout/main",
// empty fiber.Map{} is being passed due to c.Render requiring 3 parameters
func createcardetailslGet(c *fiber.Ctx) error {
	return c.Render("create", fiber.Map{}, "layout/main")
}

// Function to create a new car and return the corresponding status
func cardetails(c *fiber.Ctx) error {
	//Declaring a new car element
	var item Car

	// Writing to the item variable the values received from the client, returning the error status if any
	if err := c.BodyParser(&item); err != nil {
		return c.SendStatus(501)
	}

	// Generation of a new unique id for an car
	item.Id = uuid.New().String()

	// Write data to JSON file, return error status if any
	if err := writeDataJSON(item); err != nil {
		return c.SendStatus(501)
	}

	//Return status 200 on successful completion of the request
	return c.SendStatus(200)
}

// car list reading function
func readcardetails(c *fiber.Ctx) error {
	// Setting id = "all" by default
	id := "all"

	// Specifying the length of the id parameter passed from the client. If id exists (length is greater than 0), then update the id variable
	if len(c.Params("id")) != 0 {
		id = c.Params("id")
	}

	// Reading data from a JSON file
	data := readDataJSON(id)

	// Rendering a page with car data passed to it and using the main template - "layout/main"
	return c.Render("read", fiber.Map{
		"Data": data,
	}, "layout/main")
}

func carsearch(c *fiber.Ctx) error {
	id := "all"

	if len(c.Params("id")) != 0 {
		id = c.Params("id")
	}
	data := readDataJSON(id)
	return c.Render("index", fiber.Map{
		"Data": data,
	}, "layout/main")

}

// car update function
func updatecardetails(c *fiber.Ctx) error {
	// Getting data about an car by a specific id
	item := readDataJSON(c.Params("id"))

	// Write to the item[0] variable the values received from the client, returning the error status if any.
	// It is exactly item[0] that is passed, because slice is returned from readDataJSON
	if err := c.BodyParser(&item[0]); err != nil {
		return c.SendStatus(501)
	}

	// Updating data for a specific animal, returning an error status if there is one
	if err := updateDataJSON(item[0]); err != nil {
		return c.SendStatus(501)
	}

	// Return status 200 on successful completion of the request
	return c.SendStatus(200)
}

// output function of the car to be updated
func updateCarDetailsGet(c *fiber.Ctx) error {
	// Getting data about an animal by a specific id
	data := readDataJSON(c.Params("id"))

	// Rendering the page with data about the animal and using the main template - "layout/main"
	return c.Render("update", fiber.Map{
		"Data": data[0],
	}, "layout/main")
}

// Car details removal function
func deletecardetails(c *fiber.Ctx) error {
	// Getting data on all car
	data := readDataJSON("all")

	// Removing from the received slice data about the car that was transferred from the client
	data = removeCarFromSlice(data, c.Params("id"))

	// Converting a slice from car structures to a set of bytes, returning an error status if any
	dataBytes, err := json.Marshal(data)
	if err != nil {
		return c.SendStatus(501)
	}

	// Write data to source file data.json with access rights 0644, return error status if any
	if err := ioutil.WriteFile("./data.json", dataBytes, 0644); err != nil {
		return c.SendStatus(501)
	}

	// Return status 200 on successful completion of the request
	return c.SendStatus(200)
}
