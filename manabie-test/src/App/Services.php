<?php
declare(strict_types=1);

use App\Service\Task\TaskService;
use App\Service\User;
use Psr\Container\ContainerInterface;

$container['login_user_service'] = static fn (ContainerInterface $container): User\Login => new User\Login(
    $container->get('user_repository'),
    $container->get('redis_service')
);

$container['task_service'] = static fn (ContainerInterface $container): TaskService => new TaskService(
    $container->get('task_repository'),
    $container->get('redis_service')
);
