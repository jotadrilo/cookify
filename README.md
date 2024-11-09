# cookify

## Project Structure

This project is following the hexagonal architecture pattern.

It is structured the project as follows:

```shell
.
├── app
│  ├── adapters
│  │  ├── controllers
│  │  │  └── gin
│  │  ├── repositories
│  │  │  ├── fs
│  │  │  │  └── model
│  │  │  └── pg
│  │  │     └── model
│  │  └── usecases
│  ├── api
│  └── core
│     ├── biz
│     ├── domain
│     └── ports
│        ├── controllers.go
│        ├── repositories.go
│        └── usecases.go
├── client
├── internal
└── server
```

The `app/api` package contains the OpenAPI spec and the generated Go code
(Gin server, API client, and model type definitions).

The `app/core` package contains:

- Business logic
- Domain entities (**Domain Layer**)
- Ports (interface abstractions of the different application layers)
    - **Controllers** (driver or primary ports). They are part of the **Infrastructure Layer**
    - **Repositories** (driven or secondary ports). They are part of the **Infrastructure Layer**
    - **Use cases** (business rules). They are part of the **Application Layer**

The `app/core` package cannot depend on the `adapters` package.

The `app/adapters` package contains the concrete implementations of the ports:

- **Controllers** (adapters of the driver or primary ports). This is part of the **Infrastructure Layer**
    - Gin web framework. As simplification, I am embedding the interface generated from the API spec (in `app/api`).
- **Repositories** (adapters of the driven or secondary ports). This is part of the **Infrastructure Layer**
  - PostgreSQL repository implementation (and model)
  - File system repository implementation (and model)
- **Use Cases** (adapters of the application ports). This is part of the **Application Layer**

The `server` package contains the main program to start the server.

The `client` package contains the main program to start the client.

The `internal` package contains util packages.
