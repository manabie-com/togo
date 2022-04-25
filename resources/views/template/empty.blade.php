@section('content')
<x-block>
  <div class="app-view">
    <div class="app-view-title">
      {{$title ?? ''}}      
    </div>
    <div class="container-fluid">
      {{$content ?? url()->current()}}
    </div>
  </div>
</x-block>
@endsection
<x-layout.default />