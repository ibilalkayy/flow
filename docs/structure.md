        ./
    ├── .github
    │   └── funding.yml
    ├── cmd
    │   ├── budget
    │   │   ├── handler
    │   │   │   ├── adjust.go
    │   │   │   ├── alert.go
    │   │   │   ├── create.go
    │   │   │   ├── get.go
    │   │   │   ├── remove.go
    │   │   │   └── view.go
    │   │   ├── main.go
    │   │   └── sub_handler
    │   │       ├── msg.go
    │   │       ├── remove.go
    │   │       ├── setup.go
    │   │       ├── update.go
    │   │       └── view.go
    │   ├── init
    │   │   └── main.go
    │   ├── spend
    │   │   ├── handler
    │   │   │   └── history.go
    │   │   ├── main.go
    │   │   └── sub_handler
    │   │       ├── remove.go
    │   │       └── show.go
    │   └── total_amount
    │       ├── handler
    │       │   ├── add.go
    │       │   ├── remove.go
    │       │   ├── status.go
    │       │   ├── update.go
    │       │   └── view.go
    │       ├── main.go
    │       └── sub_handler
    │           ├── active.go
    │           ├── amount.go
    │           ├── categories.go
    │           ├── check.go
    │           └── inactive.go
    ├── common
    │   └── utils
    │       └── conversion.go
    ├── docs
    │   ├── commands.md
    │   └── structure.md
    ├── entities
    │   ├── alert.go
    │   ├── budget.go
    │   ├── email.go
    │   ├── history.go
    │   ├── init.go
    │   ├── spend.go
    │   └── total_amount.go
    ├── framework_drivers
    │   ├── db
    │   │   ├── connection.go
    │   │   ├── alert_db
    │   │   │   └── alert_db.go
    │   │   ├── budget_db
    │   │   │   ├── budget_db.go
    │   │   │   └── history_db.go
    │   │   ├── migrations
    │   │   │   ├── 001_create_budget_table.sql
    │   │   │   ├── 002_create_alert_table.sql
    │   │   │   └── 003_create_total_amount_table.sql
    │   │   └── total_amount_db
    │   │       ├── total_amount_categories.go
    │   │       └── total_amount_db.go
    │   ├── email
    │   │   ├── email.go
    │   │   └── templates
    │   │       └── alert.html
    ├── tests
    │   └── app
    │       ├── alert
    │       │   └── alert_test.go
    │       └── budget
    │           └── budget_test.go
    ├── usecases
    │   ├── app
    │   │   ├── alert
    │   │   │   └── alert.go
    │   │   ├── budget
    │   │   │   └── budget.go
    │   │   ├── init
    │   │   │   └── init.go
    │   │   ├── spend
    │   │   │   ├── history.go
    │   │   │   ├── notification.go
    │   │   │   └── spend.go
    │   │   └── total_amount
    │   │       └── total_amount.go
    │   └── middleware
    │       └── env_loader.go
    ├── .gitignore
    ├── .dockerignore
    ├── CODE_OF_CONDUCT.md
    ├── CONTRIBUTING.md
    ├── LICENSE
    ├── Dockerfile
    ├── README.md
    ├── docker-compose.yml
    ├── go.mod
    ├── go.sum
    └── main.go

## **File Details:**

### Budget command files

- **cmd/budget/main.go:** Budget service entry point.
- **cmd/budget/handler/adjust.go:** Handler for adjusting budget.
- **cmd/budget/handler/alert.go:** Handler for alerting.
- **cmd/budget/handler/create.go:** Handler for creating budget.
- **cmd/budget/handler/get.go:** Handler for getting budget.
- **cmd/budget/handler/remove.go:** Handler for removing budget.
- **cmd/budget/handler/view.go:** Handler for viewing budget.

### Budget subcommand files

