<?php
declare(strict_types=1);

namespace App\Controller\Task;

use App\Exception\TaskException;
use Slim\Http\Request;
use Slim\Http\Response;

/**
 * Class GetAllController
 * @package App\Controller\Task
 */
class GetAllController extends AbstractTaskController
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
        $userId = $this->getAndValidateUserId($params);

        $page = $request->getQueryParam('page', null);
        $perPage = $request->getQueryParam('per_page', null);
        $content = $request->getQueryParam('content', null);

        $tasks = $this->getTaskService()->getTasksByPage(
            $userId,
            (int) $page,
            (int) $perPage,
            $content
        );

        return $this->jsonResponse($response, 'success', $tasks, 200);
    }
}
