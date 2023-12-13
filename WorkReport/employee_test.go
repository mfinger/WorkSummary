package WorkReport

import (
	"log"
	"testing"
	"time"
)

func TestDoShiftsConflict2OverlapsEndOf1(t *testing.T) {
	shift1 := Shift{
		ShiftID:    1,
		EmployeeID: 1,
		StartTime:  time.Date(2009, 11, 17, 10, 0, 0, 0, time.UTC),
		EndTime:    time.Date(2009, 11, 17, 12, 0, 0, 0, time.UTC),
	}

	shift2 := Shift{
		ShiftID:    1,
		EmployeeID: 1,
		StartTime:  time.Date(2009, 11, 17, 11, 0, 0, 0, time.UTC),
		EndTime:    time.Date(2009, 11, 17, 13, 0, 0, 0, time.UTC),
	}

	result := DoShiftsConflict(shift1, shift2)

	if !result {
		t.Fatalf("TestDoShiftsConflict2Overlaps1: expected true, but got false")
	}

	result = DoShiftsConflict(shift2, shift1)

	if !result {
		t.Fatalf("TestDoShiftsConflict2Overlaps1 (inverse): expected true, but got false")
	}

}

func TestDoShiftsConflict1Overlaps2(t *testing.T) {
	shift1 := Shift{
		ShiftID:    1,
		EmployeeID: 1,
		StartTime:  time.Date(2009, 11, 17, 11, 0, 0, 0, time.UTC),
		EndTime:    time.Date(2009, 11, 17, 13, 0, 0, 0, time.UTC),
	}

	shift2 := Shift{
		ShiftID:    1,
		EmployeeID: 1,
		StartTime:  time.Date(2009, 11, 17, 10, 0, 0, 0, time.UTC),
		EndTime:    time.Date(2009, 11, 17, 12, 0, 0, 0, time.UTC),
	}

	result := DoShiftsConflict(shift1, shift2)

	if !result {
		t.Fatalf("TestDoShiftsConflict1Overlaps2: expected true, but got false")
	}

	result = DoShiftsConflict(shift2, shift1)

	if !result {
		t.Fatalf("TestDoShiftsConflict1Overlaps2 (inverse): expected true, but got false")
	}

}

func TestDoShiftsConflict2Inside1(t *testing.T) {
	shift1 := Shift{
		ShiftID:    1,
		EmployeeID: 1,
		StartTime:  time.Date(2009, 11, 17, 11, 0, 0, 0, time.UTC),
		EndTime:    time.Date(2009, 11, 17, 15, 0, 0, 0, time.UTC),
	}

	shift2 := Shift{
		ShiftID:    1,
		EmployeeID: 1,
		StartTime:  time.Date(2009, 11, 17, 12, 0, 0, 0, time.UTC),
		EndTime:    time.Date(2009, 11, 17, 13, 0, 0, 0, time.UTC),
	}

	result := DoShiftsConflict(shift1, shift2)

	if !result {
		t.Fatalf("TestDoShiftsConflict2Inside1: expected true, but got false")
	}

	result = DoShiftsConflict(shift2, shift1)

	if !result {
		t.Fatalf("TestDoShiftsConflict2Inside1 (inverse): expected true, but got false")
	}

}

func TestDoShiftsConflict1Before2(t *testing.T) {
	shift1 := Shift{
		ShiftID:    1,
		EmployeeID: 1,
		StartTime:  time.Date(2009, 11, 17, 11, 0, 0, 0, time.UTC),
		EndTime:    time.Date(2009, 11, 17, 12, 0, 0, 0, time.UTC),
	}

	shift2 := Shift{
		ShiftID:    1,
		EmployeeID: 1,
		StartTime:  time.Date(2009, 11, 17, 12, 1, 0, 0, time.UTC),
		EndTime:    time.Date(2009, 11, 17, 13, 0, 0, 0, time.UTC),
	}

	result := DoShiftsConflict(shift1, shift2)

	if result {
		t.Fatalf("TestDoShiftsConflict1Before2: expected false, but got true")
	}

	result = DoShiftsConflict(shift2, shift1)

	if result {
		t.Fatalf("TestDoShiftsConflict1Before2 (inverse): expected false, but got true")
	}

}

func TestDoShiftsConflict1after2(t *testing.T) {
	shift1 := Shift{
		ShiftID:    1,
		EmployeeID: 1,
		StartTime:  time.Date(2009, 11, 17, 11, 0, 0, 0, time.UTC),
		EndTime:    time.Date(2009, 11, 17, 12, 0, 0, 0, time.UTC),
	}

	shift2 := Shift{
		ShiftID:    1,
		EmployeeID: 1,
		StartTime:  time.Date(2009, 11, 17, 9, 0, 0, 0, time.UTC),
		EndTime:    time.Date(2009, 11, 17, 10, 0, 0, 0, time.UTC),
	}

	result := DoShiftsConflict(shift1, shift2)

	if result {
		t.Fatalf("TestDoShiftsConflict1after2: expected false, but got true")
	}

	result = DoShiftsConflict(shift2, shift1)

	if result {
		t.Fatalf("TestDoShiftsConflict1after2 (inverse): expected false, but got true")
	}

}

