package database

import "log"

type Rate struct {
	Like bool
}

func AddRateToFilm(userID int, likeStatus bool, filmID int) error {
	insert, err := db.Query("INSERT INTO View(UserID, FilmID, FilmScore) VALUES(?, ?, ?) ON DUPLICATE KEY UPDATE FilmScore = ?",
		userID, filmID, likeStatus, likeStatus)
	defer insert.Close()
	if err != nil {
		return err
	}
	return err
}

func AddRateToPlaylist(userID int, likeStatus bool, playlistID int) error {
	insert, err := db.Query("INSERT INTO PlaylistScore(UserID, PlaylistID, Score) VALUES(?, ?, ?) ON DUPLICATE KEY UPDATE Score = ?",
		userID, playlistID, likeStatus, likeStatus)
	defer insert.Close()
	if err != nil {
		return err
	}
	return err
}

func GetFilmRateByUser(userID, filmID int) (int, error) {
	res := 0
	err := db.QueryRow(
		"SELECT FilmScore FROM View WHERE UserID = ? AND FilmID = ?",
		userID, filmID).
		Scan(&res)
	if err != nil {
		log.Println("Error fetching person" + err.Error())
		return 0, err
	}
	if res == 0 {
		return -1, nil
	}
	return res, nil
}

func GetPlaylistRateByUser(userID, playlistID int) (int, error) {
	res := 0
	err := db.QueryRow(
		"SELECT Score FROM PlaylistScore WHERE UserID = ? AND PlaylistID = ?",
		userID, playlistID).
		Scan(&res)
	if err != nil {
		log.Println("Error fetching person" + err.Error())
		return 0, err
	}
	if res == 0 {
		return -1, nil
	}
	return res, nil
}
