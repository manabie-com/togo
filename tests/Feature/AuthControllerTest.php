<?php

namespace Tests\Feature;

use App\Models\User;
use Tests\TestCase;

class AuthControllerTest extends TestCase
{
    public function test_login()
    {
        $user = User::factory()->create(); 
        $data = [
            'email' => $user->email,
            'password' => 'password'
        ];

        $this->postJson('/api/login', $data)
            ->assertStatus(200);
    }

    public function test_register()
    {
        $password = $this->faker->regexify('[A-Za-z0-9]{10}');
        $data = [
            'name' => $this->faker->name,
            'email' => $this->faker->email,
            'password' => $password,
            'password_confirmation' => $password
        ];

        $this->postJson('/api/register', $data)
            ->assertStatus(201);
    }
}