func TestDoShiftsConflict1Before2Boundary(t *testing.T) {
	shift1 := Shift{
		ShiftID:    1,
		EmployeeID: 1,
		StartTime:  time.Date(2009, 11, 17, 11, 0, 0, 0, time.UTC),
		EndTime:    time.Date(2009, 11, 17, 12, 0, 0, 0, time.UTC),
	}

	shift2 := Shift{
		ShiftID:    1,
		EmployeeID: 1,
		StartTime:  time.Date(2009, 11, 17, 12, 0, 0, 0, time.UTC),
		EndTime:    time.Date(2009, 11, 17, 13, 0, 0, 0, time.UTC),
	}

	result := DoShiftsConflict(shift1, shift2)

	if result {
		t.Fatalf("TestDoShiftsConflict1Before2Boundary: expected false, but got true")
	}

	result = DoShiftsConflict(shift2, shift1)

	if result {
		t.Fatalf("TestDoShiftsConflict1Before2Boundary (inverse): expected false, but got true")
	}

}

func TestDoShiftsConflict1after2Boundary(t *testing.T) {
	shift1 := Shift{
		ShiftID:    1,
		EmployeeID: 1,
		StartTime:  time.Date(2009, 11, 17, 11, 0, 0, 0, time.UTC),
		EndTime:    time.Date(2009, 11, 17, 12, 0, 0, 0, time.UTC),
	}

	shift2 := Shift{
		ShiftID:    1,
		EmployeeID: 1,
		StartTime:  time.Date(2009, 11, 17, 9, 0, 0, 0, time.UTC),
		EndTime:    time.Date(2009, 11, 17, 11, 0, 0, 0, time.UTC),
	}

	result := DoShiftsConflict(shift1, shift2)

	if result {
		t.Fatalf("TestDoShiftsConflict1after2Boundary: expected false, but got true")
	}

	result = DoShiftsConflict(shift2, shift1)

	if result {
		t.Fatalf("TestDoShiftsConflict1after2Boundary (inverse): expected false, but got true")
	}

}

func TestDoShiftsConflict1And2Same(t *testing.T) {
	shift1 := Shift{
		ShiftID:    1,
		EmployeeID: 1,
		StartTime:  time.Date(2009, 11, 17, 11, 0, 0, 0, time.UTC),
		EndTime:    time.Date(2009, 11, 17, 12, 0, 0, 0, time.UTC),
	}

	result := DoShiftsConflict(shift1, shift1)

	if !result {
		t.Fatalf("TestDoShiftsConflict1And2Same: expected true, but got false")
	}
}

func TestNewEmployee(t *testing.T) {
	emp := NewEmployee(1234)

	if emp.ID != 1234 {
		t.Fatalf("TestNewEmployee: Expected Id to be 1234, but was %v\n", emp.ID)
	}

	if emp.Shifts == nil {
		t.Fatalf("TestNewEmployee: Expected Shifts to not be null\n")
	}

	if emp.Invalid == nil {
		t.Fatalf("TestNewEmployee: Expected Invalid to not be null\n")
	}
}

func TestEmployee_AddShift(t *testing.T) {
	emp := NewEmployee(1234)

	emp.AddShift(Shift{
		ShiftID:    1,
		EmployeeID: 1234,
		StartTime:  time.Date(2009, 11, 17, 8, 0, 0, 0, time.UTC),
		EndTime:    time.Date(2009, 11, 17, 9, 0, 0, 0, time.UTC),
	})

	emp.AddShift(Shift{
		ShiftID:    2,
		EmployeeID: 1234,
		StartTime:  time.Date(2009, 11, 17, 10, 0, 0, 0, time.UTC),
		EndTime:    time.Date(2009, 11, 17, 11, 0, 0, 0, time.UTC),
	})

	if len(emp.Shifts) != 2 {
		log.Fatalf("TestEmployee_AddShift expected shift length to be 1, but was %v\n", len(emp.Shifts))
	}

	if len(emp.Invalid) != 0 {
		log.Fatalf("TestEmployee_AddShift expected Invalid length to be 0, but was %v\n", len(emp.Invalid))
	}

	emp.AddShift(Shift{
		ShiftID:    3,
		EmployeeID: 1234,
		StartTime:  time.Date(2009, 11, 17, 10, 15, 0, 0, time.UTC),
		EndTime:    time.Date(2009, 11, 17, 11, 15, 0, 0, time.UTC),
	})

	if len(emp.Shifts) != 3 {
		log.Fatalf("TestEmployee_AddShift expected shift length to be 2, but was %v\n", len(emp.Shifts))
	}

	if len(emp.Invalid) != 2 {
		log.Fatalf("TestEmployee_AddShift expected Invalid length to be 1, but was %v\n", len(emp.Invalid))
	}

}

