package utils

import (
	"os/exec"
)

func ReadPDFText(path string) (string, error) {
	// Run pdftotext command and capture its output.
	cmd := exec.Command("pdftotext", path, "-")
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}

	return string(output), nil
}
