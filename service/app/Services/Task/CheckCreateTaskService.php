<?php

namespace App\Services\Task;

use App\Repositories\SettingRepository;
use App\Repositories\TaskRepository;
use Illuminate\Http\Response;

class CheckCreateTaskService
{
    private $taskRepository;
    private $settingRepository;
    private $code;
    private $message;

    public function __construct(
        TaskRepository $taskRepository,
        SettingRepository $settingRepository
    )
    {
        $this->taskRepository = $taskRepository;
        $this->settingRepository = $settingRepository;
    }

    public function checkCreateTask($userId)
    {

        $userSetting = $this->settingRepository->getLimit($userId);

        if (isset($userSetting->limit)) {
            $limit = $userSetting->limit;
        } else {
            $this->message = __('app.user_is_not_registered_for_task');
            $this->code = Response::HTTP_UNPROCESSABLE_ENTITY;
            return [false, $this->code, $this->message];
        }

        $date = date('Y-m-d');
        $startDate = strtotime($date);
        $endDate = $startDate + 24*60*60;
        $countTasks = $this->taskRepository->countTasks($userId, $startDate, $endDate);
        
        if ($countTasks >= $limit) {
            $this->message = __('app.max_limit_task');
            $this->code = Response::HTTP_TOO_MANY_REQUESTS;
            return  [false, $this->code, $this->message];
        }

        $this->code = Response::HTTP_OK;
        $this->message = __('app.create_task_success');

        return [true, $this->code, $this->message];

    }

}
