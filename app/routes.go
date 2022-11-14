package app

import "github.com/gofiber/fiber/v2"

func Routes(app *fiber.App) {
	app.Get("/", index) // Main Page

	app.Get("/car/create", createcardetailslGet) // // Page for creating a new car
	app.Post("/car/create", cardetails)          // Adding a new car
	app.Get("car/search", carsearch)             // Page for displaying searched data
	app.Get("/car/:id?", readcardetails)         // Page for displaying data on a set of cars, or one with saving the template

	app.Get("/car/update/:id", updateCarDetailsGet) // Data output page for updating cars data
	app.Post("/car/update/:id", updatecardetails)   // carl data update

	app.Get("/car/delete/:id", deletecardetails) //The page that is called to delete data about an cars
}
