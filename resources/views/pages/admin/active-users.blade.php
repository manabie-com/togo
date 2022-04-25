@extends('components.layout.greenland')

@section('template_title')
    {{ trans('titles.activeUsers') }}
@endsection

@section('content')
	
    <users-count :registered={{ $users }} ></users-count>

@endsection
