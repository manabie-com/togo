<?php
declare(strict_types=1);

namespace App\Service\User;

use App\Exception\UserException;
use Firebase\JWT\JWT;

/**
 * Class Login
 * @package App\Service\User
 */
class Login extends AbstractUserService
{
    /**
     * @param array $params
     * @return string
     * @throws UserException
     */
    public function login(array $params): string
    {
        $password = hash('sha512', $params['password']);

        $user = $this->userRepository->loginUser($params['id'], $password);

        $token = [
            'sub' => $user->getId(),
            'iat' => time(),
            'exp' => time() + (7 * 24 * 60 * 60),
        ];

        return JWT::encode($token, $_SERVER['SECRET_KEY']);
    }
}
