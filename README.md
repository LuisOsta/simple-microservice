# Simple Microservices
This simple project is an example of a microservice architecture. With the API gateway handling the authentication and routing, and the underlying service handling the respective business logic.

**Important Note**: In this example the same database cluster is utilized for all services. This is not a good practice, but is done for simplicity.

## Getting Started

After starting up the services and database through docker-compose, directly create a user in the `credentials` collection in the MongoDB `auth` database.

Afterwards you'll be able to test and utilize the services through the API endpoint.
```bash
docker-compose up
```

## Architectural Components

### Authentication Service
Serves as an authentication service for the users and as an API Gateway for the other services. Utilizes Gin as the request handling layer, MongoDB for data storage and retrieval, and JWT for authentication.

#### Environmental Variables
The following environmental variables are required for the services:
- MONGO_URI
- PORT
- USER_SERVICE_END_POINT
- JWT_SECRET
- DB_NAME

You can see the local configuration of these secrets in the [auth-service .env.development file](./auth-service/.env.development).

### Profile Service

Allows the user to create profiles and to update them. Can only be contacted by the authentication services.
Utilizes Gin as the request handling layer & MongoDB for data storage and retrieval.

### Environmental Variables
The following environment are required for the service:
- PORT
- MONGO_URI
- ROUTER_HOST
- DB_NAME

You can see the local configuration of these secrets in the [profile-service .env.development file](./profile-service/.env.development).