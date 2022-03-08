<?php

namespace Tests;

use Illuminate\Http\Response;
use Illuminate\Support\Facades\DB;

class TaskControllerTest extends TestCase
{

    public function testCreateTaskFailedDueToValidationErrors()
    {
        $this->json('POST', 'api/task', [
            'user_id' => 1,
        ])->assertResponseStatus(Response::HTTP_BAD_REQUEST);

        $this->json('POST', 'api/task', [
            'name' => 'test',
        ])->assertResponseStatus(Response::HTTP_BAD_REQUEST);
    }

    public function testCreateTaskFailedDueToUserOverLimit()
    {
        $this->json('POST', 'api/task', [
            'user_id' => 1,
            'name' => 'test',
        ])->assertResponseStatus(Response::HTTP_TOO_MANY_REQUESTS);
    }

    public function testCreateTaskFailedDueToUserNotRegister()
    {

        $this->json('POST', 'api/task', [
            'user_id' => 2,
            'name' => 'test',
        ])
        ->assertResponseStatus(Response::HTTP_UNPROCESSABLE_ENTITY);
    }

    public function testCreateTaskSuccess()
    {
        $this->json('POST', 'api/task', [
            'user_id' => 3,
            'name' => 'test',
        ])
        ->seeInDatabase('task', ['user_id' => 3, 'name' => 'test'])
        ->assertResponseOk();
        
        DB::table('task')->where('user_id', 3)->where('name', 'test')->delete();

    }

}
