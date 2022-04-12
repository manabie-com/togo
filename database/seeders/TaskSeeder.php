<?php

namespace Database\Seeders;

use Domain\Tasks\Models\Task;
use Domain\Users\Models\User;
use Illuminate\Database\Seeder;

class TaskSeeder extends Seeder
{
    /**
     * Run the database seeds.
     *
     * @return void
     */
    public function run()
    {
        User::query()->each(function ($user) {
            Task::factory(rand(10, 20))->create([
                'user_id' => $user->id,
            ]);
        });
    }
}
