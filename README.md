# ToDo CLI

A lightweight command-line interface (CLI) application to manage your tasks. Add, complete, delete, or list your ToDos with simple commands

---

## Features

- **Add a Task:** Add new tasks to your list.
- **Mark as Complete:** Update tasks as completed.
-**Delete a Task:** Remove tasks you no longer need.
- **List Tasks:** View all tasks along with their statuses.


---

## Installation

1. Clone the repository
    ```bash
    git clone https://github.com.Sudhir-rai07/todo-cli.git 
    ```
2. Navigate to the project directory:
    ```bash
    cd todo-cli
    ```
3. Build the application
    ```go
    go build -o todo
    ```

---
# Usage
Run the application using following commands:
 ## Add a ToDo
 ```bash
 ./todo -a "Task"
 ```
 ### Example:
 ```bash
 ./todo -a "Learn GO"
 ```
 ## Mark a ToDo as Complete
 ```bash
 ./todo -c <index-number>
 ```
 ### Example
 ```bash
 ./todo -c 2
 ```
 ## Delete a ToDo
 ```bash
 ./todo -d <index-number>
 ```
 ### Example
 ```bash
 ./todo -d 2
 ```
 ## List All ToDos
 ```bash
 ./todo -ls
 ```
 ### Example output:
 ```markdown
 | #  | Description         | Completed |
|-----|---------------------|-----------|
| 1   | Learn Go            | false     |
| 2   | Build a project     | true      |

 ```

---
 ## Data Storage
 ToDos are saved in a file named ```.todos.json``` in the current directory. This file persists your tasks between sessions.

 ## Contribution
 Feel free to fork the repository and make improvements. Submit a pull request with your changes!