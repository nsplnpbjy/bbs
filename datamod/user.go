package datamod

type User struct {
	Id          string
	Username    string
	Password    string
	Regist_time string
	Ideas_id    []string
	Comments_id []string
}

func (u *User) CheckBlank() bool {
	if u.Id == "" {
		return false
	}
	if u.Username == "" {
		return false
	}
	if u.Password == "" {
		return false
	}
	if u.Regist_time == "" {
		return false
	}
	return true
}
