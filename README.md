# Shared API

An Application Programming Interface (API) to manage information that is used in multiple applications but does not belong to one specific application.

Built with Golang and Deployed on AWS Lambda.

# How to Develop

## Running the GO API for Local Development

Either of these options starts the API at `localhost:3030`

**With Visual Studio Code Debugger**

You can use the launch.json file in this repository in lieu of `go run root/main.go` to run the API in the VSCode debugger. This takes care of the required environment variables to connect to the database.

**Without Visual Studio Code Debugger**

Set the following environment variables and type `go run root/main.go` from the top level of this repository.

    * SHARED_LAMBDACONTEXT=FALSE

Note: When running the API locally, make sure environment variable `SHARED_LAMBDACONTEXT` is either **not set** or is set to `SHARED_LAMBDACONTEXT=FALSE`.

## Running API Docs Locally

From the top level of this repository, type `make docs`. This starts a container that serves content based on "apidoc.yml" in this repository.
Open a browser and navigate to `https://localhost:4000` to view the content.
