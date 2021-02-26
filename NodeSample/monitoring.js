// https://github.com/openzipkin/zipkin-js
const {
    Tracer,
    BatchRecorder,
    jsonEncoder: {JSON_V2}
  } = require('zipkin');
const CLSContext = require('zipkin-context-cls');
const {HttpLogger} = require('zipkin-transport-http');

const recorder = new BatchRecorder({
    logger: new HttpLogger({
        endpoint: 'http://localhost:9411/api/v2/spans',
        jsonEncoder: JSON_V2
    })
})

// Setup the tracer to use http and implicit trace context
const tracer = new Tracer({
    ctxImpl: new CLSContext('zipkin'),
    recorder,
    localServiceName: 'node-sample' // name of this application
});


// now use tracer to construct instrumentation! For example, fetch

module.exports = {
    instrumentExpress : app => {
        const zipkinMiddleware = require('zipkin-instrumentation-express').expressMiddleware
        app.use(zipkinMiddleware({tracer}))        
    },
    instrumentFetch: (fetch,remoteServiceName) => {
        const wrapFetch = require('zipkin-instrumentation-fetch');
        //const remoteServiceName = 'youtube';
        const zipkinFetch = wrapFetch(fetch, {tracer, remoteServiceName});
        return zipkinFetch;
    }
}
