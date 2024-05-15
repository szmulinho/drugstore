# Drugstore API

This API allows you to manage medical drugs, providing endpoints for creating, retrieving, updating, and deleting them.

## Endpoints

### Add a New Drug

Endpoint: /drug

Method: POST

Description: Adds a new drug to the system.

### Retrieve a Drug by ID

Endpoint: /drugs/{id}

Method: GET

Description: Retrieves details of a specific drug by its ID.

## Retrieve a Drug by name

Endpoint: /drugs/{name}

Method: GET

Description: Retrieves details of a specific drug by its name.

### Retrieve All Drugs
Endpoint: /drugs
Method: GET
Description: Retrieves all drugs in the system.

### Update a Drug

Endpoint: /drugs/{id}

Method: PATCH

Description: Updates the details of a specific drug by its ID.

### Delete a Prescription

Endpoint: /drug/{id}

Method: DELETE

Description: Deletes a specific drug by its ID.

## Testing

Unit and integration tests are implemented for the API. Tests can be run individually or using Docker Compose.

### Run Tests Individually

To run the tests manually, use the following command:

```go test ./...```

### Run Tests with Docker Compose

A docker-compose.yml file is provided to facilitate testing in a containerized environment. To run the tests using Docker Compose, use the following command:

```docker-compose up``` or ```make doker-tests```

This command will build the Docker images, run the tests, and then stop the containers.

## Setup Instructions

### 1. Clone the repository:

```git clone https://github.com/szmulinho/drugstore-API.git```
```cd drugstore-api```

### 2. Install dependencies:

```go mod tidy```

### 3. Run the server:

```go run main.go```

## Contribution Guidelines

Fork the repository

Create a new branch ```git checkout -b feature/your-feature-name```

Commit your changes ```git commit -m 'Add some feature'```

Push to the branch ```git push origin feature/your-feature-name```

Create a new Pull Request

## License

This project is licensed under the MIT License - see the LICENSE file for details.

For any questions or suggestions, feel free to open an issue or contact the repository owner.