func TestEmployee_GenerateReports(t *testing.T) {
	emp := NewEmployee(1234)

	emp.AddShift(Shift{
		ShiftID:    1,
		EmployeeID: 1234,
		StartTime:  time.Date(2009, 11, 17, 8, 0, 0, 0, time.UTC),
		EndTime:    time.Date(2009, 11, 17, 9, 0, 0, 0, time.UTC),
	})

	emp.AddShift(Shift{
		ShiftID:    2,
		EmployeeID: 1234,
		StartTime:  time.Date(2009, 11, 17, 10, 0, 0, 0, time.UTC),
		EndTime:    time.Date(2009, 11, 17, 11, 0, 0, 0, time.UTC),
	})

	reports := emp.GenerateReports()

	if len(reports) != 1 {
		log.Fatalf("TestEmployee_GenerateReports expected reports length to be 1, but was %v\n", len(reports))
	}

	if reports[0].StartOfWeek != "2009-11-15" {
		log.Fatalf("TestEmployee_GenerateReports expected report start of week to be 2009-11-15, but was %v\n", reports[0].StartOfWeek)
	}

	if reports[0].RegularHours != 2 {
		log.Fatalf("TestEmployee_GenerateReports expected regular hours to be 2, but was %v\n", reports[0].RegularHours)
	}

	if reports[0].OvertimeHours != 0 {
		log.Fatalf("TestEmployee_GenerateReports expected regular hours to be 0, but was %v\n", reports[0].OvertimeHours)
	}

}

func TestEmployee_GenerateReportsOverTime(t *testing.T) {
	emp := NewEmployee(1234)

	emp.AddShift(Shift{
		ShiftID:    3,
		EmployeeID: 1234,
		StartTime:  time.Date(2009, 11, 15, 8, 0, 0, 0, time.UTC),
		EndTime:    time.Date(2009, 11, 15, 23, 0, 0, 0, time.UTC),
	})

	emp.AddShift(Shift{
		ShiftID:    4,
		EmployeeID: 1234,
		StartTime:  time.Date(2009, 11, 16, 8, 0, 0, 0, time.UTC),
		EndTime:    time.Date(2009, 11, 16, 23, 0, 0, 0, time.UTC),
	})

	emp.AddShift(Shift{
		ShiftID:    5,
		EmployeeID: 1234,
		StartTime:  time.Date(2009, 11, 18, 8, 0, 0, 0, time.UTC),
		EndTime:    time.Date(2009, 11, 18, 23, 0, 0, 0, time.UTC),
	})

	reports := emp.GenerateReports()

	if len(reports) != 1 {
		log.Fatalf("TestEmployee_GenerateReportsOverTime expected reports length to be 1, but was %v\n", len(reports))
	}

	if reports[0].StartOfWeek != "2009-11-15" {
		log.Fatalf("TestEmployee_GenerateReportsOverTime expected report start of week to be 2009-11-15, but was %v\n", reports[0].StartOfWeek)
	}

	if reports[0].RegularHours != 40 {
		log.Fatalf("TestEmployee_GenerateReportsOverTime expected regular hours to be 40, but was %v\n", reports[0].RegularHours)
	}

	if reports[0].OvertimeHours != 5 {
		log.Fatalf("TestEmployee_GenerateReportsOverTime expected regular hours to be 5, but was %v\n", reports[0].OvertimeHours)
	}

}

func TestEmployee_GenerateReportsSplit(t *testing.T) {
	emp := NewEmployee(1234)

	emp.AddShift(Shift{
		ShiftID:    3,
		EmployeeID: 1234,
		StartTime:  time.Date(2009, 11, 15, 1, 0, 0, 0, time.UTC),
		EndTime:    time.Date(2009, 11, 15, 7, 0, 0, 0, time.UTC),
	})

	reports := emp.GenerateReports()

	if len(reports) != 2 {
		log.Fatalf("TestEmployee_GenerateReportsSplit expected reports length to be 2, but was %v\n", len(reports))
	}

	if reports[0].StartOfWeek != "2009-11-08" {
		log.Fatalf("TestEmployee_GenerateReportsSplit expected first report start of week to be 2009-11-08, but was %v\n", reports[0].StartOfWeek)
	}

	if reports[0].RegularHours != 5 {
		log.Fatalf("TestEmployee_GenerateReportsSplit expected first report regular hours to be 40, but was %v\n", reports[0].RegularHours)
	}

	if reports[0].OvertimeHours != 0 {
		log.Fatalf("TestEmployee_GenerateReportsSplit expected first report  regular hours to be 5, but was %v\n", reports[0].OvertimeHours)
	}

	if reports[1].StartOfWeek != "2009-11-15" {
		log.Fatalf("TestEmployee_GenerateReportsSplit expected second report start of week to be 2009-11-15, but was %v\n", reports[0].StartOfWeek)
	}

	if reports[1].RegularHours != 1 {
		log.Fatalf("TestEmployee_GenerateReportsSplit expected first report regular hours to be 1, but was %v\n", reports[0].RegularHours)
	}

	if reports[1].OvertimeHours != 0 {
		log.Fatalf("TestEmployee_GenerateReportsSplit expected first report  regular hours to be 0, but was %v\n", reports[0].OvertimeHours)
	}

}
