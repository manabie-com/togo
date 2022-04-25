<!DOCTYPE html>
<html lang="{{ str_replace('_', '-', app()->getLocale()) }}">
    <head>
        <meta charset="utf-8">
        <meta http-equiv="X-UA-Compatible" content="IE=edge">
        <meta name="viewport" content="width=device-width, initial-scale=1">
        <meta name="csrf-token" content="{{ csrf_token() }}">
        <title>
            @hasSection('template_title')
                @yield('template_title') | 
            @endif {{ config('app.name', Lang::get('titles.app')) }}
        </title>
        <meta name="description" content="">
        <meta name="author" content="">
        <link rel="shortcut icon" href="/favicon.ico">
        <link rel="apple-touch-icon" href="{{asset('app-assets/images/ico/apple-icon-120.png')}}">
		<link rel="shortcut icon" type="image/x-icon" href="{{asset('images/warehouse-120.png')}}">
		<!-- BEGIN VENDOR CSS-->
		<link rel="stylesheet" type="text/css" href="{{asset('app-assets/css/bootstrap.css')}}">
		<link rel="stylesheet" type="text/css" href="{{asset('app-assets/fonts/feather/style.min.css')}}">
		<link rel="stylesheet" type="text/css" href="{{asset('app-assets/fonts/font-awesome/css/font-awesome.min.css')}}">
		<link rel="stylesheet" type="text/css" href="{{asset('app-assets/fonts/flag-icon-css/css/flag-icon.min.css')}}">
		<link rel="stylesheet" type="text/css" href="{{asset('app-assets/vendors/css/extensions/pace.css')}}">
		<link rel="stylesheet" type="text/css" href="{{asset('app-assets/vendors/css/charts/jquery-jvectormap-2.0.3.css')}}">
		<link rel="stylesheet" type="text/css" href="{{asset('app-assets/vendors/css/charts/morris.css')}}">
		<link rel="stylesheet" type="text/css" href="{{asset('app-assets/vendors/css/extensions/unslider.css')}}">
		<link rel="stylesheet" type="text/css" href="{{asset('app-assets/vendors/css/weather-icons/climacons.min.css')}}">
		<link rel="stylesheet" type="text/css" href="{{asset('app-assets/vendors/css/tables/datatable/dataTables.bootstrap4.min.css')}}">
		<link rel="stylesheet" type="text/css" href="{{asset('app-assets/vendors/css/extensions/sweetalert.css')}}">
		<link rel="stylesheet" type="text/css" href="{{asset('app-assets/vendors/css/pickers/daterange/daterangepicker.css')}}">
		<link rel="stylesheet" type="text/css" href="{{asset('app-assets/vendors/css/pickers/datetime/bootstrap-datetimepicker.css')}}">
		<link rel="stylesheet" type="text/css" href="{{asset('app-assets/vendors/css/pickers/pickadate/pickadate.css')}}">
		<link rel="stylesheet" type="text/css" href="{{asset('app-assets/vendors/css/forms/selects/select2.min.css')}}">
		<!-- END VENDOR CSS-->
		<!-- BEGIN STACK CSS-->
		<link rel="stylesheet" type="text/css" href="/app-assets/css/bootstrap-extended.css">
		<link rel="stylesheet" type="text/css" href="/app-assets/css/app.css">
		<link rel="stylesheet" type="text/css" href="/app-assets/css/colors.css">
		<!-- END STACK CSS-->
		<!-- BEGIN Page Level CSS-->
		<link rel="stylesheet" type="text/css" href="/app-assets/css/core/menu/menu-types/vertical-menu.css">
		<link rel="stylesheet" type="text/css" href="/app-assets/css/core/menu/menu-types/vertical-overlay-menu.css">
		<!-- END Page Level CSS-->
		<!-- BEGIN CUSTOM CSS-->
		<link rel="stylesheet" type="text/css" href="/assets/css/style.css">
        {{-- Fonts --}}
        @yield('fonts')
        {{-- Style --}}
        @yield('css')
        @yield('css-theme')
    </head>
    @yield('body')
</html>