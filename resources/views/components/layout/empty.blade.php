<!DOCTYPE html>
<html class="loading" lang="{{ str_replace('_', '-', app()->getLocale()) }}">
  <head>
    <meta charset="utf-8">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    @stack('meta')
    <title>@yield('page-title', 'Quản lý kho')</title>
    @stack('css')
    <script src="{{asset('app-assets/vendors/js/vendors.min.js')}}" type="text/javascript"></script>
  </head>
  <body {{$attributes ?? ''}} data-open="click" data-menu="vertical-menu" data-col="2-columns" class="">
    <div class="app-content content container-fluid ">
      <div class="content-wrapper">
        <div class="content-body">
          @hasSection('content') @yield('content') @else {{$content ?? $view ?? $body ?? ''}} @endif
        </div>
      </div>
    </div>
    @stack('outer')
    <x-script />
    @stack('scripts')
  </body>
</html>
