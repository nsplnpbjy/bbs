package datamod

type Idea struct {
	Id           string
	Post_time    int64
	Post_user_id string
	Comments_id  []string
	Text         string
}
