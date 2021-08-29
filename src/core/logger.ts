import { Writable } from "stream";
import { injectable } from "inversify";
import winston, { Logger as WLogger, transports } from "winston";
import { TransformableInfo } from "logform";
import DailyRotateFile from 'winston-daily-rotate-file';
import _ from "lodash";

interface ILogger {
    debug(message: any, meta?: any);

    info(message: any, meta?: any);

    warn(message: any, meta?: any);

    error(message: any, meta?: any);
}

const stream = new Writable({
    objectMode: false,
    write: raw => console.log('stream msg>>>', raw.toString()) // phat trực tiếp log tại đây
});

@injectable()
export class Logger implements ILogger {

    private logger: WLogger;

    constructor() {
        this.logger = winston.createLogger({
            level: 'verbose',
            format: winston.format.combine(
                winston.format.label({label: 'MyLife'}),
                winston.format.timestamp({format: 'YYYY-MM-DD HH:mm:ss'}),
                winston.format.metadata({fillExcept: ['message', 'level', 'timestamp', 'label', 'service']}),
                winston.format.errors({stack: true, response: true}),
                winston.format.align(),
                winston.format.printf((log: TransformableInfo) => {
                    let out: string = `${log.timestamp} [${log.label}] ${log.level.toUpperCase()}`;
                    if (log.metadata.response) return `${out}: ${log.message} ${log.metadata.response}`;
                    if (log.metadata.stack) return `${out}: ${log.metadata.stack}`;
                    if (log.metadata.data) {
                        const data = log.metadata.data;
                        if (_.isNil(data)) return `${out}: ${log.message}`;
                        if (_.isObject(data) || _.isArray(data)) return `${out}: ${log.message} ${JSON.stringify(log.metadata.data)}`;
                        return `${out}: ${log.message} ${String(log.metadata.data)}`;
                    }
                    return `${out}: ${log.message}`;
                }),
                winston.format.colorize({all: true}),
            ),
            defaultMeta: {data: null},
            transports: [
                new DailyRotateFile({
                    filename: './logs/combined/%DATE%-combined.log',
                    datePattern: 'YYYY-MM-DD',
                    zippedArchive: true,
                    maxSize: '20m',
                    maxFiles: '14d'
                }),
                new DailyRotateFile({
                    level: 'error',
                    filename: './logs/error/%DATE%-error.log',
                    datePattern: 'YYYY-MM-DD',
                    zippedArchive: true,
                    maxSize: '20m',
                    maxFiles: '14d'
                })
            ],
        });

        if (process.env.NODE_ENV !== 'production') {
            this.logger.add(new transports.Console({}));
        }
    }

    public static instance(): Logger {
        return new Logger();
    }

    debug(message: any, meta?: any) {
        this.logger.log('debug', message, this.parseMeta(meta));
    }

    info(message: any, meta?: any) {
        this.logger.log('info', message, this.parseMeta(meta));
    }

    warn(message: any, meta?: any) {
        this.logger.log('warn', message, this.parseMeta(meta));
    }

    error(message: any, meta?: any) {
        this.logger.log('error', message, this.parseMeta(meta));
    }

    private parseMeta(data?: any) {
        if (!data) return {data: null};
        if (data instanceof Error) return data;
        return {data};
    }
}
