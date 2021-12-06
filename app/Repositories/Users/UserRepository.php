<?php

namespace App\Repositories\Users;

interface UserRepository
{
    public function create(array $data);
    public function checkIfUserCanCreateTask(int $userId);
}