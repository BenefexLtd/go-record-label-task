# go-record-label-task 🎧 💿

## Dependencies

* Go 1.17+
* Docker / Docker Compose

## Task

### Brief

BenFX Records handles its release schedule via an API with a single `POST /releases` endpoint.

This endpoint consumes an array of Release objects and hands each one off for further processing.

A successful request is acknowledged with a `202 Accepted` response header, and receipt of each Release is currently added
to the system log.

Your task is to build out additional components to handle the following:

* Issue communications to the artist's fanbase to inform them of each upcoming release.

* Issue orders to the distribution warehouses for physical fulfilment of the release media.

### Requirements & Constraints

* BenFX Records is currently out to tender for the suppliers who will handle communications and distribution, so we need
to make sure we can easily swap these integrations in once these decisions have been made.
    * It is likely that CD and Vinyl orders will be dealt with by separate suppliers in order to achieve optimal rates. 

* A Docker Compose file has been supplied which provides a RabbitMQ environment that you may wish to optionally
include as part of your solution.
    * Once you are up-and-running, you can access the management console at http://localhost:15672 using the default
    username _guest_ and password _guest_.

* Your solution should prioritise maintainability as well as robustness, and utilise a queue/message broker for communications.

* You must provide your complete solution as a zip file, including the `.git` subfolder, and return it privately via email
(please do not publish your solution publicly!)
    * Please use your commits to demonstrate incremental progress, as you would whilst working on tickets as part of a sprint.

## Getting Started

From project root:

```
docker-compose up
```

This will spin up two local containers: one to run the API, and the other to provide a RabbitMQ environment with
basic/minimal configuration.

You can issue requests to `POST http://localhost:8080/releases` by running the provided script (`./curl.sh`)
