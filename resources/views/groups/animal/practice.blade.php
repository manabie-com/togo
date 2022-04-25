<meta http-equiv="refresh" content="3">
@section('template_title', trans('usersmanagement.showing-user', ['name' => @$Animal->name]))
<style type="text/css">
</style>
@section('navbar-more')
@endsection
@section('content')
  <div class="app-view">
    Luyện tập nào <span class="h5">{{$animal->user->queue_today??0}}/{{@$animal->user->queue_limit??0}}</span>
  </div>
  <p>Đang luyện tập</p>
@endsection
<x-layout.greenland />
