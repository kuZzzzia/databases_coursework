package database

import (
	"database/sql"
	"log"
)

type Playlist struct {
	ID            int
	Title         string
	Description   sql.NullString
	Rating        int
	LikeAmount    int
	DislikeAmount int
	UserName      sql.NullString
	Films         []*Film
}

func FetchPlaylistsForFilm(id int) ([]*Playlist, error) {
	var playlists []*Playlist

	results, err := db.Query(
		"SELECT p.PlaylistID, p.PlaylistTitle, getPlaylistRating(p.PlaylistID) AS rate FROM (SELECT PlaylistID AS id_int FROM Playlist_Film_INT WHERE FilmID = ?) as i LEFT JOIN Playlist AS p ON id_int = p.PlaylistID ORDER BY rate DESC",
		id)
	if err != nil {
		log.Println("Error fetching roles")
		return nil, err
	}

	defer results.Close()
	for results.Next() {
		playlist := new(Playlist)

		err = results.Scan(&playlist.ID, &playlist.Title, &playlist.Rating)
		if err != nil {
			log.Println("Error fetching roles")
			return nil, err
		}

		playlists = append(playlists, playlist)
	}

	return playlists, nil
}

func FetchPlaylist(id int) (*Playlist, error) {
	playlist := new(Playlist)

	err := db.QueryRow(
		"SELECT PlaylistID, PlaylistTitle, `Description` FROM Playlist WHERE PlaylistID = ?",
		id).
		Scan(&playlist.ID, &playlist.Title, &playlist.Description)
	if err != nil {
		log.Println("Error fetching person")
		return nil, err
	}
	err = db.QueryRow(
		"SELECT COUNT(*) FROM PlaylistScore WHERE PlaylistID = ? AND Score = TRUE", id).Scan(&playlist.LikeAmount)
	if err != nil {
		log.Println("Error fetching person" + err.Error())
		return nil, err
	}

	err = db.QueryRow(
		"SELECT COUNT(*) FROM PlaylistScore WHERE PlaylistID = ? AND Score = FALSE", id).Scan(&playlist.DislikeAmount)
	if err != nil {
		log.Println("Error fetching person")
		return nil, err
	}

	return playlist, nil
}
