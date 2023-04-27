package types

type Ticket struct {
	Id        string `json:"id"`
	UserId    string `json:"user_id"`
	EventId   string `json:"event_id"`
	CreatedAt string `json:"created_at"`
	IsValid   bool   `json:"is_valid"`
}
