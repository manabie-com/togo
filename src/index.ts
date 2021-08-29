import "reflect-metadata";
import { ApiResponse, AppContainer, BaseMiddleware, PostgresClient } from "./core";
import { HttpStatus, ServerMode } from "./config";
import dotenv from "dotenv";
import cluster, { Worker } from 'cluster';
import os from 'os';
import { InversifyExpressServer } from "inversify-express-utils";
import express, { Request, Response } from "express";
import cookieParser from 'cookie-parser';

import './controller/home.controller'
import './controller/auth.controller'
import './controller/task.controller';

export class Server {

    private readonly mode: string;
    private readonly port: number;
    private readonly host: string;

    constructor() {
        dotenv.config();
        this.mode = String(process.env.SERVER_MODE) || 'fork';
        this.host = String(process.env.SERVER_HOST) || '0.0.0.0';
        this.port = Number(process.env.SERVER_PORT) || 3033;
    }

    public init(): express.Application {
        const app: express.Application = Server.createServer();

        if (this.mode == ServerMode.fork) {
            this.runFork(app);
            return app;
        }

        if (this.mode == ServerMode.cluster) {
            cluster.isMaster ? this.runClusterMaster() : this.runFork(app);
            return app;
        }

        throw new Error("Server mode is not correct");
    }

    public static createServer(): express.Application {
        const server = new InversifyExpressServer(AppContainer.Load());
        server.setConfig((instance) => {
            instance.use('/api-docs/swagger', express.static('assets/swagger'));
            instance.use('/api-docs/swagger/assets', express.static('node_modules/swagger-ui-dist'));
            instance.use("/assets", express.static('assets'));
            instance.use("/upload", express.static('upload'));
            instance.use(express.json({limit: "5mb"}));
            instance.use(express.urlencoded({limit: "5mb", extended: false}));
            instance.use(BaseMiddleware.accessLog);
            instance.use(BaseMiddleware.dotObject);
            // instance.use(BaseMiddleware.camelcaseKey);
            instance.use(BaseMiddleware.cors);
            instance.use(cookieParser());
        });

        server.setErrorConfig((instance) => {
            instance.use((error: Error, req: Request, res: Response) => {
                return ApiResponse.create(res).status(HttpStatus.INTERNAL_SERVER_ERROR).error(error).build();
            });
        });

        return server.build();
    }

    private runFork(app: express.Application) {
        app.listen(this.port, this.host, () => {
            console.log(`Server is listening in process ${process.pid}`);
            PostgresClient.newInstance().connect().then();
        }).on('error', (err: Error) => {
            console.error(`Server fail in process ${process.pid} with error:`, err);
        });
    }

    private runClusterMaster() {
        console.log(`Master ${process.pid} is running on ${this.host}:${this.port}...`);
        let numCPUs = os.cpus().length;
        for (let i = 0; i < numCPUs; i++) {
            cluster.fork();
        }

        cluster.on('exit', (deadWorker: Worker, code: number, signal: string) => {
            // Log dead work
            const oldPID = deadWorker.process.pid;
            console.log(`Worker ${oldPID} died with code ${code} and signal ${signal}`);

            // Restart new worker and log
            const worker = cluster.fork();
            const newPID = worker.process.pid;
            console.log(`Worker ${newPID} born with code ${code} and signal ${signal}`);
        });
    }
}

new Server().init();
