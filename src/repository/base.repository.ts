import { PostgresClient } from "../core";
import _ from "lodash";
import { QueryConfig } from "pg";

interface IBasicRepository {
    find(query: any): Promise<any[]>;

    findOne(query: any): Promise<any | null>;

    create(body: any): Promise<any>;
}

export abstract class BaseRepository<T> implements IBasicRepository {

    private readonly tableName: string;

    protected constructor(tableName: string) {
        this.tableName = tableName;
    }

    public async find(query: any): Promise<any[]> {
        const client = await PostgresClient.newInstance().connectClient();
        const keys: string[] = [];
        const values: any[] = [];
        Object.keys(query).map((key, idx) => {
            keys.push(`${key} = $${idx + 1}`);
            values.push(query[key]);
        });

        let sqlQuery = `SELECT * FROM ${this.tableName}`;
        if (!_.isEmpty(query)) {
            sqlQuery = `SELECT * FROM ${this.tableName} WHERE ${keys.join(' AND ')}`;
        }

        const queryConfig: QueryConfig = {text: sqlQuery, values: values}
        const res = await client.query(queryConfig);
        client.release();
        return res.rows;
    }

    public async findOne(query: any): Promise<any | null> {
        const client = await PostgresClient.newInstance().connectClient();
        const keys: string[] = [];
        const values: any[] = [];
        Object.keys(query).map((key, idx) => {
            keys.push(`${key} = $${idx + 1}`);
            values.push(query[key]);
        });

        let sqlQuery = `SELECT * FROM ${this.tableName}`;
        if (!_.isEmpty(query)) {
            sqlQuery = `SELECT * FROM ${this.tableName} WHERE ${keys.join(' AND ')}`;
        }

        const queryConfig: QueryConfig = {text: sqlQuery + ' LIMIT 1', values: values}
        const res = await client.query(queryConfig);
        client.release();
        return res.rows[0];
    }

    public async create(body: any): Promise<any> {
        const client = await PostgresClient.newInstance().connectClient();
        const keys: string[] = [];
        const values: any[] = [];
        const idxValues: string[] = [];
        Object.keys(body).map((key, idx) => {
            keys.push(key);
            values.push(body[key]);
            idxValues.push(`$${idx + 1}`);
        });

        const queryConfig: QueryConfig = {
            text: `INSERT INTO ${this.tableName}(${keys.join(', ')}) VALUES(${idxValues.join(', ')}) RETURNING *`,
            values: values,
        }
        const res = await client.query(queryConfig);
        client.release();
        return res.rows[0];
    }

    public async delete(query: any): Promise<any> {
        const client = await PostgresClient.newInstance().connectClient();
        const keys: string[] = [];
        const values: any[] = [];
        Object.keys(query).map((key, idx) => {
            keys.push(`${key} = $${idx + 1}`);
            values.push(query[key]);
        });

        let sqlQuery = `DELETE FROM ${this.tableName} WHERE 1=1`;
        if (!_.isEmpty(query)) {
            sqlQuery = `DELETE FROM ${this.tableName} WHERE ${keys.join(' AND ')}`;
        }

        const queryConfig: QueryConfig = {text: sqlQuery, values: values}
        const res = await client.query(queryConfig);
        client.release();
        return res.rowCount;
    }
}
