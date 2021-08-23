<?php
declare(strict_types=1);

use App\Repository\TaskRepository;
use App\Repository\TaskRepositoryInterface;
use App\Repository\UserRepository;
use App\Repository\UserRepositoryInterface;
use Psr\Container\ContainerInterface;

$container['user_repository'] = static fn (ContainerInterface $container): UserRepositoryInterface => new UserRepository($container->get('db'));

$container['task_repository'] = static fn (ContainerInterface $container): TaskRepositoryInterface => new TaskRepository($container->get('db'));
