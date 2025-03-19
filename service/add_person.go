package service

import (
	"database/sql"
	"golang_task/golang_task/models"
	"golang_task/golang_task/my_sql_db"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreatePerson(c *gin.Context) {
	var req models.PersonRequest

	// Bind JSON request to struct
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Println("Invalid request body:", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"data": models.ValidationError{
				Field: "request",
				Error: err.Error(),
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
			"status": "error",
			"data": models.ValidationError{
				Field: "database",
				Error: "Database connection failed",
			},
		})
		return
	}
	defer db.Close()
	// Begin Transaction
	tx, err := db.Begin()
	if err != nil {
		log.Println("Failed to start transaction:", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "error",
			"data": models.ValidationError{
				Field: "database",
				Error: "Failed to start transaction",
			},
		})
	}
	//do the db operation
	id, err := PerformDb(tx, req)
	if err != nil || id == 0 {
		log.Println("Failed to perform database operation:", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "error",
			"data": models.ValidationError{
				Field: "database",
				Error: "Failed to perform database operation",
			},
		})
	}
	// Commit Transaction only if error don't occure
	if err := tx.Commit(); err != nil {
		log.Println("Failed to commit transaction:", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "error",
			"data": models.ValidationError{
				Field: "database",
				Error: "Failed to commit transaction",
			},
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   id,
	})

}

func PerformDb(tx *sql.Tx, req models.PersonRequest) (int, error) {
	personId, err := my_sql_db.InsertPerson(tx, req.Name)
	if err != nil {
		return 0, err
	}
	_, err = my_sql_db.InsertPhone(tx, personId, req.PhoneNumber)
	if err != nil {
		return 0, err
	}
	addressId, err := my_sql_db.InsertAddress(tx, models.Address{
		City:    req.City,
		State:   req.State,
		Street1: req.Street1,
		Street2: req.Street2,
		ZipCode: req.ZipCode,
	})
	if err != nil {
		return 0, err
	}
	joinId, err := my_sql_db.InsertAddressJoin(tx, personId, addressId)
	if err != nil {
		return 0, err
	}
	return joinId, nil
}
