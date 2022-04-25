<?php

namespace App\Exceptions;

use Illuminate\Foundation\Exceptions\Handler as ExceptionHandler;
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
    /**
     * Render an exception into an HTTP response.
     *
     * @param  \Illuminate\Http\Request  $request
     * @param  \Exception  $exception
     * @return \Illuminate\Http\Response
     */
    public function render($request, Throwable $exception)
    {

        $userLevelCheck = $exception instanceof \jeremykenedy\LaravelRoles\App\Exceptions\RoleDeniedException ||
            $exception instanceof \jeremykenedy\LaravelRoles\App\Exceptions\RoleDeniedException ||
            $exception instanceof \jeremykenedy\LaravelRoles\App\Exceptions\PermissionDeniedException ||
            $exception instanceof \jeremykenedy\LaravelRoles\App\Exceptions\LevelDeniedException;

        if ($userLevelCheck) {
            if ($request->expectsJson()) {
                return Response::json(array(
                    'error'    =>  403,
                    'message'   =>  'Không xác định được người dùng!',
                ), 403);
            }
            abort(403);
        }

        return parent::render($request, $exception);
    }
}
