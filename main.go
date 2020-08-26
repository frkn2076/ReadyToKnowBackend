package main

import (
	"github.com/gofiber/fiber" // import the fiber package
"log"
"github.com/gofiber/fiber/middleware"
"github.com/firebase007/go-rest-api-with-fiber/database"
"github.com/firebase007/go-rest-api-with-fiber/router"

_ "github.com/lib/pq"
)



// entry point to our program
func main() { 
	// Connect to database
	if err := database.Connect(); err != nil {
		log.Fatal(err)
	  }
  
   // call the New() method - used to instantiate a new Fiber App
	app := fiber.New()
  
	// Middleware
	app.Use(middleware.Logger())
  
	router.SetupRoutes(app)
  
	// listen on port 3000
	app.Listen(3000) 
  
  }

//   func main() {
// 	app := fiber.New()

// 	app.Get("/", func(c *fiber.Ctx) {
// 		c.Send("Hello, World 👋!")
// 	})

// 	app.Post("/", func(c *fiber.Ctx) {
// 		// Get raw body from POST request
// 		c.Send(c.Body()) // user=john
// 	})

// 	app.Listen(3000)
// }