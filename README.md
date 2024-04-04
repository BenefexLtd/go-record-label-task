# go-record-label-task 🎧 💿

## Example Dependencies

- Go 1.22+
- Docker / Docker Compose

## Task

### Brief

BenFX Records manages its release schedule via an API with a single `POST /releases` endpoint.

This endpoint consumes an array of Release objects and delegates each one for further processing.

Upon successful processing, a `202 Accepted` response header is issued, and the receipt of each Release is currently logged within the system.

Your task is to expand the existing components to handle the following:

- Communicate release information to the artist's fanbase to notify them of each upcoming release.
- Initiate orders to distribution warehouses for physical fulfillment of the release media.

Notes:
- You have the flexibility to extend the existing codebase or begin from scratch.
- You may opt to utilize the provided pubsub emulator or choose an alternative message broker for communication.
- When documenting, feel free to remove content from this readme.
- You're not limited to a single service.

### Requirements & Constraints

- BenFX Records is currently out to tender for the suppliers who will handle communications and distribution,
  so we need to make sure we can easily swap these integrations in once these decisions have been made.
    - It is likely that CD and Vinyl orders will be dealt with by separate suppliers in order to achieve optimal rates. 

- The supplied Docker Compose file initializes the example API code and the pubsub emulator, which you may optionally
  incorporate into your solution.
  - Additional information on the PubSub emulator can be found here: [PubSub Emulator Documentation](https://cloud.google.com/pubsub/docs/emulator)

- Your solution should prioritize maintainability and robustness, utilizing a queue/message broker for communications.

- Provide your complete solution as a zip file, and include the output of git log to show your progress along the way.
  - Use your commits to showcase incremental progress, similar to working on tickets within a sprint.
  - To write the git log to a file, run: `git --no-pager log > log.txt`

- Take time to document your approach and reasoning as you see fit. This aids in understanding your thought
  processes before a detailed discussion.

## Getting Started

From the project root:

```
docker-compose up
```

This command will initialize two local containers: one to run the API and the other to provide a pubsub emulator with basic/minimal configuration.

You can issue requests to `POST http://localhost:8080/releases` by running the provided script (`./curl.sh`).
