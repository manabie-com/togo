<?php

namespace App\Http\Controllers;

use App\Http\Controllers\BaseController;
use App\Repositories\TaskRepository;
use App\Services\Task\CheckCreateTaskService;
use Illuminate\Http\Request;
use Illuminate\Http\Response;

class TaskController extends BaseController
{
    /**
     * Create a new controller instance.
     *
     * @return void
     */

    private $taskRepository;
    private $checkCreateTaskService;

    public function __construct(
        TaskRepository $taskRepository,
        CheckCreateTaskService $checkCreateTaskService
    ) {
        $this->taskRepository = $taskRepository;
        $this->checkCreateTaskService = $checkCreateTaskService;
    }

    public function createTask(Request $request)
    {
        $userId = $request->input("user_id");
        $name = $request->input("name");

        $data = [];
        if (!$userId || !$name) {
            $this->message = __('app.invalid_params');
            $this->code = Response::HTTP_BAD_REQUEST;
            goto next;
        }

        list($checkCreateTask, $code, $message) = $this->checkCreateTaskService->checkCreateTask($userId);

        $this->code = $code;
        $this->message = $message;

        if ($checkCreateTask) {
            $dataInsertTask = [
                'user_id' => $userId,
                'name' => $name,
                'date' => time(),
                'time_created' => time()
            ];
            $this->taskRepository->insert($dataInsertTask);
        } else {
            goto next; 
        }
        
        $this->status = 'success';
        next:
        return $this->responseData($data);
    }

}
