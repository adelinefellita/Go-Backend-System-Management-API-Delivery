# System Management API Delivery

This project is a backend system for managing delivery addresses, built using the **Golang GIN framework** and **MongoDB** as the database. The system includes user roles (managers and couriers), allowing managers to manage delivery data and couriers to update delivery statuses with proof of delivery.

## Features

- **Role-based Access:**
  - **Managers:** 
    - Add, read, edit, and delete delivery addresses.
    - Default delivery status: "In Transit" (`sedang dalam pengiriman`).
  - **Couriers:** 
    - View assigned delivery addresses.
    - Update delivery status to "Completed" (`pengiriman selesai`) with proof of delivery.
- **RESTful API**: Endpoints for managing delivery data and user authentication.
- **MongoDB Integration**: Stores and retrieves delivery and user data efficiently.

## Documentation

The full API documentation is available on [Postman Documentation](https://documenter.getpostman.com/view/31688603/2sAYJ3G2Mo).

## Technologies Used

- **Golang**: GIN framework for building RESTful APIs.
- **MongoDB**: NoSQL database for storing user and delivery data.

## Installation and Setup

### Prerequisites

- **Go**: [Install Go](https://golang.org/doc/install).
- **MongoDB**: [Install MongoDB](https://www.mongodb.com/docs/manual/installation/).

### Steps

1. Clone this repository:
   ```bash
   git clone https://github.com/yourusername/Go-Backend-System-Management-API-Delivery.git
   cd Go-Backend-System-Management-API-Delivery

2. Install dependencies:
   ```go mod tidy

3. Set up environment variables in .env file:
   ```MONGO_URI=mongodb://localhost:27017
    DATABASE_NAME = admin
    JWT_SECRET = your_jwt_secret_key

4. Run the application:
   ```go run main.go

5. Access the API at http://localhost:8080

## API Endpoints

### Authentication:
- POST /login: User login.
- POST /register: User registration (restricted to managers only).

### Delivery Management:
- POST /deliveries: Add a new delivery address (managers only).
- GET /deliveries: View all delivery addresses (managers and couriers).
- PUT /deliveries/:id: Update delivery status or details (role-specific actions).
- DELETE /deliveries/:id: Delete a delivery (managers only).

Refer to the Postman Documentation for detailed request and response examples.

## Checking Data in MongoDB

To verify that the data has been successfully entered into the database, you can check it in MongoDB Compass. 

1. Open MongoDB Compass and connect to your MongoDB instance.
2. Select the `admin` database.
3. You should see the following collections:
   - `deliveries`
   - `addresses`
   - `person`
