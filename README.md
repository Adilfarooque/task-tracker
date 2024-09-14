#Task-Tracker

##Sample solution for the task-tracker challenge from roadmap.sh.

Backend Projects by roadmap.sh
This repository contains a list of project solutions for the Backend Developer roadmap. The projects are divided into three categories: beginner, intermediate, and advanced.

Run the following command to build and run the project:

go build -o task-tracker
./task-tracker --help # To see the list of available commands

# To add a task
./task-tracker add "Buy groceries"

# To update a task
./task-tracker update 1 "Buy groceries and cook dinner"

# To delete a task
./task-tracker delete 1

# To mark a task as in progress/done/todo
./task-tracker mark-in-progress 1
./task-tracker mark-done 1
./task-tracker mark-todo 1

# To list all tasks
./task-tracker list
./task-tracker list done
./task-tracker list todo
./task-tracker list in-progress