const { connectDB } = require('./config/db')
const http = require('http')
const app = require('./app')
const port = process.env.PORT
let server;

const startServer = () => {
    connectDB()
    server = http.createServer(app)
    server.listen(process.env.PORT, () => {
        try {
            console.log(`Server start at port ${port} at mode ${process.env.NODE_ENV}`)
        } catch (error) {
            console.log('App failed to start error ', error)
            process.exit(1)
        }
    })
}

startServer()