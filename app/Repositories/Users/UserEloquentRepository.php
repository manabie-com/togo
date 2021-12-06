<?php

namespace App\Repositories\Users;

use App\Models\Task;
use App\Models\User;
use App\Repositories\Users\UserRepository;
use Carbon\Carbon;

class UserEloquentRepository implements UserRepository
{
    public function create(array $data)
    {
        return User::create([
            'name' => $data['name'],
            'email' => $data['email'],
            'password' => $data['password']
        ]);
    }

    public function checkIfUserCanCreateTask(int $userId)
    {
        $taskCount = Task::where('user_id', $userId)
            ->whereDate('created_at', Carbon::today())
            ->count();

        if ($taskCount >= config('task.creation_limit_per_day')) {
            return false;
        }

        return true;
    }
}