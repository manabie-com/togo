@extends('components.form.edit')
@php
  /**
  * @todo
  */
  $pageTitle = $pageTitle ?? '';
@endphp
<style type="text/css">
  .app-view{max-width: 1141px;}
  .app-view input,.app-view select,.app-view textarea,
  .app-view .singledate-picker,
  .app-view .control-field > .form-control, .select2,
  .app-view .control-field > .form-control, .form-group {max-width: 550px!important;  }
</style>
@section('navbar-more')
  <div class=" text-uppercase">{{$pageTitle}}</div>
@endsection
@section('container-before')
  <div></div>
@endsection
@section('scripts')
  @parent
@endsection