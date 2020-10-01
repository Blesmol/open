package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Printf("No arguments provided, nothing to do\n")
		os.Exit(1)
	}

	if err := openWithDefaultViewer(os.Args[1]); err != nil {
		fmt.Printf("%v\n", err)
	}
}

// openWithDefaultViewer opens the provided file with the default viewer registered in the system for this.
func openWithDefaultViewer(file string) (err error) {
	absFile, err := filepath.Abs(file)
	if err != nil {
		return err
	}
	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", absFile).Start()
	case "windows":
		err = exec.Command(filepath.Join(os.Getenv("SYSTEMROOT"), "System32", "rundll32.exe"), "url.dll,FileProtocolHandler", absFile).Start()
	case "darwin":
		err = exec.Command("open", absFile).Start()
	default:
		err = fmt.Errorf("Unknown OS, cannot open file '%v' automatically", file)
	}

	return err
}
