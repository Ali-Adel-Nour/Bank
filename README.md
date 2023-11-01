# Bank Web Service

This repository contains the backend web service for a simple bank, developed following the Udemy course [link to the course here](https://bit.ly/backendudemy). This service provides APIs for the frontend to manage bank accounts, record balance changes, and perform money transfers.

## Features

- **Account Management:** Create, retrieve, update, and delete bank accounts.
- **Balance Tracking:** Record and view all balance changes for each account.
- **Money Transfer:** Transfer funds securely between two accounts.

## Technologies Used

- **Go:** The backend is written in Go, providing a fast and efficient server-side application.
- **PostgreSQL:** Data is stored in a PostgreSQL database, ensuring data consistency and durability.
- **Gin:** Gin is used as the HTTP web framework for routing and handling HTTP requests.
- **SQLC:** SQLC is a Go library for writing SQL queries using Go code. It helps in type-safe SQL interactions with the database.
- **Docker:** Docker is used for containerization, allowing easy deployment and scaling of the application.
- **Viper:** Viper is utilized for configuration management, enabling dynamic configuration reloads and environment-specific settings.


## Setup Instructions

1. **Clone the Repository:**
   ```bash
   git clone https://github.com/your-username/bank-web-service.git
   cd bank-web-service
   ```

2. **Database Setup:**
   - Ensure you have PostgreSQL installed and running.
   - Create a new database named `simple_bank`.
   - Update the database configuration in `config/config.yaml` with your PostgreSQL credentials.

3. **Dependencies Installation:**
   ```bash
   go mod tidy
   ```

4. **Run the Application:**
   ```bash
   go run main.go
   ```
   The server will start running at `http://localhost:8080`.

## API Endpoints

### Account Management

- **Create Account:**
  - `POST /accounts`
  - Request Body: `{ "owner": "Account Holder Name", "balance": 100.0 }`
  - Response: `{ "id": 1, "owner": "Account Holder Name", "balance": 100.0 }`

- **Get Account by ID:**
  - `GET /accounts/:id`
  - Response: `{ "id": 1, "owner": "Account Holder Name", "balance": 100.0 }`

- **Update Account:**
  - `PUT /accounts/:id`
  - Request Body: `{ "balance": 200.0 }`
  - Response: `{ "id": 1, "owner": "Account Holder Name", "balance": 200.0 }`

- **Delete Account:**
  - `DELETE /accounts/:id`

### Balance Tracking

- **Get Balance History for Account:**
  - `GET /accounts/:id/balance-history`
  - Response: `[{ "id": 1, "account_id": 1, "amount": 100.0, "transaction_type": "credit" }]`

### Money Transfer

- **Transfer Money:**
  - `POST /transfers`
  - Request Body: `{ "from_account_id": 1, "to_account_id": 2, "amount": 50.0 }`
  - Response: `{ "id": 1, "from_account_id": 1, "to_account_id": 2, "amount": 50.0 }`

## Error Handling

The API follows RESTful principles and returns appropriate HTTP status codes and error messages in case of failures.

- `400 Bad Request`: Invalid request format or missing parameters.
- `404 Not Found`: Requested resource (account, transaction, etc.) not found.
- `500 Internal Server Error`: Unexpected server errors.

## Contribution

Contributions are welcome! Feel free to open issues or create pull requests to improve the code or add new features.

## License

This project is licensed under the [MIT License](LICENSE).

---

*Note: Modify this README according to your project's specific features and requirements.*
