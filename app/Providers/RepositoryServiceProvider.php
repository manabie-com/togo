<?php

namespace App\Providers;

use Illuminate\Support\ServiceProvider;
use Repository\EloquentRepository;
use Repository\IRepositories\IEloquentRepository;
use Repository\IRepositories\ITaskRepository;
use Repository\IRepositories\IUserRepository;
use Repository\TaskRepository;
use Repository\UserRepository;

class RepositoryServiceProvider extends ServiceProvider
{
    /**
     * Register services.
     *
     * @return void
     */
    public function register()
    {
        $this->app->bind(IEloquentRepository::class, EloquentRepository::class);
        $this->app->bind(IUserRepository::class, UserRepository::class);
        $this->app->bind(ITaskRepository::class, TaskRepository::class);
    }

    /**
     * Bootstrap services.
     *
     * @return void
     */
    public function boot()
    {
        //
    }
}
