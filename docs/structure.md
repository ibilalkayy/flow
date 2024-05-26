    ./
    ├── .github
    │   └── funding.yml
    ├── CODE_OF_CONDUCT.md
    ├── CONTRIBUTING.md
    ├── Dockerfile
    ├── LICENSE
    ├── README.md
    ├── docker-compose.yml
    ├── go.mod
    ├── go.sum
    ├── main.go
    ├── cmd
    │   ├── budget
    │   │   ├── handler
    │   │   │   ├── alert.go
    │   │   │   ├── create.go
    │   │   │   ├── get.go
    │   │   │   ├── remove.go
    │   │   │   ├── update.go
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
    │   ├── root.go
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
    │   └── conversion.go
    ├── docs
    │   ├── commands.md
    │   ├── structure.md
    ├── entities
    │   ├── alert.go
    │   ├── budget.go
    │   ├── email.go
    │   ├── history.go
    │   ├── init.go
    │   ├── spend.go
    │   ├── total_amount.go
    ├── framework
    │   ├── db
    │   │   ├── alert_db
    │   │   │   └── alert_db.go
    │   │   ├── budget_db
    │   │   │   ├── budget_db.go
    │   │   │   ├── handler.go
    │   │   │   └── history_db.go
    │   │   ├── connection.go
    │   │   ├── migrations
    │   │   │   ├── 001_create_budget_table.sql
    │   │   │   ├── 002_create_alert_table.sql
    │   │   │   └── 003_create_total_amount_table.sql
    │   │   ├── total_amount_db
    │   │       ├── handler.go
    │   │       ├── total_amount_categories.go
    │   │       └── total_amount_db.go
    │   ├── email
    │   │   ├── email.go
    │   │   └── templates
    │   │       └── alert.html
    ├── handler
    │   └── handler.go
    ├── interfaces
    │   └── interfaces.go
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
    │   ├── middleware
    │       └── env_loader.go
    └── tests
        ├── app
            ├── alert
            │   └── alert_test.go
            └── budget
                └── budget_test.go

## **File Details:**

### **Root Files**

- **main.go:** Main entry point of the application.
- **CODE_OF_CONDUCT.md:** Code of conduct for contributors.
- **CONTRIBUTING.md:** Guidelines for contributing to the project.
- **Dockerfile:** Defines instructions to build Docker image of the application.
- **LICENSE:** License information for the project.
- **README.md:** Main documentation file providing information about the project.
- **docker-compose.yml:** Configuration file for Docker Compose.
- **go.mod:** Go module file specifying dependencies.
- **go.sum:** Go module file specifying exact versions of dependencies.

### **Command Files**

#### **Root Command**

- **cmd/root.go:** Root command handling the overall application logic.

#### **Budget Command**

- **cmd/budget/main.go:** Budget service entry point.
- **cmd/budget/handler/alert.go:** Handler for alerting.
- **cmd/budget/handler/create.go:** Handler for creating budget.
- **cmd/budget/handler/get.go:** Handler for getting budget.
- **cmd/budget/handler/remove.go:** Handler for removing budget.
- **cmd/budget/handler/update.go:** Handler for updating budget.
- **cmd/budget/handler/view.go:** Handler for viewing budget.

#### **Budget Sub-command**

- **cmd/budget/sub_handler/msg.go:** Handler for showing alert messages.
- **cmd/budget/sub_handler/setup.go:** Handler for setting up alert values.
- **cmd/budget/sub_handler/remove.go:** Handler for removing alert values.
- **cmd/budget/sub_handler/update.go:** Handler for updating alert values.
- **cmd/budget/sub_handler/view.go:** Handler for viewing alert values.

#### **Init Command**

- **cmd/init/main.go:** Flow initialization functionality.

#### **Spend Command**

- **cmd/spend/main.go:** Spending money on various categories.
- **cmd/spend/handler/history.go:** Show the transaction history.

#### **Spend Sub-command**

