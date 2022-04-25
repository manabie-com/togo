
  <nav class="header-navbar navbar navbar-with-menu navbar-fixed-top navbar-light navbar-border">
    <div class="navbar-wrapper">
      <div class="navbar-header" style="z-index: 9000;">
        {{-- <ul class="nav"> --}}
          {{-- <li class="nav-item"> --}}
            <a href="/" class="navbar-brand">
              <span class=" text-uppercase text-bold-700 p-1" style="color:#d32f2f; z-index: 4000">{{env('APP_NAME')??'Home'}}</span>
            </a>
          {{-- </li> --}}
        {{-- </ul> --}}
      </div>
      <div class="navbar-container content container-fluid px-1">
        <div id="navbar-mobile" class="collapse navbar-toggleable-sm">          
          <ul class="nav navbar-nav float-xs-right">
            <li class="dropdown dropdown-user nav-item">
              <a href="#" data-toggle="dropdown" class="dropdown-toggle nav-link dropdown-user-link">
                <span class="avatar avatar-online">
                  <img src="{{url('app-assets/images/portrait/small/avatar-s-1.png')}}" alt="avatar"><i></i></span>
                  <span class="user-name">{{Auth::user()->name ?? ''}}</span>
                </a>
              <div class="dropdown-menu dropdown-menu-right">
                <a href="{{route('logout')}}" class="dropdown-item"><i class="ft-power"></i>Rời khỏi</a>
              </div>
            </li>
          </ul>
        </div>
      </div>
    </div>
  </nav>