<?php

namespace App\View;

use Illuminate\Http\Request;
use Illuminate\Support\Str;

class Theme
{
    static function state()
    {
        return  empty(session('toastr')) ? null : session('toastr')['status'];
    }
    static function _js($k, $v)
    {
        return is_string($k) ?
            '<script src="' . url($k) . '" type="' . $v . '" ></script>' : '<script src="' . url($v) . '" ></script>';
    }
    static function _css($s)
    {
        return '<link href="' . url($s) . '" rel="stylesheet">';
    }


    public static function js($src)
    {
        foreach ($src as $k => $v) {
            echo static::_js($k, $v);
        }
    }
    public static function css($srcs)
    {
        foreach ($srcs as $s) {
            echo static::_css($s);
        }
    }
    public static function url($para = null)
    {
        $merg  = ($para) ? array_merge(request()->all(), $para) : request()->all();
        $query = http_build_query($merg);
        return empty($query) ? url()->current() : url()->current() . '?' . $query;
    }

    public static function toastr()
    {
        if (session('toastr')) {
            $toastr   = session('toastr');
            $mess     = "'" . $toastr['message'] . "','Message',{positionClass: 'toast-bottom-right', containerId: 'toast-bottom-right'}";
            $mess     = "toastr." . $toastr['status'] . "(" . $mess . ")";
            echo '<script>' . $mess . '</script>';
        }
    }

    public static function menuOpenSub()
    {
        if (request('open'))
            session(['menuOpenSub' => request('open')]);

        if (session('menuOpenSub')) {
            echo "<script>$('#sub-" . session('menuOpenSub') . "').addClass('open');</script>";
        }
    }

    public static function loadOr($hide)
    {
        if (static::state() ==  'success') {
            return "parent.location.reload()";
        }
        return "parent.md_hide('$hide')";
    }

    public static function title($text)
    {
        return Str::title(__($text));
    }
}
