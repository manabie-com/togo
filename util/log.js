const storeLog = ({ type, source, description, error }) => {
    // store error log to db and notify dev
    const log = {
        type,
        source,
        description,
        error
    }
    return log;
}

module.exports = { storeLog }