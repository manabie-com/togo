<?php

namespace App\Services\Contracts;

use App\Http\Requests\Tasks\TaskIndexRequest;
use App\Http\Requests\Tasks\TaskCreateRequest;

interface TaskInterface
{
    public function index(TaskIndexRequest $request);

    public function create(TaskCreateRequest $request);

    public function show(int $id);
}
