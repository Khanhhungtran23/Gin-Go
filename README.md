# Gin/Go

This is a simple CRUD application built with [Gin](https://github.com/gin-gonic/gin) and [GORM](https://gorm.io/).

## Prerequisites

Before you begin, ensure you have met the following requirements:
- You have installed [Go](https://golang.org/doc/install) (version 1.16 or later).
- You have installed [Docker](https://docs.docker.com/get-docker/).

## Installation

1. Clone the repository:
    ```sh
    git clone https://github.com/Khanhhungtran23/Gin-Go.git
    ```

2. Initialize and tidy up Go modules:
    ```sh
    go mod tidy
    ```

## Configuration

1. Set up and connect your database of your project with the following content:
    ```env
    DB_USERNAME=root
    DB_PASSWORD= 'yourpassword'
    DB_HOST=localhost
    DB_PORT=3306
    DB_NAME=gin_database
    ```
2. Start a MySQL container using Docker:
    ```sh
    docker run --name gin-mysql -e MYSQL_ROOT_PASSWORD=yourpassword -e MYSQL_DATABASE=gin_practice -p 3306:3306 -d mysql:latest
    ```
In case you do not use Docker, open your MySQL client (e.g., MySQL Workbench, command line).

## Running the Application

1. Run the application:
    ```sh
    go run main.go
    ```

2. You should see:
    ```plaintext
    Server is running on port 8080
    ```
It show that you run succesfull.

## Performance test results (Benchmark)
<img width="865" alt="image" src="https://github.com/Khanhhungtran23/Gin-Go/assets/111229500/3a6bfb7f-9573-42d9-a4b0-bc883bef9e4e">

