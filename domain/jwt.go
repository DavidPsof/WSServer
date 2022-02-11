package domain

type JwtInfo struct {
	UserID int // connecting user id
}

func (j JwtInfo) Valid() error {
	return nil
}
