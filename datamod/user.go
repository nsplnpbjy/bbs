package datamod

type User struct {
	Id          string
	Username    string
	Password    string
	Regist_time string
	Ideas_id    []string
	Comments_id []string
}

func CheckBlank(user User) bool {
	if user.Id == "" {
		return false
	}
	if user.Username == "" {
		return false
	}
	if user.Password == "" {
		return false
	}
	if user.Regist_time == "" {
		return false
	}
	return true
}
