<?php

namespace Tests\Feature;

use App\Models\Task;
use App\Models\User;
use Tests\TestCase;

class TaskControllerTest extends TestCase
{
    public function test_create()
    {
        $user = User::factory()->create()->first();
        $data = [
            'name' => $this->faker->word
        ];

        $this->actingAs($user)
            ->postJson('api/tasks', $data)
            ->assertStatus(201);
    }

    public function test_create_failure_after_reaching_limit()
    {
        $user = User::factory()->create()->first();
        $this->actingAs($user);

        for ($i = 0; $i < config('task.creation_limit_per_day'); $i++) {
            $this->postJson('api/tasks', [
                'name' => $this->faker->word
            ]);
        }

        $tasks = Task::count();
        $this->assertEquals(config('task.creation_limit_per_day'), $tasks);

        $this->postJson('api/tasks', [
                'name' => $this->faker->word
            ])
            ->assertStatus(429);
    }
}
