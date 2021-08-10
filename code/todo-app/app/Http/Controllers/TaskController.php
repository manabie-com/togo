<?php

namespace App\Http\Controllers;

use App\Constants\Result;
use Illuminate\Http\Request;
use Illuminate\Http\JsonResponse;
use App\Http\Resources\TaskResource;
use App\Services\Contracts\TaskInterface;
use App\Http\Requests\Tasks\TaskIndexRequest;
use App\Http\Requests\Tasks\TaskCreateRequest;

class TaskController extends Controller
{
    protected $taskService;

    public function __construct(TaskInterface $taskService)
    {
        $this->taskService = $taskService;
    }

    public function create(TaskCreateRequest $request): JsonResponse
    {
        $result = $this->taskService->create($request);

        if (isset($result['code']) && Result::CODE_OK == $result['code'] && !empty($result['data'])) {
            return success(['id' => $result['data']->id]);
        }

        if (isset($result['code'])) {
            return badRequest($result['code'], Result::getMessage($result['code']));
        }

        return fail();
    }

    public function show(string $id, Request $request)
    {
        $task = $this->taskService->show($id);

        if ($task) {
            $taskResource = (new TaskResource($task))->toArray($request);

            return success($taskResource);
        }

        return not_found();
    }

    public function index(TaskIndexRequest $request): JsonResponse
    {
        if (!$request->has('page')) {
            $request->merge(['page' => 1]);
        }

        if (!$request->has('limit')) {
            $request->merge(['limit' => 10]);
        }

        $tasks = $this->taskService->index($request);

        if (COUNT($tasks) > 0) {
            $taskCollection = TaskResource::collection($tasks)->toArray($request);

            return success($taskCollection);
        }

        return success();
    }
}
