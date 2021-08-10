<?php

namespace App\Services;

use App\Constants\Result;
use App\Models\User;
use Illuminate\Support\Facades\Auth;
use App\Services\Contracts\TaskInterface;
use App\Http\Requests\Tasks\TaskIndexRequest;
use App\Http\Requests\Tasks\TaskCreateRequest;
use App\Repositories\Contracts\TaskRepositoryInterface;

class TaskService implements TaskInterface
{
    protected $taskRepository;

    public function __construct(TaskRepositoryInterface $taskRepository)
    {
        $this->taskRepository = $taskRepository;
    }

    public function index(TaskIndexRequest $request)
    {
        return $this->taskRepository->orderBy('created_at', 'DESC')
                                    ->paginate($request->input('limit'));
    }

    public function create(TaskCreateRequest $request): array
    {
        /** @var User $user */
        $user = Auth::user();

        if ($this->isReachLimitPerDay($user->max_todo)) {
            return ['code' => Result::CODE_TASK_REACH_LIMIT_DAY];
        }

        $param = [
            'user_id' => Auth::id(),
            'content' => $request->input('content')
        ];

        return [
            'code' => Result::CODE_OK,
            'data' => $this->taskRepository->create($param)
        ];
    }

    public function show(int $id)
    {
        return $this->taskRepository->findWhereFirst(['id' => $id]);
    }

    protected function isReachLimitPerDay(int $limit): bool
    {
        $numberTaskToday = $this->taskRepository->countTaskToday(Auth::id());

        # Today's task number + 1 must be less than limit
        return $numberTaskToday + 1 > $limit;
    }
}
