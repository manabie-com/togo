<?php

namespace App\Providers;

use App\Repositories\Tasks\TaskEloquentRepository;
use App\Repositories\Tasks\TaskRepository;
use App\Repositories\Users\UserEloquentRepository;
use App\Repositories\Users\UserRepository;
use Illuminate\Contracts\Support\DeferrableProvider;
use Illuminate\Support\ServiceProvider;

class RepositoryServiceProvider extends ServiceProvider implements DeferrableProvider
{
    public $singletons = [
        UserRepository::class => UserEloquentRepository::class,
        TaskRepository::class => TaskEloquentRepository::class
    ];

    /**
     * Get the services provided by the provider.
     *
     * @return array
     */
    public function provides()
    {
        return [
            UserRepository::class,
            TaskRepository::class
        ];
    }
}
