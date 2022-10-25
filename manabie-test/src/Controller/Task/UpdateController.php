<?php
declare(strict_types=1);

namespace App\Controller\Task;

use App\Exception\TaskException;
use Slim\Http\Request;
use Slim\Http\Response;

/**
 * Class UpdateController
 * @package App\Controller\Task
 */
class UpdateController extends AbstractTaskController
{
    /**
     * @param Request $request
     * @param Response $response
     * @param array $args
     * @return Response
     * @throws TaskException
     */
    public function __invoke(Request $request, Response $response, array $args): Response
    {
        $params = (array) $request->getParsedBody();

        $this->validateParams($params);

        $task = $this->hydrator($params);
        $taskIdParam = $args['id'];
        $task->setId($taskIdParam );

        $task = $this->getTaskService()->update($task);

        return $this->jsonResponse($response, 'success', $task, 200);
    }
}
