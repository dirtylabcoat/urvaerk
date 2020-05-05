package storage

// PieceOfTime - a piece of time on a task in a project
type PieceOfTime struct {
	Project     string
	Task        string
	AmountInMin int
}

// ProjectSummary - a summary of an existing project
type ProjectSummary struct {
	Project    string
	NumOfTasks int
	TotalTime  int
}

// TimeHandler - perform CRUD on pieces of time
type TimeHandler interface {
	Add(time PieceOfTime) error
	// Remove(project string, task string) error
	GetProjects() []ProjectSummary
	GetProject(project string) []PieceOfTime
	GetTask(project string, task string) PieceOfTime
}
