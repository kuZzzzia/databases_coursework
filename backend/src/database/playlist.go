package database

import (
	"context"
	"database/sql"
	"log"
)

type Playlist struct {
	ID            int
	Title         string `binding:"required"`
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

func AddPlaylist(playlist *Playlist, userID int) error {
	var err error

	ctx := context.Background()
	tx, err := db.BeginTx(ctx, &sql.TxOptions{
		Isolation: sql.LevelSerializable,
	})
	if err != nil {
		log.Fatal(err)
	}
	var playlistID int
	err = tx.QueryRow(
		"SELECT `AUTO_INCREMENT` FROM information_schema.TABLES WHERE TABLE_SCHEMA = 'Film_Rec_System' AND TABLE_NAME = 'Playlist'").
		Scan(&playlistID)
	if err != nil {
		errRollback := tx.Rollback()
		if errRollback != nil {
			return errRollback
		}
		return err
	}
	if len(playlist.Description.String) == 0 {
		_, err = tx.ExecContext(ctx,
			"INSERT INTO Playlist(PlaylistTitle, UserID) VALUES (?, ?)",
			playlist.Title, userID)
	} else {
		_, err = tx.ExecContext(ctx,
			"INSERT INTO Playlist(PlaylistTitle, `Description`, UserID) VALUES (?, ?, ?)",
			playlist.Title, playlist.Description.String, userID)
	}
	if err != nil {
		errRollback := tx.Rollback()
		if errRollback != nil {
			return errRollback
		}
		return err
	}
	for _, film := range playlist.Films {
		_, err = tx.ExecContext(ctx, "INSERT INTO Playlist_Film_INT(FilmID, PlaylistID) VALUES (?, ?)",
			film.ID, playlistID)
		if err != nil {
			errRollback := tx.Rollback()
			if errRollback != nil {
				return errRollback
			}
			return err
		}
	}
	err = tx.Commit()
	return err
}
