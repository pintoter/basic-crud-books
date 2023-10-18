[![Golang](https://img.shields.io/badge/Go-v1.21-EEEEEE?logo=go&logoColor=white&labelColor=00ADD8)](https://go.dev/)
[![License](https://img.shields.io/badge/license-MIT-green)](LICENSE)

<div align="center">
    <h1>Books service</h1>
    <h5>
        A service written in the Go for create, read, update and delete books.
    </h5>
</div>

---

## Technologies used:
- [Golang](https://go.dev), [PostgreSQL](https://www.postgresql.org/)
- [Docker](https://www.docker.com/)
- [REST](https://ru.wikipedia.org/wiki/REST)

---

## Installation
```shell
git clone git@github.com/pintoter/basic-crud-books.git
```

---

## Getting started
1. **Setting up environment variables (create a .env file in the project root):**
```dotenv
# Database
export DB_HOST=
export DB_PORT=
export DB_USER=
export DB_NAME=
export DB_SSLMODE=
export DB_PASSWORD=

# Postgres service
export POSTGRES_PASSWORD=

```
> **Hint:**
if you are running the project using Docker, set `DB_HOST` to "**postgres**" (as the service name of Postgres in the docker-compose).

2. **Compile and run the project:**
```shell
make start
```

---

## [Examples of requests](./docs/examples/02-requests.md)

**[Books](./docs/examples/02-requests.md#Books)**
* [Create a book](./docs/examples/02-requests.md#1-create-a-segment)
* [Get all books](./docs/examples/02-requests.md#2-create-a-segment-with-an-indication-of-the-percentage-of-automatic-addition)
* [Get book by ID](./docs/examples/02-requests.md#2-create-a-segment-with-an-indication-of-the-percentage-of-automatic-addition)
* [Get books by Author](./docs/examples/02-requests.md#2-create-a-segment-with-an-indication-of-the-percentage-of-automatic-addition)
* [Update a book by Author and Title](./docs/examples/02-requests.md#2-create-a-segment-with-an-indication-of-the-percentage-of-automatic-addition)
* [Delete a book by ID](./docs/examples/02-requests.md#3-delete-a-segment-by-name)
* [Delete a book by Title](./docs/examples/02-requests.md#4-delete-a-segment-by-id)

**[Users](./docs/examples/02-requests.md#Users)**
* [Registration](./docs/examples/02-requests.md#1-registration)
* [Authentication](./docs/examples/02-requests.md#2-authentication)
* [Refresh token](./docs/examples/02-requests.md#3-refresh-toke)
---
