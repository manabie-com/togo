<?php

namespace Database\Seeders;

use App\Repositories\SettingRepository;
use App\Repositories\TaskRepository;
use Illuminate\Database\Seeder;

class DatabaseSeeder extends Seeder
{
    /**
     * Run the database seeds.
     *
     * @return void
     */
    private $settingRepository;
    private $taskRepository;

    public function __construct(
        SettingRepository  $settingRepository,
        TaskRepository $taskRepository
    )
    {
        $this->settingRepository = $settingRepository;
        $this->taskRepository = $taskRepository;
    }
    
    public function run()
    {
        $dataCreateSetting = [
            [
                'user_id' => 1,
                'limit' => 1,
                'time_created' => time(),
                'time_updated' => time()
            ],
            [
                'user_id' => 3,
                'limit' => 100,
                'time_created' => time(),
                'time_updated' => time()
            ]
        ];

        $this->settingRepository->insert($dataCreateSetting);

        $dataCreateTask = [
            [
                'user_id' => 1,
                'name' => 'task1',
                'time_created' => time(),
                'time_updated' => time()
            ]
        ];

        $this->taskRepository->insert($dataCreateTask);
    }
}
