<?php
declare(strict_types=1);

namespace App\Controller;

use Slim\Container;
use Slim\Http\Response;

/**
 * Class BaseController
 * @package App\Controller
 */
abstract class BaseController
{
    /**
     * @var Container
     */
    protected Container $container;

    /**
     * @param Container $container
     */
    public function __construct(Container $container)
    {
        $this->container = $container;
    }

    /**
     * @param Response $response
     * @param string $status
     * @param  $message
     * @param int $code
     * @return Response
     */
    protected function jsonResponse( Response $response, string  $status,  $message, int $code): Response
    {
        $result = [
            'code' => $code,
            'status' => $status,
            'message' => $message,
        ];

        return $response->withJson($result, $code, JSON_PRETTY_PRINT);
    }

    /**
     * @return bool
     */
    protected static function isRedisEnabled(): bool
    {
        return filter_var($_SERVER['REDIS_ENABLED'], FILTER_VALIDATE_BOOLEAN);
    }
}
