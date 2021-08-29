import { injectable } from "inversify";
import { BaseRepository } from "./base.repository";
import { User } from "../model";

interface IUserRepository {
}

@injectable()
export class UserRepository extends BaseRepository<User> implements IUserRepository {
    constructor() {
        super('users');
    }
}
