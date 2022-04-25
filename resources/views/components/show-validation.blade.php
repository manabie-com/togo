@if ($errors->any() && (env('APP_ENV') === 'local') && !empty($errors->all()))
  <div class="" style="position:fixed;z-index:3000;bottom: 0px;width: 80%;max-width: 800px;">
      <div class="card bg-warning">
          <div class="card-header p-0">
              <h4 class="card-title"><div class="h4">Kiểm thử:: {{env('APP_ENV')}}</div></h4>
              <a class="heading-elements-toggle"><i class="fa fa-ellipsis-v font-medium-3"></i></a>
            <div class="heading-elements">
                  <ul class="list-inline mb-0">
                      {{-- <li><a data-action="collapse"><i class="ft-minus"></i></a></li> --}}
                      {{-- <li><a data-action="reload"><i class="ft-rotate-cw"></i></a></li> --}}
                      <li><a data-action="expand"><i class="ft-maximize"></i></a></li>
                      <li><a data-action="close"><i class="ft-x"></i></a></li>
                  </ul>
              </div>
          </div>
          <div class="card-body" >
              <div class="card-block bg-white ">
                <ul class="">
                  @foreach ($errors->all() as $error)
                    <li>{{$error}}</li>
                  @endforeach
                </ul>
              </div>
          </div>
      </div>
  </div>

@endif
