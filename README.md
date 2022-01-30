# Simple Microservice
This simple project is an example of a microservice architecture. With the API gateway handling the authentication and routing, and the underlying service handling the respective business logic.

**Important Note**: In this example the same database cluster is utilized for all services. This is not a good practice, but is done for simplicity.

## Getting Started
You can startup the services and database through the following command
```bash
docker-compose up
```

After starting up the services and database through docker-compose, directly create a user in the `credentials` collection in the MongoDB `auth` database.

Afterwards you'll be able to test and utilize the services through the API endpoint. Just make sure to include the provided Bearer token in the 'Authorization' header for your following requests.


## Architectural Components

### Authentication Service
Serves as an authentication service for the users and as an API Gateway for the other services. Utilizes Gin as the request handling layer, MongoDB for data storage and retrieval, and JWT for authentication.

The JWT is provided in the format 'Bearer <token_string>' in the 'Authorization' header of the response to the login request (if successful)

#### Environmental Variables
The following environmental variables are required for the services:
- MONGO_URI
- PORT
- USER_SERVICE_END_POINT
- JWT_SECRET
- DB_NAME

You can see the local configuration of these secrets in the [auth-service .env.development file](./auth-service/.env.development).

#### Testing
To run the correctness tests execute `go test ./...` from the root of the [auth-service folder](./auth-service/).

To run the proxy benchmark test execute `go test -benchmem -run=^$ -bench ^BenchmarkHandleProxyRequest$ github.com/auth-service/proxy`. This test is located in the [proxy_controller_test.go file.](./auth-service/proxy/proxy_controller_test.go)
### Profile Service

Allows the user to create profiles and to update them. Can only be contacted by the authentication services.
Utilizes Gin as the request handling layer & MongoDB for data storage and retrieval.

#### Environmental Variables
The following environment are required for the service:
- PORT
- MONGO_URI
- ROUTER_HOST
- DB_NAME

You can see the local configuration of these secrets in the [profile-service .env.development file](./profile-service/.env.development).

#### Testing
To run the correctness tests execute `go test ./...` from the root of the [profile-service folder](./profile-service/).
