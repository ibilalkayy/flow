# Flow

Flow is a budget planning application designed to empower users with the ability to track,  analyze, and optimize their spending habits and financial goals. With a user-friendly CLI. It manages the finances and achieve greater financial stability  by leveraging the Formance API for a comprehensive financial management solution.

## Table of Contents

- [Clonning](#clonning)
- [Getting Started](#getting-started)
- [Installation](#installation)
- [Commands](#commands)
- [Contributing](#contributing)
- [License](#license)

## Getting Started

To get started in flow, you need to have Golang installed on your machine. Once you have installed Go, you can clone the application using the following command:

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

Flow provides a variety of commands for managing Weaviate instances. Below are some key commands:

To use the application, simply run the `flow` command followed by one of the following subcommands:

- `budget`: Makes the budget planning

Each subcommand has its own set of options and arguments. Here are some examples of how to use the application:

```bash
flow budget create --category groceries/utilities --amount 300
```

## Clonning

1. Clone the repository:

    ```bash
    git clone https://github.com/ibilalkayy/flow.git
    ```

2. Navigate to the project directory:

    ```bash
    cd flow
    ```

3. Build and install the flow binary:

    ```bash
    go install
    ```

## Contributing

We welcome contributions! If you have ideas for new features, find a bug, or want to improve documentation, feel free to open an issue or submit a pull request. Please follow our [Contribution Guidelines](CONTRIBUTING.md) for a smooth collaboration.

## License

Flow is licensed under the [Apache-2.0 License](LICENSE). Feel free to use, modify, and distribute the code as per the license terms.