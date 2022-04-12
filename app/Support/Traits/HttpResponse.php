<?php

namespace Support\Traits;

use Flugg\Responder\Http\MakesResponses;
use Illuminate\Http\JsonResponse;

trait HttpResponse
{
    use MakesResponses;

    /**
     * @param null $data
     * @param null $transformer
     * @return \Illuminate\Http\JsonResponse
     */
    public function httpOK($data = null, $transformer = null): JsonResponse
    {
        return $this->success($data, $transformer)->respond(JsonResponse::HTTP_OK);
    }

    /**
     * @param null $data
     * @param null $transformer
     * @return \Illuminate\Http\JsonResponse
     */
    public function httpCreated($data = null, $transformer = null): JsonResponse
    {
        return $this->success($data, $transformer)->respond(JsonResponse::HTTP_CREATED);
    }

    /**
     * @return \Illuminate\Http\JsonResponse
     */
    public function httpNoContent(): JsonResponse
    {
        return $this->success()->respond(JsonResponse::HTTP_NO_CONTENT);
    }
}
