<?php

namespace Tests\Feature;

use Domain\Users\Models\User;
use Illuminate\Http\JsonResponse;
use Tests\TestCase;

class AuthTest extends TestCase
{
    /**
     *  Test login
     */
    public function testLogin()
    {
        $user = User::query()->first();
        $response = $this->post('/api/auth/login', [
            'user_name' => $user->user_name,
            'password'  => 'password',
        ]);
        $response->assertStatus(JsonResponse::HTTP_OK)->assertJsonStructure([
            'status',
            'success',
            'data' => [
                'token',
            ],
        ]);
    }
}
