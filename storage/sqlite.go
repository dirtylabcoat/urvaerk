package storage

import "fmt"

// DbHandler - ...
type SqliteHandler struct {
	Filename string
}

func (sqlite SqliteHandler) Add(time PieceOfTime) error {
	fmt.Println("sql3 Add time", time.AmountInMin, "to task", time.Task, "in project", time.Project, "in file", sqlite.Filename)
	return nil
}

func (sqlite SqliteHandler) GetProjects() []ProjectSummary {
	return nil
}

func (sqlite SqliteHandler) GetProject(project string) []PieceOfTime {
	return nil
}

func (sqlite SqliteHandler) GetTask(project string, task string) PieceOfTime {
	return PieceOfTime{}
}
