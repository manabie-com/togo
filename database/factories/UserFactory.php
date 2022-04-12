<?php

namespace Database\Factories;

use Domain\Users\Models\User;
use Illuminate\Database\Eloquent\Factories\Factory;

class UserFactory extends Factory
{
    /**
     * The name of the factory's corresponding model.
     *
     * @var string
     */
    protected $model = User::class;

    /**
     * Define the model's default state.
     *
     * @return array
     */
    public function definition(): array
    {
        return [
            'user_name'  => $this->faker->unique()->userName(),
            'name'       => $this->faker->name(),
            'limit_task' => $this->faker->numberBetween(20, 50),
            'password'   => 'password',
        ];
    }
}
