# Uber-like Ride-Hailing Service API

[![Go Version](https://img.shields.io/badge/Go-1.21%2B-blue.svg)](https://golang.org/)
[![MongoDB](https://img.shields.io/badge/MongoDB-6.0%2B-green.svg)](https://www.mongodb.com/)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

A high-performance backend API for an Uber-like ride-hailing service with real-time features and payment integration.

---

## Table of Contents

1. [Features](#features)
2. [Tech Stack](#tech-stack)
3. [Installation](#installation)
   - [Prerequisites](#prerequisites)
   - [Steps](#steps)
4. [API Documentation](#api-documentation)
   - [Authentication Endpoints](#authentication-endpoints)
   - [Rides Endpoints](#rides-endpoints)
5. [Deployment](#deployment)
6. [Contributing](#contributing)
7. [License](#license)

---

## Features 🚀

- **User Authentication**: JWT-based authentication for secure login and registration.
- **Real-time Driver Matching**: Utilizes geospatial queries to match drivers with riders in real-time.
- **OTP Verification System**: Verifies users via OTP to start the ride.
- **Stripe Payment Integration**: Secure payment processing through Stripe.
- **Dynamic Surge Pricing**: Automatically adjusts pricing based on demand and location.
- **WebSocket Notifications**: Real-time notifications for ride updates.
- **Ride Tracking**: Track rides with multiple statuses like "Pending", "In Progress", and "Completed."
- **Cancellation Management**: Handle ride cancellations and related logic.
- **Driver/Rider Feedback System**: Collect and manage user feedback after each ride.

---

## Tech Stack 💻

- **Backend**: Go 1.21+
- **Framework**: Gin
- **Database**: MongoDB
- **Authentication**: JWT
- **Real-time Communication**: Gorilla WebSocket
- **Payments**: Stripe API
- **Routing**: Vincenty's Algorithm + DistanceMatrix.ai API

---

## Installation ⚙️

### Prerequisites

Before setting up the project, ensure you have the following installed:
- Go 1.21+
- MongoDB 6.0+
- Stripe account for payment integration
- DistanceMatrix.ai API key for route calculation

### Steps

1. **Clone the repository:**

    ```bash
    git clone https://github.com/yeliwzinn/uber-clone-api.git
    cd uber-clone-api
    ```

2. **Install dependencies:**

    ```bash
    go mod download
    ```

3. **Set up environment variables:**

    Create a `.env` file by copying from the example:

    ```
    cp .env.example .env
    ```
    ```
   MONGODB_URI=mongodb://localhost:27017/
   JWT_SECRET=your_jwt_secret_key
   DISTANCEMATRIXAI_API_KEY=your_api_key
   STRIPE_SECRET_KEY=your_stripe_secret_key
   DB_NAME=uber_clone
   ```

    Ensure you fill in your credentials (Stripe keys, MongoDB URI, etc.) in the `.env` file.

4. **Start the server:**

    ```bash
    go run main.go
    ```

---

## API Documentation 📚

### Authentication Endpoints

| Endpoint    | Method | Description               |
|-------------|--------|---------------------------|
| `/signup`   | POST   | Register a new user       |
| `/login`    | POST   | Authenticate a user       |

### Rides

| Method | Endpoint                    | Description                     |
|--------|-----------------------------|---------------------------------|
| POST   | /rides/                     | Request new ride                |
| GET    | /rides/{ride_id}            | Get ride details                |
| POST   | /rides/{ride_id}/verifyOTP  | Verify OTP to start ride        |
| POST   | /rides/{ride_id}/complete   | Mark ride as completed          |
| POST   | /rides/{ride_id}/cancel     | Cancel a ride                   |

### Payments

| Method | Endpoint                    | Description                     |
|--------|-----------------------------|---------------------------------|
| POST   | /rides/{ride_id}/pay        | Initiate payment                |
| POST   | /rides/{ride_id}/confirm-payment | Confirm payment completion  |

### Feedback

| Method | Endpoint                    | Description                     |
|--------|-----------------------------|---------------------------------|
| POST   | /feedback/{ride_id}         | Submit ride feedback            |

---

## WebSocket Support

The API provides real-time notifications through WebSocket connections. Clients can connect to `ws://localhost:8080/ws` after authentication.

### Supported notification types:

- **ride_request**: New ride request for drivers
- **ride_response**: Driver acceptance/rejection for riders
- **payment_requested**: Payment initiation alerts
- **ride_status**: Updates on ride state changes

---

## Deployment 🚢

To deploy the project using Docker, follow these steps:

1. **Create the Dockerfile:**

    ```dockerfile
    # Dockerfile
    FROM golang:1.21-alpine
    WORKDIR /app
    COPY . .
    RUN go mod download && go build -o uber-clone .
    CMD ["./uber-clone"]
    ```

2. **Build and run the Docker container:**

    ```bash
    docker build -t uber-clone .
    docker run -p 8080:8080 --env-file .env uber-clone
    ```

---

## Contributing 🤝

We welcome contributions to this project! Here's how you can help:

1. Fork the repository.
2. Create a new branch for your feature (`git checkout -b feature/AmazingFeature`).
3. Make your changes and commit them (`git commit -m 'Add some feature'`).
4. Push to your branch (`git push origin feature/AmazingFeature`).
5. Open a Pull Request with a description of your changes.

---

## License 📄

This project is distributed under the MIT License. See the [LICENSE](LICENSE) file for more details.

---
