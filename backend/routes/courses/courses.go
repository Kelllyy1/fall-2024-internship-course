/* In this file, I use the Courses schema to create default functions that
I will be able to use for each course I wish to create and ultimetly store in the database in the future.
*/

/* I want to be able to do these things to a course:

Create
Read
Update
Delete

*/

package courses

import (
	// "context"
	"context"
	"encoding/json"
	"final-project/backend/config"
	"final-project/backend/models/Courses"
	"fmt"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateCourse(req *http.Request, res http.ResponseWriter, ctx context.Context) {
	db := config.ConnectDB() // Get the database instance
	collection := db.Database("final-project").Collection("courses")

	// Parse the request body to create the course object
	var course Courses.CourseSchema
	err := json.NewDecoder(req.Body).Decode(&course)
	if err != nil {
		log.Printf("Error decoding request body: %v", err)
		http.Error(res, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Validate the course object
	if course.Name == nil || course.Timeline == nil || course.Deadline == nil || course.Skills == nil || course.AreaImpacted == nil {
		log.Println("Missing required fields")
		http.Error(res, "Missing required fields", http.StatusBadRequest)
		return
	}

	// Insert the new course into MongoDB
	result, err := collection.InsertOne(ctx, course)
	if err != nil {
		log.Printf("Failed to save course: %v", err)
		http.Error(res, "Failed to save course", http.StatusInternalServerError)
		return
	}

	// Retrieve the inserted document
	var createdCourse Courses.CourseSchema
	err = collection.FindOne(ctx, bson.M{"_id": result.InsertedID}).Decode(&createdCourse)
	if err != nil {
		log.Printf("Failed to retrieve created course: %v", err)
		http.Error(res, "Failed to retrieve created course", http.StatusInternalServerError)
		return
	}

	// Log the created course to the terminal with dereferenced values
	log.Printf("Created course: Name=%s, Timeline=%s, Deadline=%s, Skills=%s, AreaImpacted=%s",
		*createdCourse.Name, *createdCourse.Timeline, createdCourse.Deadline.Format("2006-01-02"), *createdCourse.Skills, *createdCourse.AreaImpacted)

	// Respond with the created course
	res.WriteHeader(http.StatusCreated)
	res.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(res)
	encoder.SetIndent("", "  ") // Set indentation for pretty-printing
	encoder.Encode(createdCourse)
}

// ReadCourse handles reading a course by ID and returns an *http.Response
func ReadCourse(req *http.Request, res http.ResponseWriter, ctx context.Context) {
	db := config.ConnectDB() // Get the database instance
	collection := db.Database("final-project").Collection("courses")

	// Get the course ID from the URL
	id := req.URL.Query().Get("id")
	if id == "" {
		http.Error(res, "Missing course ID", http.StatusBadRequest)
		return
	}

	// Convert the ID to an ObjectID
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		http.Error(res, "Invalid course ID", http.StatusBadRequest)
		return
	}

	// Retrieve the course from MongoDB
	var course Courses.CourseSchema
	err = collection.FindOne(ctx, bson.M{"_id": objectID}).Decode(&course)
	if err != nil {
		http.Error(res, "Course not found", http.StatusNotFound)
		return
	}

	// Respond with the course
	res.WriteHeader(http.StatusOK)
	res.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(res)
	encoder.SetIndent("", "  ") // Set indentation for pretty-printing
	encoder.Encode(course)

	fmt.Println("Course found!")
}

// UpdateCourse handles updating a course by ID and returns an *http.Response
func UpdateCourse(req *http.Request, res http.ResponseWriter, ctx context.Context) {
	db := config.ConnectDB() // Get the database instance
	collection := db.Database("final-project").Collection("courses")

	// Get the course ID from the URL
	id := req.URL.Query().Get("id")
	if id == "" {
		http.Error(res, "Missing course ID", http.StatusBadRequest)
		return
	}

	// Convert the ID to an ObjectID
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		http.Error(res, "Invalid course ID", http.StatusBadRequest)
		return
	}

	// Parse the request body to update the course object
	var course Courses.CourseSchema
	err = json.NewDecoder(req.Body).Decode(&course)
	if err != nil {
		http.Error(res, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Validate the course object
	if course.Name == nil || course.Timeline == nil || course.Deadline == nil || course.Skills == nil || course.AreaImpacted == nil {
		http.Error(res, "Missing required fields", http.StatusBadRequest)
		return
	}

	// Update the course in MongoDB
	update := bson.M{
		"$set": course,
	}
	_, err = collection.UpdateOne(ctx, bson.M{"_id": objectID}, update)
	if err != nil {
		http.Error(res, "Failed to update course", http.StatusInternalServerError)
		return
	}

	// Retrieve the updated document
	var updatedCourse Courses.CourseSchema
	err = collection.FindOne(ctx, bson.M{"_id": objectID}).Decode(&updatedCourse)
	if err != nil {
		http.Error(res, "Failed to retrieve updated course", http.StatusInternalServerError)
		return
	}

	// Respond with the updated course
	res.WriteHeader(http.StatusOK)
	res.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(res)
	encoder.SetIndent("", "  ") // Set indentation for pretty-printing
	encoder.Encode(updatedCourse)
	fmt.Println("Course updated successfully!")
}

// DeleteCourse handles deleting a course by ID and returns an *http.Response
func DeleteCourse(req *http.Request, res http.ResponseWriter, ctx context.Context) {
	db := config.ConnectDB() // Get the database instance
	collection := db.Database("final-project").Collection("courses")

	// Get the course ID from the URL
	id := req.URL.Query().Get("id")
	if id == "" {
		http.Error(res, "Missing course ID", http.StatusBadRequest)
		return
	}

	// Convert the ID to an ObjectID
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		http.Error(res, "Invalid course ID", http.StatusBadRequest)
		return
	}

	// Delete the course from MongoDB
	result, err := collection.DeleteOne(ctx, bson.M{"_id": objectID})
	if err != nil {
		log.Printf("Failed to delete course: %v", err)
		http.Error(res, "Failed to delete course", http.StatusInternalServerError)
		return
	}

	// Check if the course was actually deleted
	if result.DeletedCount == 0 {
		http.Error(res, "Course not found", http.StatusNotFound)
		return
	}
	fmt.Println("Course deleted successfully!")

	// Respond with a success message
	res.WriteHeader(http.StatusOK)
	res.Header().Set("Content-Type", "application/json")
	json.NewEncoder(res).Encode(map[string]string{"message": "Course deleted successfully"})
}
