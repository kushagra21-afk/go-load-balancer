# Go Load Balancer

A basic load balancer written in Go that you can use on your own projects or infrastructure.

## Features

- **Round Robin Load Balancing:** Distributes incoming requests evenly across available backend servers.
- **Health Checks:** Monitors backend server health and removes unhealthy servers from the pool.
- **Easy Configuration:** Simple to set up and deploy.
- **Lightweight:** Minimal dependencies, written in pure Go.

## Getting Started

### Prerequisites

- [Go](https://golang.org/dl/) 1.18 or newer

### Installation

Clone this repository:

```bash
git clone https://github.com/kushagra21-afk/go-load-balancer.git
cd go-load-balancer
```

Build the project:

```bash
go build -o go-load-balancer
```

### Usage

Run the load balancer with your desired configuration:

```bash
./go-load-balancer -config=config.yaml
```

Example configuration file (`config.yaml`):

```yaml
listen: ":8080"
backends:
  - "http://localhost:8081"
  - "http://localhost:8082"
health_check:
  interval_seconds: 10
  path: "/health"
```

#### Command-line Options

| Flag      | Description                      | Default      |
|-----------|----------------------------------|--------------|
| `-config` | Path to the configuration file   | `config.yaml`|

## How It Works

1. The load balancer listens on the specified address.
2. Incoming HTTP requests are forwarded to backend servers using a round-robin approach.
3. Health checks are performed at regular intervals to ensure only healthy servers receive traffic.

## Contributing

Contributions are welcome! Please open issues or submit pull requests for improvements and bug fixes.

## Author

- [@kushagra21-afk](https://github.com/kushagra21-afk)
