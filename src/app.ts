/**
 * Required External Modules
 */
import compression from "compression";
import cors from "cors";
import express from "express";
import expressJSDocSwagger from "express-jsdoc-swagger";
import morgan from "morgan";

import APIRoute from "@/routes";
import { CODE_404 } from "@/utils/responseStatus.util";

import { OPTIONS_SWAGGER } from "config/constant.config";

/**
 * App Variables
 */
const app = express();
app.disable("x-powered-by");

/**
 *  App Configuration
 */
app.use(compression());
app.use(morgan("short"));
app.use(cors());
app.use(express.json());

expressJSDocSwagger(app)(OPTIONS_SWAGGER);

/**
 * Server Activation
 */
app.use("/api/", APIRoute);

// catch 404 and forward to error handler
app.use((_, res) => res.status(404).send(CODE_404()));

export default app;