- **cmd/spend/sub_handler/remove.go:** Remove the history data.
- **cmd/spend/sub_handler/show.go:** Show the history data.

#### **Total Amount Command**

- **cmd/total_amount/main.go:** The management of the total amount to set the target.
- **cmd/total_amount/handler/add.go:** Handler for adding the total amount.
- **cmd/total_amount/handler/remove.go:** Handler for removing the total amount data.
- **cmd/total_amount/handler/status.go:** Handler for handling the total amount's status.
- **cmd/total_amount/handler/update.go:** Handler for updating the total amount data.
- **cmd/total_amount/handler/view.go:** Handler for viewing the total amount data.

#### **Total Amount Sub-command**

- **cmd/total_amount/sub_handler/active.go:** Handler for making the total amount status active.
- **cmd/total_amount/sub_handler/inactive.go:** Handler for making the total amount status inactive.
- **cmd/total_amount/sub_handler/categories.go:** Handler for handling categories related to the total amount.
- **cmd/total_amount/sub_handler/amount.go:** Handler for managing the amount-related logic.
- **cmd/total_amount/sub_handler/check.go:** Handler for checking the total amount data.

### **Common Files**

- **common/conversion.go:** Utility functions for data conversion.

### **Documentation Files**

- **docs/commands.md:** Commands of the whole application.
- **docs/structure.md:** Structure of the whole application.

### **Entity Files**

- **entities/alert.go:** Entity representing an alert.
- **entities/budget.go:** Entity representing a budget.
- **entities/email.go:** Entity representing an email.
- **entities/history.go:** Entity representing transaction history.
- **entities/init.go:** Entity representing initialization data.
- **entities/spend.go:** Entity representing spending data.
- **entities/total_amount.go:** Entity representing total amount data.

### **Framework Files**

#### **Database Files**

- **framework/db/connection.go:** Database connection setup.
- **framework/db/migrations/001_create_budget_table.sql:** SQL script for creating the budget table.
- **framework/db/migrations/002_create_alert_table.sql:** SQL script for creating the alert table.
- **framework/db/migrations/003_create_total_amount_table.sql:** SQL script for creating the total amount table.
- **framework/db/alert_db/alert_db.go:** CRUD operations for alert functionality.
- **framework/db/budget_db/budget_db.go:** CRUD operations for budget functionality.
- **framework/db/budget_db/

handler.go:** Handler for budget-related database operations.
- **framework/db/budget_db/history_db.go:** CRUD operations for budget history functionality.
- **framework/db/total_amount_db/handler.go:** Handler for total amount-related database operations.
- **framework/db/total_amount_db/total_amount_categories.go:** CRUD operations for total amount categories functionality.
- **framework/db/total_amount_db/total_amount_db.go:** CRUD operations for total amount functionality.

#### **Email Files**

- **framework/email/email.go:** Handling email functionality.
- **framework/email/templates/alert.html:** Email template for alert notifications.

### **Handler Files**

- **handler/handler.go:** General request handler logic.

### **Interface Files**

- **interfaces/interfaces.go:** Interface definitions for different modules.

### **Use Case Files**

#### **Application Logic**

- **usecases/app/alert/alert.go:** Logic for alert management functionality.
- **usecases/app/budget/budget.go:** Logic for budget management functionality.
- **usecases/app/init/init.go:** Logic for init functionality.
- **usecases/app/spend/history.go:** Logic for handling transaction history.
- **usecases/app/spend/notification.go:** Functions for setting notifications.
- **usecases/app/spend/spend.go:** Logic for transaction functionality.
- **usecases/app/total_amount/total_amount.go:** Logic for handling total amount data.

#### **Middleware**

- **usecases/middleware/env_loader.go:** Environment middleware for handling environment variables.

### **Testing Files**

- **tests/app/alert/alert_test.go:** Test file for handling alert test functions.
- **tests/app/budget/budget_test.go:** Test file for handling budget test functions.