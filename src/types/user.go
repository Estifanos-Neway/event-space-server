package types

type User struct {
	Id           string `json:"id"`
	Email        string `json:"email"`
	Name         string `json:"name"`
	Description  string `json:"description"`
	AvatarUrl    string `json:"avatarUrl"`
	PasswordHash string `json:"passwordHash"`
	MemberSince  string `json:"memberSince"`
}
