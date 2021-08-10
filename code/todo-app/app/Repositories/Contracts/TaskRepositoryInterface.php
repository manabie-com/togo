<?php

namespace App\Repositories\Contracts;

interface TaskRepositoryInterface extends RepositoryInterface
{
    public function countTaskToday(int $userId);
}
