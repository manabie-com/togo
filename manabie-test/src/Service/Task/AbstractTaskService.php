<?php
declare(strict_types=1);

namespace App\Service\Task;

use App\Entity\Task;
use App\Exception\TaskException;
use App\Repository\TaskRepository;
use App\Service\AbstractService;
use App\Service\RedisService;
use Respect\Validation\Validator as v;

/**
 * Class AbstractTaskService
 * @package App\Service\Task
 */
abstract class AbstractTaskService extends AbstractService
{
    private const REDIS_KEY = 'task:%s:user:%s';

    /**
     * @var TaskRepository
     */
    protected TaskRepository $taskRepository;

    /**
     * @var RedisService
     */
    protected RedisService $redisService;

    /**
     * @param TaskRepository $taskRepository
     * @param RedisService $redisService
     */
    public function __construct(
        TaskRepository $taskRepository,
        RedisService $redisService
    ) {
        $this->taskRepository = $taskRepository;
        $this->redisService = $redisService;
    }

    /**
     * @return TaskRepository
     */
    protected function getTaskRepository(): TaskRepository
    {
        return $this->taskRepository;
    }

    /**
     * @param string $taskId
     * @return string
     * @throws TaskException
     */
    protected static function validateTaskId(string $taskId): string
    {
        if (! v::length(1, 500)->validate($taskId)) {
            throw new TaskException('Invalid name.', 400);
        }

        return $taskId;
    }

    /**
     * @param string $taskId
     * @param string $userId
     * @return object
     * @throws TaskException
     */
    protected function getTaskFromCache(string $taskId, string $userId): object
    {
        $redisKey = sprintf(self::REDIS_KEY, $taskId, $userId);
        $key = $this->redisService->generateKey($redisKey);
        if ($this->redisService->exists($key)) {
            $task = $this->redisService->get($key);
        } else {
            $task = $this->getTaskFromDb($taskId, $userId)->toJson();
            $this->redisService->setex($key, $task);
        }

        return $task;
    }

    /**
     * @param string $taskId
     * @param string $userId
     * @return Task
     * @throws TaskException
     */
    protected function getTaskFromDb(string $taskId, string $userId): Task
    {
        return $this->getTaskRepository()->checkAndGetTask($taskId, $userId);
    }

    /**
     * @param string $taskId
     * @param string $userId
     * @param object $task
     */
    protected function saveInCache(string $taskId, string $userId, object $task): void
    {
        $redisKey = sprintf(self::REDIS_KEY, $taskId, $userId);
        $key = $this->redisService->generateKey($redisKey);
        $this->redisService->setex($key, $task);
    }

    /**
     * @param string $taskId
     * @param string $userId
     */
    protected function deleteFromCache(string $taskId, string $userId): void
    {
        $redisKey = sprintf(self::REDIS_KEY, $taskId, $userId);
        $key = $this->redisService->generateKey($redisKey);
        $this->redisService->del([$key]);
    }
}
