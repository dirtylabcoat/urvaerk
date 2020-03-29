# urvaerk
A simple time keeping CLI tool. Use for easy time reporting and such.

The name is a pun. The word "urverk" in Swedish (my native language) means "clockwork", but the word "värk" (ae = ä) means "pain", because nobody likes time reporting :-)

Usage:

	urvaerk create "My Project"
	Creates a project called "My Project".
	
	urvaerk add "My Project" "some task" 1h
	Adds one hour to the task "some task" in "My Project". If "some task" doesn't exist, it will be created.
	
	urvaerk add "My Project" 30m
	Adds thirty minutes to "My Project" without putting it on a specific task (sort of, it's actually put on the task "My Project")
	
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
