# cookify

## Project Structure

This project is following the hexagonal architecture pattern.

It is structured the project as follows:

```shell
├── api
├── adapters
│  ├── controllers
│  │  └── gin
│  ├── repositories
│  │  └── pg
│  └── usecases
└── core
   ├── biz
   ├── domain
   └── ports
      ├── repositories.go
      └── usecases.go
```

The `api` package contains the OpenAPI spec and the generated Go code
(Gin server, API client, and model type definitions).

The `core` package contains:

- Business logic
- Domain entities (**Domain Layer**)
- Ports (interface abstractions of the different application layers)
    - **Controllers** (driver or primary ports). They are part of the **Infrastructure Layer**
    - **Repositories** (driven or secondary ports). They are part of the **Infrastructure Layer**
    - **Use cases** (business rules). They are part of the **Application Layer**

The `core` package cannot depend on the `adapters` package.

The `adapters` package contains the concrete implementations of the ports:

- **Controllers** (adapters of the driver or primary ports). This is part of the **Infrastructure Layer**
    - Gin web framework. As simplification, I am embedding the interface generated from the API spec.
- **Repositories** (adapters of the driven or secondary ports). This is part of the **Infrastructure Layer**
    - PostgreSQL repository implementation (and model)
- **Use Cases** (adapters of the application ports). This is part of the **Application Layer**
