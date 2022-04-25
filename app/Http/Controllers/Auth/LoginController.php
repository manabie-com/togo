<?php

namespace App\Http\Controllers\Auth;

use App\Http\Controllers\Controller;
use Illuminate\Foundation\Auth\AuthenticatesUsers;
use Illuminate\Support\Facades\Auth;
use Illuminate\Support\Facades\Log;
use Illuminate\Support\Facades\Session;
use Illuminate\Http\Request;
use Illuminate\Validation\ValidationException;

class LoginController extends Controller
{
    /*
    |--------------------------------------------------------------------------
    | Login Controller
    |--------------------------------------------------------------------------
    |
    | This controller handles authenticating users for the application and
    | redirecting them to your home screen. The controller uses a trait
    | to conveniently provide its functionality to your applications.
    |
    */

    use AuthenticatesUsers;

    /**
     * Where to redirect users after login.
     *
     * @var string
     */
    protected $redirectAfterLogout = '/';
    /**
     * Where to redirect users after login.
     *
     * @var string
     */
    protected $username = 'email';

    /**
     * Create a new controller instance.
     *
     * @return void
     */
    public function __construct()
    {
        $this->middleware('guest', ['except' => 'logout']);
    }
    /**
     * Get the login username to be used by the controller.
     *
     * @return string
     */
    public function username()
    {
        return $this->username;
    }
    /**
     * Validate the user login request.
     *
     * @param  \Illuminate\Http\Request  $request
     * @return void
     *
     * @throws \Illuminate\Validation\ValidationException
     */
    protected function validateLogin(Request $request)
    {
        $validateUsername = 'required|string';
        /*To check the username isn't email validate*/
        if (!filter_var($request->input('email'), FILTER_VALIDATE_EMAIL)) {
            $this->username = 'name';
            $validateUsername = 'required';
        }
        $request->validate([
            $this->username() => $validateUsername,
            'password' => 'required|string',
        ]);
    }
    /**
     * Logout, Clear Session, and Return.
     *
     * @return void
     */
    public function logout()
    {
        // $user = Auth::user();
        // Log::info('User Logged Out. ', [$user]);
        Auth::logout();
        Session::flush();

        return redirect(property_exists($this, 'redirectAfterLogout') ? $this->redirectAfterLogout : '/');
    }
    /**
     * Get the failed login response instance.
     *
     * @param  \Illuminate\Http\Request  $request
     * @return \Symfony\Component\HttpFoundation\Response
     *
     * @throws \Illuminate\Validation\ValidationException
     */
    protected function sendFailedLoginResponse(Request $request)
    {
        throw ValidationException::withMessages([
            'notice'    => [trans('auth.failed_verified')],
            // 'email'     => ['Tên đăng nhập là mã nhận viên hoặc email'],
            // 'password'  => ['Mật khẩu ít nhất 6 ký tự'],
        ]);
    }
}
