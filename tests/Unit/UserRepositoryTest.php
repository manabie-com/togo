<?php

namespace Tests\Unit;

use App\Models\Task;
use App\Models\User;
use App\Repositories\Users\UserEloquentRepository;
use Illuminate\Support\Facades\Hash;
use Tests\TestCase;

class UserRepositoryTest extends TestCase
{
    public function test_create()
    {
        $password = $this->faker->regexify('[A-Za-z0-9]{10}');
        $data = [
            'name' => $this->faker->name,
            'email' => $this->faker->email,
            'password' => Hash::make($password),
        ];

        $userRepository = new UserEloquentRepository();
        $user = $userRepository->create($data);

        $this->assertDatabaseHas('users', [
            'name' => $data['name'],
            'email' => $data['email']
        ]);

        $this->assertTrue(Hash::check(
            $password,
            $user->password)
        );
    }

    public function test_user_can_create_task()
    {
        $user = User::factory()
            ->has(Task::factory()->count(1))
            ->create();

        $userRepository = new UserEloquentRepository();
        $return = $userRepository->checkIfUserCanCreateTask($user->id);

        $this->assertTrue($return);
    }

    public function test_user_cant_create_task()
    {
        $user = User::factory()
            ->has(Task::factory()->count(5))
            ->create();

        $userRepository = new UserEloquentRepository();
        $return = $userRepository->checkIfUserCanCreateTask($user->id);

        $this->assertFalse($return);
    }
}
