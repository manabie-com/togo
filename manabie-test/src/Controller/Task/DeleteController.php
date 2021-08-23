<?php
declare(strict_types=1);

namespace App\Controller\Task;

use App\Exception\TaskException;
use Slim\Http\Request;
use Slim\Http\Response;

/**
 * Class DeleteController
 * @package App\Controller\Task
 */
class DeleteController extends AbstractTaskController
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
        $input = (array) $request->getParsedBody();
        $taskId = $args['id'];

        $userId = $this->getAndValidateUserId($input);

        $this->getTaskService()->delete($taskId, $userId);

        return $this->jsonResponse($response, 'success', null, 204);
    }
}
