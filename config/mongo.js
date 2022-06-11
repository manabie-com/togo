const dotenv    = require('dotenv').config();
const MONGOOSE  = require('mongoose');

try {
    if(dotenv.error) {
        throw dotenv.error;
    }

    connection_url = process.env.ATLAS_URI;

    
    const connection = MONGOOSE.createConnection(connection_url, {
            useNewUrlParser: true,
            useUnifiedTopology: true,
        },
        async function(error) {
            if (error) {
                console.log("ERROR", error);
                console.log('Access denied');
            } else {
                console.log("Db connection already established.");
            }
        }
    );

    MONGOOSE.pluralize(null);
    MONGOOSE.con = connection;

    module.exports = MONGOOSE;
} catch (error) {
    console.log(error);
}