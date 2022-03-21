// Packages and initializing the application
const express = require('express');
const cors = require('cors');
const morgan = require('morgan')
const colors = require('colors')
const dotenv = require('dotenv')
const { notFound, errorHandler } = require('./middlewares/error')
const connectDB = require('./data/config/db')
const userRoutes = require('./routes/userRoutes')

dotenv.config()

connectDB()

const app = express()

if (process.env.NODE_ENV === 'development') {
  app.use(morgan('dev'))
}

// Middleware for JSON
app.use(express.json())
app.use(express.urlencoded({extended:true}));

app.use(cors());

// Route
app.use('/api', userRoutes)

app.use(notFound)
app.use(errorHandler)

const PORT = process.env.PORT || 3000;

app.listen(
  PORT,
  console.log(
    `Server running in ${process.env.NODE_ENV} mode on port ${PORT}`.yellow.bold
  )
)