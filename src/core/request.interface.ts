import { QueryPopulateOptions } from "mongoose";

export interface IQuery {
    condition: ICondition;
    pagination?: IPagination;
    select?: string;
    sort?: ISort;
    populate?: QueryPopulateOptions[];
    match?: ICondition;
    lean?: boolean;
    aggregations?: any[];
}

export interface IOption {
    locale?: string;
}

export interface ICondition {
    [key: string]: any;
}

export interface IPagination {
    limit: number;
    skip: number;
}

export interface ISort {
    [key: string]: 1 | -1 | 'asc' | 'desc' | '1' | '-1';
}

export interface IRangeInterface {
    $gte: string | number | Date;
    $lte: string | number | Date;
    $gt: string | number | Date;
    $lt: string | number | Date;
}

export interface IData {
    [key: string]: any;
}
