<?php

namespace App\Repositories;

use App\Models\Task;
use App\Repositories\EloquentRepository;

class TaskRepository extends EloquentRepository
{
    /**
     * get model
     * @return string
     */
    public function getModel()
    {
        return Task::class;
    }

    public function countTasks($userId, $startTime, $endTime)
    {
        $query = $this->_model->select('id')
            ->where('user_id', '=', $userId)
            ->whereBetween('time_created',[$startTime, $endTime]);
        return $query->count();
    }

}
