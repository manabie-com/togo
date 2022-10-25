<?php
declare(strict_types=1);

namespace App\Repository;

use App\Entity\Task;
use App\Exception\TaskException;

/**
 * Class TaskRepository
 * @package App\Repository
 */
class TaskRepository extends AbstractRepository implements TaskRepositoryInterface
{
    /**
     * @param string $userId
     * @return array
     */
    public function getTasksDailyByUserId(string $userId): array
    {
        $query = 'SELECT count(id) as task_count FROM `tasks` as t 
                  WHERE t.user_id = :user_id 
                  AND t.created_date >= :time_created_start AND t.created_date <= :time_created_end';

        $statement = $this->getDb()->prepare($query);
        $statement->bindParam('user_id', $userId);
        $timeCreatedStart = date("Y-m-d 00:00:00");
        $timeCreatedEnd = date("Y-m-d 23:59:59");
        $statement->bindParam('time_created_start', $timeCreatedStart);
        $statement->bindParam('time_created_end', $timeCreatedEnd);
        $statement->execute();

        $taskCount = $statement->fetch();

        return [
            'task_count_daily' => (int)$taskCount['task_count'],
            'max_todo' => $this->getLimitTaskDailyByUserId($userId)
        ];
    }

    /**
     * @param string $userId
     * @return int
     */
    private function getLimitTaskDailyByUserId(string $userId): int
    {
        $query = 'SELECT `max_todo` FROM `users` as u WHERE u.id = :id';

        $statement = $this->getDb()->prepare($query);
        $statement->bindParam('id', $userId);
        $statement->execute();

        return (int)$statement->fetch()['max_todo'];
    }

    /**
     * @return string
     */
    public function getQueryTasksByPage(): string
    {
        return "
            SELECT *
            FROM `tasks`
            WHERE `user_id` = :user_id
            AND `content` LIKE CONCAT('%', :content, '%')
            ORDER BY `id`
        ";
    }

    /**
     * @param string $userId
     * @param int $page
     * @param int $perPage
     * @param string|null $content
     * @return array
     */
    public function getTasksByPage(string $userId, int $page, int $perPage, ?string $content): array
    {
        $params = [
            'user_id' => $userId,
            'content' => is_null($content) ? '' : $content,
        ];

        $query = $this->getQueryTasksByPage();
        $statement = $this->database->prepare($query);

        $statement->bindParam('user_id', $params['user_id']);
        $statement->bindParam('content', $params['content']);

        $statement->execute();

        $total = $statement->rowCount();

        return $this->getResultsWithPagination($query, $page, $perPage, $params, $total);
    }

    /**
     * @param string $taskId
     * @param string $userId
     * @return Task
     * @throws TaskException
     */
    public function checkAndGetTask(string $taskId, string $userId): Task
    {
        $query = 'SELECT * FROM `tasks` WHERE `id` = :id AND `user_id` = :user_id';

        $statement = $this->getDb()->prepare($query);
        $statement->bindParam('id', $taskId);
        $statement->bindParam('user_id', $userId);

        $statement->execute();

        $task = $statement->fetchObject(Task::class);

        if (!$task) {
            throw new TaskException('Task not found.', 404);
        }

        return $task;
    }

    /**
     * @return array
     */
    public function getAllTasks(): array
    {
        $query = 'SELECT * FROM `tasks` ORDER BY `id`';
        $statement = $this->getDb()->prepare($query);
        $statement->execute();

        return (array)$statement->fetchAll();
    }

    /**
     * @param Task $task
     * @return Task
     * @throws TaskException
     */
    public function create(Task $task): Task
    {
        $query = '
            INSERT INTO `tasks`
                (`id`, `content`, `created_date`, `user_id`)
            VALUES
                (:id, :content, :created_date, :user_id)
        ';

        $params = [
            'id' => uniqid('task'),
            'created_date' => date("Y-m-d H:i:s"),
            'content' => $task->getContent(),
            'user_id' => $task->getUserId()
        ];

        $statement = $this->getDb()->prepare($query);

        $statement->bindParam('id', $params['id']);
        $statement->bindParam('content', $params['content']);
        $statement->bindParam('created_date', $params['created_date']);
        $statement->bindParam('user_id', $params['user_id']);

        $statement->execute();

        return $this->checkAndGetTask($params['id'], $params['user_id']);
    }

    /**
     * @param Task $task
     * @return Task
     * @throws TaskException
     */
    public function update(Task $task): Task
    {
        $query = '
            UPDATE `tasks`
            SET `content` = :content
            WHERE `id` = :id AND `user_id` = :user_id
        ';
        $statement = $this->getDb()->prepare($query);
        $id = $task->getId();

        $content = $task->getContent();
        $userId = $task->getUserId();

        $statement->bindParam('id', $id);
        $statement->bindParam('content', $content);
        $statement->bindParam('user_id', $userId);
        $statement->execute();

        return $this->checkAndGetTask($id, $userId);
    }

    /**
     * @param string $taskId
     * @param string $userId
     */
    public function delete(string $taskId, string $userId): void
    {
        $query = 'DELETE FROM `tasks` WHERE `id` = :id AND `user_id` = :user_id';

        $statement = $this->getDb()->prepare($query);
        $statement->bindParam('id', $taskId);
        $statement->bindParam('user_id', $userId);

        $statement->execute();
    }
}
