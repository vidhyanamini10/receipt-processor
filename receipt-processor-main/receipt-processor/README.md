# Receipt Processor

Receipt Processor is a simple Go-based application that processes receipts and calculates reward points based on predefined rules.

## Prerequisites

Before you can run the application, ensure you have the following prerequisites installed:

- Go (Golang)
- Docker (Optional, for running the application in a Docker container)

## Getting Started...

### 1. Clone the Repository

Clone this repository to your local machine:

```shell
git clone https://github.com/Saireddy09/receipt-processor.git
cd receipt-processor
```

### 2. Build the Application (Go Installation)

If you have Go installed locally, you can build and run the application directly.

```shell
go build -o receipt-processor
./receipt-processor
```

### 3. Build and Run with Docker (Optional)

If you prefer to use Docker, you can build and run the application in a Docker container.

```shell
docker build -t receipt-processor .
docker run -p 8080:8080 receipt-processor
```

The application will be accessible at `http://localhost:8080`.

## Usage

1. Submit a receipt for processing by making a POST request to `http://localhost:8080/receipts/process` with a JSON payload containing the receipt details. See the API documentation for details on the request format.

2. Retrieve the points awarded for a receipt by making a GET request to `http://localhost:8080/receipts/{id}/points`, replacing `{id}` with the ID returned when processing the receipt.

## API Documentation

For detailed API documentation, refer to the `api.yml` file in the project root.