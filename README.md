# Mock API Server

This project is a **Mock API Server** built with **Go (Golang)** and **MongoDB**. It allows you to define custom API endpoints dynamically, store them in MongoDB, and serve mock responses based on the stored configurations.

## Features
- **Dynamic API Routing**: Define API endpoints dynamically.
- **MongoDB Storage**: Store and retrieve mock API configurations.
- **File Upload Support**: Upload mock routes via JSON files.
- **Dockerized Deployment**: Run the project using Docker and Docker Compose.

## Project Structure
```
mock-api-server/
│── config/         # Configuration for database and environment variables
│── handlers/       # API endpoint handlers
│── middlewares/    # Logging and request-handling middleware
│── models/         # Data models
│── routes/         # Router setup
│── main.go         # Main entry point
│── Dockerfile      # Docker configuration
│── docker-compose.yml # Docker Compose configuration
│── README.md       # Project documentation
```

## Prerequisites
- Docker
- Docker Compose

## Getting Started
### 1. Clone the Repository
```sh
git clone https://github.com/Shashwat-Sharma3108/mock-api-go-lang.git
cd mock-api-server
```

### 2. Set Up Environment Variables
Create a `.env` file in the project root and define the necessary environment variables:
```
MONGO_URI=mongodb://admin:admin@mongo:27017/
PORT=3002
DB_NAME=MOCK_DB
```

### 3. Build and Run with Docker
#### Using Docker
1. Build the Docker image:
    ```sh
    docker build -t mock-api-server .
    ```
2. Run the container:
    ```sh
    docker run -p 3002:3002 --env-file .env mock-api-server
    ```

#### Using Docker Compose
1. Start the containers (API server + MongoDB):
    ```sh
    docker-compose up --build
    ```
2. Stop the containers:
    ```sh
    docker-compose down
    ```

## API Endpoints
### 1. Upload Mock API Endpoints
- **Endpoint**: `POST /upload`
- **Description**: Upload mock routes in JSON format.
- **Example Request (Raw JSON in Postman)**:
  ```json
  [
    {
      "url": "/cartoons/tom-and-jerry",
      "method": "GET",
      "response": {
        "title": "Tom and Jerry",
        "year": 1940,
        "characters": ["Tom", "Jerry", "Spike"],
        "genre": "Comedy"
      },
      "status_code": 200
    }
  ]
  ```

### 2. Fetch Mock API Responses
- **Example Request**:
  ```sh
  curl -X GET http://localhost:3002/cartoons/tom-and-jerry
  ```
- **Example Response**:
  ```json
  {
    "title": "Tom and Jerry",
    "year": 1940,
    "characters": ["Tom", "Jerry", "Spike"],
    "genre": "Comedy"
  }
  ```

## Troubleshooting
### "Invalid WriteHeader code 0" Error
- Ensure that all mock endpoints have a valid `status_code` (e.g., `200`).
- Update the `ServeMockEndpoint` function to use a default status code if `0` is found.

## Contributing
Feel free to submit issues and pull requests to improve this project!

## License
This project is licensed under the MIT License.

