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
  - `transaction`: Transaction service
- **Flags**:
  - `-h, --help`: Display help information
  - `-t, --toggle`: Display help message for toggle
  - `-v, --version`: Display version information

### 2. `flow budget`
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

### 3. `flow budget adjust`
- **Description**: Adjusts the budget details such as changing the amount or category name.
- **Usage**: `flow budget adjust [flags]`
- **Flags**:
  - `-a, --amount string`: New amount of the category to adjust
  - `-n, --newcategory string`: New category name to allocate
  - `-o, --oldcategory string`: Old category name to adjust

### 4. `flow budget alert`
- **Description**: Command for the alert notification.
- **Usage**: 
  - `flow budget alert [flags]`
  - `flow budget alert [command]`
- **Available Commands**:
  - `msg`: The message of alert notifications
  - `setup`: Setup for alert notification

### 5. `flow budget alert setup`
- **Description**: Sets up alerts to get notifications when the budget (for a specific category amount or the whole amount) is passed.
- **Usage**: 
  - `flow budget alert setup [flags]`
- **Flags**:
  - `-c, --category string`: Category name to monitor for budget
  - `-f, --frequency string`: Frequency of notifications (e.g., hourly, daily, weekly, monthly)
  - `-m, --method string`: Preferred method of notification [email or CLI] message
  - `-t, --total string`:  Total budget amount to set the alert

### 6. `flow budget alert msg`
- **Description**: Show the alert messages that were set.
- **Usage**: 
  - `flow budget alert msg [flags]`
- **Flags**:
  - `-h, --help`: Display help information

### 7. `flow budget create`
- **Description**: Creates a budget for different spending categories.
- **Usage**: `flow budget create [flags]`
- **Flags**:
  - `-a, --amount string`: Total amount for the category
  - `-c, --category string`: Category name (e.g., groceries, utilities, etc.)

### 8. `flow budget get`
- **Description**: Retrieves the budget data in CSV format.
- **Usage**: `flow budget get [flags]`
- **Flags**:
  - `-n, --filename string`: CSV file name to store the data
  - `-p, --filepath string`: File path to store the data

### 9. `flow budget remove`
- **Description**: Removes budget details for a specific category.
- **Usage**: `flow budget remove [flags]`
- **Flags**:
  - `-c, --category string`: Category name to remove

### 10. `flow budget view`
- **Description**: Views the budget details, optionally for a specific category.
- **Usage**: `flow budget view [flags]`
- **Flags**:
  - `-c, --category string`: Category name to show specific details

### 11. `flow completion`
- **Description**: Generates autocompletion scripts for various shells.
- **Usage**: `flow completion [command]`
- **Available Commands**:
  - `bash`: Generates autocompletion script for Bash
  - `fish`: Generates autocompletion script for Fish
  - `powershell`: Generates autocompletion script for PowerShell
  - `zsh`: Generates autocompletion script for Zsh
- **Flags**:
  - `-h, --help`: Display help information

### 12. `flow help`
- **Description**: Provides help about the Flow application or specific commands.
- **Usage**: `flow help [command]`

### 13. `flow transaction`
- **Description**: Provides transaction services (not further detailed in provided output).