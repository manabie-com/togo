<?php

namespace Repository;

use Domain\Tasks\Models\Task;
use Repository\IRepositories\ITaskRepository;

class TaskRepository extends EloquentRepository implements ITaskRepository
{
    public function getModel(): string
    {
        return Task::class;
    }
}
