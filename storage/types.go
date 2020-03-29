package storage

// PieceOfTime - a piece of time on a task in a project
type PieceOfTime struct {
	Project     string
	Task        string
	AmountInMin int32
}

// TimeHandler - perform CRUD on pieces of time
type TimeHandler interface {
	Create(projectName string, taskName string) error
	// Add(time PieceOfTime) error
	// Remove(projectName string, taskName string) error
	// GetAllProjects() []string
	// GetProject(name string) []PieceOfTime
	// GetTask(name string) []PieceOfTime
}
