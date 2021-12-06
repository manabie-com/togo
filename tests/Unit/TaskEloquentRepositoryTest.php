<?php

namespace Tests\Unit;

use App\Models\User;
use App\Models\Task;
use App\Repositories\Tasks\TaskEloquentRepository;
use App\Repositories\Tasks\TaskRepository;
use Mockery\MockInterface;
use Tests\TestCase;

class TaskEloquentRepositoryTest extends TestCase
{
    public function test_create()
    {
        $user = User::factory()->create();
        $data = [
            'name' => $this->faker->word,
            'user_id' => $user->id
        ];
        $taskRepository = new TaskEloquentRepository();
        $task = $taskRepository->create($data);

        $this->assertDatabaseHas('tasks', [
            'name' => $data['name']
        ]);
    }
}
