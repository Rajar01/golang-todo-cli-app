# Todo CLI Application

This is a simple Todo List Command Line Interface (CLI) application written in Go. It allows users to add, list, update, and remove todos from the list.

## Features

- Add new todos with a title.
- List all todos with their status.
- Remove todos by their index.
- Update the title and status of todos.
- Cross-platform terminal clearing for a better CLI experience.

## Prerequisites

- Go 1.16+ installed on your machine.
- Git installed (optional, for cloning the repository).

## Installation

1. **Clone the repository**:

    ```bash
    git clone https://github.com/Rajar01/golang-todo-cli-app.git
    cd golang-todo-cli-app
    ```

2. **Navigate to the project directory**:

    ```bash
    cd golang-todo-cli-app
    ```

3. **Build the application**:

    ```bash
    go build -o todo-cli
    ```

    This will create an executable named `todo-cli`.

## Usage

1. **Run the application**:

    ```bash
    ./todo-cli
    ```

    If you're on Windows, run:

    ```cmd
    todo-cli.exe
    ```

2. **Select from the menu**:

    - `1`: Add a new todo.
    - `2`: List all todos.
    - `3`: Remove a todo by its index.
    - `4`: Update a todo's title and status.
    - `5`: Exit the application.

3. **Follow the prompts**:

    - For adding a todo, enter a title.
    - For removing or updating a todo, enter the index of the todo.
    - For updating, you will be prompted to enter a new title and status (`true` or `false`).

## Notes

- The terminal will clear before and after each operation for a better user experience.
- Use numbers to select options in the main menu.
- Use `true` or `false` to update the todo status.

## Troubleshooting

- **Error reading input**: Make sure you are entering valid input as expected by the prompts.
- **Index out of range**: Ensure that you are entering a valid index number when removing or updating a todo.

## Contributing

If you'd like to contribute to this project, feel free to fork the repository and submit a pull request.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for more details.

## Author

This project is maintained by [Rajar01](https://github.com/Rajar01).
