const 
    express = require("express"),
    app = express(),
    monitoring = require("./monitoring")
    PORT = 3000,
    starWarsRepo = require("./starWarsRepo")
;

monitoring.instrumentExpress(app)

app.get("/",(req, res) => {
    res.json({status: "ok"})
        .status(200)
        .end()
})

app.get("/people/:id",async (req, res) => {
    try {
        const id = req.params.id
        const people = await starWarsRepo.findPeople(id)
        res.send(people).end()
    } catch (error){
        res.json(Error).status(500).end()
    }
})

app.listen(PORT, () => {
    console.log(`App started ${PORT}`)
})
