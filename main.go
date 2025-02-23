/* In this file, it all comes together.
I use the functions from the routes to create the endpoints that I will use to interact with my database.
*/

package main

import (
	"final-project/backend/routes/courses" // Import the routes for the courses, to utilize their functions
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

// var client *mongo.Client

func main() {
	// client = config.ConnectDB()

	// Create a chi router
	r := chi.NewRouter()

	// Define the routes for each CRUD operation, and call the respective functions
	r.Post("/create-course", func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		courses.CreateCourse(r, w, ctx)
	})

	r.Get("/read-course", func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		courses.ReadCourse(r, w, ctx)
	})

	r.Put("/update-course", func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		courses.UpdateCourse(r, w, ctx)
	})

	r.Delete("/delete-course", func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		courses.DeleteCourse(r, w, ctx)
	})

	// Start the server - running on localhost
	log.Println("Starting server on :3000")
	log.Fatal(http.ListenAndServe(":3000", r))

}
