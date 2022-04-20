<?php

namespace Api\Controllers;

use Api\Requests\Auth\LoginRequest;
use Illuminate\Auth\AuthenticationException;
use Illuminate\Http\JsonResponse;
use Repository\IRepositories\IUserRepository;

class AuthController extends ApiController
{
    private IUserRepository $userRepository;

    public function __construct(IUserRepository $userRepository)
    {
        $this->userRepository = $userRepository;
    }

    /**
     * Login
     *
     * @param LoginRequest $request
     * @return JsonResponse
     * @throws AuthenticationException
     */
    public function login(LoginRequest $request): JsonResponse
    {
        $loginData = $request->validated();
        $token = $this->userRepository->login($loginData);
        if (! $token) {
            throw new AuthenticationException();
        }

        return $this->httpOK($token);
    }
}
