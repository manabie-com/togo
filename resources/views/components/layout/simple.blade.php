<!DOCTYPE html>
<html class="loading"  lang="{{ str_replace('_', '-', app()->getLocale()) }}">
  <head>
    <meta charset="utf-8">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    @stack('meta')
    <title>@yield('page-title', 'TekTop')</title>
    @stack('css')
  </head>

  <body {{ $attributes??null }} >
    @yield('header')
    @yield('body')
    @yield('footer')
    @stack('outer')
    <x-script />
    @stack('script')

  </body>
</html>
