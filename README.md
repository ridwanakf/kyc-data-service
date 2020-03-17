# KYC Data Service
Service for processing users' data such as ID Card data.

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing
purposes.

### Prerequisites

1. Clone repository: `git clone git@github.com:ridwanakf/kyc-data-service.git`
2. Run dep: `dep ensure -v`
3. Install postgresql: [Postgresql](https://www.postgresql.org/download/)
4. Run database migrations: refer below

### Run Project

Running project:

```$xslt
make run
```

### Migrations

When running in Local, you need to run the db-migrations to setup the app's database for your local machine.

1. Go to directory `kyc-data-service/migrations`
2. Run `go run *.go up`

## Directory Structure

This repository is organized in the following directory structure.

```
kyc-data-service
|-- cmd                                    # Contains executable codes and serves as entry point of the app
|   |-- main                               # entry point of the app
|-- config                                 # Configuration files needed for deployment
|-- constant                               # Collections of constants file for each module
|-- internal                               # Go files in this folder represent the Big-Pictures and Contracts of the system
|   |-- app                                # Contains constructions of the app and other app's related configs
|   |   |-- config                         # Configuration struct for the app
|   |-- delivery                           # Delivery layer of the app
|   |   |-- rest                           # HTTP REST API delivery of the app
|   |   |-- <other_delivery_mechanisms>    # Other delivery mechanisms of the app (eg. GRPC, Console, Web, etc.)
|   |
|   |-- entity                             # Enterprise Data structures
|   |   |-- ktp.go                         # Data structure for KTP module
|   |   |-- <other_entities>.go            # Other data structures, preferrably 1 struct 1 file
|   | 
|   |-- helper                             # Collections of Helper/Common function
|   |   |-- helpers.go                     # RenderJSON func
|   |   |-- <other_helpers>.go             # Helper implementations to other.
|   |
|   |-- repo                               # Implementations of Repository-pattern to data-sources
|   |   |-- db                             # Implementations of the repositories with Postgres database
|   |   |-- <other_repos>                  # Other Repositories implementations based on interfaces on folder internal.
|   |
|   |-- usecase                            # Usecases implementations for Application Business Logic
|   |   |-- ktp.go                         # Other use-case implementations based on interfaces on folder internal.
|   |
|   |-- repo.go                            # Interfaces / Contracts of all the repositories (Repository Pattern)
|   |-- usecase.go                         # Interfaces / Contracts of all the use-cases (Application Business Logic)
|
|-- migrations                             # Contains Database migration files or the system
|-- vendor                                 # Dependencies folder that's maintained by dep tool https://golang.github.io/dep/
|-- Gopkg.lock                             # https://golang.github.io/dep/docs/Gopkg.lock.html
|-- Gopkg.toml                             # https://golang.github.io/dep/docs/Gopkg.toml.html

```

## Tech Stacks

- Golang
- Postgresql
