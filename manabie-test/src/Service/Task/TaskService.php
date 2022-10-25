<?php
declare(strict_types=1);

namespace App\Service\Task;

use App\Entity\Task;
use App\Exception\TaskException;

/**
 * Class TaskService
 * @package App\Service\Task
 */
class TaskService extends AbstractTaskService
{
    /**
     * @param string $userId
     * @param int $page
     * @param int $perPage
     * @param string|null $content
     * @return array
     */
    public function getTasksByPage(string $userId, int $page, int $perPage, ?string $content): array
    {
        if ($page < 1) {
            $page = 1;
        }

        if ($perPage < 1) {
            $perPage = self::DEFAULT_PER_PAGE_PAGINATION;
        }

        return $this->getTaskRepository()->getTasksByPage($userId, $page, $perPage, $content);
    }

    /**
     * @return array
     */
    public function getAllTasks(): array
    {
        return $this->getTaskRepository()->getAllTasks();
    }

    /**
     * @param string $taskId
     * @param string $userId
     * @return object
     * @throws TaskException
     */
    public function getOne(string $taskId, string $userId): object
    {
        if (self::isRedisEnabled() === true) {
            return $this->getTaskFromCache($taskId, $userId);
        }

        return $this->getTaskFromDb($taskId, $userId)->toJson();
    }

    /**
     * @param Task $task
     * @return object
     * @throws TaskException
     */
    public function create(Task $task): object
    {
        $task = $this->getTaskRepository()->create($task);

        if (self::isRedisEnabled() === true) {
            $this->saveInCache($task->getId(), $task->getUserId(), $task->toJson());
        }

        return $task->toJson();
    }

    /**
     * @param string $userId
     * @throws TaskException
     */
    public function validateTaskLimitDailyByUserId(string $userId): void
    {
        $tasksDailyData = $this->getTaskRepository()->getTasksDailyByUserId($userId);

        if ($tasksDailyData['task_count_daily'] >= $tasksDailyData['max_todo']) {
            throw new TaskException(sprintf('You only create %d tasks each day!', $tasksDailyData['max_todo']), 400);
        }
    }

    /**
     * @param Task $task
     * @return object
     * @throws TaskException
     */
    public function update(Task $task): object
    {
        $this->getTaskFromDb($task->getId(), $task->getUserId());

        $task = $this->getTaskRepository()->update($task);

        if (self::isRedisEnabled() === true) {
            $this->saveInCache($task->getId(), $task->getUserId(), $task->toJson());
        }

        return $task->toJson();
    }

    /**
     * @param string $taskId
     * @param string $userId
     * @throws TaskException
     */
    public function delete(string $taskId, string $userId): void
    {
        $this->getTaskFromDb($taskId, $userId);
        $this->getTaskRepository()->delete($taskId, $userId);

        if (self::isRedisEnabled() === true) {
            $this->deleteFromCache($taskId, $userId);
        }
    }
}
