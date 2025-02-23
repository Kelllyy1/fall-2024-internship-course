/* In this file, I use the Goals schema to create default functions that
I will be able to use for each goal I wish to create and ultimetly store in the database in the future.
*/

/*
	I want to be able to do these things to a goal:

Create
Read
Update
Delete
*/
package goals

import (
	// "context"
	// "final-project/backend/models/Courses"
	"fmt"
	// "log"
	"net/http"
	// "time"
	// "github.com/gin-gonic/gin" // an alternative router to "github.com/go-chi/chi/v5"
	// "github.com/go-chi/v5"
	// "go.mongodb.org/mongo-driver/mongo"
	// "go.mongodb.org/mongo-driver/mongo/options"
)

// Create a new course
func CreateGoal(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Implement the same functions used in \"course.go\" in order to handle goals!\n This process can be repeated for any additional objects you wish to create.")
}
