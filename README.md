# Events Service
The project contains the backend code for logging historical events built with `GoLang`, `GoKit` (https://gokit.io/) and `gRPC` as the transport layer. `PostgreSQL` was chosen as the choice database for this application 

## Getting Started

These instructions will get you a copy of the app up and running on your local machine for development and testing purposes.

### Prerequisites

Things you need to install the software and how to install them
- Ensure you have `Docker` installed on your local machine
- Ensure you have `Go` installed as well
- Ensure you have `Docker Compose` installed by checking the current version like so

```
docker-compose --version
```

### Installing

This is a step by step series of examples to get the project up and running on your local machine

First thing to do is clone the github repository like so

```
git clone https://github.com/johnchuks/events-service.git
```

After cloning the repository, we will need to enter the current directory of the project with this command

```
cd events-service
```

Create a `.env` file and add a `DATABASE_USER` and `DATABASE_PASSWORD` of your choice as shown in the `.env-sample` file. The PostgreSQL docker image automatically creates a database user and password on startup with the provided `DATABASE_USER` and `DATABASE_PASSWORD`.

We need to build and run the `docker` images for the `grpc` and `database`. Our containers are currently orchestrated using `docker-compose`.

To build and run the container, we use this command below

```
docker-compose up --build
```

Finally we can go to `0.0.0.0:50052` to communicate with our gRPC server.

To bring down all running containers and network, Run `docker-compose down`.


## Testing

The Backend adheres to Separation of Concerns which makes it very easy to test. Every layer in the backend will be unit tested. Also, integration tests will be utilized to test the system in general.

For naive testing or manual testing of gRPC, there is a great tool I normally use which I highly recommend. It exposes a GUI for your services based on the `.proto` file. `BloomRPC` (https://github.com/uw-labs/bloomrpc) is very easy to setup and start testing immediately. 

Ensure the correct URL is passed to the URL field on the app.


## Architecture
The backend uses a 3 layer architecture pattern which comprises of a `Transport`, `Service` and `Endpoint` layer.

`Transport` - The transport layer defines the protocol i.e. `grpc` which services communicate with each other. The backend can support a number of other protocols such as `pub/sub` and  `HTTP`

`Service` - The service layer in this architecture is where the core business logic is located. The service layer is connected to the database and we are able to make queries with `gorm` ORM.

`Endpoint` - The endpoint layer defines an HTTP request handler or in our case an RPC method. Based on the proto file we currently have two RPC methods.

- Create : Adds a new historical Event

- Retrieve: Retrieves all historical events that match a certain criteria or filters. All filters are optional and represented in this order {"text": "hello", "email": "example@test.com"}

## Questions

- Would any Cache be used? 

    - I suggest adding caching to this architecture in the service layer. A `cache-aside` strategy is highly recommended for retrieving data from the service because our request will be directed to the cache first before making any database query. If the data is changed in the database, the cache is freed of the old data and subsequent requests ensures the data is added back to the cache.


- Are there any kind of strategy for saving to the DB? 
    - To increase performance, an efficient approach would be to batch the write queries to the database instead of hitting the database for every single write query. Another optimization strategy is having a `Master-slave` replication where all writes request are handled by the master DB and the Reads queries are shared among the slave nodes.