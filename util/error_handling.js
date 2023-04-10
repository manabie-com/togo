

const NODE_ENV = process.env.NODE_ENV;
const { storeLog } = require('./log')
const { COMMON_ERROR_CODE } = require('./constant')

/**
 * Error handling middleware.
 */

module.exports = (err, req, res, next) => {
    // console.log('ERROR ', err)
    if (!res.headersSent) { // response not sent yet
        if (err.status_code && COMMON_ERROR_CODE[err.status_code]) {
            return res.status(err.status_code).send({ error_msg: err.error_msg });
        }
        // save unexpected error
        storeLog({
            type: 'ERROR',
            source: `${req.path}`,
            description: `API Unexpted ERROR`,
            err
        })
        return res.status(500).json({ error_msg: ['CONTACT_SUPPORT'] });
    }
};