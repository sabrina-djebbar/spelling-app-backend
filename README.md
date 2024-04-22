# spelling-app-backend

## Running the system

### running locally with main.go

#### First set up the databases, run the following commands:

```shell
  docker run --name pg-container -e POSTGRES_PASSWORD=secret -p 5432:5432 -d postgres
  docker exec -ti pg-container createdb -U postgres users
  docker exec -ti pg-container psql -U postgres
```
this creates a postgres instance with a database names user with the user postgres

to connect to the user db run 
```
\c users
```
from here you can run any postsql query

#### Now to the microservice

to run a single microservice, in the root directory run 
```shell
go run main.go [microservice_name] api
```

then run any http command within the services, http file with the url `http:localhost:{{port}}/{{endpoint}}`