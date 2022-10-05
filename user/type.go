package user

type User struct {
	Id   int64  `json:"id,string" db:"id"`
	Name string `json:"name" db:"name"`
}

type createRequest struct {
	Name string `json:"name"`
}

type editRequest struct {
	Name string `json:"name"`
}
