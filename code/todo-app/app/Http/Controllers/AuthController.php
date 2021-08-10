<?php

namespace App\Http\Controllers;

use App\Models\User;
use Tymon\JWTAuth\Factory;
use Illuminate\Http\Request;
use Illuminate\Http\JsonResponse;
use App\Http\Requests\LoginRequest;
use App\Http\Resources\UserResource;

class AuthController extends Controller
{
    public function login(LoginRequest $request)
    {
        $credentials = $request->only(['username', 'password']);

        if (!$token = auth()->attempt($credentials)) {
            return unauthorized();
        }

        /** @var Factory $athFactory */
        $athFactory = auth()->factory('');

        return success([
            'access_token' => $token,
            'token_type'   => User::TOKEN_TYPE,
            'expires_in'   => $athFactory->getTTL() * 60
        ]);
    }

    public function logout(): JsonResponse
    {
        auth()->logout();

        return success();
    }

    public function me(Request $request): JsonResponse
    {
        $userResource = (new UserResource(auth()->user()))->toArray($request);

        return success($userResource);
    }
}
