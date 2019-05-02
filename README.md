# Spotcheck Go

docker run -p 8080:8080 -it --rm --name spotcheck-go-run spotcheck-go-image

### DB
Entities:
User
Spot

spotchecker
users
spots

user:
- id
- email
- username
- created
- updated

#### "Seed"
```
CREATE TABLE users(
  id SERIAL NOT NULL PRIMARY KEY,
  email text not null unique,
  username text
);

CREATE TABLE spots(
  id SERIAL NOT NULL PRIMARY KEY,
  name text,
  image_url text,
  description text,
  address text,
  lat DECIMAL,
  lng DECIMAL,
  user_id INTEGER REFERENCES users(id)
);

INSERT INTO users (email, username) VALUES ('scruffmcgruff@gmail.com', 'smooth_operator');
INSERT INTO users (email, username) VALUES ('cheezgawd@gmail.com', 'icecube');

SELECT * from users;

INSERT INTO spots (name, image_url, description, address, lat, lng, user_id) VALUES (
  'terrible down ledge',
  NULL,
  'a very bad down ledge into a door',
  NULL,
  40.702590, 
  -73.992887,
  1
), (
  'ledge on pier',
  NULL,
  'an okay ledge next to some basketball courts',
  NULL,
  40.699387, 
  -73.998415,
  1
), (
  'step ledge',
  NULL,
  'a two stair manny pad and jank ledge',
  NULL,
  40.680627, 
  -73.991448,
  1
);

```

### GoogleMaps API
- LatLng type
https://godoc.org/google.golang.org/genproto/googleapis/type/latlng

- Postgres geographic data points
https://tapoueh.org/blog/2018/05/postgresql-data-types-point/

### Docker
> Compose is a tool for defining and running multi-container Docker applications. With Compose, you use a YAML file to configure your application’s services. Then, with a single command, you create and start all the services from your configuration. 

```
docker-compose up
```
### Go
> Slices can be created with the built-in make function; this is how you create dynamically-sized arrays.
> The make function allocates a zeroed array and returns a slice that refers to that array
https://tour.golang.org/moretypes/13

- Golang - Asterisk and Ampersand Cheatsheet
https://gist.github.com/josephspurrier/7686b139f29601c3b370

### Resources
#### Go
* https://medium.com/@rrgarciach/bootstrapping-a-go-application-with-docker-47f1d9071a2a
* [interfaces](https://stackoverflow.com/questions/23148812/whats-the-meaning-of-interface)

#### Go + Postgres
* https://medium.com/@vptech/complexity-is-the-bane-of-every-software-engineer-e2878d0ad45a
* https://flaviocopes.com/golang-sql-database/
* https://medium.com/@beld_pro/postgres-with-golang-3b788d86f2ef
* [Why does Go treat a Postgresql numeric & decimal columns as []uint8?
](https://stackoverflow.com/questions/31946344/why-does-go-treat-a-postgresql-numeric-decimal-columns-as-uint8)
* https://medium.com/aubergine-solutions/how-i-handled-null-possible-values-from-database-rows-in-golang-521fb0ee267
> sql.NullString
> https://golang.org/src/database/sql/sql.go?s=4941:5029#L177
* [Nullable Json types](https://gist.github.com/rsudip90/022c4ef5d98130a224c9239e0a1ab397)


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
* https://dba.stackexchange.com/questions/5222/why-shouldnt-we-allow-nulls

#### Docker
* [How do I pass environment variables to Docker containers?](https://stackoverflow.com/questions/30494050/how-do-i-pass-environment-variables-to-docker-containers)
* [Environment variables in Compose](https://docs.docker.com/compose/environment-variables/)


#### Google Maps
Using MySQL and PHP with Google Maps:  
* https://developers.google.com/maps/documentation/javascript/mysql-to-maps
