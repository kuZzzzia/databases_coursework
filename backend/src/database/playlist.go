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

const (
	playlistsForFilm      = "SELECT p.PlaylistID, p.PlaylistTitle, getPlaylistRating(p.PlaylistID) AS rate FROM Playlists_For_Film as p WHERE p.FilmID = ? ORDER BY rate DESC"
	PlaylistsForProfile   = "SELECT PlaylistID, PlaylistTitle, getPlaylistRating(PlaylistID) FROM Playlist WHERE UserID = ? ORDER BY PlaylistID DESC"
	playlistLikeAmount    = "SELECT COUNT(*) FROM PlaylistRating WHERE PlaylistID = ? AND Rating = TRUE"
	playlistDislikeAmount = "SELECT COUNT(*) FROM PlaylistRating WHERE PlaylistID = ? AND Rating = FALSE"
)

func FetchPlaylists(query string, id int) ([]*Playlist, error) {
	var playlists []*Playlist

	results, err := db.Query(
		query,
		id)
	if err != nil {
		log.Println("Error fetching playlists: " + err.Error())
		return nil, err
	}

	defer results.Close()
	for results.Next() {
		playlist := new(Playlist)

		err = results.Scan(&playlist.ID, &playlist.Title, &playlist.Rating)
		if err != nil {
			log.Println("Error fetching playlists: " + err.Error())
			return nil, err
		}

		playlists = append(playlists, playlist)
	}

	return playlists, nil
}

func FetchPlaylist(id int) (*Playlist, error) {
	playlist := new(Playlist)

	err := db.QueryRow(
		"SELECT PlaylistID, PlaylistTitle, `Description`, Username FROM Playlist_With_Username WHERE PlaylistID = ?",
		id).
		Scan(&playlist.ID, &playlist.Title, &playlist.Description, &playlist.UserName)
	if err != nil {
		log.Println("Error fetching playlist: " + err.Error())
		return nil, err
	}
	err = getRating(playlistLikeAmount, id, &playlist.LikeAmount)
	if err != nil {
		return nil, err
	}

	err = getRating(playlistDislikeAmount, id, &playlist.DislikeAmount)
	if err != nil {
		return nil, err
	}

	results, err := db.Query(
		"SELECT inter.FilmID, f.FullName, f.ProductionYear, getFilmRating(f.FilmID) FROM (SELECT FilmID FROM Playlist_Film_INT WHERE PlaylistID = ?) AS inter LEFT JOIN Film AS f ON inter.FilmID = f.FilmID",
		id)
	if err != nil {
		log.Println("Error fetching playlist's films: " + err.Error())
		return nil, err
	}

	defer results.Close()
	for results.Next() {
		film := new(Film)

		err = results.Scan(&film.ID, &film.Name, &film.Year, &film.Rating)
		if err != nil {
			log.Println("Error fetching playlist's films: " + err.Error())
			return nil, err
		}

		playlist.Films = append(playlist.Films, film)
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
		log.Println("Error adding playlist: " + err.Error())
		return err
	}
	var playlistID int
	err = tx.QueryRow(
		"SELECT `AUTO_INCREMENT` FROM information_schema.TABLES WHERE TABLE_SCHEMA = 'Film_Rec_System' AND TABLE_NAME = 'Playlist'").
		Scan(&playlistID)
	if err != nil {
		log.Println("Error adding playlist: " + err.Error())
		errRollback := tx.Rollback()
		if errRollback != nil {
			log.Println("Error adding playlist: " + errRollback.Error())
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
		log.Println("Error adding playlist: " + err.Error())
		errRollback := tx.Rollback()
		if errRollback != nil {
			log.Println("Error adding playlist: " + errRollback.Error())
			return errRollback
		}
		return err
	}
	for _, film := range playlist.Films {
		_, err = tx.ExecContext(ctx, "INSERT INTO Playlist_Film_INT(FilmID, PlaylistID) VALUES (?, ?)",
			film.ID, playlistID)
		if err != nil {
			log.Println("Error adding playlist: " + err.Error())
			errRollback := tx.Rollback()
			if errRollback != nil {
				log.Println("Error adding playlist: " + errRollback.Error())
				return errRollback
			}
			return err
		}
	}
	err = tx.Commit()
	return err
}

func DeletePlaylist(playlistId, userId int) error {
	_, err := db.Exec("DELETE FROM Playlist WHERE PlaylistID = ? AND UserID = ?",
		playlistId, userId)
	if err != nil {
		log.Println("Error deleting playlist: " + err.Error())
		return err
	}
	return err
}
