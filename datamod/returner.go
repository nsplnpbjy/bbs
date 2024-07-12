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

func (i *Idea) SuccessReturner(token string) *Returner {
	r := Returner{"ok", 0, token, i}
	return &r
}

func (i *Idea) FailReturner() *Returner {
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

func (i *Idea) IdeaInsertedSuccess(token string) *Returner {
	r := Returner{"idea insert success", 0, token, i}
	return &r
}

func (i *Idea) IdeaInsertedFailed() *Returner {
	r := Returner{"idea insert failed", 2, "", nil}
	return &r
}

func (i *Idea) IdeaDeleteFailed() *Returner {
	r := Returner{"idea delete failed", 2, "", nil}
	return &r
}

func (i *Idea) IdeaDeleteSuccess(token string) *Returner {
	r := Returner{"idea delete success", 0, token, i}
	return &r
}

type Ideas []Idea

func (is *Ideas) IdeasSelectSuccess(token string) *Returner {
	r := Returner{"ideas select success", 0, token, is}
	return &r
}

func (is *Ideas) IdeasSelectFailed() *Returner {
	r := Returner{"ideas select failed", 2, "", nil}
	return &r
}
