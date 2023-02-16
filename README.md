# Platform Engineer Challenge

## Overview

Inside the Online Tools Group we make information about the various development boards that our partners offer available on our websites. These development boards are used by embedded developers when building new products. You can explore this data for yourself at https://www.keil.arm.com/boards/. The data about these boards comes from a HTTP API.

The goal of this challenge is to demonstrate your skills as a platform engineer by deploying a simplified version of this HTTP API, contained in this directory, to a cloud environment. You should aim to automate the process of deploying this service, following CI/CD and Infrastructure as Code (IaC) practices.

## The Challenge

You have been tasked with deploying this service, a HTTP server and Postgres database, to a cloud environment so that the HTTP API is publicly accessible. You should aim to automate the process of deploying this service, following CI/CD and Infrastructure as Code (IaC) practices.

- You can use any cloud provider or platform, many offer a free trial, for example (but not limited to):
  - https://aws.amazon.com/free
  - https://azure.microsoft.com/account/free
  - https://cloud.google.com/free
  - https://try.digitalocean.com/freetrialoffer/
- You can use any technologies and tools
- You should not need to make any changes to the Go code or Postgres data, but feel free to make or suggest changes that help your solution

### Evaluation Criteria

We'll discuss your solution with you during the interview. Having the API running and publicly accessible is desirable but not essential; having something representative but not deployed is also acceptable. We would like to see the configuration and automation for deploying the service, but you can make use of pseudocode and diagrams if it allows you to focus on designing the overall solution in the time given.

We will be evaluating your solution on:

- Readability and explanation of your solution
- Reasoning behind the decisions made
- Use of automation
- Use of Infrastructure as Code
- Adherence to security best practices

### Time Limit

We do not expect you to spend more than two hours on this challenge.
It may not be possible to finish everything you would like inside this timeframe and an incomplete solution is not a failure.
During the interview there will be an opportunity to discuss why you prioritised what you did, and what your next steps would be.

### Submission

Please provide a link to a repository containing the code along with anything you have added to deploy the service or explain your solution.

## The Service

The HTTP service in this repository has two parts: a small Go HTTP server, and a PostgreSQL database.

### Go HTTP Server

You can find the implementation for the Go HTTP server in the `main.go` and `get_board_endpoint.go` files.
The HTTP server has two endpoints, both of which respond to `GET` requests:

- `/board/{id}`
- `/health`

The `/board/{id}` endpoint returns data about a board from the database with the given `id`.
For example, sending a `GET` request to `/board/1` returns a HTTP `200` response containing the board data as JSON:

```json
{
  "id": 1,
  "name": "B7-400X",
  "vendor": "Boards R Us",
  "core": "Cortex-M7"
}
```

Requesting data about a board that doesn't exist will result in a HTTP `404` response.

The `/health` endpoint always returns a HTTP `200` response as long as the server is running.

### Postgres Database

The database has just a single table called `boards` which has four columns: `id`, `name`, `vendor`, `core`.
These columns match the fields returned by the HTTP API.

To help create and seed this database `pg_dump.sql` has been provided which can be imported into a Postgres database:

```
psql -h hostname -U username -f pg_dump.sql
```

Alternatively, the commands used to create and seed this database were:

```sql
CREATE TABLE BOARDS(
  ID SERIAL PRIMARY KEY,
  NAME TEXT NOT NULL,
  VENDOR TEXT NOT NULL,
  CORE TEXT NOT NULL
);
INSERT INTO BOARDS (NAME, VENDOR, CORE) VALUES ('B7-400X', 'Boards R Us', 'Cortex-M7');
INSERT INTO BOARDS (NAME, VENDOR, CORE) VALUES ('Low_Power', 'Tech Corp.', 'Cortex-M0+');
INSERT INTO BOARDS (NAME, VENDOR, CORE) VALUES ('D4-200S', 'Boards R Us', 'Cortex-M4');
```

As you can see, sample data is provided for three boards:

```sql
SELECT * FROM boards;

 id |   name    |   vendor    |    core
----+-----------+-------------+------------
  1 | B7-400X   | Boards R Us | Cortex-M7
  2 | Low_Power | Tech Corp.  | Cortex-M0+
  3 | D4-200S   | Boards R Us | Cortex-M4
```

## Running the Service

With the Golang toolchain installed you can compile the Go server with:

```sh
go build -o simple-http-server
```

And to run the server:

```sh
./simple-http-server
```

By default the HTTP server runs on port `8080`, however this can be changed via configuration.
Also you will need to provide the address to a running Postgres via configuration, else you will likely receive an error.

### Configuration

The server allows configuration to be provided via the command line, a configuration file (`simple-http-server.yaml`), and environment variables.

| CLI Flag | Environment Variable | Config file | Description                                                                                   |
| -------- | -------------------- | ----------- | --------------------------------------------------------------------------------------------- |
| `--db`   | `DB`                 | `db`        | The postgres connection string e.g. `postgresql://<username>:<password>@<address>/<database>` |
| `--port` | `PORT`               | `port`      | The port the server will run on (default 8080)                                                |

### Example

Assuming the Postgres database is running on localhost, the service can be run with:

```sh
./simple-http-server --db "postgresql://postgres:mypassword@localhost/postgres?sslmode=disable"
```

And now the HTTP API can be accessed:

```sh
curl http://localhost:8080/board/1
```

### Versions

This was tested with the following versions:

- Golang 1.19.4
- Postgres 13.9

## Good luck!

We are looking forward to seeing your solution and discussing it with you.
