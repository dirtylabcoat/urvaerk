package storage

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

// TxtHandler - ...
type TxtHandler struct {
	Filename string
}

// TasksAndTime - ...
type TasksAndTime struct {
	Tasks int
	Time  int
}

// Add - add time to a task on a project
func (txt TxtHandler) Add(time PieceOfTime) error {
	fileHandle, err := os.OpenFile(txt.Filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer fileHandle.Close()
	line := time.Project + ";" + time.Task + ";" + strconv.Itoa(time.AmountInMin) + "\n"
	if _, err := fileHandle.WriteString(line); err != nil {
		log.Println(err)
	}
	return nil
}

// Remove - remove task or project
func (txt TxtHandler) Remove(project string, task string) error {
	tmpFilename := txt.Filename + ".tmp"
	rm := project + ";"
	if len(task) != 0 {
		rm = rm + task + ";"
	}
	fileHandle, err := os.Open(txt.Filename)
	if err != nil {
		log.Fatalf("Failed to open file: %s", err)
	}
	// File handle for tmp-file
	newFileHandle, err := os.OpenFile(tmpFilename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	linesRemoved := 0
	scanner := bufio.NewScanner(fileHandle)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		line := scanner.Text()
		if !strings.HasPrefix(line, rm) {
			if _, err := newFileHandle.WriteString(line + "\n"); err != nil {
				log.Println(err)
			}
		} else {
			linesRemoved++
		}
	}
	defer fileHandle.Close()
	defer newFileHandle.Close()
	if linesRemoved > 0 {
		os.Rename(tmpFilename, txt.Filename)
	}
	return nil
}

// GetProjects - ...
func (txt TxtHandler) GetProjects() []ProjectSummary {
	fileHandle, err := os.Open(txt.Filename)
	if err != nil {
		log.Fatalf("Failed to open file: %s", err)
	}
	scanner := bufio.NewScanner(fileHandle)
	scanner.Split(bufio.ScanLines)
	var projectSummaries []ProjectSummary
	var projects []string
	for scanner.Scan() {
		line := scanner.Text()
		pieces := strings.Split(line, ";")
		projectName := pieces[0]
		if !arrContainsStr(projects, projectName) {
			projects = append(projects, pieces[0])
			projectTaskSummaries := txt.GetProject(projectName)
			sum := 0
			for _, p := range projectTaskSummaries {
				sum += p.AmountInMin
			}
			projectSummaries = append(projectSummaries, ProjectSummary{Project: projectName, NumOfTasks: len(projectTaskSummaries), TotalTime: sum})
		}
	}
	defer fileHandle.Close()
	return projectSummaries
}

// GetProject - ...
func (txt TxtHandler) GetProject(project string) []PieceOfTime {
	fileHandle, err := os.Open(txt.Filename)
	if err != nil {
		log.Fatalf("Failed to open file: %s", err)
	}
	scanner := bufio.NewScanner(fileHandle)
	scanner.Split(bufio.ScanLines)
	var piecesOfTime []PieceOfTime
	var tasks []string
	for scanner.Scan() {
		line := scanner.Text()
		pieces := strings.Split(line, ";")
		task := pieces[1]
		if pieces[0] == project && !arrContainsStr(tasks, task) {
			tasks = append(tasks, task)
			piecesOfTime = append(piecesOfTime, txt.GetTask(project, task))
		}
	}
	defer fileHandle.Close()
	return piecesOfTime
}

// GetTask - ...
func (txt TxtHandler) GetTask(project string, task string) PieceOfTime {
	fileHandle, err := os.Open(txt.Filename)
	if err != nil {
		log.Fatalf("Failed to open file: %s", err)
	}
	scanner := bufio.NewScanner(fileHandle)
	scanner.Split(bufio.ScanLines)
	var piecesOfTime []PieceOfTime
	for scanner.Scan() {
		line := scanner.Text()
		pieces := strings.Split(line, ";")
		if pieces[0] == project && pieces[1] == task {
			mins, _ := strconv.Atoi(pieces[2])
			piecesOfTime = append(piecesOfTime, PieceOfTime{Project: pieces[0], Task: pieces[1], AmountInMin: mins})
		}
	}
	defer fileHandle.Close()
	sum := 0
	for _, p := range piecesOfTime {
		sum += p.AmountInMin
	}
	return PieceOfTime{Project: project, Task: task, AmountInMin: sum}
}

func arrContainsStr(arr []string, str string) bool {
	for _, n := range arr {
		if str == n {
			return true
		}
	}
	return false
}
