@extends('layouts.default')

@section('template_title')
  {{ $user->name }}'s Profile
@endsection

@section('navbar-more')
  Thông tin chi tiết
@endsection

@php
  $currentUser = Auth::user()
@endphp
<style type="text/css">
  .user-info dt {margin-bottom: 5px;}
  .user-info dd{padding: 5px 10px; border: 0.5px solid #B5CBE8; border-radius: 6px;}
  /*.user-info dd a {padding: 5px;}*/
</style>
@section('content')
<div class="app-view">  
  <div class="container-fluid">
    <div class="row">
      <div class="col-xl-4 col-12 pl-xl-0 form-group d-none">
        <div class="card">
          <div class="card-body text-center">
            <div class="dz-preview"></div>
            {!! Form::open(array('method' => 'POST', 'name' => 'avatarDropzone','id' => 'avatarDropzone', 'class' => 'form single-dropzone dropzone single', 'files' => true)) !!}
              <img id="user_selected_avatar" class="user-avatar" src="@if ($user->profile->avatar ?? '' != NULL) {{ $user->profile->avatar }} @endif" alt="{{ $user->name }}" width="80%" style="min-height: 270px">
            {!! Form::close() !!}
          </div>
        </div>
      </div>
      <div class="col-xl-6 col-12 form-group">
        <div class="card">
          <div class="card-body p-2">
            <dl class="user-info">
              @if ($user->last_name && ($currentUser->id == $user->id || $currentUser->hasRole('admin')))
                <dt>{{ trans('profile.showProfileLastName') }}</dt>
                <dd>{{ $user->last_name }}</dd>
              @endif
              <dt>{{ trans('profile.showProfileFirstName') }}</dt>
              <dd>{{ $user->first_name }}</dd>

              <dt>{{ trans('profile.showProfileUsername') }}</dt>
              <dd>{{ $user->name }}</dd>

              @if ($user->email && ($currentUser->id == $user->id || $currentUser->hasRole('admin')))
                <dt>{{ trans('profile.showProfileEmail') }}</dt>
                <dd>{{ $user->email }}</dd>
              @endif

              {{-- @if ($user) --}}
                @if ($user->code)
                  <dt>{{ trans('profile.showProfileUsernameId') }}</dt>
                  <dd>
                    {{ $user->code }} <br />
                  </dd>
                @endif
              {{-- @endif --}}
            </dl>
            <div class="mx-auto" style="max-width: 300px">
              @if ($user->profile)
                @if ($currentUser->id == $user->id)
                  <a class="btn btn-small btn-warning btn-block" href="{{url('/profile/'.$currentUser->name.'/edit')}}">
                    {{-- <i class="fa fa-fw fa-cog"></i> --}}
                    {{__('titles.editProfile')}}
                  </a>
                  {{-- {!! HTML::icon_link(URL::to('/profile/'.$currentUser->name.'/edit'), 'fa fa-fw fa-cog', trans('titles.editProfile'), array('class' => 'btn btn-small btn-info btn-block')) !!} --}}
                @endif
              @else
                <p>{{ trans('profile.noProfileYet') }}</p>
                <a class="btn btn-small btn-info btn-block" href="{{url('/profile/'.$currentUser->name.'/edit')}}">
                  <i class="fa fa-fw fa-plus "></i>
                  {{__('titles.createProfile')}}
                </a>
                {{-- {!! HTML::icon_link(URL::to('/profile/'.$currentUser->name.'/edit'), 'fa fa-fw fa-plus ', trans('titles.createProfile'), array('class' => 'btn btn-small btn-info btn-block')) !!} --}}
              @endif
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</div>
@endsection

@section('footer_scripts')

  @if(config('settings.googleMapsAPIStatus'))
    @include('scripts.google-maps-geocode-and-map')
  @endif

  {{-- @include('scripts.user-avatar-dz') --}}
@endsection
