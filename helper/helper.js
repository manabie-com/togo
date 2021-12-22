`use strict`;

// response object
async function createResponseObject(event, flag, data, message) {
    if (!message) {
        message = "";
    }
    if (!data) {
        data = {};
    }
    return {
        event, flag, data, message
    };
};

module.exports = {
    createResponseObject
};