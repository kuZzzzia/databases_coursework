package database

import "log"

type Film struct {
	ID          int
	Name        string
	AltName     string
	Poster      string
	Duration    int
	Description string
	Year        int
}

func FetchFilms(pattern string) ([]*Film, error) {
	var films []*Film

	results, err := db.Query(
		"SELECT FilmID, FullName, AlternativeName, Poster, Duration, ProductionYear FROM Film WHERE FullName = ? OR AlternativeName = ?",
		pattern, pattern)
	if err != nil {
		log.Println("Error fetching films")
		return nil, err
	}

	defer results.Close()

	for results.Next() {
		film := new(Film)

		err = results.Scan(&film.ID, &film.Name, &film.AltName, &film.Poster, &film.Duration, &film.Year)
		if err != nil {
			log.Println("Error fetching films")
			return nil, err
		}

		films = append(films, film)
	}

	return films, nil
}
