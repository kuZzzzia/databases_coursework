package database

type Rate struct {
	Like bool
}

func AddRateToFilm(user *User, likeStatus bool, filmID int) error {
	insert, err := db.Query("INSERT INTO View(UserID, FilmID, FilmScore) VALUES(?, ?, ?) ON DUPLICATE KEY UPDATE FilmScore = ?",
		user.ID, filmID, likeStatus, likeStatus)
	defer insert.Close()
	if err != nil {
		return err
	}
	return err
}
