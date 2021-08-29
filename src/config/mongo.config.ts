import mongoose, { ConnectionOptions, ModelUpdateOptions, QueryFindOneAndUpdateOptions } from 'mongoose';
import path from "path";
import { config } from "migrate-mongo/lib/migrate-mongo";
import { BasePropOptions, IModelOptions } from "@typegoose/typegoose/lib/types";
import { Severity } from "@typegoose/typegoose";
import { StringHelper } from "../shared";
import _ from "lodash";
import moment from "moment";

export const CONNECTION_OPTIONS: ConnectionOptions = {
    useUnifiedTopology: true,
    useNewUrlParser: true,
    useFindAndModify: false,
    useCreateIndex: true,
    reconnectTries: Number.MAX_VALUE,
    reconnectInterval: 500,
    connectTimeoutMS: 10 * 1000
};

export const SCHEMA_OPTIONS: mongoose.SchemaOptions = {
    versionKey: false,
    strict: true,
    strictQuery: true,
    id: false,
    minimize: true,
    toObject: {virtuals: true},
    toJSON: {virtuals: true},
    timestamps: true
};

export const ENTITY_OPTIONS = (collection?: string, transform?: (doc: any, ret: any, options: any) => any): IModelOptions => {
    return {
        schemaOptions: {
            collection,
            versionKey: false,
            strict: true,
            strictQuery: true,
            id: false,
            minimize: true,
            timestamps: true,
            toObject: {virtuals: true},
            toJSON: {virtuals: true, transform}
        },
        options: {
            allowMixed: Severity.ALLOW
        }
    }
};

export const OBJECT_OPTIONS = (transform?: (doc: any, ret: any, options: any) => any): IModelOptions => {
    return {
        schemaOptions: {
            versionKey: false,
            strict: true,
            strictQuery: true,
            _id: false,
            id: false,
            minimize: true,
            timestamps: false,
            toObject: {virtuals: true},
            toJSON: {virtuals: true, transform}
        },
        options: {
            allowMixed: Severity.ALLOW
        }
    }
};

export const SLUG_OPTIONS: any = {
    slugPaddingSize: 4,
    unique: true,
    slugOn: {
        findOneAndUpdate: false,
        update: false,
        updateOne: false
    }
};

export const UPDATE_MANY_OPTIONS: ModelUpdateOptions = {
    runValidators: true
};

export const UPDATE_OPTIONS: QueryFindOneAndUpdateOptions = {
    runValidators: true,
    new: true
};

export const DELETE_OPTIONS: QueryFindOneAndUpdateOptions = {
    rawResult: false
};

export const SAVE_OPTIONS: QueryFindOneAndUpdateOptions = {
    upsert: true,
    new: true,
    runValidators: true,
    setDefaultsOnInsert: true
};

export const MIGRATE_CONFIG: config.Config = {
    mongodb: {
        databaseName: '',
        options: {
            useUnifiedTopology: true,
            useNewUrlParser: true,
            connectTimeoutMS: 5 * 60 * 1000,
            socketTimeoutMS: 5 * 60 * 1000,
        }
    },
    migrationsDir: path.resolve('migrations'),
    changelogCollectionName: "migrations",
    migrationFileExtension: ".js"
};
