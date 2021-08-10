<?php

namespace App\Providers;

use App\Services\TaskService;
use App\Services\UserService;
use Illuminate\Support\ServiceProvider;
use App\Services\Contracts\UserInterface;
use App\Services\Contracts\TaskInterface;

class AppServiceProvider extends ServiceProvider
{
    /**
     * Register any application services.
     *
     * @return void
     */
    public function register()
    {
        $this->app->bind(UserInterface::class, UserService::class);
        $this->app->bind(TaskInterface::class, TaskService::class);
    }

    /**
     * Bootstrap any application services.
     *
     * @return void
     */
    public function boot()
    {
        //
    }
}
