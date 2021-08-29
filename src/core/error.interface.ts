export interface IErrorName {
    name: string;
    message?: string;
}

export interface IErrorConfig {
    status: number;
    names: IErrorName[];
}
