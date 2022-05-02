# FaceIT Go coding Challenge v1

Resolution of the proposed technical test.

Created using go version `1.17`.

## How to run it?

Clone the project, get into the correct folder and execute the command:

```bash
docker-compose -f server/docker-compose.yml up --build
```
This command will launch `docker-compose` and create four services:
  * rpc API to support add, update, delete and list actions
  * MySQL instance as storage service
  * Kafka + Zookeeper as service to handle notifications

Because the database takes a few seconds to start up, the application must be made to try to connect recurrently for a period of time (20 seconds). Same thing for Kakfa (60 seconds max).

When the connection is ready and the application is active, you will see the following entry in the docker-compose logs:

```sh
faceit          | ***** Database is ready *****
faceit          | ***** Kafka is ready *****
```

## How to turn it off completely?

```bash
docker-compose -f server/docker-compose.yml down -v
```

## How to run the tests?

```sh
go test ./...
```

## How to use it?

To invoke the exported rpc endpoints, use the client component. This client will consume all the available methods exported by the server.

```sh
cd client
go run main.go
```

Output with the values returned by the server will be printed.

## How to check the kafka topic with event notifications?

To check the event notifications created during the consumption of server exposed methods, you can run this command to get the elements added to the Kafka topic.

```sh
docker exec --interactive --tty broker \
kafka-console-consumer --bootstrap-server broker:9092 \
                       --topic user_events \
                       --from-beginning
```

You will see an output like this:

```sh
User Add: Id -> 11d4d09a-e9e1-46d4-94f6-0bb89a72a3b8
User Update: Id -> 11d4d09a-e9e1-46d4-94f6-0bb89a72a3b8
User Delete: Id -> 11d4d09a-e9e1-46d4-94f6-0bb89a72a3b8
```

## Explanation and assumptions

* I have tried to implement what was requested by following the instructions.

* The data validation is not exhaustive, I have only considered a couple of representative cases (uuid validation, email validation and firstname, nickname and password should not be empty).

* The tests do not cover absolutely all the cases that could occur, it is a sample of how they would be implemented in a real environment.

* For logging I used the `log` package, it could also be improved.

* As storage method I used a MySQL database version 5.7.

* I use Kafka as service to handle the publication of notifications.