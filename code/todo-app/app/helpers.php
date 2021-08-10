<?php

use App\Constants\Result;
use Illuminate\Http\JsonResponse;
use Symfony\Component\HttpFoundation\Response;

function success(array $data = [], int $code = Result::CODE_OK, string $message = 'success', int $httpCode = Response::HTTP_OK, array $headers = []): JsonResponse
{
    $body = respond($code, $message, $data);

    return response()->json($body, $httpCode, $headers);
}

function fail(int $code = Result::CODE_FAIL, string $message = 'fail', int $httpCode = Response::HTTP_INTERNAL_SERVER_ERROR, array $headers = []): JsonResponse
{
    $body = respond($code, $message);

    return response()->json($body, $httpCode, $headers);
}

function badRequest(int $code = Result::CODE_FAIL, $message = 'fail', int $httpCode = Response::HTTP_BAD_REQUEST, array $headers = []): JsonResponse
{
    $body = respond($code, $message);

    return response()->json($body, $httpCode, $headers);
}

function unauthorized()
{
    return response('Unauthorized', Response::HTTP_UNAUTHORIZED);
}

function not_found()
{
    return response('Not found', Response::HTTP_NOT_FOUND);
}

function respond($code, $message, $data = null): array
{
    $response = [
        'code' => $code,
        'message' => $message
    ];

    if (!is_null($data)) {
        $response['data'] = $data;
    }

    return $response;
}
