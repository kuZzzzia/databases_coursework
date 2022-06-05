package database

import "log"

type Person struct {
	ID      int
	Name    string
	AltName string
	Photo   string
	Date    string
}

func FetchPeople(pattern string) (*[]Person, error) {
	people := new([]Person)

	results, err := db.Query(
		"SELECT PersonID, FullName, AlternativeName, Photo, DateOfBirth FROM Person WHERE FullName = ? OR AlternativeName = ?",
		pattern)
	if err != nil {
		log.Println("Error fetching people")
		return nil, err
	}

	defer results.Close()

	for results.Next() {
		var person Person

		err = results.Scan(&person.ID, &person.Name, &person.AltName, &person.Photo, &person.Date)
		if err != nil {
			log.Println("Error fetching people")
			return nil, err
		}

		*people = append(*people, person)
	}

	return people, nil
}
