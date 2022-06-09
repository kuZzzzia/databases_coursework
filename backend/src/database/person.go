package database

import (
	"database/sql"
	"log"
)

type Person struct {
	ID      int
	Name    string
	AltName sql.NullString
	Photo   string
	Date    sql.NullString
}

type Role struct {
	FilmID     int
	Name       sql.NullString
	FilmName   string
	Year       sql.NullInt16
	FilmRating int
}

func FetchPeople(pattern string) ([]*Person, error) {
	var people []*Person

	results, err := db.Query(
		"SELECT PersonID, FullName, AlternativeName, Photo, DateOfBirth FROM Person WHERE FullName = ? OR AlternativeName = ?",
		pattern, pattern)
	if err != nil {
		log.Println("Error fetching people: " + err.Error())
		return nil, err
	}

	defer results.Close()

	for results.Next() {
		person := new(Person)

		err = results.Scan(&person.ID, &person.Name, &person.AltName, &person.Photo, &person.Date)
		if err != nil {
			log.Println("Error fetching people: " + err.Error())
			return nil, err
		}

		people = append(people, person)
	}

	return people, nil
}

func FetchRoles(id int) ([]*Role, error) {
	var roles []*Role

	results, err := db.Query(
		"SELECT FilmID, CharacterName, FilmName, ProductionYear, getFilmRating(FilmID) From Film_Cast WHERE PersonID = ? ORDER BY ProductionYear DESC",
		id)
	if err != nil {
		log.Println("Error fetching roles: " + err.Error())
		return nil, err
	}

	defer results.Close()
	for results.Next() {
		role := new(Role)

		err = results.Scan(&role.FilmID, &role.Name, &role.FilmName, &role.Year, &role.FilmRating)
		if err != nil {
			log.Println("Error fetching roles: " + err.Error())
			return nil, err
		}

		roles = append(roles, role)
	}
	return roles, nil
}

func FetchPerson(id int) (*Person, []*Role, []*Film, error) {
	person := new(Person)

	err := db.QueryRow(
		"SELECT PersonID, FullName, AlternativeName, Photo, DateOfBirth FROM Person WHERE PersonID = ?",
		id).Scan(&person.ID, &person.Name, &person.AltName, &person.Photo, &person.Date)
	if err != nil {
		log.Println("Error fetching person: " + err.Error())
		return nil, nil, nil, err
	}

	films, err := FetchFilmByDirector(id)
	if err != nil {
		return nil, nil, nil, err
	}

	roles, err := FetchRoles(id)
	if err != nil {
		return nil, nil, nil, err
	}

	return person, roles, films, nil
}
