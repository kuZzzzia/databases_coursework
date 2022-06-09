package database

import (
	"database/sql"
	"log"
)

type Film struct {
	ID            int `binding:"required"`
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

func FetchFilms(search *Search) ([]*Film, error) {
	var (
		films   []*Film
		results *sql.Rows
		args    []any
		err     error
	)
	query := "SELECT FilmID, FullName, AlternativeName, Poster, Duration, ProductionYear, getFilmRating(FilmID) FROM Film AS f"
	if len(search.Pattern) != 0 {
		query += " WHERE FullName = ? OR AlternativeName = ?"
		args = append(args, search.Pattern, search.Pattern)
		if len(search.Genre) != 0 {
			query += " AND EXISTS(SELECT * FROM Film_Genres AS g WHERE GenreName = ? AND g.FilmID = f.FilmID)"
			args = append(args, search.Genre)
		}
		if len(search.Country) != 0 {
			query += " AND EXISTS(SELECT * FROM Film_Countries AS c WHERE CountryName = ? AND c.FilmID = f.FilmID)"
			args = append(args, search.Country)
		}
	} else if len(search.Genre) != 0 {
		query += " WHERE EXISTS(SELECT * FROM Film_Genres AS g WHERE GenreName = ? AND g.FilmID = f.FilmID"
		args = append(args, search.Genre)
		if len(search.Country) != 0 {
			query += " AND EXISTS(SELECT * FROM Film_Countries AS c WHERE CountryName = ? AND c.FilmID = f.FilmID)"
			args = append(args, search.Country)
		}
	} else if len(search.Country) != 0 {
		query += " WHERE EXISTS(SELECT * FROM Film_Countries AS c WHERE CountryName = ? AND c.FilmID = f.FilmID)"
		args = append(args, search.Country)
	}
	if len(args) != 0 {
		results, err = db.Query(query, args...)
	} else {
		results, err = db.Query(query)
	}
	if err != nil {
		log.Println("Error fetching films")
		return nil, err
	}

	defer results.Close()

	for results.Next() {
		film := new(Film)

		err = results.Scan(&film.ID, &film.Name, &film.AltName, &film.Poster, &film.Duration, &film.Year, &film.Rating)
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

	playlists, _ := FetchPlaylists(playlistsForFilm, id)

	discussion, err := FetchDiscussionsForFilm(id)
	if err != nil {
		return nil, nil, nil, nil, err
	}

	return film, cast, playlists, discussion, nil
}
