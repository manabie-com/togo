const { ERROR, TOKEN_EXPIRED_MSG } = require('./constant')
const { decode } = require('./jwt')

const authen = async (req, res, next) => {
    try {
        const header = req.headers.authorization
        if (!header) return res.status(ERROR.UNAUTHORIZED.status_code).send(ERROR.UNAUTHORIZED)
        const tokenArray = header.split(' ')
        if (tokenArray.length < 1) {
            return res.status(ERROR.UNAUTHORIZED.status_code).send(ERROR.UNAUTHORIZED)
        }

        const authRes = decode(tokenArray[1])
        req.user = authRes
        return next()
    } catch (error) {
        if (error.code === TOKEN_EXPIRED_MSG) {
            return res.status(ERROR.TOKEN_EXPIRED.status_code).send(ERROR.TOKEN_EXPIRED)
        }
        return res.status(ERROR.UNAUTHORIZED.status_code).send(ERROR.UNAUTHORIZED)
    }
}
module.exports = { authen } 