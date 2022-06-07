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

func FetchFilm(id int) (*Film, []*CastItem, []*Playlist, []*Message, error) {
	film := new(Film)

	err := db.QueryRow(
		"SELECT * FROM Film_With_Director WHERE FilmID = ?",
		id).
		Scan(&film.ID, &film.Name, &film.AltName,
			&film.Poster, &film.Description, &film.Duration,
			&film.Year, &film.DirectorID, &film.Director)
	if err != nil {
		log.Println("Error fetching person")
		return nil, nil, nil, nil, err
	}
	err = db.QueryRow(
		"SELECT COUNT(*) FROM View WHERE FilmID = ? AND FilmScore = TRUE", id).Scan(&film.LikeAmount)
	if err != nil {
		log.Println("Error fetching person" + err.Error())
		return nil, nil, nil, nil, err
	}

	err = db.QueryRow(
		"SELECT COUNT(*) FROM View WHERE FilmID = ? AND FilmScore = FALSE", id).Scan(&film.DislikeAmount)
	if err != nil {
		log.Println("Error fetching person")
		return nil, nil, nil, nil, err
	}

	var cast []*CastItem
	results, err := db.Query(
		"SELECT PersonID, FullName, CharacterName FROM Film_Cast WHERE FilmID = ?",
		id)
	if err != nil {
		log.Println("Error fetching films" + err.Error())
		return nil, nil, nil, nil, err
	}
	defer results.Close()
	for results.Next() {
		castItem := new(CastItem)

		err = results.Scan(&castItem.ID, &castItem.Name, &castItem.Character)
		if err != nil {
			log.Println("Error fetching films" + err.Error())
			return nil, nil, nil, nil, err
		}

		cast = append(cast, castItem)
	}

	playlists, err := FetchPlaylistsForFilm(id)
	if err != nil {
		return nil, nil, nil, nil, err
	}

	discussion, err := FetchDiscussionsForFilm(id)
	if err != nil {
		return nil, nil, nil, nil, err
	}

	return film, cast, playlists, discussion, nil
}
