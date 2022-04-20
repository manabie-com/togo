<?php

namespace App\Http\Middleware;

use Illuminate\Auth\Middleware\Authenticate as Middleware;

class Authenticate extends Middleware
{
    /**
     * Get the path the user should be redirected to when they are not authenticated.
     *
     * @param \Illuminate\Http\Request $request
     * @return string|null
     */
    protected function redirectTo($request)
    {
        $middlewares = $request->route()->gatherMiddleware();
        if ($middlewares && in_array('api', $middlewares)) {
            if ($request->headers->get('Cookie')) {
                $request->headers->remove('Cookie');
            }

            if ($request->get('Accept') !== 'application/json') {
                $request->headers->set('Accept', 'application/json');
            }
        }
        if (! $request->expectsJson()) {
            return route('login');
        }
    }
}
