# gRPC Authorization service

This project is a gRPC-based authorization service written in Go. It provides a secure and efficient way to manage access control across various services within a distributed system.

## Contents

- [Getting Started](#getting-started)
- [Prerequisites](#prerequisites)
- [Installation](#installation)
- [Usage](#usage)
  -- [Migrations](#migrations)
- [License](#license)

## Getting Started
To run the application:
```bash
make build && make run
```

### Prerequisites

Make sure you have the following software installed:

- [Go](https://golang.org/doc/install)
- [Docker](https://docs.docker.com/get-docker/)
- [Docker Compose](https://docs.docker.com/compose/install/)

### Installation

1. Clone the repository:

    ```bash
    git clone https://github.com/rtzgod/auth-service.git
    ```

2. Create a `.env` file with PostgreSQL environment variables:

    ```env
    POSTGRES_USER=myuser
    POSTGRES_PASSWORD=mypassword
    POSTGRES_DB=mydatabase
    ```

3. Configure your local configs in configs/local.yml file according to your PostgresSQL settings

### Usage

After starting app you can access gRPC endpoints on the host you configured 

(default: localhost:44044)

#### Migrations

App performs db migrations in-code, you can create your own migration schemas in db/migrations

```bash
# You can choose your own path for migration folder in -dir
migrate create -ext sql -dir ./db/migrations -seq init
```

or you can create migration files manually and set versions of migrations by yourself

Also after creating migrations you can perform any capable operations with it
```bash
# For detailed usage of migrate command
migrate -help
```

### License
This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.