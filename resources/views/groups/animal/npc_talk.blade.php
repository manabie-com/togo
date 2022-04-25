@push('css')
<style type="text/css">
  .npc_talk{max-width: 600px;}
</style>
<style>
  @import url("https://fonts.googleapis.com/css?family=Raleway:900&display=swap");

body {
    margin: 0px;
}

  #container {
      position: absolute;
      margin: auto;
      width: 100vw;
      height: 80pt;
      top: 0;
      bottom: 0;

      filter: url(#threshold) blur(0.6px);
  }

  #text1,
  #text2 {
      position: absolute;
      width: 100%;
      display: inline-block;

      font-family: "Raleway", sans-serif;
      font-size: 80pt;

      text-align: center;

      user-select: none;
  }
</style>
@endpush
@section('content')
  <div class="app-view">
    <div class="npc_talk">
      <div>{{@$topic}}</div>
      <p class="form-control" id="text2">{{@$message}}</p>
    </div>
  </div>
@endsection
<x-layout.greenland />
