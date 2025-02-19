# Mock API Server

This project is a **Mock API Server** built with **Go (Golang)** and **MongoDB**. It allows you to define custom API endpoints dynamically, store them in MongoDB, and serve mock responses based on the stored configurations.

## Overview
The Mock API Server is a Golang-based server that allows developers to create and manage mock API endpoints dynamically. It integrates with MongoDB for persistent storage of endpoints and includes a frontend for managing mock APIs.


## Features
- **Dynamic API Routing**: Define API endpoints dynamically.
- **MongoDB Storage**: Store and retrieve mock API configurations.
- **File Upload Support**: Upload mock routes via JSON files.
- **Dockerized Deployment**: Run the project using Docker and Docker Compose.
- **Frontend**: A React-based frontend for easy management

## Project Structure
```
mock-api-server/
├── Dockerfile
├── README.md
├── config
│   └── db.go
├── docker-compose.yml
├── go.mod
├── go.sum
├── handlers
│   └── endpoints.go
├── helper_files
│   ├── Mock GO API.postman_collection.json
│   └── mock_routes.json
├── main.go
├── middlewares
│   └── logging.go
├── mock_server_frontend
│   ├── Dockerfile
│   ├── README.md
│   ├── index.html
│   ├── package.json
│   ├── public
│   ├── src
│   ├── tailwind.config.js
│   ├── vite.config.js
│   └── yarn.lock
├── models
│   ├── endpoints.go
│   └── request_log.go
└── routes
    └── router.go
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
    docker-compose down --remove-orphans
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
### 3. Delete a Mock API
- **Example Request**:
  ```sh
  curl -X DELETE http://localhost:3002/{id}/
  ```

## Troubleshooting
### "Invalid WriteHeader code 0" Error
- Ensure that all mock endpoints have a valid `status_code` (e.g., `200`).
- Update the `ServeMockEndpoint` function to use a default status code if `0` is found.

## Technologies Used
- Backend: Golang (mux, MongoDB driver)
- Database: MongoDB
- Frontend: React, Vite, Tailwind CSS
- Containerization: Docker, Docker Compose

## Contributing
Feel free to submit issues and pull requests to improve this project!
