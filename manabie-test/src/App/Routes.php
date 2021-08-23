<?php
declare(strict_types=1);

use App\Controller\Task;
use App\Controller\User\LoginController;
use App\Middleware\AuthMiddleware;
use Slim\App;

/** @var App $app */

$app->get('/', 'App\Controller\DefaultController:getHelp');
$app->get('/status', 'App\Controller\DefaultController:getStatus');
$app->post('/login', LoginController::class);

$app->group('/api/v1', function () use ($app): void {
    $app->group('/tasks', function () use ($app): void {
        $app->get('', Task\GetAllController::class);
        $app->post('', Task\CreateController::class);
        $app->get('/{id}', Task\GetOneController::class);
        $app->put('/{id}', Task\UpdateController::class);
        $app->delete('/{id}', Task\DeleteController::class);
    })->add(new AuthMiddleware());
});
