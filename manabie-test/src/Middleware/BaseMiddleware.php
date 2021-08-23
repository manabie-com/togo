<?php
declare(strict_types=1);

namespace App\Middleware;

use App\Exception\Auth;
use Firebase\JWT\JWT;

/**
 * Class BaseMiddleware
 * @package App\Middleware
 */
abstract class BaseMiddleware
{
    /**
     * @param string $token
     * @return object
     * @throws Auth
     */
    protected function checkToken(string $token): object
    {
        try {
            return JWT::decode($token, $_SERVER['SECRET_KEY'], ['HS256']);
        } catch (\UnexpectedValueException $exception) {
            throw new Auth('Forbidden: you are not authorized.', 403);
        }
    }
}
