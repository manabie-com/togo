<?php
declare(strict_types=1);

namespace App\Controller\Task;

use App\Controller\BaseController;
use App\Entity\Task;
use App\Exception\TaskException;
use App\Service\Task\TaskService;

/**
 * Class AbstractTaskController
 * @package App\Controller\Task
 */
abstract class AbstractTaskController extends BaseController
{
    /**
     * @return TaskService
     */
    protected function getTaskService(): TaskService
    {
        return $this->container->get('task_service');
    }

    /**
     * @param array $params
     * @return string
     * @throws TaskException
     */
    protected function getAndValidateUserId(array $params): string
    {
        if (isset($params['decoded']) && isset($params['decoded']->sub)) {
            return $params['decoded']->sub;
        }

        throw new TaskException('Invalid user. Permission failed.', 400);
    }

    /**
     * @param array $params
     * @throws TaskException
     */
    protected function validateParams(array $params)
    {
        $data = json_decode((string)json_encode($params), false);

        if (!isset($data->content)) {
            throw new TaskException('The field "content" is required.', 400);
        }
    }

    /**
     * @param array $params
     * @return Task
     */
    protected function hydrator(array $params): Task
    {
        $task = new Task();
        $task->setContent($params['content']);
        $task->setUserId($params['decoded']->sub);

        return $task;
    }
}
