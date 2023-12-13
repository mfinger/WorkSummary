package WorkReport

import "time"

type Shift struct {
	ShiftID    int64
	EmployeeID int64
	StartTime  time.Time
	EndTime    time.Time
}
