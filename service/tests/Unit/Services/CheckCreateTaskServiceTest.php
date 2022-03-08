<?php

namespace Tests;

use App\Services\Task\CheckCreateTaskService;
use Illuminate\Http\Response;

class CheckCreateTaskServiceTest extends TestCase
{
    
    private CheckCreateTaskService $checkCreateTaskService;

    protected function setUp(): void
    {
        parent::setUp();
        $this->checkCreateTaskService = app(CheckCreateTaskService::class);
    }

    public function testCheckCreateTaskFailedDueToUserOverLimit()
    {
        list($checkCreateTask, $code, $message) = $this->checkCreateTaskService->checkCreateTask(1);
        $this->assertFalse($checkCreateTask);
        $this->assertSame(Response::HTTP_TOO_MANY_REQUESTS, $code);
        $this->assertSame(__('app.max_limit_task'), $message);
    }

    public function testCheckCreateTaskFailedDueToUserNotRegister()
    {
        list($checkCreateTask, $code, $message) = $this->checkCreateTaskService->checkCreateTask(2);
        $this->assertFalse($checkCreateTask);
        $this->assertSame(Response::HTTP_UNPROCESSABLE_ENTITY, $code);
        $this->assertSame(__('app.user_is_not_registered_for_task'), $message);
    }

    public function testCheckCreateTaskSuccess()
    {
        list($checkCreateTask, $code, $message) = $this->checkCreateTaskService->checkCreateTask(3);
        $this->assertTrue($checkCreateTask);
        $this->assertSame(Response::HTTP_OK, $code);
        $this->assertSame(__('app.create_task_success'), $message);
    }

}
