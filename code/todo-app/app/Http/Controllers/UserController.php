<?php

namespace App\Http\Controllers;

use Illuminate\Http\JsonResponse;
use App\Services\Contracts\UserInterface;
use App\Http\Requests\UserRegisterRequest;

class UserController extends Controller
{
    protected $userService;

    public function __construct(UserInterface $userService)
    {
        $this->userService = $userService;
    }

    public function register(UserRegisterRequest $request): JsonResponse
    {
        $user = $this->userService->register($request->input('username'), $request->input('password'));

        if ($user) {
            return success(['id' => $user->id]);
        }

        return fail();
    }
}
