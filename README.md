
## Getting Started

### Installation
The Money Transfer API will be available at `http://localhost:8090`.

### Configuration

Configuration is managed using environment variables. You can set these variables in the [`.env`](command:_github.copilot.openRelativePath?%5B%7B%22scheme%22%3A%22file%22%2C%22authority%22%3A%22%22%2C%22path%22%3A%22%2FUsers%2Fhafiedh%2Fgo%2Fsrc%2Fside_projects%2Fmoney-transfer%2F.env%22%2C%22query%22%3A%22%22%2C%22fragment%22%3A%22%22%7D%2C%22cb46b11c-ca19-40e9-ba04-d48960866c0c%22%5D "/Users/hafiedh/go/src/side_projects/money-transfer/.env") file.

### Database

The project uses PostgreSQL as the database. The database schema can be found in the [`money-transfer.sql`](command:_github.copilot.openRelativePath?%5B%7B%22scheme%22%3A%22file%22%2C%22authority%22%3A%22%22%2C%22path%22%3A%22%2FUsers%2Fhafiedh%2Fgo%2Fsrc%2Fside_projects%2Fmoney-transfer%2Fmoney-transfer.sql%22%2C%22query%22%3A%22%22%2C%22fragment%22%3A%22%22%7D%2C%22cb46b11c-ca19-40e9-ba04-d48960866c0c%22%5D "/Users/hafiedh/go/src/side_projects/money-transfer/money-transfer.sql") file.

### Running the Application

To run the application locally without Docker, use the following commands:

1. Install dependencies:
    ```sh
    go mod download
    ```

2. Run the application:
    ```sh
    go run main.go
    ```

### Project Structure

- [`cmd/`](command:_github.copilot.openRelativePath?%5B%7B%22scheme%22%3A%22file%22%2C%22authority%22%3A%22%22%2C%22path%22%3A%22%2FUsers%2Fhafiedh%2Fgo%2Fsrc%2Fside_projects%2Fmoney-transfer%2Fcmd%2F%22%2C%22query%22%3A%22%22%2C%22fragment%22%3A%22%22%7D%2C%22cb46b11c-ca19-40e9-ba04-d48960866c0c%22%5D "/Users/hafiedh/go/src/side_projects/money-transfer/cmd/"): Entry point of the application.
- [`internal/config/`](command:_github.copilot.openRelativePath?%5B%7B%22scheme%22%3A%22file%22%2C%22authority%22%3A%22%22%2C%22path%22%3A%22%2FUsers%2Fhafiedh%2Fgo%2Fsrc%2Fside_projects%2Fmoney-transfer%2Finternal%2Fconfig%2F%22%2C%22query%22%3A%22%22%2C%22fragment%22%3A%22%22%7D%2C%22cb46b11c-ca19-40e9-ba04-d48960866c0c%22%5D "/Users/hafiedh/go/src/side_projects/money-transfer/internal/config/"): Configuration management.
- [`internal/domain/entities/`](command:_github.copilot.openRelativePath?%5B%7B%22scheme%22%3A%22file%22%2C%22authority%22%3A%22%22%2C%22path%22%3A%22%2FUsers%2Fhafiedh%2Fgo%2Fsrc%2Fside_projects%2Fmoney-transfer%2Finternal%2Fdomain%2Fentities%2F%22%2C%22query%22%3A%22%22%2C%22fragment%22%3A%22%22%7D%2C%22cb46b11c-ca19-40e9-ba04-d48960866c0c%22%5D "/Users/hafiedh/go/src/side_projects/money-transfer/internal/domain/entities/"): Domain entities.
- [`internal/domain/repositories/`](command:_github.copilot.openRelativePath?%5B%7B%22scheme%22%3A%22file%22%2C%22authority%22%3A%22%22%2C%22path%22%3A%22%2FUsers%2Fhafiedh%2Fgo%2Fsrc%2Fside_projects%2Fmoney-transfer%2Finternal%2Fdomain%2Frepositories%2F%22%2C%22query%22%3A%22%22%2C%22fragment%22%3A%22%22%7D%2C%22cb46b11c-ca19-40e9-ba04-d48960866c0c%22%5D "/Users/hafiedh/go/src/side_projects/money-transfer/internal/domain/repositories/"): Repository interfaces and implementations.
- [`internal/infrastructure/container/`](command:_github.copilot.openRelativePath?%5B%7B%22scheme%22%3A%22file%22%2C%22authority%22%3A%22%22%2C%22path%22%3A%22%2FUsers%2Fhafiedh%2Fgo%2Fsrc%2Fside_projects%2Fmoney-transfer%2Finternal%2Finfrastructure%2Fcontainer%2F%22%2C%22query%22%3A%22%22%2C%22fragment%22%3A%22%22%7D%2C%22cb46b11c-ca19-40e9-ba04-d48960866c0c%22%5D "/Users/hafiedh/go/src/side_projects/money-transfer/internal/infrastructure/container/"): Dependency injection container.
- [`internal/infrastructure/postgres/`](command:_github.copilot.openRelativePath?%5B%7B%22scheme%22%3A%22file%22%2C%22authority%22%3A%22%22%2C%22path%22%3A%22%2FUsers%2Fhafiedh%2Fgo%2Fsrc%2Fside_projects%2Fmoney-transfer%2Finternal%2Finfrastructure%2Fpostgres%2F%22%2C%22query%22%3A%22%22%2C%22fragment%22%3A%22%22%7D%2C%22cb46b11c-ca19-40e9-ba04-d48960866c0c%22%5D "/Users/hafiedh/go/src/side_projects/money-transfer/internal/infrastructure/postgres/"): PostgreSQL database connection.
- `pkg/`: Shared constants and utilities.
- `server/`: HTTP server and handlers.
- `usecase/`: Use case implementations.
- `mocks`: Mocks for testing.
postman.mocks.url=https://3f3cb471-213f-445e-a1b3-b21b6aecaed0.mock.pstmn.io
postman.mocks.checkAccount=/banks/account-check
postman.mocks.transfer=/banks/transfer