import {createLogger, transports, format} from 'winston'
import configs from '../Configs'

const {printf, colorize} = format;
const myFormat = printf( info => {
    return `[${info.timestamp}] [${info.level}]: ${info.message}`
} );
const logger = createLogger( {
    transports: [
        // info console log
        new transports.Console( {
            level: 'info',
            name: 'info-console',
            format: format.combine(
                format.timestamp( {
                    format: 'YYYY-MM-DD HH:mm:ss' // Optional for choosing your own timestamp format.
                } ),
                colorize(),
                myFormat
            )

        } ),
        // info log file
        new transports.File( {
            level: 'info',
            name: 'info-file',
            filename: configs.pathLogAccess,
            format: format.combine(
                format.timestamp( {
                    format: 'YYYY-MM-DD HH:mm:ss' // Optional for choosing your own timestamp format.
                } ),
                myFormat
            ),
            json: false
        } ),
        // errors console log
        new transports.Console( {
            level: 'error',
            name: 'error-console',
            format: format.combine(
                format.timestamp( {
                    format: 'YYYY-MM-DD HH:mm:ss' // Optional for choosing your own timestamp format.
                } ),
                colorize(),
                myFormat
            )
        } ),
        // errors log file
        new transports.File( {
            level: 'error',
            name: 'error-file',
            filename: configs.pathLogError,
            format: format.combine(
                format.timestamp( {
                    format: 'YYYY-MM-DD HH:mm:ss' // Optional for choosing your own timestamp format.
                } ),
                myFormat
            )
        } ),
        // errors console log
        new transports.Console( {
            level: 'warn',
            name: 'warn-console',
            colorize: true,
            format: format.combine(
                format.timestamp( {
                    format: 'YYYY-MM-DD HH:mm:ss' // Optional for choosing your own timestamp format.
                } ),
                colorize(),
                myFormat
            )
        } ),
        // errors log file
        new transports.File( {
            level: 'warn',
            name: 'warn-file',
            filename: configs.pathLogWarning,
            format: format.combine(
                format.timestamp( {
                    format: 'YYYY-MM-DD HH:mm:ss' // Optional for choosing your own timestamp format.
                } ),
                myFormat
            ),
            json: false
        } )
    ]
} );
export default logger;
