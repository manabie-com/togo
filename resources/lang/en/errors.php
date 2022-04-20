<?php

return [

    /*
    |--------------------------------------------------------------------------
    | Error Message Language Lines
    |--------------------------------------------------------------------------
    |
    | The following language lines are used by the Laravel Responder package.
    | When it generates error responses, it will search the messages array
    | below for any key matching the given error code for the response.
    |
    */

    'unauthenticated'    => 'You are not authenticated for this request.',
    'unauthorized'       => 'You are not authorized for this request.',
    'page_not_found'     => 'The requested page does not exist.',
    'relation_not_found' => 'The requested relation does not exist.',
    'validation_failed'  => 'The given data failed to pass validation.',
    'err_400'            => 'Validation error.',
    'err_401'            => 'Authentication failed.',
    'err_403'            => 'Access is denied.',
    'err_404'            => 'No data was found.',
    'err_406'            => 'Not Acceptable.',
    'unknown_sorts'      => 'Requested sort(s) :attribute is not allowed.',
];
