# QuickMovies
## Welcome to QuickMovies

### About QuickMovies
QuickMovies is the backend which allows you to create and manage your own library of movies and users.

### Setup
Please keep in mind that it's still work in progress

Database:
To make life easier I suggest exporting the DSN locally with `export DSN=postgres://devuser:password@localhost/go_movies?sslmode=disable` <br>
Run `docker-compose up -d` to create a local postgresql DB in docker. <br>
Then you can use a tool like migrate found here `https://github.com/golang-migrate/migrate/releases` to create the tables with `migrate -path ./migrations -database $DSN up`

### Starting the server
There are several flags that can be passed to change things like the default port, environment, database connection info ect.<br>
It is best to configure these directly in the provided makefile, which currently uses the defaults.

* `make start` will start the server.
* `make restart` will restart the server.
* `make stop` will stop the server. 

Once the server is up you can use Postman, or curl to send requests. A frontend written in either Vue or React is also in the works & will be committed to the project.

## Available endpoints (WIP, more endpoints will be added and or endpoints changed.)

## GET
`/healthcheck` returns status info <br>

`/v1/movies/:id` returns a movie by ID <br>


## POST

## PATCH

## DELETE