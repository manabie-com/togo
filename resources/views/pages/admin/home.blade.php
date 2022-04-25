@extends('components.layout.greenland')

@section('template_title')
    Welcome {{ Auth::user()->name }}
@endsection

@section('head')
@endsection

@section('content')

    @include('panels.dashboard')

@endsection
