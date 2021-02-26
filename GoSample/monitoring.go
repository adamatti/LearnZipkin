package main

import (
	"log"
	"net/http/httptest"

	mux "github.com/gorilla/mux"
	zipkin "github.com/openzipkin/zipkin-go"
	zipkinhttp "github.com/openzipkin/zipkin-go/middleware/http"
	reporter "github.com/openzipkin/zipkin-go/reporter"
	reporterhttp "github.com/openzipkin/zipkin-go/reporter/http"
)

const (
	zipkinHTTPEndpoint = "http://localhost:9411/api/v2/spans"
	serviceName        = "go-sample"
)

type ZipkinObjects struct {
	client *zipkinhttp.Client
	tracer *zipkin.Tracer
}

func buildReporter() reporter.Reporter {
	// set up a span reporter
	//reporter := logreporter.NewReporter(log.New(os.Stderr, "", log.LstdFlags))
	return reporterhttp.NewReporter(
		zipkinHTTPEndpoint,
		reporterhttp.RequestCallback(requestCallback),
	)
	//defer reporter.Close()
}

func monitoring(router *mux.Router) ZipkinObjects {
	// set up a span reporter
	//reporter := logreporter.NewReporter(log.New(os.Stderr, "", log.LstdFlags))
	reporter := buildReporter()
	//defer reporter.Close()

	// create our local service endpoint
	endpoint, err := zipkin.NewEndpoint(serviceName, "localhost:0")
	if err != nil {
		log.Fatalf("unable to create local endpoint: %+v\n", err)
	}

	// Sampler tells you which traces are going to be sampled or not. In this case we will record 100% (1.00) of traces.
	sampler, err := zipkin.NewCountingSampler(1)
	if err != nil {
		log.Fatalf("unable to create local sampler: %+v\n", err)
	}

	// initialize our tracer
	tracer, err := zipkin.NewTracer(
		reporter,
		zipkin.WithSampler(sampler),
		zipkin.WithLocalEndpoint(endpoint),
	)

	if err != nil {
		log.Fatalf("unable to create tracer: %+v\n", err)
	}

	// create global zipkin http server middleware
	serverMiddleware := zipkinhttp.NewServerMiddleware(
		tracer,
		zipkinhttp.TagResponseSize(true),
	)

	// create global zipkin traced http client
	client, err := zipkinhttp.NewClient(tracer, zipkinhttp.ClientTrace(true))
	if err != nil {
		log.Fatalf("unable to create client: %+v\n", err)
	}

	// start web service with zipkin http server middleware
	if router != nil {
		ts := httptest.NewServer(serverMiddleware(router))
		defer ts.Close()
	}

	ctx := ZipkinObjects{}
	ctx.client = client
	ctx.tracer = tracer
	return ctx
}
