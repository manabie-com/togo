
@extends('layouts.default')
@section('template_title')
  {{ trans('profile.templateTitle') }}
@endsection
@section('navbar-more')
  <div class="text-uppercase">{{ trans('profile.showProfileTitle',['username' => $user->name]) }}</div>
@endsection
<style type="text/css">
  
  .app-view .container-fluid .has-feedback >label {padding: 12px;}
  .app-view .container-fluid .card.card-avatar{max-width: 400px; border-radius: 6px;}
  .app-view .container-fluid .card.images-avatar{height: 107px;width: 107px;}
  .app-view .container-fluid .card .info label{font-size: 14px;color: #7D9AC0;min-width: 150px;}
  .app-view .container-fluid .card .info span{font-size: 14px;color: #223E62;font-weight: 400;}
  .nav.nav-tabs.nav-linetriangle .nav-item a.nav-link.active {
    border-bottom: none!important;
    color: #24C4A4!important;
  }
  .nav.nav-tabs.nav-linetriangle.nav-justified {border-bottom: none!important;width: min-content!important;}
  .nav.nav-tabs.nav-linetriangle .nav-item a.nav-link.active:before,
  .nav.nav-tabs.nav-linetriangle .nav-item a.nav-link.active:after {display: none!important;}
</style>
@section('content')
<div class="app-view" style="max-width: 1141px">
  <div class="app-page-title">
  </div>
  <div class="container-fluid mb-4">
    <div class="row">
      <div class="col-sm-12 col-md-4">
        <div class="card card-avatar box-shadow-1 mx-1 mt-2 round">
            <div class="text-xs-center">
                <div class="card-block images-avatar">
                    <img src="{{asset('images/avatars/avatar_default_noname.png')}}" class="rounded " height="107" width="107" alt="">
                </div>
                <div class="card-block">
                    <h4 class="">{{$user->getFullname()}}</h4>
                    <h4 class="">MNV:{{$user->getCode()}}</h4>
                    <h6 class="card-subtitle text-muted"></h6>
                </div>
                <div class="text-xs-center d-none">
                    <a href="#" class="btn btn-social-icon mr-1 mb-1 btn-outline-facebook"><span class="fa fa-facebook"></span></a>
                    <a href="#" class="btn btn-social-icon mr-1 mb-1 btn-outline-twitter"><span class="fa fa-twitter"></span></a>
                    <a href="#" class="btn btn-social-icon mb-1 btn-outline-linkedin"><span class="fa fa-linkedin font-medium-4"></span></a>
                </div>
            </div>
            <div class="row px-3 info">
              <div class="col-12">
                  <label>Tên tài khoản:</label>
                  <span>{{$user->name}}</span>
              </div>
              <div class="col-12">
                  <label>Email:</label>
                  <span>{{$user->email}}</span>
              </div>
              <div class="col-12">
                  <label>Vai trò:</label>
                  <span>{{$user->role}}</span>
              </div>
            </div>
        </div>
      </div>
      <div class="col-sm-12 col-md-8">
        <div class="card mb-0">
          <div class="card-header d-none">
            <h4 class="card-title"></h4>
          </div>
          <div class="card-body">
            <div class="card-block px-1 py-0">
              <ul class="nav nav-tabs nav-linetriangle no-hover-bg nav-justified">
                <li class="nav-item">
                  <a class="nav-link active" id="active-tab3" data-toggle="tab" href="#active3" aria-controls="active3" aria-expanded="true">Thông tin chung</a>
                </li>
                <li class="nav-item">
                  <a class="nav-link" id="link-tab3" data-toggle="tab" href="#link3" aria-controls="link3" aria-expanded="false">Thay đổi mật khẩu</a>
                </li>
              </ul>
              <div class="tab-content px-1 pt-1">
                <div role="tabpanel" class="tab-pane fade active in" id="active3" aria-labelledby="active-tab3" aria-expanded="true">
                  <p>
                    {!! Form::model($user, array('action' => array('ProfilesController@updateUserAccount', $user->id), 'method' => 'PUT', 'id' => 'user_basics_form')) !!}
              {!! csrf_field() !!}

              <div class="pt-4 pr-3 pl-2 form-group has-feedback row {{ $errors->has('name') ? ' has-error ' : '' }}">
              {!! Form::label('name', trans('forms.create_user_label_username'), array('class' => 'col-md-4 control-label')); !!}
              <div class="col-md-8">
                <div class="input-group">
                {!! Form::text('name-', $user->name, array('id' => 'name', 'class' => 'form-control','readonly'=>true, 'placeholder' => trans('forms.create_user_ph_username'))) !!}
                <div class="input-group-append input-group-addon">
                  <label class="input-group-text" for="name">
                  <i class="fa fa-fw {{ trans('forms.create_user_icon_username') }}" aria-hidden="true"></i>
                  </label>
                </div>
                </div>
                @if($errors->has('name'))
                <span class="help-block">
                  <strong>{{ $errors->first('name') }}</strong>
                </span>
                @endif
              </div>
              </div>

              <div class="pr-3 pl-2 form-group has-feedback row {{ $errors->has('email') ? ' has-error ' : '' }}">
              {!! Form::label('email', trans('forms.create_user_label_email'), array('class' => 'col-md-4 control-label')); !!}
              <div class="col-md-8">
                <div class="input-group">
                {!! Form::text('email-', $user->email, array('id' => 'email', 'class' => 'form-control','readonly'=>true, 'placeholder' => trans('forms.create_user_ph_email'))) !!}
                <div class="input-group-append input-group-addon">
                  <label for="email" class="input-group-text">
                  <i class="fa fa-fw {{ trans('forms.create_user_icon_email') }}" aria-hidden="true"></i>
                  </label>
                </div>
                </div>
                @if ($errors->has('email'))
                <span class="help-block">
                  <strong>{{ $errors->first('email') }}</strong>
                </span>
                @endif
              </div>
              </div>

              <div class="pr-3 pl-2 form-group has-feedback row {{ $errors->has('first_name') ? ' has-error ' : '' }}">
              {!! Form::label('first_name', trans('forms.create_user_label_firstname'), array('class' => 'col-md-4 control-label')); !!}
              <div class="col-md-8">
                <div class="input-group">
                {!! Form::text('first_name', $user->first_name, array('id' => 'first_name', 'class' => 'form-control', 'placeholder' => trans('forms.create_user_ph_firstname'))) !!}
                <div class="input-group-append input-group-addon">
                  <label class="input-group-text" for="first_name">
                  <i class="fa fa-fw {{ trans('forms.create_user_icon_firstname') }}" aria-hidden="true"></i>
                  </label>
                </div>
                </div>
                @if($errors->has('first_name'))
                <span class="help-block">
                  <strong>{{ $errors->first('first_name') }}</strong>
                </span>
                @endif
              </div>
              </div>

              <div class="pr-3 pl-2 form-group has-feedback row {{ $errors->has('last_name') ? ' has-error ' : '' }}">
              {!! Form::label('last_name', trans('forms.create_user_label_lastname'), array('class' => 'col-md-4 control-label')); !!}
              <div class="col-md-8">
                <div class="input-group">
                {!! Form::text('last_name', $user->last_name, array('id' => 'last_name', 'class' => 'form-control', 'placeholder' => trans('forms.create_user_ph_lastname'))) !!}
                <div class="input-group-append input-group-addon">
                  <label class="input-group-text" for="last_name">
                  <i class="fa fa-fw {{ trans('forms.create_user_icon_lastname') }}" aria-hidden="true"></i>
                  </label>
                </div>
                </div>
                @if($errors->has('last_name'))
                <span class="help-block">
                  <strong>{{ $errors->first('last_name') }}</strong>
                </span>
                @endif
              </div>
              </div>

              {{-- <div class="d-none pr-3 pl-2 form-group has-feedback row {{ $errors->has('bio') ? ' has-error ' : '' }}">
              {!! Form::label('bio', trans('profile.label-bio') , array('class' => 'col-md-4 control-label')); !!}
              <div class="col-md-8">
                {!! Form::textarea('bio', $user->profile->bio, array('id' => 'bio', 'class' => 'form-control', 'placeholder' => trans('profile.ph-bio'))) !!}
                <span class="glyphicon glyphicon-pencil form-control-feedback" aria-hidden="true"></span>
              </div>
              @if ($errors->has('bio'))
                <span class="help-block">
                <strong>{{ $errors->first('bio') }}</strong>
                </span>
              @endif
              </div> --}}

              <div class="row form-group pr-3 pl-2">
              <div class="col-md-4"></div>
              <div class="col-md-8 text-xs-right">
                <input type="reset" name="" value="Hủy" class="btn mx-1" style="min-width: 160px;">
                {!! Form::button(
                '<i class="fa fa-fw fa-save" aria-hidden="true"></i> ' . trans('profile.submitProfileButton'),
                array(
                  'class'     => 'btn btn-success disabled',
                  'id'    => 'account_save_trigger',
                  'type'    => 'submit',
                )
                )!!}
              </div>
              </div>
            {!! Form::close() !!}
                  </p>
                </div>
                <div class="tab-pane fade" id="link3" role="tabpanel" aria-labelledby="link-tab3" aria-expanded="false">
                  <p>
                    {!! Form::model($user, array('action' => array('ProfilesController@updateUserPassword', $user->id), 'method' => 'PUT', 'autocomplete' => 'password')) !!}

                <div class="pw-change-container margin-bottom-2">

                <div class="form-group has-feedback row {{ $errors->has('password') ? ' has-error ' : '' }}">
                  {!! Form::label('password', trans('forms.create_user_label_password_current'), array('class' => 'col-md-4 control-label')); !!}
                  <div class="col-md-8">
                  {!! Form::password('password', array('id' => 'password', 'class' => 'form-control ', 'placeholder' => trans('forms.create_user_label_password_current'), 'autocomplete' => 'password', 'required'=>'')) !!}
                  @if ($errors->has('password'))
                    <span class="help-block">
                    <strong>{{ $errors->first('password') }}</strong>
                    </span>
                  @endif
                  </div>
                </div>
                <div class="form-group has-feedback row {{ $errors->has('password_new') ? ' has-error ' : '' }}">
                  {!! Form::label('password_new', trans('forms.create_user_label_password_new'), array('class' => 'col-md-4 control-label')); !!}
                  <div class="col-md-8">
                  {!! Form::password('password_new', array('id' => 'password_new', 'class' => 'form-control ', 'placeholder' => trans('forms.create_user_label_password_new'), 'autocomplete' => 'new-password','required'=>'')) !!}
                  @if ($errors->has('password_new'))
                    <span class="help-block">
                    <strong>{{ $errors->first('password_new') }}</strong>
                    </span>
                  @endif
                  </div>
                </div>

                <div class="form-group has-feedback row {{ $errors->has('password_confirmation') ? ' has-error ' : '' }}">
                  {!! Form::label('password_confirmation', trans('forms.create_user_label_password_confirm'), array('class' => 'col-md-4 control-label')); !!}
                  <div class="col-md-8">
                  {!! Form::password('password_confirmation', array('id' => 'password_confirmation', 'class' => 'form-control', 'placeholder' => trans('forms.create_user_label_password_confirm'), 'required'=>'')) !!}
                  <span id="pwcf_status"></span>
                  @if ($errors->has('password_confirmation'))
                    <span class="help-block">
                    <strong>{{ $errors->first('password_confirmation') }}</strong>
                    </span>
                  @endif
                  </div>
                </div>
                </div>
                <div class="form-group row">
                <div class="col-md-4"></div>
                <div class="col-md-8 text-xs-right">
                  {!! Form::button(
                  '<i class="fa fa-fw fa-save" aria-hidden="true"></i> ' . trans('profile.submitPWButton'),
                   array(
                    'class'     => 'btn btn-success',
                    'id'    => 'pw_save_trigger',
                    'type'    => 'submit',
                    'data-submit'   => trans('profile.submitButton'),
                  )) !!}
                </div>
                </div>
              {!! Form::close() !!}
                  </p>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
    
    <div class="card border-0 d-none">
      <div class="card-body p-0">
      {{-- @if (session('success'))
        <div class="alert alert-success m-3">
        {!! session('success') !!}
        </div>
      @endif --}}
      <script type="text/javascript">
        $(document).ready(function() {
        @if (session('open-tab'))
          $('.nav-link').eq(1).click();
        @endif
        @if (session('update-success'))
          swalAlert('{{session('update-success')}}')
        @endif
        @if (session('success'))
          swalAlert('{{session('success')}}', );
        @endif
        @if (session('failed'))
          @if($errors->has('password'))
          @php 
            $msg = str_replace('.', '</br>', $errors->get('password')[0] ?? ''); 
          @endphp
          swalAlert('', 'error',{html:'{!!$msg!!}'});
          
          @else
          swalAlert('{{session('failed')}}', 'error');
          @endif
        @endif
        })
      </script>
      @if ($user->profile)
        @if (Auth::user()->id == $user->id)
        <div class="row">
          <div class="col-12 col-sm-4 col-md-3 profile-sidebar text-white rounded-left-sm-up">
          <div class="nav flex-column nav-pills" id="v-pills-tab" role="tablist" aria-orientation="vertical">
            <a class="nav-link active" data-toggle="pill" href=".edit-settings-tab" role="tab" aria-controls="edit-settings-tab" aria-selected="true">
            {{ trans('profile.editAccountTitle') }}
            </a>
            <a class="nav-link" data-toggle="pill" href=".edit-account-tab" role="tab" aria-controls="edit-settings-tab" aria-selected="false">
            {{ trans('profile.editAccountAdminTitle') }}
            </a>
          </div>
          </div>
          <div class="col-12 col-sm-8 col-md-9">
          <div class="tab-content" id="v-pills-tabContent">
            <div class="tab-pane fade show active edit-settings-tab" role="tabpanel" aria-labelledby="edit-settings-tab">
            {!! Form::model($user, array('action' => array('ProfilesController@updateUserAccount', $user->id), 'method' => 'PUT', 'id' => 'user_basics_form')) !!}

              {!! csrf_field() !!}

              <div class="pt-4 pr-3 pl-2 form-group has-feedback row {{ $errors->has('name') ? ' has-error ' : '' }}">
              {!! Form::label('name', trans('forms.create_user_label_username'), array('class' => 'col-md-3 control-label')); !!}
              <div class="col-md-9">
                <div class="input-group">
                {!! Form::text('name-', $user->name, array('id' => 'name', 'class' => 'form-control','readonly'=>true, 'placeholder' => trans('forms.create_user_ph_username'))) !!}
                <div class="input-group-append input-group-addon">
                  <label class="input-group-text" for="name">
                  <i class="fa fa-fw {{ trans('forms.create_user_icon_username') }}" aria-hidden="true"></i>
                  </label>
                </div>
                </div>
                @if($errors->has('name'))
                <span class="help-block">
                  <strong>{{ $errors->first('name') }}</strong>
                </span>
                @endif
              </div>
              </div>

              <div class="pr-3 pl-2 form-group has-feedback row {{ $errors->has('email') ? ' has-error ' : '' }}">
              {!! Form::label('email', trans('forms.create_user_label_email'), array('class' => 'col-md-3 control-label')); !!}
              <div class="col-md-9">
                <div class="input-group">
                {!! Form::text('email-', $user->email, array('id' => 'email', 'class' => 'form-control','readonly'=>true, 'placeholder' => trans('forms.create_user_ph_email'))) !!}
                <div class="input-group-append input-group-addon">
                  <label for="email" class="input-group-text">
                  <i class="fa fa-fw {{ trans('forms.create_user_icon_email') }}" aria-hidden="true"></i>
                  </label>
                </div>
                </div>
                @if ($errors->has('email'))
                <span class="help-block">
                  <strong>{{ $errors->first('email') }}</strong>
                </span>
                @endif
              </div>
              </div>

              <div class="pr-3 pl-2 form-group has-feedback row {{ $errors->has('first_name') ? ' has-error ' : '' }}">
              {!! Form::label('first_name', trans('forms.create_user_label_firstname'), array('class' => 'col-md-3 control-label')); !!}
              <div class="col-md-9">
                <div class="input-group">
                {!! Form::text('first_name', $user->first_name, array('id' => 'first_name', 'class' => 'form-control', 'placeholder' => trans('forms.create_user_ph_firstname'))) !!}
                <div class="input-group-append input-group-addon">
                  <label class="input-group-text" for="first_name">
                  <i class="fa fa-fw {{ trans('forms.create_user_icon_firstname') }}" aria-hidden="true"></i>
                  </label>
                </div>
                </div>
                @if($errors->has('first_name'))
                <span class="help-block">
                  <strong>{{ $errors->first('first_name') }}</strong>
                </span>
                @endif
              </div>
              </div>

              <div class="pr-3 pl-2 form-group has-feedback row {{ $errors->has('last_name') ? ' has-error ' : '' }}">
              {!! Form::label('last_name', trans('forms.create_user_label_lastname'), array('class' => 'col-md-3 control-label')); !!}
              <div class="col-md-9">
                <div class="input-group">
                {!! Form::text('last_name', $user->last_name, array('id' => 'last_name', 'class' => 'form-control', 'placeholder' => trans('forms.create_user_ph_lastname'))) !!}
                <div class="input-group-append input-group-addon">
                  <label class="input-group-text" for="last_name">
                  <i class="fa fa-fw {{ trans('forms.create_user_icon_lastname') }}" aria-hidden="true"></i>
                  </label>
                </div>
                </div>
                @if($errors->has('last_name'))
                <span class="help-block">
                  <strong>{{ $errors->first('last_name') }}</strong>
                </span>
                @endif
              </div>
              </div>

              {{-- <div class="d-none pr-3 pl-2 form-group has-feedback row {{ $errors->has('bio') ? ' has-error ' : '' }}">
              {!! Form::label('bio', trans('profile.label-bio') , array('class' => 'col-md-3 control-label')); !!}
              <div class="col-md-9">
                {!! Form::textarea('bio', $user->profile->bio, array('id' => 'bio', 'class' => 'form-control', 'placeholder' => trans('profile.ph-bio'))) !!}
                <span class="glyphicon glyphicon-pencil form-control-feedback" aria-hidden="true"></span>
              </div>
              @if ($errors->has('bio'))
                <span class="help-block">
                <strong>{{ $errors->first('bio') }}</strong>
                </span>
              @endif
              </div> --}}

              <div class="form-group row">
              <div class="col-md-9 offset-md-3">
                {!! Form::button(
                '<i class="fa fa-fw fa-save" aria-hidden="true"></i> ' . trans('profile.submitProfileButton'),
                array(
                  'class'     => 'btn btn-success disabled',
                  'id'    => 'account_save_trigger',
                  'type'    => 'submit',
                )
                )!!}
              </div>
              </div>
            {!! Form::close() !!}
            </div>

            <div class="tab-pane fade edit-account-tab" role="tabpanel" aria-labelledby="edit-account-tab">
            <div class="tab-content">
              <div id="changepw" class="pt-4 pr-3 pl-2 tab-pane fade show active">

              {!! Form::model($user, array('action' => array('ProfilesController@updateUserPassword', $user->id), 'method' => 'PUT', 'autocomplete' => 'password')) !!}

                <div class="pw-change-container margin-bottom-2">

                <div class="form-group has-feedback row {{ $errors->has('password') ? ' has-error ' : '' }}">
                  {!! Form::label('password', trans('forms.create_user_label_password_current'), array('class' => 'col-md-3 control-label')); !!}
                  <div class="col-md-9">
                  {!! Form::password('password', array('id' => 'password', 'class' => 'form-control ', 'placeholder' => trans('forms.create_user_label_password_current'), 'autocomplete' => 'password', 'required'=>'')) !!}
                  @if ($errors->has('password'))
                    <span class="help-block">
                    <strong>{{ $errors->first('password') }}</strong>
                    </span>
                  @endif
                  </div>
                </div>
                <div class="form-group has-feedback row {{ $errors->has('password_new') ? ' has-error ' : '' }}">
                  {!! Form::label('password_new', trans('forms.create_user_label_password_new'), array('class' => 'col-md-3 control-label')); !!}
                  <div class="col-md-9">
                  {!! Form::password('password_new', array('id' => 'password_new', 'class' => 'form-control ', 'placeholder' => trans('forms.create_user_label_password_new'), 'autocomplete' => 'new-password','required'=>'')) !!}
                  @if ($errors->has('password_new'))
                    <span class="help-block">
                    <strong>{{ $errors->first('password_new') }}</strong>
                    </span>
                  @endif
                  </div>
                </div>

                <div class="form-group has-feedback row {{ $errors->has('password_confirmation') ? ' has-error ' : '' }}">
                  {!! Form::label('password_confirmation', trans('forms.create_user_label_password_confirm'), array('class' => 'col-md-3 control-label')); !!}
                  <div class="col-md-9">
                  {!! Form::password('password_confirmation', array('id' => 'password_confirmation', 'class' => 'form-control', 'placeholder' => trans('forms.create_user_label_password_confirm'), 'required'=>'')) !!}
                  <span id="pwcf_status"></span>
                  @if ($errors->has('password_confirmation'))
                    <span class="help-block">
                    <strong>{{ $errors->first('password_confirmation') }}</strong>
                    </span>
                  @endif
                  </div>
                </div>
                </div>
                <div class="form-group row">
                <div class="col-md-9 offset-md-3">
                  {!! Form::button(
                  '<i class="fa fa-fw fa-save" aria-hidden="true"></i> ' . trans('profile.submitPWButton'),
                   array(
                    'class'     => 'btn btn-warning',
                    'id'    => 'pw_save_trigger',
                    'type'    => 'submit',
                    'data-submit'   => trans('profile.submitButton'),
                  )) !!}
                </div>
                </div>
              {!! Form::close() !!}
              </div>
            </div>
            </div>
          </div>
          </div>
        </div>
        @else
        <p>{{ trans('profile.notYourProfile') }}</p>
        @endif
      @else
        <p>{{ trans('profile.noProfileYet') }}</p>
      @endif

      </div>
    </div>
  </div>
</div>
@endsection

{{-- @include('modals.modal-form') --}}

@section('scripts')

  @include('scripts.form-modal-script')

  @include('scripts.user-avatar-dz')

  <script type="text/javascript">

  $('.dropdown-menu li a').click(function() {
    $('.dropdown-menu li').removeClass('active');
  });

  $('.profile-trigger').click(function() {
    $('.panel').alterClass('card-*', 'card-default');
  });

  $('.settings-trigger').click(function() {
    $('.panel').alterClass('card-*', 'card-info');
  });

  $('.warning-pill-trigger').click(function() {
    $('.panel').alterClass('card-*', 'card-warning');
  });

  $('.danger-pill-trigger').click(function() {
    $('.panel').alterClass('card-*', 'card-danger');
  });

  $('#user_basics_form').on('keyup change', 'input, select, textarea', function(){
    $('#account_save_trigger').attr('disabled', false).removeClass('disabled').show();
  });

  $("#password_confirmation").keyup(function() {
    checkPasswordMatch();
  });

  $("#password_new, #password_confirmation").keyup(function() {
    enableSubmitPWCheck();
  });

  $('#password_new, #password_confirmation').hidePassword(true);

  $('#password_new').password({
    shortPass: 'The password is too short',
    badPass: 'Weak - Try combining letters & numbers',
    goodPass: 'Medium - Try using special charecters',
    strongPass: 'Strong password',
    containsUsername: 'The password contains the username',
    enterPass: false,
    showPercent: false,
    showText: true,
    animate: true,
    animateSpeed: 50,
    username: false, // select the username field (selector or jQuery instance) for better password checks
    usernamePartialMatch: true,
    minimumLength: 6
  });

  function checkPasswordMatch() {
    var passwordNew = $("#password_new").val();
    var confirmPassword = $("#password_confirmation").val();
    if (passwordNew.length < 6) {
    $("#pwcf_status").html("Mật khẩu ít nhất 6 ký tự!");
    $("#password_new").focus();
    } else if (passwordNew != confirmPassword) {
    $("#pwcf_status").html("Mật khẩu xác nhận không chính xác!");
    } else {
    $("#pwcf_status").html("");
    }
  }

  function enableSubmitPWCheck() {
    var password = $("#password_new").val();
    var confirmPassword = $("#password_confirmation").val();
    var submitChange = $('#pw_save_trigger');
    if (password != confirmPassword) {
    submitChange.attr('disabled', true);
    }
    else {
    submitChange.attr('disabled', false);
    }
  }
  </script>
@endsection
