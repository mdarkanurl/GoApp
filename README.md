# GoApp

A Go-based RSS feed aggregator API that allows users to register, follow RSS feeds, and automatically scrape posts on a scheduled basis.

## Features

- **User Management**: Create accounts and authenticate via API keys
- **RSS Feed Following**: Follow any RSS feed and have it automatically scraped
- **Automatic Scraping**: Background workers periodically fetch and parse RSS feeds into posts
- **Post Aggregation**: View all posts from feeds you follow
- **RESTful API**: Full CRUD operations for users, feeds, and feed follows

## Tech Stack

- **Language**: Go 1.27
- **Web Framework**: [chi](https://github.com/go-chi/chi) router
- **Database**: PostgreSQL with [sqlc](https://sqlc.dev/) for type-safe queries
- **RSS Parsing**: Built-in XML/RSS feed parser
- **Environment**: [godotenv](https://github.com/joho/godotenv) for configuration
- **CORS**: [go-chi/cors](https://github.com/go-chi/cors) middleware

## Prerequisites

- Go 1.27+
- PostgreSQL
- `sqlc` CLI (for generating database code, if modifying queries)

## Setup

1. **Clone the repository**:
   ```bash
   git clone <repo-url>
   cd GoApp
   ```

2. **Set up PostgreSQL**:
   Create a database named `GoApp` and configure the connection string in `.env`.

3. **Configure environment**:
   Copy `.env.example` or set the following environment variables:
   ```env
   PORT=8080
   DATABASE_URL=postgresql://postgres:postgres@localhost:5432/GoApp?sslmode=disable
   ```

4. **Initialize the database**:
   Run the SQL schema migrations (located in `sql/schema/`) to create the required tables:
   - `users`
   - `feeds`
   - `feed_follows`
   - `posts`

5. **Install dependencies**:
   ```bash
   go mod tidy
   ```

6. **Build and run**:
   ```bash
   go build -o GoApp .
   ./GoApp
   ```

## API Endpoints

All endpoints are prefixed with `/v1`.

### Public

| Method | Endpoint       | Description           |
|--------|----------------|-----------------------|
| GET    | `/ready`       | Health check          |
| POST   | `/users`       | Create a new user     |
| GET    | `/feeds`       | Get all feeds         |

### Authenticated (requires `Authorization: ApiKey <key>` header)

| Method | Endpoint                         | Description                    |
|--------|----------------------------------|--------------------------------|
| GET    | `/users`                         | Get current user               |
| POST   | `/feeds`                         | Create a new feed              |
| GET    | `/posts`                         | Get posts from followed feeds  |
| POST   | `/feed_follows`                  | Follow a feed                  |
| GET    | `/feed_follows`                  | Get all followed feeds         |
| DELETE | `/feed_follows/{feedFollowID}`   | Unfollow a feed                |

### Authentication

All authenticated endpoints require an API key passed in the `Authorization` header:

```
Authorization: ApiKey <your-api-key>
```

## Project Structure

```
.
├── main.go                  # Entry point, server initialization, routing
├── models.go                # Data models (User, Feed, FeedFollow, Post)
├── scraper.go               # Background RSS scraping logic
├── rss.go                   # RSS/XML parsing utilities
├── json.go                  # HTTP response helpers
├── middleware_auth.go       # API key authentication middleware
├── handler_*.go             # HTTP request handlers
├── .env                     # Environment configuration
├── sqlc.yml                 # sqlc configuration
├── sql/
│   ├── schema/              # Database migration files (goose)
│   └── queries/             # SQL queries for sqlc
├── internal/
│   ├── database/            # Generated sqlc database code
│   └── auth/                # Authentication utilities
└── vendor/                  # Vendored dependencies
```

## Database Schema

The project uses PostgreSQL with the following tables:

- **users** — User accounts with UUID-based IDs
- **feeds** — RSS feed sources (linked to users)
- **feed_follows** — Many-to-many relationship between users and feeds
- **posts** — Scraped articles from feeds (linked to feeds)

## Development

- **Database queries** are managed via [sqlc](https://sqlc.dev/). Modify `.sql` files in `sql/queries/` and regenerate code with `sqlc generate`.
- **Migrations** are written in `sql/schema/` using goose-style up/down statements.
- **Dependencies** are vendored in the `vendor/` directory.

## License

MIT
