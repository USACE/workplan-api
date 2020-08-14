# Shared API

An Application Programming Interface (API) to manage information that is used in multiple applications but does not belong to one specific application.

Built with Golang and Deployed on AWS Lambda.

# How to Develop

## Running the GO API for Local Development

Either of these options starts the API at `localhost:3030`. The API uses JSON Web tokens (JWT) for authorization by default. To disable JWT for testing or development, you can set the environment variable `JWT_DISABLED=TRUE`.

**With Visual Studio Code Debugger**

You can use the launch.json file in this repository in lieu of `go run root/main.go` to run the API in the VSCode debugger. This takes care of the required environment variables to connect to the database.

**Without Visual Studio Code Debugger**

Set the following environment variables and type `go run root/main.go` from the top level of this repository.

    * DB_USER=postgres
    * DB_PASS=postgres
    * DB_NAME=postgres
    * DB_HOST=localhost
    * DB_SSLMODE=disable
    * LAMBDA=FALSE
    * JWT_DISABLED=FALSE

Note: When running the API locally, make sure environment variable `LAMBDA` is either **not set** or is set to `LAMBDA=FALSE`.

## Running API Docs Locally

From the top level of this repository, type `make docs`. This starts a container that serves content based on "apidoc.yml" in this repository.
Open a browser and navigate to `https://localhost:4000` to view the content.

# How To Deploy

## Postgres Database on AWS Relational Database Service (RDS)

Database should be initialized with the following SQL files in the order listed:

1. rds_init.sql (install PostGIS extensions)

1. tables.sql (create tables for application)

1. roles.sql (database roles, grants, etc.)

   Note: Change 'password' in roles.sql to a real password for the `instrumentation_user` account.
