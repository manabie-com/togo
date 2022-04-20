<?php

namespace App\Exceptions;

use Illuminate\Auth\Access\AuthorizationException;
use Illuminate\Auth\AuthenticationException;
use Illuminate\Database\Eloquent\ModelNotFoundException;
use Illuminate\Foundation\Exceptions\Handler as ExceptionHandler;
use Illuminate\Validation\UnauthorizedException;
use Illuminate\Validation\ValidationException;
use Spatie\QueryBuilder\Exceptions\InvalidSortQuery;
use Support\Traits\HandleErrorException;
use Symfony\Component\HttpKernel\Exception\MethodNotAllowedHttpException;
use Symfony\Component\HttpKernel\Exception\NotFoundHttpException;
use Throwable;

class Handler extends ExceptionHandler
{
    use HandleErrorException;

    /**
     * A list of the exception types that are not reported.
     *
     * @var array<int, class-string<Throwable>>
     */
    protected $dontReport = [
        //
    ];

    /**
     * A list of the inputs that are never flashed for validation exceptions.
     *
     * @var array<int, string>
     */
    protected $dontFlash = [
        'current_password',
        'password',
        'password_confirmation',
    ];

    /**
     * Register the exception handling callbacks for the application.
     *
     * @return void
     */
    public function register()
    {
        $this->reportable(function (Throwable $e) {
            //
        });
    }

    /**
     * @param \Illuminate\Http\Request $request
     * @param \Throwable $e
     * @return \Illuminate\Http\JsonResponse|\Symfony\Component\HttpFoundation\Response
     */
    public function render($request, Throwable $e)
    {
        switch (true) {
            case $e instanceof ValidationException:
                return $this->validationError($e);
            case $e instanceof InvalidSortQuery:
                return $this->invalidSortQuery($e);
            case $e instanceof UnauthorizedException:
            case $e instanceof AuthenticationException:
            case $e instanceof AuthorizationException:
                return $this->unauthorized();
            case $e instanceof NotFoundHttpException:
            case $e instanceof ModelNotFoundException:
            case $e instanceof MethodNotAllowedHttpException:
                return $this->notFound();
            default:
                return $this->serverError($e->getMessage());
        }
    }
}
