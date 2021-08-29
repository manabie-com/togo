import { Response } from 'express';
import { ErrorConfig, HttpStatus, MessageConfig } from "../config";
import moment from "moment";

interface IBodyBuilder {
    status(status: number): ApiResponse;

    message(message: string): ApiResponse;

    data(data: any): ApiResponse;

    error(error: Error | any): ApiResponse;

    build(): void;
}

interface IResponse {
    timestamp: Date;
    status: boolean;
    message: string;
    data?: any;
    error?: Error | any;
    stack?: any;
}

export class ApiResponse implements IBodyBuilder {

    private response: Response;
    private _status!: number;
    private _message!: string;
    private _data!: any;
    private _error!: Error | any;
    private _stack?: any;

    private constructor(response: Response) {
        this.response = response;
    }

    static create(response: Response): ApiResponse {
        return new ApiResponse(response);
    }

    status(status: number): ApiResponse {
        this._status = status;
        return this;
    }

    message(message: string): ApiResponse {
        this._message = message;
        return this;
    }

    data(data: any | null = null): ApiResponse {
        this._data = data;
        return this;
    }

    ok(message?: string): ApiResponse {
        this._status = HttpStatus.OK;
        this._message = message ?? MessageConfig.SUCCESS;
        return this;
    }

    notOk(message: string): ApiResponse {
        this._status = HttpStatus.BAD_REQUEST;
        this._message = message;
        return this;
    }

    notFound(message: string = MessageConfig.NOT_FOUND): ApiResponse {
        this._status = HttpStatus.NOT_FOUND;
        this._message = message;
        return this;
    }

    forbidden(message: string = MessageConfig.FORBIDDEN): ApiResponse {
        this._status = HttpStatus.FORBIDDEN;
        this._message = message;
        return this;
    }

    unauthorized(message: string = MessageConfig.UNAUTHORIZED): ApiResponse {
        this._status = HttpStatus.UNAUTHORIZED;
        this._message = message;
        return this;
    }

    error(error: Error | any = null): ApiResponse {
        if (!error) {
            this._status = HttpStatus.UNKNOWN_ERROR;
            this._message = MessageConfig.UNKNOWN_ERROR;
            return this;
        }

        if (error.response) {
            this._status = HttpStatus.BAD_REQUEST;
            this._message = error.response.data.message || '';
            this._error = error.response;
            this._data = null;
            if (this._message.includes('There are not any pos session of')) {
                this._message = 'POS_SESSION_NOT_AVAILABLE';
            }
            return this;
        }

        this._status = HttpStatus.EXPECTATION_FAILED;
        this._message = error.message ?? error.name;
        this._data = null;
        this._error = error;

        for (let errorConfig of ErrorConfig) {
            for (let errorName of errorConfig.names) {
                if (errorName.name.toLowerCase() === error.name?.toLowerCase()) {
                    this._status = errorConfig.status;
                    if (errorName.message) this._message = errorName.message;
                }
            }
        }

        if (error.stack) {
            this._stack = error.stack;
        }

        return this;
    }

    build() {
        let body: IResponse = {
            timestamp: new Date(),
            status: this._status === HttpStatus.OK,
            message: this._message
        };

        if (this._data) {
            body.data = this._data;
        }

        if (this._error) {
            body.error = this._error;
        }

        if (this._stack) {
            body.stack = this._stack;
        }

        if (this._status !== HttpStatus.OK) {
            return this.response.status(this._status).json({error: body.message}).end();
        }

        this.response.set('Content-Language', moment.locale());
        return this.response.status(this._status).json(body.data).end();
    }
}
