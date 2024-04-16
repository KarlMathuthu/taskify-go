# Taskify-Go

Taskify is a beginner-friendly project aimed at providing a platform for GoLang enthusiasts to test and improve their skills. It's a simple task manager application developed entirely in Go, offering a hands-on experience for beginners to dive into GoLang development.

## Table of Contents

- [Introduction](#introduction)
- [Features](#features)
- [Installation](#installation)
- [Usage](#usage)
- [Contributing](#contributing)
- [License](#license)

## Introduction

Taskify is designed to be a simple yet functional task manager application. It allows users to add, list, update, and remove tasks from their list. The project is structured in a beginner-friendly manner, making it easy for those new to GoLang to understand and contribute.

## Features

- **Add Task**: Users can add tasks to their list, providing a title and description for each task.
- **List Tasks**: Users can view a list of all their tasks.
- **Update Task**: Users can update existing tasks, modifying their title or description.
- **Remove Task**: Users can remove tasks from their list.

## Installation

To get started with Taskify, follow these simple steps:

1. Ensure you have Go installed on your system. If not, you can download and install it from the [official GoLang website](https://go.dev/).
2. Clone this repository to your local machine:

    ```bash
    git clone https://github.com/KarlMathuthu/taskify-go.git
    ```

3. Navigate to the project directory:

    ```bash
    cd Taskify
    ```

4. Build the project:

    ```bash
    go build
    ```

5. Run the executable:

    ```bash
    ./Taskify
    ```

## Usage.

Once the application is running, you can interact with it using the following endpoints:

- `add`: Add a new task to the list.
- `list`: List all tasks.
- `update`: Update an existing task.
- `remove`: Remove a task from the list.

Example usage:

- `GET` http://localhost:8080/tasks/
- `POST` http://localhost:8080/tasks/
- `PUT` http://localhost:8080/tasks/:id
- `DELETE ` http://localhost:8080/tasks/:id

## Contributing

Contributions to Taskify are welcome! Whether you're a beginner looking to improve your GoLang skills or an experienced developer wanting to support beginners, there are many ways to contribute:

- **Bug Fixes**: Find and fix bugs within the application.
- **Feature Implementation**: Add new features to enhance the functionality of the task manager.
- **Code Refactoring**: Improve the codebase for better readability and maintainability.
- **Documentation**: Enhance the README or provide inline comments for better understanding.
- **Testing**: Write tests to ensure the reliability of the application.

To contribute, fork this repository, make your changes, and submit a pull request. Your contributions will be greatly appreciated!

Developed with ❤️ by Karl Kiyotaka.

---
