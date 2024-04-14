# Flow Command Reference

Welcome to the Flow command reference! This document provides an overview of the commands available in the Flow budget planning application. Flow is designed to empower users with the ability to track, analyze, and optimize their spending habits and financial goals through a user-friendly command-line interface (CLI).

## Introduction to Flow

Flow helps users manage their finances and achieve greater financial stability by leveraging the Formance API for a comprehensive financial management solution. With Flow, users can:

- **Track Spending**: Easily track their spending across different categories.
- **Set Budgets**: Create budgets for various spending categories to control expenses.
- **Receive Alerts**: Get notified when they exceed their budget for a specific category.
- **View and Adjust Budgets**: View, adjust, and remove budget allocations as needed.
- **Export Data**: Export budget data in CSV format for further analysis.

The Flow CLI provides intuitive commands and options to perform these tasks efficiently, making it easier for users to manage their finances effectively.

## Command Reference

Below is a detailed overview of each command available in the Flow CLI:

Sure, let's break down each command along with its structure for better understanding:

### 1. `flow`
- **Description**: The main command for the Flow application. It provides an overview of the application's purpose and available commands.
- **Usage**: `flow [command]`
- **Available Commands**:
  - `budget`: Manage your budget
  - `completion`: Generate the autocompletion script for the specified shell
  - `help`: Help about any command
  - `init`: Initialize your flow application
  - `transaction`: Transaction service
- **Flags**:
  - `-h, --help`: Display help information
  - `-t, --toggle`: Display help message for toggle
  - `-v, --version`: Display version information

### 2. `flow init`
- **Description**: Initalize your flow application before running other commands.
- **Note**: All these commands are required to enter.
- **Usage**: `flow budget flow [flags]`
- **Flags**:
  - `-a, --app_password string`   Write the App Password of your Gmail account
  - `-d, --dbname string`         Write the PostgreSQL DB name
  - `-g, --gmail string`          Write your Gmail address for alert notifications
  - `-h, --help`                  help for init
  - `-o, --host string`           Write the PostgreSQL host
  - `-w, --password string`       Write the PostgreSQL password
  - `-p, --port string`           Write the PostgreSQL port
  - `-s, --sslmode string`        Write the PostgreSQL SSLMode
  - `-u, --user string`           Write the PostgreSQL username
  - `-n, --username string`       Write your username

### 3. `flow budget`
- **Description**: Allows users to manage their budgetary allocations for different spending categories. Users can create, view, adjust, remove, and get budget details.
- **Usage**: 
  - `flow budget [flags]`
  - `flow budget [command]`
- **Available Commands**:
  - `adjust`: Adjust the budget details
  - `alert`: Get notification once you pass the budget
  - `create`: Create the budget of different categories
  - `get`: Get the budget data in CSV
  - `remove`: Remove the budget details
  - `view`: View the budget details
- **Flags**:
  - `-h, --help`: Display help information

### 4. `flow budget adjust`
- **Description**: Adjusts the budget details such as changing the amount or category name.
- **Usage**: `flow budget adjust [flags]`
- **Flags**:
  - `-a, --amount string`: New amount of the category to adjust
  - `-n, --newcategory string`: New category name to allocate
  - `-o, --oldcategory string`: Old category name to adjust

### 5. `flow budget alert`
- **Description**: Command for the alert notification.
- **Usage**: 
  - `flow budget alert [flags]`
  - `flow budget alert [command]`
- **Available Commands**:
  - `msg`: The message of alert notifications
  - `setup`: Setup for alert notification

### 6. `flow budget alert setup`
- **Description**: Sets up alerts to get notifications when the budget (for a specific category amount or the whole amount) is passed.
- **Usage**: 
  - `flow budget alert setup [flags]`
- **Flags**:
  - `-c, --category string`: Category name to monitor for budget
  - `-f, --frequency string`: Frequency of notifications (e.g., hourly, daily, weekly, monthly)
  - `-m, --method string`: Preferred method of notification [email or CLI] message

### 7. `flow budget alert msg`
- **Description**: Show the alert messages that were set.
- **Usage**: 
  - `flow budget alert msg [flags]`
- **Flags**:
  - `-h, --help`: Display help information

### 8. `flow budget create`
- **Description**: Creates a budget for different spending categories.
- **Usage**: `flow budget create [flags]`
- **Flags**:
  - `-a, --amount string`: Total amount for the category
  - `-c, --category string`: Category name (e.g., groceries, utilities, etc.)

### 9. `flow budget get`
- **Description**: Retrieves the budget data in CSV format.
- **Usage**: `flow budget get [flags]`
- **Flags**:
  - `-n, --filename string`: CSV file name to store the data
  - `-p, --filepath string`: File path to store the data

### 10. `flow budget remove`
- **Description**: Removes budget details for a specific category.
- **Usage**: `flow budget remove [flags]`
- **Flags**:
  - `-c, --category string`: Category name to remove

### 11. `flow budget view`
- **Description**: Views the budget details, optionally for a specific category.
- **Usage**: `flow budget view [flags]`
- **Flags**:
  - `-c, --category string`: Category name to show specific details

### 12. `flow completion`
- **Description**: Generates autocompletion scripts for various shells.
- **Usage**: `flow completion [command]`
- **Available Commands**:
  - `bash`: Generates autocompletion script for Bash
  - `fish`: Generates autocompletion script for Fish
  - `powershell`: Generates autocompletion script for PowerShell
  - `zsh`: Generates autocompletion script for Zsh
- **Flags**:
  - `-h, --help`: Display help information

### 13. `flow help`
- **Description**: Provides help about the Flow application or specific commands.
- **Usage**: `flow help [command]`

### 14. `flow spend`
- **Description**: Provides spending services on various categories.
- **Usage**: `flow spend [flags]`
- **Flags**:
  - `-a, --amount string`: Write the spending amount for a category
  - `-c, --category string`: Write the category name to spend the money on
  - `-h, --help`: Display help for spend

### 15. `flow total-amount`
- **Description**: Manage your total amount.
- **Usage**: `flow total-amount [flags]`
- **Available Commands**:
  - `set`: Set the total amount data
  - `remove`: Remove the total amount data
  - `update`: Update the total amount data
  - `view`: View the total amount data
- **Flags**:
  - `-a, --active string`: Make the total amount active
  - `-h, --help`: help for total-amount
  - `-i, --inactive string`: Make the total amount inactive

### 16. `flow total-amount active`
- **Description**: Make the total amount active.
- **Usage**: `flow total-amount active [flags]`

### 17. `flow total-amount inactive`
- **Description**: Make the total amount inactive.
- **Usage**: `flow total-amount inactive [flags]`

### 18. `flow total-amount set`
- **Description**: Set the total amount data.
- **Usage**: `flow total-amount set [flags]`
- **Flags**:
  - `-a, --amount string`: Write the total amount that you want to set
  - `-e, --exclude string`: Specify a category to exclude from the total amount
  - `-i, --include string`: Specify a category to include in the total amount
  - `-l, --label string`: Provide a label for setting up your total amount
  - `-h, --help`: help for set

### 19. `flow total-amount update`
- **Description**: Update the total amount data.
- **Usage**: `flow total-amount update [flags]`
- **Flags**:
  - `-a, --amount string`: Write the total amount that you want to update
  - `-l, --label string`: Write the label that you want to update
  - `-h, --help`: help for set

### 20. `flow total-amount remove`
- **Description**: Remove the total amount data.
- **Usage**: `flow total-amount remove [flags]`

### 21. `flow total-amount view`
- **Description**: View the total amount data.
- **Usage**: `flow total-amount view [flags]`