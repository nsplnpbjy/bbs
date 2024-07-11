package datamod

type Returner struct {
	Msg          string
	Service_code int
	Token        string
	Data         interface{}
}

func (u *User) SuccessReturner(token string) *Returner {
	r := Returner{"ok", 0, token, u}
	return &r
}

func (u *User) FailReturner() *Returner {
	r := Returner{"fail", 2, "", nil}
	return &r
}

func (u *User) UsernameAlreadyExist() *Returner {
	r := Returner{"username already exist", 2, "", nil}
	return &r
}

func (i *Ideas) SuccessReturner(token string) *Returner {
	r := Returner{"ok", 0, token, i}
	return &r
}

func (i *Ideas) FailReturner() *Returner {
	r := Returner{"fail", 2, "", nil}
	return &r
}

func (c *Comments) SuccessReturner(token string) *Returner {
	r := Returner{"ok", 0, token, c}
	return &r
}

func (c *Comments) FailReturner() *Returner {
	r := Returner{"fail", 2, "", nil}
	return &r
}

func BlankTokenReturner() *Returner {
	r := Returner{"blank token", 2, "", nil}
	return &r
}

func InvalidTokenReturner() *Returner {
	r := Returner{"invalid token", 2, "", nil}
	return &r
}

func TimeOutTokenReturner() *Returner {
	r := Returner{"timeout token", 2, "", nil}
	return &r
}
