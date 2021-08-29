import { injectable } from "inversify";
import { BaseRepository } from "./base.repository";
import { Task } from "../model";

interface ITaskRepository {
}

@injectable()
export class TaskRepository extends BaseRepository<Task> implements ITaskRepository {
    constructor() {
        super('tasks');
    }
}
