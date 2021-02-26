const fetch = require("node-fetch"),
      monitoring = require("./monitoring")
;

exports.findPeople = async id => {
    const url = `https://swapi.dev/api/people/${id}/`
    const caller = monitoring.instrumentFetch(fetch,"star-wars")
    return (await caller(url)).json()
} 

