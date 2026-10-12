# Backend for Disbursement Order

## Pre-requisites
1. Install [sqlc](https://github.com/sqlc-dev/sqlc/blob/main/docs/overview/install.md)

    Quick install command:
    ```
    go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
    ```

## Development

1. Start database server: `docker-compose up -d`
2. Copy `config.example.yaml` to `config.yaml`
3. Fetch Dependency: `go mod tidy`
4. Vendoring Dependency: `go mod vendor`
5. Generate sqlc: `make generate-queries`
6. Update database: `go run main.go migrate`
7. Start server: `go run main.go start`

Whenever there is changes in queries, re-run step number 5.

Whenever there is changes in db_schema, re-run step number 5 and the re-run step number 6.

## Accessing the database from the terminal

Start PostgreSQL with Docker Compose:

```sh
docker-compose up -d postgres
```

Connect to the database from inside the PostgreSQL container:

```sh
docker-compose exec postgres psql -U disburseorder -d disburseorder
```

At the `psql` prompt, query the orders table:

```sql
SELECT * FROM orders LIMIT 20;
```

Exit `psql` with `\q`. If the `orders` table does not exist yet, apply the migrations with `go run main.go migrate`.

The Compose port mapping exposes PostgreSQL on port `5438` on the host. If `psql` is installed on your machine, connect directly with:

```sh
psql -h localhost -p 5438 -U disburseorder -d disburseorder
```

## Sample request

Submit a disbursement order request:
- get the invoice_id param from table you have queried before
- status can be fill with approve / reject.

```sh
curl --location 'localhost:3000/v1/order/disbursement' \
--header 'Content-Type: application/json' \
--data '{
    "invoice_id": "0b3d1447-ffd1-45cb-99a2-6ea36c6c3308",
    "status": "approve"
}'
```