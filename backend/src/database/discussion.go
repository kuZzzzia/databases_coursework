package database

import (
	"database/sql"
	"log"
)

type DiscussionItem struct {
	ID       int
	Date     string
	Review   sql.NullString
	UserID   sql.NullInt16
	UserName sql.NullString
}

func FetchDiscussionsForFilm(id int) ([]*DiscussionItem, error) {
	var discussion []*DiscussionItem

	results, err := db.Query(
		"SELECT DiscussionID, `Date`, Review, u.UserID, Username FROM (SELECT DiscussionID, `Date`, Review, UserID FROM Discussion WHERE FilmID = ?) AS d LEFT JOIN User AS u ON d.UserID = u.UserID ORDER BY d.Date",
		id)
	if err != nil {
		log.Println("Error fetching roles")
		return nil, err
	}

	defer results.Close()
	for results.Next() {
		discussionItem := new(DiscussionItem)

		err = results.Scan(&discussionItem.ID, &discussionItem.Date, &discussionItem.Review, &discussionItem.UserID, &discussionItem.UserName)
		if err != nil {
			log.Println("Error fetching roles")
			return nil, err
		}

		discussion = append(discussion, discussionItem)
	}
	return discussion, nil
}
