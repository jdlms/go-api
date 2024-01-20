# Basic Go api to learn file structure conventions

your-project/
├── cmd/ # Application entry points
│ └── myapp/
│ └── main.go # Main application
├── pkg/ # Library code
│ ├── api/ # API-specific logic, possibly including route definitions
│ │ └── api.go
│ ├── db/ # Database-specific code
│ │ └── db.go
│ ├── models/ # Data models
│ │ └── models.go
│ └── handlers/ # HTTP handlers
│ └── handlers.go
├── migrations/ # Database migration scripts
│ ├── 1_users.up.sql
│ └── 2_books.up.sql
├── Dockerfile
├── docker-compose.yml
├── go.mod
├── go.sum
└── README.md
