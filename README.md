# Go Load Balancer with Gin

A sophisticated load balancer implementation in Go using the Gin framework. This project demonstrates how to efficiently rotate through a series of API keys for external service calls, implementing the round-robin load balancing algorithm.

## 🌟 Features

- **Round-robin load balancing** of API keys
- **Gin-powered HTTP server** with fast and efficient request handling
- **RESTful JSON API** with multiple endpoints
- **Dockerized application** for easy deployment
- **Environment variable configuration** for flexible setup
- **Concurrent request simulation** for load testing
- **Clean architecture** with separation of concerns

## 📋 Table of Contents

- [Architecture](#architecture)
- [Installation](#installation)
- [Usage](#usage)
- [API Endpoints](#api-endpoints)
- [Docker Deployment](#docker-deployment)
- [Environment Variables](#environment-variables)
- [Load Balancer Simulation](#load-balancer-simulation)
- [Contributing](#contributing)
- [License](#license)

## 🏗️ Architecture

The application follows a clean, modular architecture:

- `main.go` - Application entry point
- `setup.go` - Initialization and configuration
- `server.go` - Gin server setup and routing
- `handlers.go` - API endpoint handlers
- `loadbalancer.go` - Load balancer core implementation
- `simulator.go` - Load balancer simulation utilities
- `config.go` - Configuration management
- `globals.go` - Global variables

## 🔧 Installation

### Prerequisites

- Go 1.18+ (application tested with Go 1.24)
- Git

### Local Setup

1. Clone the repository:

```bash
git clone https://github.com/szeyu/go-loadbalancer.git
cd go-loadbalancer
```

2. Install dependencies:

```bash
go mod download
```

3. Create a `.env` file in the project root (or use the provided one):

```
PORT=8080
API_KEYS=key1,key2,key3,key4,key5
```

4. Build the application:

```bash
go build
```

5. Run the application:

```bash
./go-loadbalancer
```

## 🚀 Usage

After starting the application, you can access the API endpoints at `http://localhost:8080`.

## 🌐 API Endpoints

- **GET /health**
  - Health check endpoint
  - Returns: `{"status": "ok", "message": "Server is healthy"}`

- **GET /api/key**
  - Get the next API key from the load balancer
  - Returns: `{"api_key": "key1", "message": "This API key would be used for external service authentication"}`

- **GET /simulator/load-balancer?n=10**
  - Run a load balancer simulation with `n` iterations (default: 10)
  - Returns: `{"message": "Running load balancer simulation"}`

- **GET /simulator/load-balancer-concurrent?n=10**
  - Run a concurrent load balancer simulation with `n` iterations (default: 10)
  - Returns: `{"message": "Running concurrent load balancer simulation"}`

## 🐳 Docker Deployment

### Build the Docker Image

```bash
docker build -t go-loadbalancer .
```

### Run the Container

```bash
# Using environment variables from host
docker run -d -p 8080:8080 --env-file .env --name loadbalancer-app go-loadbalancer

# Using environment variables baked into the image
docker run -d -p 8080:8080 --name loadbalancer-app go-loadbalancer
```

### Docker Commands

- View container logs: `docker logs loadbalancer-app`
- Stop container: `docker stop loadbalancer-app`
- Remove container: `docker rm loadbalancer-app`

## 🔑 Environment Variables

| Variable | Description | Default |
|----------|-------------|---------|
| PORT     | Port to run the server on | 8080 |
| API_KEYS | Comma-separated list of API keys | key1,key2,key3,key4,key5 |

## 🔄 Load Balancer Simulation

The application includes utilities to demonstrate load balancing in action:

- **Simple simulation**: Sequential requests using different API keys
- **Concurrent simulation**: Parallel requests demonstrating thread-safe load balancing

Run the simulations by accessing the simulation endpoints or inspect the code in `simulator.go`.

## 🤝 Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.