- **cmd/budget/sub_handler/msg.go:** Handler for showing the alert messages.
- **cmd/budget/sub_handler/setup.go:** Handler for setting up the alert values.
- **cmd/budget/sub_handler/remove.go:** Handler for removing the alert values.
- **cmd/budget/sub_handler/update.go:** Handler for updating the alert values.
- **cmd/budget/sub_handler/view.go:** Handler for viewing the alert values.

### Total amount command files

- **cmd/total_amount/main.go:** The management of the total amount to set the target.
- **cmd/total_amount/handler/add.go:** Handler for adding the total amount.
- **cmd/total_amount/handler/remove.go:** Handler for removing the total amount data.
- **cmd/total_amount/handler/status.go:** Handler for handling the total amount's status.
- **cmd/total_amount/handler/update.go:** Handler for updating the total amount data.
- **cmd/total_amount/handler/view.go:** Handler for viewing the total amount data.

### Total amount subcommand files

- **cmd/total_amount/sub_handler/active.go:** Handler for making the total amount status active.
- **cmd/total_amount/sub_handler/inactive.go:** Handler for making the total amount status inactive.
- **cmd/total_amount/sub_handler/categories.go:** Handler for making the total amount status active.
- **cmd/total_amount/sub_handler/amount.go:** Handler for making the total amount status inactive.

### Other command files

- **cmd/init/main.go:** Flow initialization functionality.
- **cmd/spend/main.go:** Spending money on various categories.
- **cmd/spend/handler/history.go:** Show the transaction history.
- **cmd/spend/sub_handler/show.go:** Show the history data.
- **cmd/spend/sub_handler/remove.go:** Remove the history data.

### Database files

- **framework_drivers/db/connection.go:** Database connection setup.
- **framework_drivers/db/migrations/001_create_budget_table.sql:** SQL script for creating budget table.
- **framework_drivers/db/migrations/002_create_alert_table.sql:** SQL script for creating alert table.
- **framework_drivers/db/alert_db/alert_db.go:** CRUD operation for the alert functionality.
- **framework_drivers/db/budget_db/budget_db.go:** CRUD operation for the budget functionality.
- **framework_drivers/db/total_amount_db/total_amount_db.go:** CRUD operation for the total amount functionality.

### Documentation files

- **docs/commands.md:** Commands of the whole application.
- **docs/structure.md:** Structure of the whole application.

### Email files

- **framework_drivers/email/email.go:** Handling the email functionality.
- **framework_drivers/email/templates/alert.html:** Email template for alert notification.

### App logic files

- **usecases/app/alert/alert.go:** Logic for alert management functionality.
- **usecases/app/budget/budget.go:** Logic for budget management functionality.
- **usecases/app/init/init.go:** Logic for init functionality.
- **usecases/app/spend/history.go:** Logic for handling the history.
- **usecases/app/spend/spend.go:** Logic for transaction functionality.
- **usecases/app/spend/notification.go:** Functions for setting the hourly, daily and more notifications.
- **usecases/app/total_amount/total_amount.go:** Logic for handling the total amount data.

### Middleware files

- **usecases/middleware/env_loader.go:** Environment middleware for handling environment variables.

### Testing files

- **tests/app/alert/alert_test.go:** Test file for handling the alert test functions.
- **tests/app/budget/budget_test.go:** Test file for handling the budget test functions.

### Root files

- **.gitignore:** Specifies intentionally untracked files that Git should ignore.
- **.dockerignore:** Specifies files and directories that should be ignored when building Docker images.
- **CODE_OF_CONDUCT.md:** Code of conduct for contributors.
- **CONTRIBUTING.md:** Guidelines for contributing to the project.
- **Dockerfile:** Defines instructions to build Docker image of the application.
- **README.md:** Main documentation file providing information about the project.
- **docker-compose.yml:** Configuration file for Docker Compose.
- **go.mod:** Go module file specifying dependencies.
- **go.sum:** Go module file specifying exact versions of dependencies.
- **main.go:** Main entry point of the application.