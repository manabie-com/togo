<?php

namespace Repository;

use Domain\Users\Models\User;
use Illuminate\Support\Facades\Hash;
use Repository\IRepositories\IUserRepository;

class UserRepository extends EloquentRepository implements IUserRepository
{
    public function getModel(): string
    {
        return User::class;
    }

    public function login(array $authData): ?array
    {
        $user = $this->model::query()->login($authData['user_name'])->first();
        if ($user && Hash::check($authData['password'], $user->password)) {
            $token = $user->createToken('authToken');

            return [
                'token' => $token->plainTextToken,
            ];
        }

        return null;
    }
}
