<?php


namespace App\Http\Controllers;

use App\Http\Repositories\TodoCounterRepository;
use App\Http\Repositories\TodoRepository;
use App\Http\Services\TodoCounterService;
use App\Http\Services\TodoService;
use Illuminate\Http\Request;
use Illuminate\Support\Facades\Validator;

class TodoController
{
    private $todoService;
    private $todoCounterService;

    private $todoCounterRepository;
    private $todoRepository;

    public function __construct(
        TodoService $todoService,
        TodoCounterService $todoCounterService,
        TodoCounterRepository $todoCounterRepository,
        TodoRepository $todoRepository
    )
    {
        $this->todoService = $todoService;
        $this->todoCounterService = $todoCounterService;
        $this->todoCounterRepository = $todoCounterRepository;
        $this->todoRepository = $todoRepository;
    }

    public function store(Request $request)
    {
        // VALIDATE INPUT
        $validator = Validator::make($request->all(), [
            'content' => 'required|max:255',
        ]);

        if ($validator->fails()) {
            return response()->json([
                'error' => 'Invalid input'
            ], 400);
        }

        // ADD TASK
        return $this->addTask($request->get('content'));
    }

    public function addTask($content) {
        // Get task counter
        $todoCounter = $this->todoCounterRepository->getByField(['user_id' => auth()->user()->id]);
        // if updated_at is today
        if($this->todoCounterService->checkUpdatedToday($todoCounter))
        {
            // if reached maximum tasks per user
            if($this->todoCounterService->hasReachedMaximum($todoCounter)) {
                return response()->json([
                    'error' => 'Reached maximum'
                ], 400);
            }
            else {
                // set task counter +1
                $this->todoCounterRepository->updateByField(
                    ['user_id' =>  auth()->user()->id],
                    ['count_today' => $todoCounter->count_today + 1]
                );
            }
        }
        // if updated_at is not today
        else {
            // set task counter = 1
            $this->todoCounterRepository->updateByField(
                ['user_id' =>  auth()->user()->id],
                ['count_today' => 1]
            );
        }

        // add task
        $this->todoRepository->create([
            'content' => $content,
            'user_id' => auth()->user()->id,
        ]);

        return response()->json([
            'message' => 'added task'
        ], 200);
    }


}
