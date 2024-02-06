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
	UpdatedAt  string
}
type GetStaffs struct {
	Limit int
	Page  int
}

type GetEntryResponse struct {
	Entry []Entry `json:"entry"`
	Count int     `json:"count"`
}

type Entry struct {
	Id           string `json:"id"`
	StaffId      string `json:"staff_id"`
	ActivityType string `json:"activity_type"`
	Date         string `json:"date"`
	City         string `json:"city"`
	UpdatedAt    string `json:"updated_at"`
}

type StaffLeaveList struct {
	Count int `json:"count"`
	Leave []LeaveRequest `json:"leave"`
}
type LeaveRequest struct {
	Id           string `json:"id"`
	StaffId      string `json:"staff_id"`
	Reason       string `json:"reason"`
	Read         bool   `json:"read"`
	CreatedAt    string `json:"created_at"`
	LeaveDate    string `json:"leave_date"`
	Approved     bool   `json:"approved"`
	ApprovedAt   string `json:"approved_at"`
	UpdatedAt    string `json:"updated_at"`
	ReadTime     string `json:"read_time"`
	ApprovedTime string `json:"approved_time"`
}
type GetStaffEntries struct {
	Id      string
	StaffID string
	Date    string
	Limit   int
	Page    int
	From    string
	To      string
}
type GetStaffLeavesRequest struct {
	StaffID string
	Limit   int
	Page    int
	From    string
	To      string
	Id      string
}
