package storage

import "fmt"

// TxtHandler - ...
type TxtHandler struct {
	Filename string
}

// Create - create new project or task in text-file
func (txt TxtHandler) Create(projectName string, taskName string) error {
	fmt.Println("Create TXT", projectName, taskName)
	return nil
}
