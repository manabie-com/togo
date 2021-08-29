import { ErrorName, MessageConfig } from "./message.config";
import { HttpStatus } from "./http.status";
import { IErrorConfig } from "../core";

export const ErrorConfig: IErrorConfig[] = [
    {
        status: HttpStatus.UNAUTHORIZED,
        names: [
            {
                name: MessageConfig.UNAUTHORIZED,
                message: MessageConfig.UNAUTHORIZED
            },
            {
                name: ErrorName.JsonWebTokenError,
                message: MessageConfig.JSON_WEB_TOKEN_ERROR
            },
            {
                name: ErrorName.TokenExpiredError,
                message: MessageConfig.TOKEN_EXPIRED_ERROR
            }
        ]
    },
    {
        status: HttpStatus.FORBIDDEN,
        names: [
            {name: MessageConfig.FORBIDDEN},
            {name: MessageConfig.PERMISSION_DENIED}
        ]
    },
    {
        status: HttpStatus.BAD_REQUEST,
        names: [
            {name: MessageConfig.VALIDATION_ERROR},
            {name: MessageConfig.CAST_ERROR},
            {name: MessageConfig.MONGO_ERROR},
            {name: MessageConfig.MULTER_ERROR},
        ]
    },
    {
        status: HttpStatus.UNKNOWN_ERROR,
        names: [{
            name: MessageConfig.UNKNOWN_ERROR
        }]
    }
];
