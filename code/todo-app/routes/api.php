<?php

use Illuminate\Support\Facades\Route;
use App\Http\Controllers\TaskController;
use App\Http\Controllers\UserController;
use App\Http\Controllers\AuthController;

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

Route::post('users/register', [UserController::class, 'register']);
Route::post('auth/login', [AuthController::class, 'login']);

Route::group(['middleware' => 'auth:api'], function() {
    Route::group(['prefix' => 'auth'], function() {
        Route::get('me', [AuthController::class, 'me']);
        Route::post('logout', [AuthController::class, 'logout']);
    });

    Route::group(['prefix' => 'tasks'], function () {
        Route::get('', [TaskController::class, 'index']);
        Route::post('', [TaskController::class, 'create']);
        Route::get('/{id}', [TaskController::class, 'show']);
    });
});
