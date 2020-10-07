# urvaerk
A simple time keeping CLI tool. Use for easy time reporting and such.

The name is a pun. The word "urverk" in Swedish (my native language) means "clockwork", but the word "värk" (ae = ä) means "pain", because nobody likes time reporting :-)

Usage:
	
	urvaerk start "My Project" "some task"
	Start time counting on a task in a project. If the project or the task doesn't exist, it will be created. If there is a time count going on on another task this will be stopped and time recorded before starting the time count on the new task.
	
	urvaerk stop
	Stop time counting on current task in a project and record the time counted.
	
	urvaerk add "My Project" "some task" 60
	Adds 60 minutes to the task "some task" in "My Project". If the project or the task doesn't exist, it will be created.
	
	urvaerk add "My Project" 30
	Adds 30 minutes to "My Project" without putting it on a specific task (sort of, it's actually put on the task "My Project"). If the project doesn't exist it will be created.
	
	urvaerk remove "My Project" "some task"
	Removes the task "some task" in "My Project".
	
	urvaerk remove "My Project"
	Removes "My Project" and all of it's associated tasks.

	urvaerk show
	Shows all projects with a short summation of each one.

	urvaerk show "My Project"
	Shows all the tasks in a project and the time spent on them.

	urvaerk show "My Project" "some task"
	Shows all the time entries for a task in a project.

Parameters:

	--storage TYPE - Decides between storage in text-format (txt) or in an SQLite3 database (sql3). Defaults to txt.

	--filename FILENAME - FILENAME is the full path to the text-file or SQLite3 database used for storage. Defaults to $HOME/.urvaerk.dat and $HOME/.urvaerk.db respectively.
	

If you decide to use an SQLite3 database for storage you can create and initialize such a database-file with the supplied script *initdb.sh*.

***NOTE: Database storage is currently not implemented.***
