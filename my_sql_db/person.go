package my_sql_db

import (
	"database/sql"
	"golang_task/golang_task/models"
	"log"
)

func GetPersonFromDataBase(db *sql.DB, personId int) (models.PersonResponse, error) {
	query := `
		SELECT p.name, ph.number, a.city, a.state, a.street1, a.street2, a.zip_code 
		FROM person p
		JOIN phone ph ON p.id = ph.person_id
		JOIN address_join aj ON p.id = aj.person_id
		JOIN address a ON aj.address_id = a.id
		WHERE p.id = ?
	`
	var person models.PersonResponse
	err := db.QueryRow(query, personId).Scan(&person.Name, &person.PhoneNumber, &person.City, &person.State, &person.Street1, &person.Street2, &person.ZipCode)
	if err != nil {
		log.Println("Error fetching data:", err)
		return models.PersonResponse{}, err
	}

	return person, nil
}

func InsertPerson(db *sql.Tx, name string) (int, error) {
	// Insert into 'person' table
	result, err := db.Exec("INSERT INTO person (name) VALUES (?)", name)
	if err != nil {
		log.Println("Error inserting into person table:", err)
		return 0, err
	}
	personId, err := result.LastInsertId()
	if err != nil {
		log.Println("Error getting last inserted ID:", err)
		return 0, err
	}
	return int(personId), nil
}

func InsertPhone(db *sql.Tx, personId int, PhoneNumber string) (int, error) {
	// Insert into 'person' table
	result, err := db.Exec("INSERT INTO phone (person_id, number) VALUES (?, ?)", personId, PhoneNumber)
	if err != nil {
		log.Println("Error inserting into phone table:", err)
		return 0, err
	}
	phoneId, err := result.LastInsertId()
	if err != nil {
		log.Println("Error getting last inserted ID:", err)
		return 0, err
	}
	return int(phoneId), nil
}

func InsertAddress(db *sql.Tx, req models.Address) (int, error) {
	// Insert into 'person' table
	result, err := db.Exec("INSERT INTO address (city, state, street1, street2, zip_code) VALUES (?, ?, ?, ?, ?)",
		req.City, req.State, req.Street1, req.Street2, req.ZipCode)
	if err != nil {
		log.Println("Error inserting into address table:", err)
		return 0, err
	}
	adressId, err := result.LastInsertId()
	if err != nil {
		log.Println("Error getting last inserted ID:", err)
		return 0, err
	}
	return int(adressId), nil
}

func InsertAddressJoin(db *sql.Tx, personId int, addressId int) (int, error) {
	// Insert into 'person' table
	result, err := db.Exec("INSERT INTO address_join (person_id, address_id) VALUES (?, ?)", personId, addressId)
	if err != nil {
		log.Println("Error inserting into address_join table:", err)
		return 0, err
	}
	joinId, err := result.LastInsertId()
	if err != nil {
		log.Println("Error getting last inserted ID:", err)
		return 0, err
	}
	return int(joinId), nil
}
