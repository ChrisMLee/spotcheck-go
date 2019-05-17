# Spotcheck Go

### To run:  
`go run ./src/main.go`

~docker run -p 8080:8080 -it --rm --name spotcheck-go-run spotcheck-go-image~


### Golang tutorial thoughts
- importing modules
- attaching a DB instance to context

### Running with Realize



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

### GraphQL: Notes
#### Fetching data at a field-level
> child fields are responsible for fetching their own data
> This code is easy to reason about
> This code is more testable
> To some, the getEvent duplication might look like a code smell. But, having code that is simple, easy to reason about, and is more testable is worth a little bit of duplication.
via [GraphQL Resolvers: Best Practices
](https://medium.com/paypal-engineering/graphql-resolvers-best-practices-cd36fdbcef55)


### Resources
#### Go
* https://medium.com/@rrgarciach/bootstrapping-a-go-application-with-docker-47f1d9071a2a
* [interfaces](https://stackoverflow.com/questions/23148812/whats-the-meaning-of-interface)
* [packages](https://medium.com/rungo/everything-you-need-to-know-about-packages-in-go-b8bac62b74cc)
* [go packages on SO](https://stackoverflow.com/a/44016468)
* [Golang how to import local packages without gopath?](https://stackoverflow.com/questions/17539407/golang-how-to-import-local-packages-without-gopath)


View all installed go packages:  
`godoc --http :6060`  

You have to develop under ``GOPATH/src/github.com/username/project-name` if you want to develop locally and have your imports work.

View your gopath:
`go env GOPATH`


#### Go + Postgres
* https://medium.com/@vptech/complexity-is-the-bane-of-every-software-engineer-e2878d0ad45a
* https://flaviocopes.com/golang-sql-database/
* https://medium.com/@beld_pro/postgres-with-golang-3b788d86f2ef
* [Why does Go treat a Postgresql numeric & decimal columns as []uint8?
](https://stackoverflow.com/questions/31946344/why-does-go-treat-a-postgresql-numeric-decimal-columns-as-uint8)
* [How I handled possible null values from database rows in Golang?](https://medium.com/aubergine-solutions/how-i-handled-null-possible-values-from-database-rows-in-golang-521fb0ee267)
> sql.NullString
> https://golang.org/src/database/sql/sql.go?s=4941:5029#L177
* [Nullable Json types](https://gist.github.com/rsudip90/022c4ef5d98130a224c9239e0a1ab397)
* [Go database/sql tutorial](http://go-database-sql.org/index.html)
* [Inserting records into a PostgreSQL database with Go's database/sql package](https://www.calhoun.io/inserting-records-into-a-postgresql-database-with-gos-database-sql-package/)
* [Golang Insert NULL into sql instead of empty string](https://stackoverflow.com/questions/40266633/golang-insert-null-into-sql-instead-of-empty-string)

#### Go + GraphQL
* https://medium.com/@bradford_hamilton/building-an-api-with-graphql-and-go-9350df5c9356
* [Why we moved our graphQL server from Node.js to Golang](https://medium.com/safetycultureengineering/why-we-moved-our-graphql-server-from-node-js-to-golang-645b00571535)
* [Concurrent resolvers](https://github.com/graphql-go/graphql/tree/master/examples/concurrent-resolvers)
* [Getting Started With GraphQL Using Golang](https://www.thepolyglotdeveloper.com/2018/05/getting-started-graphql-golang/)
* [Dataloader Go Implementation](https://github.com/graph-gophers/dataloader)
* [GraphQL with Golang: A Deep Dive From Basics To Advanced](https://medium.freecodecamp.org/deep-dive-into-graphql-with-golang-d3e02a429ac3)
* [My first GraphQL server in Go responding to Apollo](https://medium.com/@maumribeiro/my-first-graphql-server-in-go-responding-to-apollo-bd1c11426572)

```
// Trivial resolver example
https://github.com/graphql-go/graphql/blob/796d22335788bc53c1921db39d262f5eccf5a1e2/examples/modify-context/main.go


// Instead of trying to modify context within a resolve function, use:
// `graphql.Params.RootObject` is a mutable optional variable and available on
// each resolve function via: `graphql.ResolveParams.Info.RootValue`.


// https://github.com/graphql-go/graphql/blob/199d20bbfed70dae8c7d4619d4e0d339ce738b43/definition.go
// ResolveParams Params for FieldResolveFn()
type ResolveParams struct {
	// Source is the source value
	Source interface{}

	// Args is a map of arguments for current GraphQL request
	Args map[string]interface{}

	// Info is a collection of information about the current execution state.
	Info ResolveInfo

	// Context argument is a context value that is provided to every resolve function within an execution.
	// It is commonly
	// used to represent an authenticated user, or request-specific caches.
	Context context.Context
}

type ResolveInfo struct {
	FieldName      string
	FieldASTs      []*ast.Field
	Path           *ResponsePath
	ReturnType     Output
	ParentType     Composite
	Schema         Schema
	Fragments      map[string]ast.Definition
	RootValue      interface{}
	Operation      ast.Definition
	VariableValues map[string]interface{}
}
```

#### Go + docker + postgres + elastic
https://medium.com/@leo_hetsch/local-development-with-go-postgresql-and-elasticsearch-in-docker-61bc8a0d5e66

* [Go tutorial: REST API backed by PostgreSQL](https://flaviocopes.com/golang-tutorial-rest-api/)

#### GraphQL
* [GraphQL Resolvers: Best Practices](https://medium.com/paypal-engineering/graphql-resolvers-best-practices-cd36fdbcef55)
> Fetching data at a field-level 
* [Apollo: Fetching data with resolvers](https://www.apollographql.com/docs/apollo-server/essentials/data)

> The context is how you access your shared connections and fetchers in resolvers to get data.

> The context is the third argument passed to every resolver. It is useful for passing things that any resolver may need, like authentication scope, database connections, and custom fetch functions. Additionally, if you're using dataloaders to batch requests across resolvers, you can attach them to the context as well.

> As a best practice, context should be the same for all resolvers, no matter the particular query or mutation, and resolvers should never modify it. This ensures consistency across resolvers, and helps increase development velocity.

* [Apollo: Designing GraphQL Mutations](https://blog.apollographql.com/designing-graphql-mutations-e09de826ed97)
> Specific mutations that correspond to semantic user actions are more powerful than general mutations. This is because specific mutations are easier for a UI developer to write, they can be optimized by a backend developer, and only providing a specific subset of mutations makes it much harder for an attacker to exploit your API.
> Mutations should only ever have one input argument. That argument should be named input and should have a non-null unique input object type.

#### DB
* https://medium.com/@kimtnguyen/relational-database-schema-design-overview-70e447ff66f9

#### SQL
* https://stackoverflow.com/questions/4111594/why-always-close-database-connection
> 1. Open connections as late as possible
> 2. Close connections as soon as possible
> The connection itself is returned to the connection pool. Connections are a limited and relatively expensive resource. Any new connection you establish that has exactly the same connection string will be able to reuse the connection from the pool.
* https://dba.stackexchange.com/questions/5222/why-shouldnt-we-allow-nulls


* [Permission denied for relation <table>](https://dba.stackexchange.com/questions/53914/permission-denied-for-relation-table)
```
GRANT ALL PRIVILEGES ON ALL TABLES IN SCHEMA public TO spotcheck_db_dev;
GRANT ALL PRIVILEGES ON ALL SEQUENCES IN SCHEMA public TO spotcheck_db_dev;
```

#### Docker
* [How do I pass environment variables to Docker containers?](https://stackoverflow.com/questions/30494050/how-do-i-pass-environment-variables-to-docker-containers)
* [Environment variables in Compose](https://docs.docker.com/compose/environment-variables/)


#### Google Maps
Using MySQL and PHP with Google Maps:  
* https://developers.google.com/maps/documentation/javascript/mysql-to-maps
