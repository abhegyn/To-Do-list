# To-Do List CLI

A simple command-line to-do list application written in Go using [Cobra](https://github.com/spf13/cobra).

Tasks are stored locally in a tab-separated file named `tasks.csv`.

## Features

- Add tasks
- View all tasks
- Update task status
- Store tasks locally in a file
- Use command-line flags

## Requirements

- Go 1.20 or newer

## Installation

Clone the repository:

```bash
git clone https://github.com/your-username/your-repository.git
cd your-repository
```

Install dependencies:

```bash
go mod tidy
```

## Usage

Display available commands:

```bash
go run . --help
```

## View Tasks

Display all tasks:

```bash
go run . view
```

Example output:

```text
ID  Task             Status
1   Buy groceries    pending
2   Read a book      completed
3   Clean the room   pending
```

## Update a Task

Set a task to completed:

```bash
go run . update --taskid 3 --newstatus
```

Using short flags:

```bash
go run . update -i 3 -n
```

Set a task to pending:

```bash
go run . update --taskid 3
```

## Update Flags

| Flag | Short Form | Description |
|---|---|---|
| `--taskid` | `-i` | ID of the task to update |
| `--newstatus` | `-n` | Set the task status to completed |

If `--newstatus` is omitted, the task status is set to `pending`.

## Data Format

Tasks are stored in `tasks.csv` using tab-separated values:

```text
ID	Task	Status
1	Buy groceries	pending
2	Read a book	completed
3	Clean the room	pending
```

Although the file is named `tasks.csv`, its fields are separated by tabs instead of commas.

Each row contains:

1. `ID`
2. `Task`
3. `Status`

## Build the Application

Build an executable:

```bash
go build -o tasks
```

Run the executable:

```bash
./tasks
```

On Windows:

```powershell
go build -o tasks.exe
.\tasks.exe
```

## Development Commands

Format the code:

```bash
go fmt ./...
```

Run tests:

```bash
go test ./...
```

Run static checks:

```bash
go vet ./...
```

## Project Structure

```text
.
├── cmd/
│   ├── add.go
│   ├── root.go
│   ├── update.go
│   └── view.go
├── tasks.csv
├── go.mod
├── go.sum
└── main.go
```

## Future Improvements

- Delete tasks
- Search and filter tasks
- Add due dates
- Prevent duplicate task IDs
- Add automated tests
- Support JSON or database storage

## License

This project is available under the MIT License.
