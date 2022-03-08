<?php

namespace App\Http\Controllers;

use Illuminate\Support\Facades\App;
use Illuminate\Http\Response as HttpResponse;

class BaseController extends Controller
{
    protected $lang_id;
    protected $is_vn = false;
    protected $ip;
    protected $lang;
    protected $limit = 20;
    protected $status = 'fail';

    protected $message = '';

    protected $code = 200;

    public function __construct(
    ){
        App::setLocale("vi");
    }

    protected function responseData($data = [], $more = '', $code = 200)
    {
        if(!$data) {
            $data = (object) $data;
        }
        $res = [
            'status' => $this->status,
            'message' => $this->message,
            'code' => $this->code,
            'data' => $data
        ];
        if ($more) {
            $res = array_merge($res, $more);
        }
        return response()->json($res, $this->code);
    }

    protected function infoUser()
    {
        return $this->request()->attributes->get('userInfo');
    }

    protected function infoUserId()
    {
        $user = $this->infoUser();
        return $user['id'];
    }

    protected function request($key = null, $default = null)
    {
        if (is_null($key)) {
            return app('request');
        }

        if (is_array($key)) {
            return app('request')->only($key);
        }

        return data_get(app('request')->all(), $key, $default);
    }

}
