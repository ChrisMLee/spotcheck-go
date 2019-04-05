# Spotcheck Go

docker run -p 8080:8080 -it --rm --name spotcheck-go-run spotcheck-go-image

### DB
Entities:
User
Spot

spotchecker
users
spots

### Docker
> Compose is a tool for defining and running multi-container Docker applications. With Compose, you use a YAML file to configure your application’s services. Then, with a single command, you create and start all the services from your configuration. 

```
docker-compose up
```
### Go
> Slices can be created with the built-in make function; this is how you create dynamically-sized arrays.
> The make function allocates a zeroed array and returns a slice that refers to that array
https://tour.golang.org/moretypes/13

### Resources
#### Go
* https://medium.com/@rrgarciach/bootstrapping-a-go-application-with-docker-47f1d9071a2a

#### Go + Postgres
* https://medium.com/@vptech/complexity-is-the-bane-of-every-software-engineer-e2878d0ad45a
* https://flaviocopes.com/golang-sql-database/
* https://medium.com/@beld_pro/postgres-with-golang-3b788d86f2ef

#### Go + GraphQL
* https://medium.com/@bradford_hamilton/building-an-api-with-graphql-and-go-9350df5c9356

#### Go + docker + postgres + elastic
https://medium.com/@leo_hetsch/local-development-with-go-postgresql-and-elasticsearch-in-docker-61bc8a0d5e66

* [Go tutorial: REST API backed by PostgreSQL](https://flaviocopes.com/golang-tutorial-rest-api/)

#### DB
* https://medium.com/@kimtnguyen/relational-database-schema-design-overview-70e447ff66f9

#### SQL
* https://stackoverflow.com/questions/4111594/why-always-close-database-connection
> 1. Open connections as late as possible
> 2. Close connections as soon as possible
> The connection itself is returned to the connection pool. Connections are a limited and relatively expensive resource. Any new connection you establish that has exactly the same connection string will be able to reuse the connection from the pool.

#### Docker
* [How do I pass environment variables to Docker containers?](https://stackoverflow.com/questions/30494050/how-do-i-pass-environment-variables-to-docker-containers)
* [Environment variables in Compose](https://docs.docker.com/compose/environment-variables/)


#### Google Maps
Using MySQL and PHP with Google Maps:  
* https://developers.google.com/maps/documentation/javascript/mysql-to-maps
