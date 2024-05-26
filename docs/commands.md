# Flow Command Reference

Welcome to the Flow command reference! This document provides an overview of the commands available in the Flow budget planning application. Flow is designed to empower users with the ability to track, analyze, and optimize their spending habits and financial goals through a user-friendly command-line interface (CLI).

## Introduction to Flow

Flow helps users manage their finances and achieve greater financial stability by leveraging different payment APIs for a comprehensive financial management solution. With Flow, users can:

- **Track Spending**: Easily track their spending across different categories.
- **Set Budgets**: Create budgets for various spending categories to control expenses.
- **Receive Alerts**: Get notified when they exceed their budget for a specific category.
- **View, Update and Remove Budgets**: View, update, and remove budget allocations as needed.
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
  - `-t, --toggle`: Display help message for toggle
  - `-v, --version`: Display version information
  - `-h, --help`: Display help information

### 2. `flow init`
- **Description**: Initalize your flow application before running other commands.
- **Note**: All these commands are required to enter.
- **Usage**: `flow budget init [flags]`
- **Flags**:
  - `-a, --app-password string`: The App Password of your Gmail account
  - `-d, --db-name string`: The PostgreSQL DB name
  - `-g, --gmail string`: Your Gmail address for alert notifications
  - `-o, --db-host string`: The PostgreSQL host
  - `-w, --db-password string`: The PostgreSQL password
  - `-p, --db-port string`: The PostgreSQL port
  - `-s, --sslmode string`: The PostgreSQL SSLMode
  - `-u, --db-user string`: The PostgreSQL username
  - `-n, --db-user string`: Your username
  - `-h, --help`: Display help information

### 3. `flow budget`
- **Description**: Allows users to manage their budgetary allocations for different spending categories. Users can create, view, update, remove, and get budget details.
- **Usage**: 
  - `flow budget [flags]`
  - `flow budget [command]`
- **Available Commands**:
  - `update`: Update the budget details
  - `alert`: Get notification once you pass the budget
  - `create`: Create the budget of different categories
  - `get`: Get the budget data in CSV
  - `remove`: Remove the budget details
  - `view`: View the budget details
- **Flags**:
  - `-h, --help`: Display help information

### 4. `flow budget update`
- **Description**: Updates the budget details such as changing the amount or category name.
- **Usage**: `flow budget update [flags]`
- **Flags**:
  - `-a, --amount string`: New amount of the category to update
  - `-n, --new-category string`: New category name to allocate
  - `-o, --old-category string`: Old category name to update

### 5. `flow budget alert`
- **Description**: Command for the alert notification.
- **Usage**: 
  - `flow budget alert [flags]`
  - `flow budget alert [command]`
- **Available Commands**:
  - `setup`: Setup for alert notification
  - `msg`: Get alert notifications in your email
  - `remove`: Remove the alert values
  - `update`: Update the alert values for notification
  - `view`: View the alert notifications

### 6. `flow budget alert setup`
- **Description**: Sets up alerts to get notifications when the budget (for a specific category amount or the whole amount) is passed.
- **Usage**: 
  - `flow budget alert setup [flags]`
- **Flags**:
  - `-c, --category string`: Category name to take its budget amount
  - `-f, --frequency string`: Frequency of notifications (e.g., hourly, daily, weekly, monthly)
  - `-m, --method string`: Preferred method of notification [email or CLI] message
  - `-d, --day string`: A day to set the notification
  - `-o, --hour string`: An Hour to set the notification
  - `-t, --method string`: A preferred method of notification [email or CLI] message
  - `-m, --minute string`: The minute to set the notification
  - `-s, --second string`: The second to set the notification
  - `-w, --weekday string`: Write a weekday to set the notification
  - `-h, --help`: Display help information

### 7. `flow budget alert msg`
- **Description**: Show the alert messages that were set.
- **Usage**: 
  - `flow budget alert msg [flags]`
- **Flags**:
  - `-c, --category string`: Write the category to get the notification
  - `-h, --help`: Display help information

### 8. `flow budget alert remove`
- **Description**: Remove the alert values.
- **Usage**: 
  - `flow budget alert remove [flags]`
- **Flags**:
  - `-c, --category string`: Write the category to remove the alert values
  - `-h, --help`: Display help information

### 9. `flow budget alert update`
- **Description**: Update the alert values for notification.
- **Usage**: 
  - `flow budget alert update [flags]`
- **Flags**:
  - `-c, --category string`: Category name to take its budget amount
  - `-f, --frequency string`: Frequency of notifications (e.g., hourly, daily, weekly, monthly)
  - `-m, --method string`: Preferred method of notification [email or CLI] message
  - `-d, --day string`: A day to set the notification
  - `-o, --hour string`: An Hour to set the notification
  - `-t, --method string`: A preferred method of notification [email or CLI] message
  - `-m, --minute string`: The minute to set the notification
  - `-s, --second string`: The second to set the notification
  - `-w, --weekday string`: Write a weekday to set the notification
  - `-h, --help`: Display help information

### 10. `flow budget alert view`
- **Description**: View the alert notification values.
- **Usage**: 
  - `flow budget alert view [flags]`
