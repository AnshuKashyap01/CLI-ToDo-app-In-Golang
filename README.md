# CLI Todo App

A simple and efficient command-line Todo application built with Go. This project allows users to manage tasks directly from the terminal with persistent storage and a clean, minimal interface.

## Features

* Add new tasks
* List all tasks
* Mark tasks as completed
* Delete tasks
* Persistent task storage using JSON
* Lightweight and fast
* Cross-platform support

## Project Structure

```text
.
├── main.go
├── storage.go
├── todo.go
├── tasks.json
└── README.md
```

## Installation

### Prerequisites

* Go 1.18 or later

### Clone the Repository

```bash
git clone https://github.com/your-username/cli-todo.git
cd cli-todo
```

### Build the Application

```bash
go build -o todo
```

## Usage

### Add a Task

```bash
./todo add "Learn Go Generics"
```

### List All Tasks

```bash
./todo list
`````

### Complete a Task

```bash
./todo complete 1
```

### Delete a Task

```bash
./todo delete 1
```

## Data Storage

Tasks are stored locally in a JSON file, ensuring data persists between application runs.

Example:

```json
[
  {
    "id": 1,
    "title": "Learn Go Generics",
    "completed": false
  }
]
```

## Technologies Used

* Go
* Standard Library

  * `encoding/json`
  * `os`
  * `fmt`

## Future Improvements

* Task priorities
* Due dates
* Categories and tags
* Search and filtering
* Colored terminal output
* Export/import functionality

## Contributing

Contributions are welcome. Feel free to fork the repository, create a feature branch, and submit a pull request.

## License

This project is licensed under the MIT License.
