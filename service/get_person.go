package service

import (
	"golang_task/golang_task/models"
	"golang_task/golang_task/my_sql_db"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetPersonInfo(c *gin.Context) {
	// Extract the person ID from the request URL parameter
	personIDFromParam := c.Param("person_id")

	// Convert the person ID from string to integer
	personId, err := strconv.Atoi(personIDFromParam)
	if err != nil {
		// Log the error and return a bad request response if the ID is not a valid integer
		log.Println("Invalid person ID:", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"data": models.ValidationError{
				Field: "person_id",
				Error: "Invalid person ID",
			},
		})
		return
	}

	// Connect to the MySQL database
	db, err := my_sql_db.ConnectDB()
	if err != nil {
		// Log the error and return an internal server error if the database connection fails
		log.Println("Database connection failed:", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"data": models.ValidationError{
				Field: "database",
				Error: "Database connection failed",
			},
		})
		return
	}
	// Ensure the database connection is closed when the function returns
	defer db.Close()

	// Fetch person details from the database using the provided person ID
	person, err := my_sql_db.GetPersonFromDataBase(db, personId)
	if err != nil {
		// Log the error and return an internal server error if fetching data fails
		log.Println("Error fetching person data:", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"data": models.ValidationError{
				Field: "person",
				Error: "Error fetching person data",
			},
		})
		return
	}

	// Return the retrieved person details as a JSON response
	c.JSON(http.StatusOK, person)
}
