package storage

import "fmt"

// DbHandler - ...
type SqliteHandler struct {
	Filename string
}

// Create - create new project or task in database
func (sqlite SqliteHandler) Create(projectName string, taskName string) error {
	fmt.Println("Create DB", projectName, taskName)
	return nil
}
