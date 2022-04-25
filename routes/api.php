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
/**
 * @author toannguyen.dev
 * @todo API extends
 * @var 2021-08-11
 */
// Route::group(['middleware' => []], function () {
//     /*internal*/
//     Route::get('api/get/province/{option?}', 'APIController@get_province')->name('api.get.province');
//     Route::get('api/get/district', 'APIController@get_district')->name('api.get.district');
//     Route::get('api/get/ward', 'APIController@get_ward')->name('api.get.ward');
// });
//     /*ROLE::elf*/
Route::middleware(['role:elf','throttle:10000,1'])->group(function ($route) {
    Route::prefix('elf')->namespace('Group\Elf\\')->group(function () {
        Route::get('animals/generic', 'AnimalController@generic')->name('animals.generic');
        Route::get('animals/book', 'AnimalController@book')->name('animals.book');
        /*Set resource*/
        Route::resources([
            'animals'       => AnimalController::class,
            'practices'    => PracticeController::class,
        ], ['as', 'prefix']); 
    });
});

Route::group(['middleware' => []], function () {
    Route::get('api/{object}/{action?}/{option?}', 'APIController@api');
});