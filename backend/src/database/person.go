package database

import "log"

type Person struct {
	ID      int
	Name    string
	AltName string
	Photo   string
	Date    string
}

type Role struct {
	FilmID   int
	Name     string
	FilmName string
	Year     int
}

func FetchPeople(pattern string) ([]*Person, error) {
	var people []*Person

	results, err := db.Query(
		"SELECT PersonID, FullName, AlternativeName, Photo, DateOfBirth FROM Person WHERE FullName = ? OR AlternativeName = ?",
		pattern, pattern)
	if err != nil {
		log.Println("Error fetching people")
		return nil, err
	}

	defer results.Close()

	for results.Next() {
		person := new(Person)

		err = results.Scan(&person.ID, &person.Name, &person.AltName, &person.Photo, &person.Date)
		if err != nil {
			log.Println("Error fetching people")
			return nil, err
		}

		people = append(people, person)
	}

	return people, nil
}

func FetchPerson(id int) (*Person, []*Role, []*Film, error) {
	person := new(Person)

	err := db.QueryRow(
		"SELECT PersonID, FullName, AlternativeName, Photo, DateOfBirth FROM Person WHERE PersonID = ?",
		id).Scan(&person.ID, &person.Name, &person.AltName, &person.Photo, &person.Date)
	if err != nil {
		log.Println("Error fetching person")
		return nil, nil, nil, err
	}

	var films []*Film

	results, err := db.Query(
		"SELECT FilmID, FullName, ProductionYear FROM Film WHERE PersonID = ? ORDER BY ProductionYear DESC",
		id)
	if err != nil {
		log.Println("Error fetching films")
		return nil, nil, nil, err
	}

	defer results.Close()
	for results.Next() {
		film := new(Film)

		err = results.Scan(&film.ID, &film.Name, &film.Year)
		if err != nil {
			log.Println("Error fetching films")
			return nil, nil, nil, err
		}

		films = append(films, film)
	}

	var roles []*Role

	results, err = db.Query(
		"SELECT r.FilmID, r.CharacterName, f.FullName, f.ProductionYear FROM (SELECT FilmID, CharacterName From Role WHERE PersonID = ?) AS r LEFT JOIN Film AS f ON f.FilmID = r.FilmID ORDER BY f.ProductionYear DESC",
		id)
	if err != nil {
		log.Println("Error fetching roles")
		return nil, nil, nil, err
	}

	defer results.Close()
	for results.Next() {
		role := new(Role)

		err = results.Scan(&role.FilmID, &role.Name, &role.FilmName, &role.Year)
		if err != nil {
			log.Println("Error fetching roles")
			return nil, nil, nil, err
		}

		roles = append(roles, role)
	}

	return person, roles, nil, nil
}
