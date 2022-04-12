<?php

namespace Tests\Feature;

use Illuminate\Http\JsonResponse;
use Tests\TestCase;

class UserTest extends TestCase
{
    /**
     *  Get users list
     */
    public function testGetAllUser()
    {
        $response = $this->get('/api/users');
        $response->assertStatus(JsonResponse::HTTP_OK)->assertJsonStructure([
            'status',
            'success',
            'data' => [
                [
                    'id',
                    'user_name',
                    'name',
                    'limit_task',
                    'created_at',
                ],
            ],
        ]);
    }
}
