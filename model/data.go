package model

// Data : transfer
type Data struct {
	ID        string   `json:"id"`
	IP        string   `json:"ip"`
	User      string   `json:"user"`
	From      string   `json:"from"`
	Type      string   `json:"type"`
	Group     string   `json:"group"`
	Content   string   `json:"content"`
	UserList  []string `json:"user_list"`
	Ping      string   `json:"ping"`
	TimeStamp string   `json:"time_stamp"`
}
