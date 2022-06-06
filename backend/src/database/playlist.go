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
	Films         []*Film
}

func FetchPlaylistsForFilm(id int) ([]*Playlist, error) {
	var playlists []*Playlist

	//TODO: write the select query
	results, err := db.Query(
		"",
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
