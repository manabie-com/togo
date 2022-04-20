<?php

namespace Tests\Unit;

use Domain\Users\Models\User;
use Tests\TestCase;

class TaskTest extends TestCase
{
    public function testCheckMaxWhenUserCreateNewTask()
    {
        $user = User::query()->inRandomOrder()->first();
        $limit = $user->limit_task;
        $currentTask = $user->tasks()->count();
        $this->assertTrue($limit >= $currentTask);
        if ($limit >= $currentTask) {
            $this->assertTrue(true);
        } else {
            $this->assertFalse(false);
        }
    }
}
