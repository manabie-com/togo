import { ConnectOptions } from "mongoose";

export const URL = `${process.env.DATABASE_URL}?retryWrites=true&w=majority`;

export const OPTIONS: ConnectOptions = {};
