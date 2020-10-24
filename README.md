# simpleapi

This repo contains the code to run a REST API server. The availables endpoints are the following: 

- GET /hello -> returns a string containing the message "Hello, world"
- POST /shortenen -> returns the shortened version of a url using https://cleanuri.com. 

The project runs on localhost:8080.

Clone the project into your workspace directory, this project is written in Go, using modules. You must have a module compatible version of Go in order to run this project (1.11+). 

To run the project: 
```go
  go run main.go
```

Use a tool like postman or curl to test the API. 

For example, to test GET /hello
![Image of hello endopint](https://raw.githubusercontent.com/jegacs/simpleapi/main/hello.jpeg)

Otherwise, to test the POST /shortenen endpoint:
![Image of shortenen endopint](https://raw.githubusercontent.com/jegacs/simpleapi/main/shortenen.jpeg)

To develop the REST API, net/package was used. To test it, net/http/httptest native package was used to mocke requests onto the API. Also, the package smartystreets/goconvey was used to help with some assertions. 
