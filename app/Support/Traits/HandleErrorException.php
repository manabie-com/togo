<?php

namespace Support\Traits;

use Illuminate\Http\JsonResponse;
use Illuminate\Validation\ValidationException;
use Spatie\QueryBuilder\Exceptions\InvalidSortQuery;

trait HandleErrorException
{
    /**
     * @param ValidationException $exception
     * @return JsonResponse
     */
    public function validationError(ValidationException $exception): JsonResponse
    {
        $errors = $this->convertErrors($exception->errors());

        return $this->makeErrorResponse(JsonResponse::HTTP_BAD_REQUEST, trans('errors.err_400'), $errors);
    }

    /**
     * @param InvalidSortQuery $exception
     * @return JsonResponse
     */
    public function invalidSortQuery(InvalidSortQuery $exception): JsonResponse
    {
        $unknownSorts = $exception->unknownSorts->filter()->map(fn($item) => trim($item))->join(', ');

        return $this->makeErrorResponse(JsonResponse::HTTP_NOT_ACCEPTABLE, trans('errors.err_406'), trans('errors.unknown_sorts', ['attribute' => $unknownSorts]));
    }

    /**
     * @return JsonResponse
     */
    public function unauthorized(): JsonResponse
    {
        return $this->makeErrorResponse(JsonResponse::HTTP_UNAUTHORIZED, trans('errors.err_401'), trans('errors.unauthorized'));
    }

    /**
     * @return JsonResponse
     */
    public function forbidden(): JsonResponse
    {
        return $this->makeErrorResponse(JsonResponse::HTTP_FORBIDDEN, trans('errors.err_403'), trans('message.detail_403'));
    }

    /**
     * @return JsonResponse
     */
    public function notFound(): JsonResponse
    {
        return $this->makeErrorResponse(JsonResponse::HTTP_NOT_FOUND, trans('errors.err_404'), trans('errors.page_not_found'));
    }

    /**
     * @param $message
     * @return JsonResponse
     */
    public function serverError($message): JsonResponse
    {
        return $this->makeErrorResponse(JsonResponse::HTTP_INTERNAL_SERVER_ERROR, trans('errors.err_500'), $message);
    }

    /**
     * @param $status
     * @param $error
     * @param $detail
     * @return JsonResponse
     */
    private function makeErrorResponse($status, $error, $detail): JsonResponse
    {
        $data = [
            'status'  => $status,
            'success' => false,
            'errors'  => [
                'message' => $error,
                'detail'  => $detail,
            ],
        ];

        return response()->json($data, $status);
    }

    /**
     * @param $errors
     * @return array
     */
    private function convertErrors($errors): array
    {
        $result = [];
        foreach ($errors as $field => $error) {
            $result[] = [
                'field'  => $field,
                'detail' => last($error),
            ];
        }

        return $result;
    }
}
