<?php
declare(strict_types=1);

namespace App\Handler;

use Psr\Http\Message\ResponseInterface as Response;
use Psr\Http\Message\ServerRequestInterface as Request;

/**
 * Class ApiError
 * @package App\Handler
 */
class ApiError extends \Slim\Handlers\Error
{
    /**
     * @param Request $request
     * @param Response $response
     * @param \Exception $exception
     * @return Response
     */
    public function __invoke(Request $request, Response $response, \Exception $exception): Response
    {
        $statusCode = $this->getStatusCode($exception);
        $className = new \ReflectionClass(get_class($exception));

        $data = [
            'message' => $exception->getMessage(),
            'class' => $className->getName(),
            'status' => 'error',
            'code' => $statusCode,
        ];
        $body = json_encode($data, JSON_UNESCAPED_SLASHES | JSON_PRETTY_PRINT);
        $response->getBody()->write((string)$body);

        return $response
            ->withStatus($statusCode)
            ->withHeader('Content-type', 'application/problem+json');
    }

    /**
     * @param \Exception $exception
     * @return int
     */
    private function getStatusCode(\Exception $exception): int
    {
        $statusCode = 500;

        if (is_int($exception->getCode()) && $exception->getCode() >= 400 && $exception->getCode() <= 500) {
            $statusCode = $exception->getCode();
        }

        return $statusCode;
    }
}
