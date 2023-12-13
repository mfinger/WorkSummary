package WorkReport

import (
	"fmt"
	"time"
)

type Employee struct {
	ID      int64
	Shifts  []Shift
	Invalid map[int64]interface{}
}

// DoShiftsConflict - Determine if two shifts conflict with each other
func DoShiftsConflict(shift1 Shift, shift2 Shift) bool {
	result := true
	if shift1.StartTime.Compare(shift2.EndTime) >= 0 || shift1.EndTime.Compare(shift2.StartTime) <= 0 {
		result = false
	}
	return result
}

// NewEmployee Create a new Employee object
func NewEmployee(id int64) *Employee {
	employee := new(Employee)

	employee.ID = id
	employee.Shifts = []Shift{}
	employee.Invalid = make(map[int64]interface{})

	return employee
}

// AddShift Add a shift to an employees record, checking for conflicts
func (e *Employee) AddShift(newShift Shift) {
	for _, shift := range e.Shifts {
		if DoShiftsConflict(newShift, shift) {
			e.Invalid[shift.ShiftID] = nil
			e.Invalid[newShift.ShiftID] = nil
		}
	}
	e.Shifts = append(e.Shifts, newShift)
}

// GenerateReports Generate the list of weekly reports for this employee
func (e Employee) GenerateReports() []WeeklyReport {
	location, _ := time.LoadLocation("CST6CDT")
	reports := make(map[string]*WeeklyReport)
	shifts := e.Shifts
	for i := 0; i < len(shifts); i++ {
		shift := shifts[i]
		startTimeCentral := shift.StartTime.In(location)
		endTimeCentral := shift.EndTime.In(location)
		hours := float32(endTimeCentral.Sub(startTimeCentral).Seconds()) / 3600

		startWeekForStart := startTimeCentral.AddDate(0, 0, -int(startTimeCentral.Weekday()))
		startWeekTagForStart := fmt.Sprintf("%04d-%02d-%02d", startWeekForStart.Year(), int(startWeekForStart.Month()), startWeekForStart.Day())
		startWeekForEnd := endTimeCentral.AddDate(0, 0, -int(endTimeCentral.Weekday()))
		startWeekTagForEnd := fmt.Sprintf("%04d-%02d-%02d", startWeekForEnd.Year(), int(startWeekForEnd.Month()), startWeekForEnd.Day())

		// Check for splits
		if startWeekTagForStart != startWeekTagForEnd {
			newEndTime := time.Date(startWeekForEnd.Year(), startWeekForEnd.Month(), startWeekForEnd.Day(), 0, 0, 0, 0, location)
			hours = float32(newEndTime.Sub(startTimeCentral).Seconds()) / 3600
			splitShift := Shift{
				EmployeeID: shift.EmployeeID,
				ShiftID:    shift.ShiftID,
				StartTime:  newEndTime,
				EndTime:    shift.EndTime,
			}

			shifts = append(shifts, splitShift)
		}

		report, found := reports[startWeekTagForStart]
		if !found {
			report = &WeeklyReport{EmployeeID: e.ID, StartOfWeek: startWeekTagForStart}
			reports[startWeekTagForStart] = report
		}
		if _, found := e.Invalid[shift.ShiftID]; !found {
			report.RegularHours += hours
			if report.RegularHours > 40 {
				report.OvertimeHours += report.RegularHours - 40
				report.RegularHours = 40
			}
		} else {
			report.InvalidShifts = append(report.InvalidShifts, shift.ShiftID)
		}
	}

	reportSlice := []WeeklyReport{}
	for _, report := range reports {
		reportSlice = append(reportSlice, *report)
	}
	return reportSlice
}