- **Flags**:
  - `-c, --category string`: Write the category to see the alert notification values
  - `-h, --help`: Display help information

### 11. `flow budget create`
- **Description**: Creates a budget for different spending categories.
- **Usage**: `flow budget create [flags]`
- **Flags**:
  - `-a, --amount string`: Total amount for the category
  - `-c, --category string`: Category name (e.g., groceries, utilities, etc.)

### 12. `flow budget view`
- **Description**: Views the budget details, optionally for a specific category.
- **Usage**: `flow budget view [flags]`
- **Flags**:
  - `-c, --category string`: Category name to show specific details

### 13. `flow budget get`
- **Description**: Retrieves the budget data in CSV format.
- **Usage**: `flow budget get [flags]`
- **Flags**:
  - `-n, --filename string`: CSV file name to store the data
  - `-p, --filepath string`: File path to store the data

### 14. `flow budget remove`
- **Description**: Removes the budget details for a specific category.
- **Usage**: `flow budget remove [flags]`
- **Flags**:
  - `-c, --category string`: Category name to remove

### 15. `flow completion`
- **Description**: Generates autocompletion scripts for various shells.
- **Usage**: `flow completion [command]`
- **Available Commands**:
  - `bash`: Generates autocompletion script for Bash
  - `fish`: Generates autocompletion script for Fish
  - `powershell`: Generates autocompletion script for PowerShell
  - `zsh`: Generates autocompletion script for Zsh
- **Flags**:
  - `-h, --help`: Display help information

### 16. `flow help`
- **Description**: Provides help about the Flow application or specific commands.
- **Usage**: `flow help [command]`

### 17. `flow spend`
- **Description**: Provides spending services on various categories.
- **Usage**: `flow spend [flags]`
- **Available Commands**:
  - `history`: Show the transaction history
- **Flags**:
  - `-a, --amount string`: Write the spending amount for a category
  - `-c, --category string`: Write the category name to spend the money on
  - `-h, --help`: Display help information

### 18. `flow spend history`
- **Description**: Provides spending services on various categories.
- **Usage**: 
    - `flow spend history [flags]`
    - `flow spend history [command]`
- **Flags**:
  - `-h, --help`: Display help information

### 19. `flow spend history remove`
- **Description**: Removes the history data.
- **Usage**: `flow spend history remove [flags]`
- **Flags**:
  - `-c, --category string`: Write the category name to remove it's history
  - `-h, --help`: Display help information

### 20. `flow spend history show`
- **Description**: Show the history data.
- **Usage**: `flow spend history show [flags]`
- **Flags**:
  - `-c, --category string`: Write the category to show it's history
  - `-h, --help`: Display help information

### 21. `flow total-amount`
- **Description**: Manage your total amount.
- **Usage**: `flow total-amount [flags]`
- **Available Commands**:
  - `add`: Add the total amount data
  - `remove`: Remove the total amount data
  - `status`: Get teh status of the total amount
  - `update`: Update the total amount data
  - `view`: View the total amount data

### 22. `flow total-amount add`
- **Description**: Add the total amount data.
- **Usage**: `flow total-amount add [flags]`
- **Flags**:
  - `-a, --amount string`: Write the total amount that you want to set
  - `-c, --category string`: Specify a category to include in the total amount
  - `-l, --label string`: Provide a label for setting up your total amount
  - `-h, --help`: Display help information

### 23. `flow total-amount status`
- **Description**: Handle the total amount status.
- **Usage**: `flow total-amount status [flags]`
- **Available Commands**:
  - `active`: Make the total amount active
  - `inactive`: Make the total amount inactive
  - `check`: Check the status of the total amount
- **Flags**:
  - `-h, --help`: Display help information

### 24. `flow total-amount status active`
- **Description**: Make the total amount active.
- **Usage**: `flow total-amount status active [flags]`
- **Flags**:
  - `-h, --help`: Display help information

### 25. `flow total-amount status inactive`
- **Description**: Make the total amount inactive.
- **Usage**: `flow total-amount status inactive [flags]`
- **Flags**:
  - `-h, --help`: Display help information

### 26. `flow total-amount status check`
- **Description**: Check the status of total amount.
- **Usage**: `flow total-amount status check [flags]`
- **Flags**:
  - `-h, --help`: Display help information

### 27. `flow total-amount update`
- **Description**: Update the total amount data.
- **Usage**: `flow total-amount update [flags]`
- **Flags**:
  - `-a, --amount string`: Write the total amount that you want to update
  - `-l, --label string`: Write the label that you want to update
  - `-n, --new-category string`: Write the new category to update with
  - `-o, --old-category string`: Write the old category that you want to update
  - `-h, --help`: Display help information

### 28. `flow total-amount view`
- **Description**: View the total amount data.
- **Usage**: `flow total-amount view [flags]`
- **Available Commands**:
  - `amount`: View the total amount
  - `categories`: View the categories in total amount
- **Flags**:
  - `-h, --help`: Display help information

### 29. `flow total-amount remove`
- **Description**: Remove the total amount data.
- **Usage**: `flow total-amount remove [flags]`
- **Flags**:
  - `-c, --category string`: Write the category to remove it's date
  - `-h, --help`: Display help information