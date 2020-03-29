#!/bin/bash
# Create and initialize SQLite3 database for use with urvaerk
# https://github.com/dirtylabcoat/urvaerk
# Pre-requisites: SQLite3 (duh!)

if [ "$1" = "" ]; then
    DB=urvaerk.db
else
    DB="$1"
fi
sqlite3 "$DB" "CREATE TABLE project(id INTEGER AUTO_INCREMENT PRIMARY_KEY, name TEXT NOT NULL UNIQUE);"
sqlite3 "$DB" "CREATE TABLE task(id INTEGER AUTO_INCREMENT PRIMARY_KEY, name TEXT NOT NULL UNIQUE, project_id INTEGER NOT NULL, FOREIGN KEY (project_id) REFERENCES project (id) ON DELETE CASCADE);"
sqlite3 "$DB" "CREATE TABLE time(id INTEGER AUTO_INCREMENT PRIMARY_KEY, amount INTEGER NOT NULL, project_id INTEGER NOT NULL, task_id INTEGER NOT NULL, FOREIGN KEY (project_id) REFERENCES project (id) ON DELETE CASCADE, FOREIGN KEY (task_id) REFERENCES task (id) ON DELETE CASCADE);"
echo "Database [./$DB] has been created and initialized."
echo "Don't forget to put it where urvaerk can find it. Default is [$HOME/.urvaerk.db]."
