import express from 'express'
import bodyParser from 'body-parser'
import fs from 'fs'
import path from 'path'
import dotenv from 'dotenv'
import helmet from 'helmet'
import hpp from 'hpp'
import hsts from 'hsts'
import cors from 'cors'
import restFul from './Routes/api'
import https from 'https'
import http from 'http'
import toobusy from './Helpers/toobusy'
import useragent from 'express-useragent'
// import {requestLog} from './Helpers/middlewares'
import expressip from 'express-ip'
import dir from 'node-dir'
import multer from 'multer'
//call env
dotenv.config(path.resolve(process.cwd(), '.env'))
const app = express()
//security
if (process.env.TOOBUSY) {
    app.use(toobusy)
}
app.use(hpp())
if (process.env.NODE_ENV !== 'development') {
    app.use(
        hsts({
            maxAge: 300,
            includeSubDomains: true,
            preload: true,
        }),
    )
}
//app.use(helmet.contentSecurityPolicy())
app.use(helmet.hidePoweredBy())
app.use(helmet.xssFilter())
app.use(helmet.frameguard('deny'))
app.use(helmet.ieNoOpen())
app.use(helmet.noSniff())

app.use(cors())//hỗ trợ crossdomain
app.use(useragent.express())
app.use(expressip().getIpInfoMiddleware)
// app.use("*", (req, res, next) => {
//     //requestLog(req)
//     next()
// })

const controllerPath = {
    graphql: path.resolve(__dirname, '../app/Controllers/graphql'),
    restFul: path.resolve(__dirname, '../app/Controllers/api'),
    web: path.resolve(__dirname, '../app/Controllers/web'),
}

/* API restful */
app.use(bodyParser.json())
app.use(bodyParser.raw())
app.use(bodyParser.text({ type: 'text/html' }))
app.use(multer().array())
app.use(bodyParser.urlencoded({extended: true}))

const staticPath = path.resolve(process.cwd(), 'public/')
app.use(express.static(staticPath))

const restFulControllers = dir.files(controllerPath.restFul, {sync: true})
if (restFulControllers.length > 0) {
    let RC = []
    restFulControllers.forEach(file => {
        let filename = path.basename(file)
        let name = filename.split(".")
        if (!['Controller.js'].includes(filename) && name.pop() === 'js') {
            const objClass = require(path.resolve(__dirname, file))
            if (objClass)
                RC[name.join('.').replace('Controller', '')] = objClass.objClass
        }
    })
    restFul(app, RC)
}

/**
 * permission all
 */
app.get('*', (req, res) => {
    return res.json({
        "message": "restricted access"
    })
})

//run schedule
//runSchedule()

//start app
const port = process.env.NODE_PORT || 3000
if (process.env.NODE_ENV === "development") {
    if (process.env.ENABLED_HTTPS) {
        let key = fs.readFileSync(path.resolve(process.cwd(), 'selfsigned.key'))
        let cert = fs.readFileSync(path.resolve(process.cwd(), 'selfsigned.crt'))
        let options = {
            key: key,
            cert: cert,
        }
        let server = https.createServer(options, app)
        server.listen(port, () => console.log(`Now browse to https://localhost:${port}`))
    }
} else {
    app.listen(port, () => console.log(`Now browse to http://localhost:${port}`))
}

export default app

