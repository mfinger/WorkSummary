package WorkReport

import (
	"encoding/json"
	"io"
	"os"
)

type WorkReport struct {
	employees map[int64]*Employee
	shifts    []Shift
}

func NewWorkReport() *WorkReport {
	workSummary := new(WorkReport)

	workSummary.employees = make(map[int64]*Employee)

	return workSummary
}

func (ws WorkReport) processShifts() {
	for _, shift := range ws.shifts {
		var emp *Employee
		var found bool
		if emp, found = ws.employees[shift.EmployeeID]; !found {
			emp = NewEmployee(shift.EmployeeID)
			ws.employees[emp.ID] = emp
		}
		emp.AddShift(shift)
	}
}

// Load the list of lottery entries from a file
func (ws *WorkReport) Load(filename string) error {

	if f, err := os.Open(filename); err != nil {
		return err
	} else {
		defer f.Close()
		if err = ws.load(f); err != nil {
			return err
		} else {
			ws.processShifts()
			ws.shifts = nil
		}
	}
	return nil
}

// Internal load method, separated for better testing.
func (ws *WorkReport) load(reader io.Reader) error {

	if buff, err := io.ReadAll(reader); err != nil {
		return err
	} else {
		if err := json.Unmarshal(buff, &ws.shifts); err != nil {
			return err
		}
	}
	return nil
}

func (ws WorkReport) GenerateReports(filename string) error {
	reports := []WeeklyReport{}
	for _, emp := range ws.employees {
		reports = append(reports, emp.GenerateReports()...)
	}

	return ws.SaveReports(filename, reports)
}

func (ws WorkReport) SaveReports(filename string, reports []WeeklyReport) error {
	f, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	defer f.Close()
	if err != nil {
		return err
	}

	buff, err := json.MarshalIndent(reports, "", "  ")
	if err != nil {
		return err
	}

	if _, err := f.Write(buff); err != nil {
		return err
	}
	return nil

}
