<?php

namespace Tests\Unit\Repositories;

use App\Http\Repositories\TodoCounterRepository;
use App\Http\Repositories\UserRepository;
use App\Models\TodoCounter;
use App\Models\User;
use Faker\Factory as Faker;
use Illuminate\Support\Facades\Config;
use Illuminate\Support\Facades\DB;
use Tests\TestCase;

class TodoCounterTest extends TestCase
{

    private $user;
    private $todoCounter;

    private $todoCounterRepo;
    private $userRepo;

    public function setUp() : void
    {
        parent::setUp();
        DB::beginTransaction();

        $faker = Faker::create();

        // initialize repo
        $this->todoCounterRepo = new TodoCounterRepository(new TodoCounter());
        $this->userRepo = new UserRepository(new User());

        // prepare test data
        $this->user = $this->userRepo->create([
            'name' =>  $faker->unique()->name,
            'email' =>  $faker->unique()->email,
            'password' => bcrypt($faker->unique()->password)
        ]);
        $this->todoCounter = $this->todoCounterRepo->create([
            'user_id' => $this->user->id,
            'max_count' => Config::get('constants.default_counter')
        ]);

    }

    public function tearDown() : void
    {
        DB::rollback();
        parent::tearDown();
    }

    // TEST GET TO DO COUNTER BY FIELD
    public function testGetByField() {
       // call get by field
       $todoCounter = $this->todoCounterRepo->getByField(['user_id' => $this->user->id]);
        // check returned object is an instance of To do Class
        $this->assertInstanceOf(TodoCounter::class, $todoCounter);
        // check returned data from the object
        $this->assertEquals($this->user->id, $todoCounter->user_id);
        // check whether the data exists in db or not
        $this->assertDatabaseHas('todo_counters', $todoCounter->toArray());
    }

    // TEST UPDATE BY FIELD
    public function testUpdateByField()  {
        $faker = Faker::create();
        $newCounter = $faker->numberBetween(1,5);
        // call update by field
        $this->todoCounterRepo->updateByField(
            ['user_id' => $this->user->id],
            ['count_today' => $newCounter]
        );
        // call get by field
        $this->todoCounter = $this->todoCounterRepo->getByField(['user_id' => $this->user->id]);
        // test
        $this->assertEquals($newCounter, $this->todoCounter->count_today);

    }
}
