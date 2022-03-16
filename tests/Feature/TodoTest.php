<?php


namespace Tests\Feature;


use App\Http\Repositories\TodoCounterRepository;
use App\Http\Repositories\UserRepository;
use App\Models\TodoCounter;
use App\Models\User;
use Faker\Factory as Faker;
use Illuminate\Support\Facades\Config;
use Illuminate\Support\Facades\DB;
use Laravel\Passport\Passport;
use Tests\TestCase;

class TodoTest extends TestCase
{
    private $user;
    private $todoCounter;

    private $userRepo;
    private $todoCounterRepo;

    public function setUp() : void
    {
        parent::setUp();
        DB::beginTransaction();

        $faker = Faker::create();

        $this->userRepo = new UserRepository(new User());
        $this->todoCounterRepo = new TodoCounterRepository(new TodoCounter());

        $this->user = $this->userRepo->create([
            'name' =>  $faker->unique()->name,
            'email' =>  $faker->unique()->email,
            'password' => bcrypt($faker->unique()->password)
        ]);
        $this->todoCounter = $this->todoCounterRepo->create([
            'user_id' => $this->user->id,
            'max_count' => Config::get('constants.default_counter')
        ]);
        Passport::actingAs($this->user);

    }

    public function tearDown() : void
    {
        DB::rollback();
        parent::tearDown();
    }

    public function testAddTask() {
        $tokenResult = $this->user->createToken('Personal Access Token');
        $response = $this->post(
            '/api/v1/todo',
            [
                'content' => "asfsdfsdf"
            ],
            [
                'Authorization' => "Bearer $tokenResult->token"
            ],

        );


        $response->assertStatus(200);
    }

}
