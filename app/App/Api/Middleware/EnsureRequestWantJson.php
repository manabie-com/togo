<?php

namespace Api\Middleware;

use Closure;
use Illuminate\Http\Request;

class EnsureRequestWantJson
{
    /**
     * Handle an incoming request.
     *
     * @param Request $request
     * @param \Closure $next
     * @return mixed
     */
    public function handle(Request $request, Closure $next)
    {
        if ($request->headers->get('Cookie')) {
            $request->headers->remove('Cookie');
        }

        if ($request->get('Accept') === 'application/json') {
            return $next($request);
        }
        $request->headers->set('Accept', 'application/json');

        return $next($request);
    }
}
