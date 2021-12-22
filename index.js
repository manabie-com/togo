const express = require('express')
var bodyParser = require('body-parser')
var methodOverride = require('method-override')
const route = require('./src/routers/index')
const port = process.env.PORT
require('./src/db/db')
const app = express()

app.use(bodyParser.urlencoded({
    extended: true
  }))
app.use(bodyParser.json())
app.use(methodOverride())

app.use(express.json())
route(app);

app.listen(port, () => {
    console.log(`Server running on port ${port}`)
})