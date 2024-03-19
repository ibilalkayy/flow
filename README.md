# Flow

Flow is a budget planning application designed to empower users with the ability to track,  analyze, and optimize their spending habits and financial goals. 

With a user-friendly CLI. It manages the finances and achieve greater financial stability  by leveraging the Formance API for a comprehensive financial management solution.

## Table of Contents

- [Documentation](#documentation)
- [What You Need?](#what-you-need)
- [Clonning](#clonning)
- [Installation](#installation)
- [Commands](#commands)
- [Run The App](#run-the-app)
- [Contributing](#contributing)
- [License](#license)


## Documentation

To know about the application in detail, you can visit the [docs](https://github.com/ibilalkayy/flow/tree/main/docs) to understand the application in a better way.

## What You Need?

To get started in flow, you need to have two applications installed on your machine.

1. [Golang](https://go.dev/dl/)
2. [Docker](https://www.docker.com/get-started/)
3. [PostgreSQL](https://www.postgresql.org/)

## Installation

You can install the application in your local machine using the following command:

  ```bash
  go install github.com/ibilalkayy/flow@latest
  ```

Verify the installation through the following command:

```bash
flow --version
```

This will display the installed flow version.

## Commands

Flow provides a variety of commands for managing the budget. Below are some key commands:

To use the application, simply run the `flow` command followed by one of the following subcommands:

- `budget`: Makes the budget planning

Each subcommand has its own set of options and arguments. Here are some examples of how to use the application:

```bash
# Create a budget
flow budget create --category groceries/utilities --amount 300

# View the budget info
flow budget view
```

## Clonning

Clone the repository:

```bash
git clone https://github.com/ibilalkayy/flow.git
```

Navigate to the project directory:

```bash
cd flow
```

Create a `.env` file to put all your PostgresSQL credentials.

## Run the App

There are two ways through which you can run this clonned application.

1. Build and install the flow binary through Golang:

    ```bash
    go build
    ```

    ```bash
    go install
    ```
2. Use the docker command to run it:

    ```bash
    docker compose up -d
    ```

    ```bash
    1. docker exec -it flow-app-1 ./flow budget create
    2. docker exec -it flow-app-1 ./flow budget view
    ...
    ```

## Contributing

We welcome contributions! If you have ideas for new features, find a bug, or want to improve documentation, feel free to open an issue or submit a pull request. Please follow our [Contribution Guidelines](CONTRIBUTING.md) for a smooth collaboration.

## License

Flow is licensed under the [Apache-2.0 License](LICENSE). Feel free to use, modify, and distribute the code as per the license terms.