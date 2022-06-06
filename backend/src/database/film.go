package database

import (
	"database/sql"
	"log"
)

type Film struct {
	ID            int
	Name          string
	AltName       sql.NullString
	Poster        string
	Duration      sql.NullInt16
	Description   sql.NullString
	Year          sql.NullInt16
	Director      sql.NullString
	DirectorID    sql.NullInt16
	Rating        int
	LikeAmount    int
	DislikeAmount int
}

type CastItem struct {
	ID        int
	Name      string
	Character sql.NullString
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

func FetchFilmByDirector(id int) ([]*Film, error) {
	var films []*Film

	results, err := db.Query(
		"SELECT FilmID, FullName, ProductionYear, getFilmRating(FilmID) FROM Film WHERE PersonID = ? ORDER BY ProductionYear DESC",
		id)
	if err != nil {
		log.Println("Error fetching films")
		return nil, err
	}

	defer results.Close()
	for results.Next() {
		film := new(Film)

		err = results.Scan(&film.ID, &film.Name, &film.Year, &film.Rating)
		if err != nil {
			log.Println("Error fetching films")
			return nil, err
		}

		films = append(films, film)
	}
	return films, nil
}

func FetchFilm(id int) (*Film, []*CastItem, []*Playlist, []*DiscussionItem, error) {
	film := new(Film)

	err := db.QueryRow(
		"SELECT f.FilmID, f.FullName, f.AlternativeName, f.Poster, f.`Description`, f.Duration, f.ProductionYear, p.FullName, p.PersonID FROM (SELECT * FROM Film WHERE FilmID = ?) AS f LEFT JOIN Person AS p on f.PersonID = p.PersonID",
		id).Scan(&film.ID, &film.Name, &film.AltName, &film.Poster, &film.Description, &film.Duration, &film.Year, &film.Director, &film.DirectorID)
	if err != nil {
		log.Println("Error fetching person")
		return nil, nil, nil, nil, err
	}
	err = db.QueryRow(
		"SELECT COUNT(*) FROM View WHERE FilmID = ? AND FilmScore = TRUE").Scan(&film.LikeAmount)
	if err != nil {
		log.Println("Error fetching person")
		return nil, nil, nil, nil, err
	}

	err = db.QueryRow(
		"SELECT COUNT(*) FROM View WHERE FilmID = ? AND FilmScore = FALSE").Scan(&film.DislikeAmount)
	if err != nil {
		log.Println("Error fetching person")
		return nil, nil, nil, nil, err
	}

	var cast []*CastItem

	results, err := db.Query(
		"SELECT r.PersonID, p.FullName, r.CharacterName FROM (SELECT PersonID, CharacterName FROM Role WHERE FilmID = ?) AS r LEFT JOIN Person AS p ON r.PersonID = p.PersonID",
		id)
	if err != nil {
		log.Println("Error fetching films")
		return nil, nil, nil, nil, err
	}

	defer results.Close()
	for results.Next() {
		castItem := new(CastItem)

		err = results.Scan(&castItem.ID, castItem.Name, castItem.Character)
		if err != nil {
			log.Println("Error fetching films")
			return nil, nil, nil, nil, err
		}

		cast = append(cast, castItem)
	}

	discussion, err := FetchDiscussionsForFilm(id)
	if err != nil {
		return nil, nil, nil, nil, err
	}

	return film, cast, nil, discussion, nil
}
