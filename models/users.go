package models

type Staff struct {
	Id         string `json:"id"`
	Name       string `json:"first_name"`
	Email      string `json:"email"`
	Phone      string `json:"phone"`
	Password   string `json:"password"`
	Lastname   string `json:"last_name"`
	BranchName string `json:"branch_name"`
	Blocked    bool
	BlockedAt  string
	CreatedAt  string
}
