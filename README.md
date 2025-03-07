# POS Backend System

A complete Point of Sale (POS) backend system built with Go (Golang) and Gin framework. This backend provides all the essential features needed for a modern POS system, including authentication, product management, and sales transaction processing.

## Features

- **Authentication** - Secure JWT-based authentication system
- **Product Management** - Complete CRUD operations for product inventory
- **Sales Transactions** - Record and retrieve sales data
- **Image Storage** - Local storage for product images
- **Authorization Middleware** - Route protection with JWT validation
- **Containerization** - Docker and Kubernetes (Minikube) support
- **Monitoring** - Prometheus and Grafana integration
- **CI/CD** - Jenkins pipeline configuration

## Project Structure

```
pos-backend/
│── main.go
│── config/
│   ├── database.go
│── models/
│   ├── user.go
│   ├── product.go
│   ├── sale.go
│── routes/
│   ├── user_routes.go
│   ├── product_routes.go
│   ├── sale_routes.go
│── controllers/
│   ├── auth_controller.go
│   ├── product_controller.go
│   ├── sale_controller.go
│── middlewares/
│   ├── jwt_middleware.go
│── storage/
│   ├── uploads/ (Stores product images)
│── .env
│── go.mod
│── go.sum
│── Dockerfile
│── k8s/
│   ├── deployment.yaml
│   ├── service.yaml
│   ├── prometheus.yaml
│   ├── grafana.yaml
│   ├── jenkins.yaml
```

## Prerequisites

- Go 1.19 or later
- PostgreSQL
- Docker (optional, for containerization)
- Kubernetes/Minikube (optional, for deployment)

## Installation

1. Clone the repository:
   ```
   git clone https://github.com/yourusername/pos-backend.git
   cd pos-backend
   ```

2. Install dependencies:
   ```
   go mod tidy
   ```

3. Set up environment variables in `.env` file:
   ```
   PORT=8080
   DATABASE_URL=postgres://user:password@localhost:5432/pos_db
   ```

4. Create the PostgreSQL database:
   ```
   createdb pos_db
   ```

## Running the Application

### Local Development

Run the application locally:
```
go run main.go
```

The server will start on http://localhost:8080 (or the port specified in your .env file).

### Using Docker

1. Build the Docker image:
   ```
   docker build -t pos-backend:latest .
   ```

2. Run the container:
   ```
   docker run -p 8080:8080 -d pos-backend:latest
   ```

### Using Kubernetes (Minikube)

1. Start Minikube:
   ```
   minikube start
   ```

2. Deploy the application:
   ```
   kubectl apply -f k8s/deployment.yaml
   kubectl apply -f k8s/service.yaml
   ```

3. For monitoring, deploy Prometheus and Grafana:
   ```
   kubectl apply -f k8s/prometheus.yaml
   kubectl apply -f k8s/grafana.yaml
   ```

## API Endpoints

### Authentication
- `POST /register` - Register a new user
- `POST /login` - Login and receive JWT token

### Products
- `GET /products` - Get all products
- `POST /products` - Create a new product (requires authentication)

### Sales
- `GET /sales` - Get all sales (requires authentication)
- `POST /sales` - Create a new sale (requires authentication)

## Monitoring

The application includes Prometheus metrics and Grafana dashboards for monitoring performance and usage.

To access:
1. Prometheus: `http://localhost:9090` (after port-forwarding)
2. Grafana: `http://localhost:3000` (after port-forwarding)

## CI/CD

The project includes a Jenkins configuration file for continuous integration and deployment.

To set up Jenkins:
```
kubectl apply -f k8s/jenkins.yaml
```

## Future Improvements

- Add user roles and permissions
- Implement more comprehensive authentication
- Add reporting features
- Develop a front-end client

## License

[MIT License](LICENSE)