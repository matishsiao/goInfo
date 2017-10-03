package goInfo

import (
	"bytes"
	"encoding/csv"
	"os"
	"os/exec"
	"runtime"
)

func GetInfo() *GoInfoObject {
	// Run the systeminfo command, specifying the output to be in CSV (/FO), with no header (/NH)
	cmd := exec.Command("systeminfo", "/NH", "/FO", "CSV")
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		panic(err)
	}
	// Parse the output as a single CSV
	r := csv.NewReader(&out)
	records, err := r.ReadAll()

	if err != nil {
		panic(err)
	}

	// If there are less than 1 rows in the CSV, there has been some kind of mistake
	if len(records) < 1 {
		panic("Failed to parse the output of systeminfo as a CSV")
	}

	// TODO: We should stop cheating with the lack of CSV header and instead parse the CSV into a real map or similar so we can be sure we have the right index always
	gio := &GoInfoObject{
		Kernel:   "windows",        // windowss
		Core:     records[0][2],    // 10.0.15063 N/A Build 15063
		Platform: records[0][13],   // x64-based PC
		OS:       records[0][1],    // Microsoft Windows 10 Pro
		GoOS:     runtime.GOOS,     // windows
		CPUs:     runtime.NumCPU(), // 4
	}

	// Set the hostname
	gio.Hostname, err = os.Hostname()
	if err != nil {
		panic(err)
	}

	return gio
}
