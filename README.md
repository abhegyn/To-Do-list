# To-Do List CLI

A simple command-line to-do list application written in Go using [Cobra](https://github.com/spf13/cobra). Tasks are stored locally in a `tasks.csv` file.

## Features

- Add tasks from the command line
- Store each task with:
  - ID
  - Task name
  - Status
- Automatically assign the next available ID
- Append tasks to an existing CSV file
- Mark tasks as pending or completed
- Support for long and short command-line flags

## Requirements

- [Go](https://go.dev/dl/) installed on your system

## Installation

Clone the repository:

```bash
git clone https://github.com/abhegyn/To-Do-list.git
cd To-Do-list
