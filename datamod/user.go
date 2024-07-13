package datamod

type User struct {
	Id          string
	Username    string
	Password    string
	Regist_time int64
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
	return true
}

func (u *User) DePassword() *User {
	u.Password = ""
	return u
}

func (u *User) InfoForOtherUsers() *User {
	u.Password = ""
	u.Ideas_id = nil
	u.Comments_id = nil
	return u
}
