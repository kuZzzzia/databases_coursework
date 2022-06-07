package database

import (
	"database/sql"
	"log"
)

type Message struct {
	ID       int
	Date     string
	Review   string
	UserID   sql.NullInt16
	UserName sql.NullString
	FilmID   int
}

func FetchDiscussionsForFilm(id int) ([]*Message, error) {
	var discussion []*Message

	results, err := db.Query(
		"SELECT DiscussionID, `Date`, Review, UserID, Username FROM Film_Discussion_With_Users WHERE FilmID = ? ORDER BY `Date`",
		id)
	if err != nil {
		log.Println("Error fetching roles")
		return nil, err
	}

	defer results.Close()
	for results.Next() {
		discussionItem := new(Message)

		err = results.Scan(&discussionItem.ID, &discussionItem.Date, &discussionItem.Review, &discussionItem.UserID, &discussionItem.UserName)
		if err != nil {
			log.Println("Error fetching roles")
			return nil, err
		}

		discussion = append(discussion, discussionItem)
	}
	return discussion, nil
}

func AddMessage(user *User, message *Message, filmID int) error {
	message.UserID = sql.NullInt16{
		Int16: int16(user.ID),
		Valid: true,
	}
	message.UserName = sql.NullString{
		String: user.Username,
		Valid:  true,
	}
	insert, err := db.Query("INSERT INTO Discussion(`Date`, Review, UserID, FilmID) VALUES (NOW(), ?, ?, ?)",
		message.Review, message.UserID, filmID)
	defer insert.Close()
	if err != nil {
		return err
	}
	return err
}
