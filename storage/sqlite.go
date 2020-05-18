package storage

import "fmt"

// SqliteHandler - ...
type SqliteHandler struct {
	Filename string
}

// Remove - ...
func (sqlite SqliteHandler) Remove(project string, task string) error {
	return nil
}

// Add - ...
func (sqlite SqliteHandler) Add(time PieceOfTime) error {
	fmt.Println("sql3 Add time", time.AmountInMin, "to task", time.Task, "in project", time.Project, "in file", sqlite.Filename)
	return nil
}

// GetProjects - ...
func (sqlite SqliteHandler) GetProjects() []ProjectSummary {
	return nil
}

// GetProject - ...
func (sqlite SqliteHandler) GetProject(project string) []PieceOfTime {
	return nil
}

// GetTask - ...
func (sqlite SqliteHandler) GetTask(project string, task string) PieceOfTime {
	return PieceOfTime{}
}
