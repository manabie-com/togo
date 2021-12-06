<?php

namespace App\Http\Controllers;

use App\Http\Requests\Tasks\CreateTaskRequest;
use App\Repositories\Tasks\TaskRepository;
use App\Repositories\Users\UserRepository;
use Illuminate\Http\JsonResponse;
use Illuminate\Support\Facades\Auth;

class TaskController extends Controller
{
    /**
     * Create Task
     * 
     * @param CreateTaskRequest $request
     * Responsible for validating the request
     * 
     * @param TaskRepository $taskRepository
     * Responsible for data access related to tasks
     * 
     * @param UserRepository $userRepository
     * Responsible for data access related to users
     */
    public function create(
        CreateTaskRequest $request,
        TaskRepository $taskRepository,
        UserRepository $userRepository
    ) {
        if (! $userRepository->checkIfUserCanCreateTask(Auth::id())) {
            return new JsonResponse([
                'error' => 'Task creation limit exceeded'
            ], 429);
        }

        $data = $request->all();
        $data['user_id'] = Auth::id();

        $task = $taskRepository->create($data);
        return new JsonResponse([
            'success' => 'Task successfully created',
            'task' => $task
        ], 201);
    }
}
