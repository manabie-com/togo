<?php

namespace Tests\Unit\Repositories;

use App\Http\Repositories\TodoRepository;
use App\Models\Todo;
use Illuminate\Support\Facades\DB;
use Tests\TestCase;
use Faker\Factory as Faker;


class TodoTest extends TestCase
{
    protected $todo;
    protected $todoRepository;

    public function setUp() : void
    {
        parent::setUp();
        DB::beginTransaction();
        $faker = Faker::create();
        // prepare test data
        $this->todo = [
            'user_id' =>  $faker->unique()->numberBetween(1, 20),
            'content' => $faker->realText(10, 2),
        ];
        // initialize repo
        $this->todoRepository = new TodoRepository(new Todo());
    }

    public function tearDown() : void
    {
        DB::rollback();
        parent::tearDown();
    }

    // TEST CREATE TO DO
    public function testStore()
    {
        // call create function
        $todo = $this->todoRepository->create($this->todo);
        // check returned object is an instance of To do Class
        $this->assertInstanceOf(Todo::class, $todo);
        // check returned data from the object
        $this->assertEquals($this->todo['user_id'], $todo->user_id);
        $this->assertEquals($this->todo['content'], $todo->content);
        // check whether the data exists in db or not
        $this->assertDatabaseHas('todos', $todo->toArray());
    }
}
