<?php
declare(strict_types=1);

namespace App\Controller\Task;

use App\Exception\TaskException;
use Slim\Http\Request;
use Slim\Http\Response;

/**
 * Class CreateController
 * @package App\Controller\Task
 */
class CreateController extends AbstractTaskController
{
    /**
     * @param Request $request
     * @param Response $response
     * @return Response
     * @throws TaskException
     */
    public function __invoke(Request $request, Response $response): Response
    {
        $params = (array) $request->getParsedBody();

        $this->validateParams($params);
        
        $task = $this->hydrator($params);

        $this->getTaskService()->validateTaskLimitDailyByUserId($task->getUserId());
        
        $task = $this->getTaskService()->create($task);

        return $this->jsonResponse($response, 'success', $task, 201);
    }
}
