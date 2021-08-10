<?php

namespace App\Exceptions;

use App\Constants\Result;
use Illuminate\Auth\Access\AuthorizationException;
use Illuminate\Auth\AuthenticationException;
use Illuminate\Foundation\Exceptions\Handler as ExceptionHandler;
use Illuminate\Validation\ValidationException;
use Symfony\Component\HttpFoundation\Response;
use Symfony\Component\HttpKernel\Exception\MethodNotAllowedHttpException;
use Symfony\Component\HttpKernel\Exception\NotFoundHttpException;
use Throwable;

class Handler extends ExceptionHandler
{
    /**
     * A list of the exception types that are not reported.
     *
     * @var array
     */
    protected $dontReport = [
        //
    ];

    /**
     * A list of the inputs that are never flashed for validation exceptions.
     *
     * @var array
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

    public function render($request, Throwable $e)
    {
        if ($request->is('api/*') || $request->is('api')) {
            if ($e instanceof MethodNotAllowedHttpException) {
                return response('Method not allow', Response::HTTP_METHOD_NOT_ALLOWED);
            } else if ($e instanceof ValidationException) {
                return badRequest(Result::CODE_FAIL, $e->errors());
            } else if ($e instanceof NotFoundHttpException) {
                return not_found();
            } else if ($e instanceof AuthenticationException) {
                return unauthorized();
            } else if ($e instanceof AuthorizationException) {
                return response('Forbidden', Response::HTTP_FORBIDDEN);
            }

            return fail();
        }

        return parent::render($request, $e);
    }
}
