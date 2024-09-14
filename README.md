# Task-Tracker CLI
This is a simple command-line interface (CLI) application designed to help you manage your tasks efficiently. The Task Tracker allows you to add, update, delete tasks, and track their status (todo, in-progress, or done) using a JSON file to store the data.

## Sample Solution for Task-Tracker Challenge from roadmap.sh
This repository provides a solution for the [Task Tracker Challenge](https://roadmap.sh/projects/task-tracker) by roadmap.sh. The challenge is part of the backend projects roadmap, designed to enhance backend development skills through building a task management CLI tool.

View the challenge on roadmap.sh

## Project Overview
#### The Task Tracker CLI application lets you:

Add new tasks with descriptions.
Update the descriptions of existing tasks.
Delete tasks by their unique IDs.
Mark tasks as "todo", "in-progress", or "done".
List tasks based on their status or show all tasks.

## Commands and Usage
### Build and Run the Project
To build and run the project, use the following commands:

go build -o task-tracker
./task-tracker --help    # Display available commands

## Adding a Task
./task-tracker add "Buy groceries"

## Updating a Task
./task-tracker update 1 "Buy groceries and cook dinner"

## Deleting a Task
./task-tracker delete 1

## Marking a Task (as todo, in-progress, or done)
./task-tracker mark-todo 1
./task-tracker mark-in-progress 1
./task-tracker mark-done 1

## Listing Tasks
./task-tracker list           # List all tasks
./task-tracker list todo      # List tasks that are todo
./task-tracker list in-progress # List tasks that are in-progress
./task-tracker list done      # List tasks that are done
