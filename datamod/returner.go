package datamod

type returner struct {
	Msg          string
	Service_code int
	Data         interface{}
}

func (u *User) SuccessReturner() *returner {
	r := returner{"ok", 0, u}
	return &r
}

func (u *User) FailReturner() *returner {
	r := returner{"fail", 2, nil}
	return &r
}

func (i *Ideas) SuccessReturner() *returner {
	r := returner{"ok", 0, i}
	return &r
}

func (i *Ideas) FailReturner() *returner {
	r := returner{"fail", 2, nil}
	return &r
}

func (c *Comments) SuccessReturner() *returner {
	r := returner{"ok", 0, c}
	return &r
}

func (c *Comments) FailReturner() *returner {
	r := returner{"fail", 2, nil}
	return &r
}
