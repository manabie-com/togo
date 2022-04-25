
@section('template_title', trans('usersmanagement.showing-user', ['name' => @$Animal->name]))
<style type="text/css">
</style>
@section('navbar-more')
@endsection
@section('content')
  <div class="app-view">
    <div class="card card-outline-info round p-2">
      <div class="card-title ">        
        Bạn là <span class="text-bold-700">{{$animal->user->name}}</span>
      </div>
      <div class="card-body">
        <div class="card-block">
          <div>
            Số lần luyện tập trong ngày là: <span class="h5">{{$animal->user->queue_today??0}}/{{@$animal->user->queue_limit??0}}</span>
          </div>          
        </div>
      </div>
    </div>
    <div class="h4">
    </div>
  </div>
@endsection
<x-layout.greenland />
