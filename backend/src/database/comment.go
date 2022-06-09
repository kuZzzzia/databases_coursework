package database

import (
	"database/sql"
	"log"
)

type Comment struct {
	ID       int
	Review   string `binding:"required,min=5,max=63"`
	UserName sql.NullString
	FilmID   int
}

func FetchCommentsForFilm(id int) ([]*Comment, error) {
	var discussion []*Comment

	results, err := db.Query(
		"SELECT DiscussionID, Review, Username FROM Film_Discussion_With_Users WHERE FilmID = ? ORDER BY DiscussionID DESC",
		id)
	if err != nil {
		log.Println("Error fetching comments: " + err.Error())
		return nil, err
	}

	defer results.Close()
	for results.Next() {
		comment := new(Comment)

		err = results.Scan(&comment.ID, &comment.Review, &comment.UserName)
		if err != nil {
			log.Println("Error fetching comments: " + err.Error())
			return nil, err
		}

		discussion = append(discussion, comment)
	}
	return discussion, nil
}

func AddComment(user *User, comment *Comment, filmID int) error {
	comment.UserName = sql.NullString{
		String: user.Username,
		Valid:  true,
	}
	_, err := db.Exec("INSERT INTO Discussion(`Date`, Review, UserID, FilmID) VALUES (NOW(), ?, ?, ?)",
		comment.Review, user.ID, filmID)
	if err != nil {
		log.Println("Error adding comment to film: " + err.Error())
		return err
	}
	return err
}
