
@section('template_title', trans('usersmanagement.showing-user', ['name' => $user->name]))
<style type="text/css">
</style>
@section('navbar-more')
@endsection
@section('content')
  <div class="app-view">
    "My home"
  </div>
@endsection
<x-layout.greenland />
