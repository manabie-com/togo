<?php

namespace Tests\Feature;

use Domain\Tasks\Models\Task;
use Domain\Users\Models\User;
use Illuminate\Http\JsonResponse;
use Illuminate\Support\Carbon;
use Illuminate\Support\Str;
use Tests\TestCase;

class TaskTest extends TestCase
{
    /**
     *  Test if user not login
     */
    public function testIfUserNotLogin()
    {
        $response = $this->get('/api/tasks');
        $response->assertStatus(JsonResponse::HTTP_UNAUTHORIZED)
                 ->assertJsonStructure([
                     'status',
                     'success',
                     'errors' => [
                         'message',
                         'detail',
                     ],
                 ]);
    }

    /**
     * Get user's task
     */
    public function testGetTasks()
    {
        $user = User::query()->first();
        $response = $this->actingAs($user)->get('/api/tasks');

        $response->assertStatus(JsonResponse::HTTP_OK)
                 ->assertJsonStructure([
                     'status',
                     'success',
                     'data'       => [
                         [
                             'id',
                             'name',
                             'user_id',
                             'description',
                             'created_at',
                         ],
                     ],
                     'pagination' => [
                         'count',
                         'total',
                         'perPage',
                         'currentPage',
                         'totalPages',
                         'links' => [],
                     ],
                 ]);
    }

    /**
     *  Test validate
     */
    public function testRequireName()
    {
        $user = User::query()->first();
        $task = [
            'name'        => '',
            'description' => $this->faker->realText(250),
        ];
        $response = $this->actingAs($user)->post('/api/tasks', $task);

        $response->assertStatus(JsonResponse::HTTP_BAD_REQUEST)
                 ->assertJson([
                     'status'  => JsonResponse::HTTP_BAD_REQUEST,
                     'success' => false,
                     'errors'  => [
                         'message' => __('errors.err_400'),
                         'detail'  => [
                             [
                                 'field'  => 'name',
                                 'detail' => __('validation.required', ['attribute' => 'name']),
                             ],
                         ],
                     ],
                 ]);
    }

    /**
     *  Test full task
     */
    public function testFullTaskToDay()
    {
        // Get user
        $user = User::query()->inRandomOrder()->first();
        // Get limit_task
        $limit = $user->limit_task;
        // Current task
        $currentTask = $user->tasks()->count();
        // Add dummy tasks to full
        $left = $limit - $currentTask;
        $task = [];
        for ($i = 0; $i < $left; $i++) {
            $task[] = [
                'id'          => Str::orderedUuid(),
                'name'        => $this->faker->realText(50),
                'description' => $this->faker->realText(250),
                'user_id'     => $user->id,
                'created_at'  => Carbon::now(),
            ];
        }
        Task::query()->insert($task);

        // Body
        $task = [
            'name'        => $this->faker->realText(50),
            'description' => $this->faker->realText(250),
        ];
        $response = $this->actingAs($user)->post('/api/tasks', $task);

        $response->assertStatus(JsonResponse::HTTP_BAD_REQUEST)
                 ->assertJson([
                     'status'  => JsonResponse::HTTP_BAD_REQUEST,
                     'success' => false,
                     'errors'  => [
                         'message' => __('errors.err_400'),
                         'detail'  => [
                             [
                                 'field'  => 'task_number',
                                 'detail' => __('validation.max.numeric', [
                                     'attribute' => 'task number',
                                     'max'       => $limit,
                                 ]),
                             ],
                         ],
                     ],
                 ]);
    }

    /**
     *  Test create new task
     */
    public function testCreateNewTask()
    {
        // Get user
        $user = User::query()->inRandomOrder()->first();

        // Body
        $task = [
            'name'        => $this->faker->realText(50),
            'description' => $this->faker->realText(250),
        ];
        $response = $this->actingAs($user)->post('/api/tasks', $task);

        $response->assertStatus(JsonResponse::HTTP_OK)
                 ->assertJsonStructure([
                     'status',
                     'success',
                     'data' => [
                         'id',
                         'name',
                         'description',
                         'created_at',
                     ],
                 ]);
    }

    /**
     *  Test update task
     */
    public function testUpdateTask()
    {
        // Get user
        $user = User::query()->inRandomOrder()->first();
        // Task
        $task = $user->tasks()->first();
        // Body
        $newName = $this->faker->realText(50);
        $data = [
            'name' => $newName,
        ];
        $response = $this->actingAs($user)->put('/api/tasks/'.$task->id, $data);

        $response->assertStatus(JsonResponse::HTTP_OK)
                 ->assertJson([
                     'status'  => JsonResponse::HTTP_OK,
                     'success' => true,
                     'data'    => [
                         'id'          => $task->id,
                         'name'        => $newName,
                         'description' => $task->description,
                         'created_at'  => $task->created_at,
                     ],
                 ]);
    }

    /**
     *  Test delete task
     */
    public function testDeleteTask()
    {
        // Get user
        $user = User::query()->inRandomOrder()->first();
        // Task
        $task = $user->tasks()->first();

        $response = $this->actingAs($user)->delete('/api/tasks/'.$task->id);

        $response->assertStatus(JsonResponse::HTTP_OK)
                 ->assertJson([
                     'status'  => JsonResponse::HTTP_OK,
                     'success' => true,
                     'data'    => [],
                 ]);
    }
}
