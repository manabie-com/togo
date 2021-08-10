<?php

namespace App\Services;

use Exception;
use App\Models\User;
use Illuminate\Support\Facades\Log;
use App\Services\Contracts\UserInterface;

class UserService implements UserInterface
{
    public function register(string $username, string $password): ?User
    {
        try {
            $user = new User();
            $user->username = $username;
            $user->password = bcrypt($password);
            $resultCreate = $user->save();

            if ($resultCreate) {
                return $user;
            }

            Log::warning('Register user fail, param: ' . json_encode($user));

            return null;
        } catch (Exception $exception) {
            Log::error($exception);
            return null;
        }
    }
}
