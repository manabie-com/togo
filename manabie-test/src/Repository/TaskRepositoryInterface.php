<?php
declare(strict_types=1);

namespace App\Repository;

use App\Entity\Task;

/**
 * Interface TaskRepository
 * @package App\Repository
 */
interface TaskRepositoryInterface
{
    /**
     * @param string $userId
     * @return array
     */
    public function getTasksDailyByUserId(string $userId): array;

    /**
     * @param string $userId
     * @param int $page
     * @param int $perPage
     * @param string|null $content
     * @return array
     */
    public function getTasksByPage(string $userId, int $page, int $perPage, ?string $content): array;

    /**
     * @param string $taskId
     * @param string $userId
     * @return Task
     */
    public function checkAndGetTask(string $taskId, string $userId): Task;

    /**
     * @param Task $task
     * @return Task
     */
    public function create(Task $task): Task;

    /**
     * @param Task $task
     * @return Task
     */
    public function update(Task $task): Task;

    /**
     * @param string $taskId
     * @param string $userId
     */
    public function delete(string $taskId, string $userId): void;
}
