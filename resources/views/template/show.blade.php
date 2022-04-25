@php
  /**
  * @todo
  */
  $pageTitle = $pageTitle ?? '';
  $classSplitCol = ' col-md-'.(12 / 1);
  $addLink = $addLink ?? '';
  $editLink = $editLink ?? $viewModel->getLink('edit');
@endphp
@section('navbar-more')
  <div class="text-uppercase" >{{$pageTitle}}</div>
@endsection
  <style type="text/css">    
  </style>
<div class="app-view">
  <div class="row px-2">
    <div class="col-12 col-md-7">
      
    </div>
    <div class="col-12 col-md-5">      
    </div>
  </div>
</div>
@section('scripts')
  @parent
  <script type="text/javascript"></script>
@endsection  
