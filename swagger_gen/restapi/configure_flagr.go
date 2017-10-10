// Code generated by go-swagger; DO NOT EDIT.

package restapi

import (
	"crypto/tls"
	"net/http"
	"os"

	errors "github.com/go-openapi/errors"
	runtime "github.com/go-openapi/runtime"
	"github.com/rs/cors"
	graceful "github.com/tylerb/graceful"
	"github.com/urfave/negroni"

	"github.com/checkr/flagr/pkg/config"
	"github.com/checkr/flagr/pkg/handler"
	"github.com/checkr/flagr/swagger_gen/restapi/operations"
)

// This file is safe to edit. Once it exists it will not be overwritten

//go:generate swagger generate server --target ../swagger_gen --name  --spec ../swagger.yml

var (
	pwd, _     = os.Getwd()
	enableCORS = config.Config.CORS.Enabled
)

func configureFlags(api *operations.FlagrAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.FlagrAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.JSONConsumer = runtime.JSONConsumer()
	api.JSONProducer = runtime.JSONProducer()
	api.ServerShutdown = func() {}

	handler.Setup(api)
	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix"
func configureServer(s *graceful.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	if enableCORS {
		c := cors.New(cors.Options{
			AllowedOrigins: []string{"*"},
			AllowedHeaders: []string{"Content-Type", "Accepts"},
			AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
		})
		handler = c.Handler(handler)
	}

	n := negroni.New()
	ui := negroni.NewStatic(http.Dir(pwd + "/browser/flagr-ui/dist/"))
	n.Use(ui)
	n.Use(negroni.NewLogger())

	n.UseHandler(handler)
	return n
}
