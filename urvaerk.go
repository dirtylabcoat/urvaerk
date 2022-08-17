package main

import (
	"fmt"
	"log"
	"os"
	"os/user"
	"strconv"
	"time"

	"dirtylabcoat.org/urvaerk/storage"
	"github.com/urfave/cli/v2"
)

var storageHandler storage.TimeHandler

func main() {
	usr, _ := user.Current()
	storageType := "txt"
	timeCountFile := usr.HomeDir + "/.urvaerk_timecount.dat"
	txtDefaultFile := usr.HomeDir + "/.urvaerk.dat"
	sql3DefaultFile := usr.HomeDir + "/.urvaerk.db"
	storageFilename := ""
	app := &cli.App{
		Name:        "urvaerk",
		Usage:       "simple time keeping",
		Version:     "0.1",
		Compiled:    time.Now(),
		UsageText:   "",
		Description: "A simple time keeping CLI tool",
		ArgsUsage:   "",
		Authors: []*cli.Author{
			&cli.Author{
				Name:  "Fighter Hayabusa",
				Email: "fighterhayabusa@dirtylabcoat.org",
			},
		},
		Copyright: "(c) 2020 Fighter Hayabusa <fighterhayabusa@dirtylabcoat.org>",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "storage",
				Aliases:     []string{"s"},
				Usage:       "Use `TYPE` for storage. Choose between text-format (txt) or SQLite3 (sql3). Default is 'txt'.",
				Value:       "txt",
				Destination: &storageType,
			},
			&cli.StringFlag{
				Name:        "filename",
				Aliases:     []string{"f"},
				Usage:       "Store your data in `FILENAME`. Defaults to $HOME/.urvaerk.dat for text-format and $HOME/.urvaerk.db for SQLite3.",
				Value:       storageFilename,
				Destination: &storageFilename,
			},
		},
		Commands: []*cli.Command{
			&cli.Command{
				Name:        "start",
				Aliases:     []string{"s"},
				Usage:       "Start counting time on task in project, creates project/task if it doesn't exist",
				UsageText:   "start - start counting time",
				Description: "no really, there is a lot of starting to be done",
				Action:      start,
			},
			&cli.Command{
				Name:        "stop",
				Aliases:     []string{"p"},
				Usage:       "Stop counting time on current task and record the time counted",
				UsageText:   "stop - stop counting time",
				Description: "no really, there is a lot of adding to be done",
				Action:      stop,
			},
			&cli.Command{
				Name:        "add",
				Aliases:     []string{"a"},
				Usage:       "Add time to task in project, creates project/task if it doesn't exist",
				UsageText:   "add - does the adding",
				Description: "no really, there is a lot of adding to be done",
				Action:      add,
			},
			&cli.Command{
				Name:        "remove",
				Aliases:     []string{"rm"},
				Usage:       "Remove task or project",
				UsageText:   "remove - does the removing",
				Description: "no really, there is a lot of removing to be done",
				Action:      remove,
			},
			&cli.Command{
				Name:        "show",
				Aliases:     []string{"ls", "list"},
				Usage:       "Show projects and tasks",
				UsageText:   "show - does the showing",
				Description: "no really, there is a lot of showing to be done",
				Action:      show,
			},
		},
		Before: func(c *cli.Context) error {
			if storageType == "txt" {
				if storageFilename == "" {
					storageFilename = txtDefaultFile
				}
				storageHandler = storage.TxtHandler{Filename: storageFilename}
			} else if storageType == "sql3" {
				if storageFilename == "" {
					storageFilename = sql3DefaultFile
				}
				storageHandler = storage.SqliteHandler{Filename: storageFilename}
			}
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func start(c *cli.Context) error {
	project := c.Args().Get(0)
	var task string
	if c.Args().Len() == 1 {
		task = project
	} else if c.Args().Len() == 2 {
		task = c.Args().Get(1)
	} else {
		log.Fatal("Command add takes 1 or 2 arguments.")
	}
	return nil
}

func stop(c *cli.Context) error {
	return nil
}

func add(c *cli.Context) error {
	project := c.Args().Get(0)
	var task string
	var time int
	if c.Args().Len() == 2 {
		task = project
		time, _ = strconv.Atoi(c.Args().Get(1))
	} else if c.Args().Len() == 3 {
		task = c.Args().Get(1)
		time, _ = strconv.Atoi(c.Args().Get(2))
	} else {
		log.Fatal("Command add takes 2 or 3 arguments.")
	}
	timePiece := storage.PieceOfTime{Project: project, Task: task, AmountInMin: time}
	storageHandler.Add(timePiece)
	return nil
}

func remove(c *cli.Context) error {
	fmt.Printf("Remove %q %q\n", c.Args().Get(0), c.Args().Get(1))
	storageHandler.Remove(c.Args().Get(0), c.Args().Get(1))
	return nil
}

func show(c *cli.Context) error {
	if c.Args().Len() == 0 {
		projects := storageHandler.GetProjects()
		for _, p := range projects {
			fmt.Printf("%s : %d tasks : %d minutes\n", p.Project, p.NumOfTasks, p.TotalTime)
		}
	} else if c.Args().Len() == 1 {
		projectSummary := storageHandler.GetProject(c.Args().Get(0))
		for _, p := range projectSummary {
			fmt.Printf("%s : %d minutes\n", p.Task, p.AmountInMin)
		}
	} else if c.Args().Len() == 2 {
		taskSummary := storageHandler.GetTask(c.Args().Get(0), c.Args().Get(1))
		fmt.Printf("%s : %s : %d minutes\n", taskSummary.Project, taskSummary.Task, taskSummary.AmountInMin)
	}
	return nil
}
