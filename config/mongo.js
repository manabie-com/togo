const dotenv    = require('dotenv').config();
const MONGOOSE  = require('mongoose');
// require('mongoose-double')(MONGOOSE);

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
                console.log('access denied');
            } else {
                console.log("access granted");
            }
        }
    );

    MONGOOSE.pluralize(null);
    MONGOOSE.con = connection;

    module.exports = MONGOOSE;
} catch (error) {
    console.log(error);
}