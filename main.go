package main

import (
	"WorkSummary/WorkReport"
	"fmt"
	"log"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Fprint(os.Stderr, "Usage: WorkReport input.json output.json\n")
		os.Exit(1)
	}

	ws := WorkReport.NewWorkReport()

	if err := ws.Load(os.Args[1]); err != nil {
		log.Fatalf("Failed to load shifts file \"%s\": %v\n", os.Args[1], err)
	}

	if err := ws.GenerateReports(os.Args[2]); err != nil {
		log.Fatalf("Failed to save reports to \"%s\": %v\n", os.Args[2], err)
	}

}
