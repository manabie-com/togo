<?php
declare(strict_types=1);

namespace App\Middleware;

use App\Exception\Auth;
use Psr\Http\Message\ResponseInterface;
use Slim\Http\Request;
use Slim\Http\Response;
use Slim\Route;

/**
 * Class AuthMiddleware
 * @package App\Middleware
 */
class AuthMiddleware extends BaseMiddleware
{
    /**
     * @param Request $request
     * @param Response $response
     * @param Route $next
     * @return ResponseInterface
     * @throws Auth
     */
    public function __invoke(
        Request $request,
        Response $response,
        Route $next
    ): ResponseInterface {
        $jwtHeader = $request->getHeaderLine('Authorization');

        if (! $jwtHeader) {
            throw new Auth('JWT Token required.', 400);
        }
        $jwt = explode('Bearer ', $jwtHeader);

        if (!isset($jwt[1])) {
            throw new Auth('JWT Token invalid.', 400);
        }

        $decoded = $this->checkToken($jwt[1]);
        $object = (array) $request->getParsedBody();
        $object['decoded'] = $decoded;

        return $next($request->withParsedBody($object), $response);
    }
}
