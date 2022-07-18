<?php

namespace Tests\Feature;

use App\Models\TodoList;
use Illuminate\Foundation\Testing\RefreshDatabase;
use Illuminate\Foundation\Testing\WithFaker;
use Tests\TestCase;

class ListApiTest extends TestCase
{
    use RefreshDatabase;
    /**
     * A basic feature test example.
     *
     * @return void
     */
    public function test_example()
    {
        $response = $this->get('/');

        $response->assertStatus(200);
    }

    /**
     * @test
     */
    public function test_api_list_store()
    {
        $this->withoutExceptionHandling();
        $lists = TodoList::factory()->make();
        $response = $this->post(route('api.lists.submit'), array(
            'task'=>$lists->task,
            'description'=>$lists->description,
            'is_complete'=>$lists->is_complete
        ));
        $response->assertStatus(201);
        //db has one 1 recode just created
        $this->assertCount(1, TodoList::all());
//        $response->assertJson((array)$lists);
    }
}
