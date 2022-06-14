import express from 'express'
import bodyParser from 'body-parser'
import cookieSession from 'cookie-session'
import 'express-async-errors'
import { authRoutes } from './module/auth/routes'
import { taskRoutes } from './module/task/routes'
import { errorHandler } from './middlewares/error-handler'
import { NotFoundError } from './errors/not-found-error'



const app = express()
app.use(bodyParser.urlencoded({
    extended: true
}))
app.use(bodyParser.json())
app.set('trust proxy', true)
app.use(cookieSession({
    signed: false,
    secure: false
}))
app.use(authRoutes)
app.use(taskRoutes)
app.all('*', async (req, res) => {
    throw new NotFoundError()
})


app.use(errorHandler)

export {app}
