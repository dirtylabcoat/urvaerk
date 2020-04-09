package storage

// PieceOfTime - a piece of time on a task in a project
type PieceOfTime struct {
	Project     string
	Task        string
	AmountInMin int
}

type ProjectSummary struct {
	Project    string
	NumOfTasks int
	TotalTime  int
}

// TimeHandler - perform CRUD on pieces of time
type TimeHandler interface {
	Create(projectName string) error
	Add(time PieceOfTime) error
	// Remove(projectName string, taskName string) error
	GetProjects() []ProjectSummary
	GetProject(projectName string) []PieceOfTime
	GetTask(projectName string, taskName string) PieceOfTime
}
