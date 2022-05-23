import { connect } from "mongoose";
import "dotenv/config";

import app from "@/app";
import { PORT } from "@/utils/constant.util";

import { OPTIONS, URL } from "config/database.config";

/**
 * Function used to start the application, the connect database must be connected and then the API and kafka will be start
 */
export const startApplication = async () => {
   const listenPort = parseInt(PORT, 10);
   await connect(URL, OPTIONS);

   app.listen(listenPort, () => {
      console.info(`Listening Port: ${listenPort}`);
   });
};

startApplication().catch((error) => {
   console.error(`\nStart app : ${error}`);
   process.exit(1);
});
