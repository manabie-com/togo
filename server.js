const app = require('./src/app')
const config = require('./src/config')
const { sequelize } = require('./src/models')

// app.listen(config.PORT, () => {
//     console.log(`App listening at http://localhost:${config.PORT}`)
// })

sequelize.sync().then(() => {
    app.listen(config.PORT, () => {
        console.log(`App listening at http://localhost:${config.PORT}`)
    })
})