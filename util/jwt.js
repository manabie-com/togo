require('../config/environment')
const jwt = require('jsonwebtoken')

const JWT_SECRET = process.env.JWT_SECRET;
const JWT_REFRESH_SECRET = process.env.JWT_REFRESH_SECRET;

const jwtService = {
    encode: (message, refresh = false) => {
        if (refresh) {
            return jwt.sign(message, JWT_REFRESH_SECRET, { algorithm: 'HS256' }, { expiresIn: '1y' })
        } else {
            return jwt.sign(message, JWT_SECRET, { algorithm: 'HS256' }, { expiresIn: '1m' }) // I put 1M for longer testing CURL
        }
    },
    decode: (token, refresh = false) => {
        if (refresh) {
            return jwt.verify(token, JWT_REFRESH_SECRET)
        } else {
            return jwt.verify(token, JWT_SECRET)
        }
    },
    generateJwtToken(message, refresh = false) {
        return jwtService.encode(message, refresh)
    }
}

module.exports = jwtService

