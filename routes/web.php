<?php

use Illuminate\Support\Facades\Route;
use Illuminate\Support\Facades\Auth;

/*
|--------------------------------------------------------------------------
| Web Routes
|--------------------------------------------------------------------------
|
| Here is where you can register web routes for your application. These
| routes are loaded by the RouteServiceProvider within a group which
| contains the "web" middleware group. Now create something great!
|
| Middleware options can be located in `app/Http/Kernel.php`
|
*/
    
/*Authentication Routes*/
Auth::routes();
 /**
 * @author toannguyen
 * @todo as dev mode
 * @var 
 */
Route::group(['middleware' => ['auth']], function () {
    /*logger*/
    Route::get('logs', '\Rap2hpoutre\LaravelLogViewer\LogViewerController@index');
    /*route*/
    Route::get('routes', 'AdminDetailsController@listRoutes');
});
 /**
 * @author 
 * @todo  
 * @var
 */
/*ROLE::Administrator System*/
Route::group(['middleware' => ['auth', 'role:admin']], function () {
    Route::get('users/export', 'UserController@index')->name('users.export');
    Route::resource('users', 'UserController');
});
// xdebug_print_function_stack('hi men');
// dd(xdebug_memory_usage());

    /*ROLE::elf*/
    Route::middleware(['auth', 'role:elf','throttle:1000,'.(60*24)])->group(function ($route) {
        Route::prefix('elf')->namespace('Group\Elf\\')->group(function () {
            Route::get('animals/generic', 'AnimalController@generic')->name('animals.generic');
            Route::get('animals/book', 'AnimalController@book')->name('animals.book');
            // Route::put('home/', '\App\Http\Controllers\UserController@update')->name('elf.update');
            /*Set resource*/
            Route::resources([
                'animals'       => AnimalController::class,
                'practices'    => PracticeController::class,
            ], ['as', 'prefix']); 
        });
    });

Route::middleware(['auth'])->group(function ($route) {
    /*ROLE::elf*/
    Route::prefix('anm')->middleware(['role:animal'])->namespace('Group\Animal\\')->group(function () {
        Route::get('animals/practice', 'AnimalController@practice')->name('animal.practice');
        Route::get('animals/book', 'AnimalController@book')->name('animal.book');
        // Route::put('home/', '\App\Http\Controllers\UserController@update')->name('dog.update');
        /*Set resource*/
        Route::resources([
            'animal'         => AnimalController::class,
        ], ['as', 'prefix']); 
    });

});
 /**
 * @author toannguyen
 * @todo  as the users was only authenticated. 
 * @var
 */ 
Route::group(['middleware' => ['auth']], function () {
    /* index page*/
    Route::get('/', 'UserController@dashboard');
    Route::get('/home', 'UserController@dashboard');
    Route::get('/logout', ['uses' => 'Auth\LoginController@logout'])->name('logout');
    Route::resource('users', 'UserController')->only(['show', 'edit', 'update']);
});
/**
 *  @todo As guest
 */
Route::group(['middleware' => []], function () {
    /*index page*/
});