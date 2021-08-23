<?php
declare(strict_types=1);

namespace App\Controller\User;

use App\Exception\UserException;
use Slim\Http\Request;
use Slim\Http\Response;

/**
 * Class LoginController
 * @package App\Controller\User
 */
class LoginController extends AbstractUserController
{
    /**
     * @param Request $request
     * @param Response $response
     * @return Response
     * @throws UserException
     */
    public function __invoke(Request $request, Response $response): Response
    {
        $params = (array) $request->getParsedBody();

        $this->validateParams($params);

        $jwt = $this->getLoginUserService()->login($params);

        $message = [
            'Authorization' => 'Bearer ' . $jwt,
        ];

        return $this->jsonResponse($response, 'success', $message, 200);
    }

    /**
     * @param array $params
     * @throws UserException
     */
    private function validateParams(array $params): void
    {
        if (! isset($params['id'])) {
            throw new UserException('The field "id" is required.', 400);
        }

        if (! isset($params['password'])) {
            throw new UserException('The field "password" is required.', 400);
        }
    }
}
