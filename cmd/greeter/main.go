package main

import (
	"flag"
	"fmt"
	"local/custom-server/gen/restapi"
	"local/custom-server/gen/restapi/operations"
	"log"

	"github.com/go-openapi/loads"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/runtime/middleware"
)

var portFlag = flag.Int("port", 3000, "Port to run this service on")

func main() {
	swaggerSpec, err := loads.Analyzed(restapi.SwaggerJSON, "")
	if err != nil {
		log.Fatalln(err)
	}

	api := operations.NewGreeterAPI(swaggerSpec)
	server := restapi.NewServer(api)
	defer server.Shutdown()

	flag.Parse()
	server.Port = *portFlag

	api.GetGreetingHandler = operations.GetGreetingHandlerFunc(
		func(params operations.GetGreetingParams) middleware.Responder {
			name := swag.StringValue(params.Name)
			if name == "" {
				name = "World"
			}
	
			greeting := fmt.Sprintf("Hello, %s!", name)
			return operations.NewGetGreetingOK().WithPayload(greeting)
		})

	if err := server.Serve(); err != nil {
		log.Fatalln(err)
	}
}