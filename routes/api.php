<?php

use Illuminate\Http\Request;
use Illuminate\Support\Facades\Route;



/*
|--------------------------------------------------------------------------
| API Routes
|--------------------------------------------------------------------------
|
| Here is where you can register API routes for your application. These
| routes are loaded by the RouteServiceProvider within a group which
| is assigned the "api" middleware group. Enjoy building your API!
|
*/

Route::group(['prefix' => 'v1'], function () {
    // AUTHENTICATE
    Route::post('login', 'App\Http\Controllers\PassportController@login');
    Route::post('register', 'App\Http\Controllers\PassportController@register');


    Route::middleware('auth:api')->group(function () {
        Route::get('test', 'App\Http\Controllers\PassportController@test');
        // TO-DO
        Route::resource('todo', \App\Http\Controllers\TodoController::class)->only([
            'store'
        ]);

    });

});
