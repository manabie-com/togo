<?php

namespace App\Repositories\Tasks;

use App\Models\Task;
use App\Repositories\Tasks\TaskRepository;

class TaskEloquentRepository implements TaskRepository
{
    public function create(array $data)
    {
        return Task::create([
            'name' => $data['name'],
            'user_id' => $data['user_id']
        ]);
    }
}