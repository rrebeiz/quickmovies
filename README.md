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
`/v1/movies` creates a new movie <br>
## PATCH

## DELETE



## Endpoints WIP
### Show Movie
Returns json data about a single movie
* URL: `/v1/movies/:id`
* Method: GET
* URL Params:
  * Required: id=[int]
* Body Params: None
* Success Response:
  * Code: 200
  * Content: {"movie":{"id":1,"title":"test","runtime":100,"year":2020,"genres":["action", "adventure"]}}
* Error Response:
* Code: 500
* Content: {"error": "internal server error"


### Create Movie
Creates a new movie.
* URL: `/v1/movies`
* Method: POST
* URL Params: None
* Body Params:
  * Required:
    * `{"title":"test", "runtime":100, "year":2020, "genres":["action","adventure"]}`
* Success Response:
  * Code: 200
  * Content: {"movie":{"id":1, "title":"test"...}}
* Error Response:
  * Code: 400
  * Content: {"error": "body must not be empty"}
  * Code: 422
  * Content: {"error": {"title":"should not be empty","runtime":"should not be empty"...}}
  * Code: 500
  * Content: {"error": "internal server error"}