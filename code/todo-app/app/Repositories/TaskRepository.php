<?php

namespace App\Repositories;

use Carbon\Carbon;
use App\Models\Task;
use App\Repositories\Contracts\TaskRepositoryInterface;

class TaskRepository extends BaseRepository implements TaskRepositoryInterface
{
    public function model(): string
    {
        return Task::class;
    }

    public function countTaskToday(int $userId): int
    {
        return $this->count([
            'user_id' => $userId,
            ['created_at', '>=', Carbon::now()->startOfDay()],
            ['created_at', '<=', Carbon::now()->endOfDay()]
        ]);
    }
}
