<?php

namespace Api\Controllers;

use Api\Requests\Task\CreateTaskRequest;
use Api\Requests\Task\UpdateTaskRequest;
use Domain\Tasks\Models\Task;
use Domain\Tasks\Transformers\TaskTransformer;
use Domain\Users\Filters\UserFilter;
use Domain\Users\Sorts\UserSort;
use Illuminate\Http\JsonResponse;
use Repository\IRepositories\ITaskRepository;

class TaskController extends ApiController
{
    private ITaskRepository $taskRepository;

    /**
     * UserController constructor.
     *
     * @param ITaskRepository $taskRepository
     */
    public function __construct(ITaskRepository $taskRepository)
    {
        $this->taskRepository = $taskRepository;
    }

    /**
     * Display a listing of the resource.
     *
     * @param UserFilter $filter
     * @param UserSort $sort
     * @return JsonResponse
     */
    public function index(UserFilter $filter, UserSort $sort): JsonResponse
    {
        $task = $this->taskRepository->getList($filter, $sort);

        return $this->httpOK($task, TaskTransformer::class);
    }

    /**
     * Store a newly created resource in storage.
     *
     * @param CreateTaskRequest $createTaskRequest
     * @return JsonResponse
     */
    public function store(CreateTaskRequest $createTaskRequest): JsonResponse
    {
        $attributes = $createTaskRequest->validated();
        $attributes['user_id'] = auth()->user()->id;
        $task = $this->taskRepository->create($attributes);

        return $this->httpOK($task, TaskTransformer::class);
    }

    /**
     * Display the specified resource.
     *
     * @param Task $task
     * @return JsonResponse
     */
    public function show(Task $task): JsonResponse
    {
        return $this->httpOK($task, TaskTransformer::class);
    }

    /**
     * Update the specified resource in storage.
     *
     * @param UpdateTaskRequest $updateTaskRequest
     * @param Task $task
     * @return JsonResponse
     */
    public function update(UpdateTaskRequest $updateTaskRequest, Task $task): JsonResponse
    {
        $attributes = $updateTaskRequest->validated();
        $task = $this->taskRepository->update($task, $attributes);

        return $this->httpOK($task, TaskTransformer::class);
    }

    /**
     * Remove the specified resource from storage.
     *
     * @param Task $task
     * @return JsonResponse
     */
    public function destroy(Task $task): JsonResponse
    {
        $this->taskRepository->delete($task);

        return $this->httpOK($task, TaskTransformer::class);
    }
}
