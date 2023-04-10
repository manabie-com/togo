const express = require('express')
const app = express()
const cookieParser = require('cookie-parser')
const bodyParser = require('body-parser')
const cors = require('cors')
const swaggerUi = require('swagger-ui-express');
const YAML = require('yamljs')
const swaggerDocument = YAML.load('./swagger.yaml')
const errorHandling = require('./util/error_handling');

const swaggerOption = {
    swaggerswaggerOption: {
        supportedSubmitMethods: ['get', 'put', 'post', 'delete']
    }
}

app.use(
    bodyParser.urlencoded({
        extended: false
    })
)
app.use(bodyParser.json())
app.use(cors())
app.use(express.json())
app.use(express.urlencoded({ extended: false }))
app.use(cookieParser())
app.use('/api-docs', swaggerUi.serve, swaggerUi.setup(swaggerDocument, swaggerOption))
app.use('/', require('./route/index'))
app.use(errorHandling);
module.exports = app