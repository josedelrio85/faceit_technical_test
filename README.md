# FaceIT Go coding Challenge

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
cd server
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

## Exposed methods

### Add user endpoint

Add an entity (user) into the database if validations is OK.

A valid User entity parameter should be provided. The response will be stored in a Response struct.

```sh
# User entity
User{
  Id:        "d2a7924e-765f-4949-bc4c-219c956d0f8b",
  FirstName: "foo",
  LastName:  "bar",
  NickName:  "foobar",
  Password:  "foobarpassword",
  Email:     "foo@bar.test",
  Country:   "UK",
  CreatedAt: time.Now(),
  UpdatedAt: time.Now(),
}
```
You will get a response with this format if the process succeed:

```json
{
  "status":200,
  "success":true,
  "data":{
    "userid":"d2a7924e-765f-4949-bc4c-219c956d0f8b"
  }
}
```

### Update user endpoint

Update an entity (user) into the database if validations is OK.

A valid User entity parameter should be provided. The response will be stored in a Response struct.

```sh
# User entity
User{
  Id:        "d2a7924e-765f-4949-bc4c-219c956d0f8b",
  FirstName: "foo",
  LastName:  "bar",
  NickName:  "foobar",
  Password:  "foobarpassword",
  Email:     "foo@bar.test",
  Country:   "UK",
  CreatedAt: time.Now(),
  UpdatedAt: time.Now(),
}
```
You will get a response with this format if the process succeed:

```json
{
  "status":200,
  "success":true,
  "data":{
    "userid":"d2a7924e-765f-4949-bc4c-219c956d0f8b"
  }
}
```

### Delete user endpoint

Delete an entity (user) from the database if the provided user identifier previously exists in the database.

A valid user id value (uuid) should be provided in the URL path.

```sh
# User entity
User{
  Id:        "d2a7924e-765f-4949-bc4c-219c956d0f8b",
}
```

You will get a response with this format if the process succeed:

```json
{
  "status":200,
  "success":true,
  "data":{
    "userid":"d2a7924e-765f-4949-bc4c-219c956d0f8b"
  }
}
```

### List user endpoint

List endpoint retrieves a paginated list of entities (user) using the criteria provided in the pagination struct property.

The pagination property has this format:

```json
{
  "id": "",
  "pagination": {
    "search_by": "country",
    "search_value": "UK",
    "search_results_per_page": 3,
    "offset": 0
  }
}
```

The criteria applied for pagination is:
  - if no search_by param is provided, returns all data (1=1).
  - if no search_value param is provided, returns all data (1=1).
  - if no results_per_page param is provided, returns the default limit number (10)
  - if no offset param is provided, returns from offset 0.

You will get a response with this format if the process succeed:

```json
{
  "status":200,
  "success":true,
  "data":{
    "userlist":[
      {
        "id":"0818cf0f-298f-4f40-bcdb-8a13ff0fc11f",
        "first_name":"f7055ec08a4d5ff6fd3a4d10e18a6bd9f222bdd8",
        "last_name":"84051c5e5bee03551db528eac31ffad9c53bc371",
        "nickname":"2f98f4e2a41292dcbe97e956bc5d14cd17b0ce0c",
        "password":"519b2b6af20123fbf1aed2500d2804ca95705997",
        "email":"dddddd@bob.com",
        "country":"UK",
        "created_at":"2022-04-29T12:13:55+02:00",
        "updated_at":"2022-04-29T12:13:55+02:00"
      },
      {
        "id":"0a56e44b-a6ca-4dc2-858e-8da078b78b6f",
        "first_name":"f5151d739f9605ed5b35a39b93168fc6cb15015e",
        "last_name":"98661c46c282e1db7f480824cd8c661758bc9757",
        "nickname":"733517675ef37c4e9641c42268dd74a6b26c6e78",
        "password":"f62a659309a894ef18e35501d088cd331005683e",
        "email":"aaaaa@bob.com",
        "country":"UK",
        "created_at":"2022-04-29T11:13:47+02:00",
        "updated_at":"2022-04-29T11:13:47+02:00"
      },
      {
        "id":"12c8a195-c3b2-42db-8582-326d8c88e210",
        "first_name":"d4138258b604d021b41a0cb03bfc7598d9bf9346",
        "last_name":"3040f4dde1634ef1877dc011d022e501e83e1e53",
        "nickname":"c8e9279351c516a80d8752e3e298a8dd6066793c",
        "password":"a652b4f0bd2b91cb2f9bb7a83137dd4b0c5e6125",
        "email":"alice@bob.com",
        "country":"UK",
        "created_at":"2019-10-12T09:20:51+02:00",
        "updated_at":"2019-10-12T09:20:51+02:00"
      },
    ]
  }
}
```

## Explanation and assumptions

* I have tried to implement what was requested by following the instructions.

* The data validation is not exhaustive, I have only considered a couple of representative cases (uuid validation, email validation and firstname, nickname and password should not be empty).

* The tests do not cover absolutely all the cases that could occur, it is a sample of how they would be implemented in a real environment.

* For logging I used the `log` package, it could also be improved.

* As storage method I used a MySQL database version 5.7.

* I use Kafka as service to handle the publication of notifications.