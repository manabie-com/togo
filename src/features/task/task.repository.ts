import { EntityRepository, Repository } from "typeorm";
import { TaskEntity } from "./entities/task.entity";

@EntityRepository(TaskEntity)
export class TaskRepository extends Repository<TaskEntity> {
    async getNumberOfTaskByUserId(userId: number) {
        const [{ numberoftasks }] = await TaskEntity.query(`
            SELECT COUNT(id) as numberOfTasks 
            FROM tasks
            WHERE "userId" = $1
            AND "createdAt"::date = now()::date
        `, [userId]);

        return numberoftasks;
    }
}