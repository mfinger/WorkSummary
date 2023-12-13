package WorkReport

type WeeklyReport struct {
	EmployeeID    int64
	StartOfWeek   string
	RegularHours  float32
	OvertimeHours float32
	InvalidShifts []int64 `json:",omitempty""`
}
