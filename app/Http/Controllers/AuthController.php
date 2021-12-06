<?php

namespace App\Http\Controllers;

use App\Http\Requests\Auth\LoginRequest;
use App\Http\Requests\Auth\RegisterRequest;
use App\Repositories\Users\UserRepository;
use Illuminate\Http\Request;
use Illuminate\Http\JsonResponse;
use Illuminate\Support\Facades\Auth;
use Illuminate\Support\Facades\Hash;

class AuthController extends Controller
{
    /**
     * Login
     * 
     * @param LoginRequest $request
     * Responsible for validating the request
     */
    public function login(LoginRequest $request)
    {
        $authResult = Auth::attempt([
            'email' => $request->email,
            'password' => $request->password
        ]);

        if (! $authResult) {
            return new JsonResponse([
                'error' => 'Incorrect Email or Password'
            ], 422);
        }

        $token = $request->user()
            ->createToken('API Token')
            ->plainTextToken;

        return new JsonResponse([
            'token' => $token
        ]);
    }


    /**
     * Logout
     */
    public function logout(Request $request)
    {
        $request->user()->currentAccessToken()->delete();
        return new JsonResponse([], 204);
    }


    /**
     * Register
     * 
     * @param RegisterRequest $request
     * Responsible for validating the request
     * 
     * @param UserRepository $userRepository
     * Responsible for data access related to users
     */
    public function register(
        RegisterRequest $request,
        UserRepository $userRepository
    ) {
        $data = $request->all();
        $data['password'] = Hash::make($data['password']);
        $user = $userRepository->create($data); 
        return new JsonResponse([
            'success' => 'User successfully registered',
            'user' => $user
        ], 201);
    }
}
