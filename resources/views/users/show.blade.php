{{-- @extends('greenland') --}}
@section('template_title', trans('usersmanagement.showing-user', ['name' => $user->name]))
<style type="text/css">
  .app-view-show .card .card-title{font-size: 24px;font-weight: 700;}
  .app-view-show .nav{border-radius: 0px;}
  .app-view-show .nav.nav-tabs.nav-underline {position: relative;border-bottom: 0px solid #FAA227;}
  .app-view-show .nav.nav-tabs.nav-underline .nav-item a.nav-link {color: #333F48;}
  .app-view-show .nav.nav-tabs.nav-underline .nav-item a.nav-link:before {background: #FAA227;}
  .app-view-show .nav.nav-tabs.nav-underline .nav-item a.nav-link.active:focus,
  .app-view-show .nav.nav-tabs.nav-underline .nav-item a.nav-link.active:hover {color: #FAA227!important;}
</style>
@section('navbar-more')
  <svg width="24" height="24" viewBox="0 0 24 24" fill="none" xmlns="">
  <path d="M12 11C14.2091 11 16 9.20914 16 7C16 4.79086 14.2091 3 12 3C9.79086 3 8 4.79086 8 7C8 9.20914 9.79086 11 12 11Z" stroke="#223E62" stroke-linecap="round" stroke-linejoin="round"/>
  <path d="M16.0319 21.1421C17.4199 20.8922 18.5204 20.5115 19.0319 19.9999C19.1038 19.353 19.0319 18.3332 19.0319 18.3332C19.0319 17.4492 18.6105 16.6013 17.8604 15.9762C15.4914 14.0021 7.57247 14.0021 5.20352 15.9762C4.45337 16.6013 4.03195 17.4492 4.03195 18.3332C4.03195 18.3332 3.96007 19.353 4.03195 19.9999C4.54351 20.5115 5.64392 20.8922 7.03194 21.1421" stroke="#223E62" stroke-linecap="round" stroke-linejoin="round"/>
  </svg>
  <h4 class="text-uppercase d-inline">HỒ SƠ CỦA BẠN</h4>
@endsection
@section('content')
  <div class="app-view app-view-show">
    <div class="app-view-title row">
      <div class="col-12col-md-6 float-xs-left">
        <span class="h4 text-uppercase">
        </span>
      </div>
      <div class="col-xs float-xs-right"></div>
    </div>
    <div class="container-fluid">
      <div class="row">
          <div class="col-xl-6 col-lg-12">
            <div class="card m-1 profile" style="min">
              <div class="card-header pb-0">
                <h4 class="card-title">Tài khoản của bạn</h4>
              </div>
              <div class="card-body">
                <div class="card-block p-0">
                  <ul class="nav nav-tabs nav-underline no-hover-bg nav-justified">
                    <li class="nav-item">
                      <a class="nav-link active" id="active-tab32" data-toggle="tab" href="#active32" aria-controls="active32" aria-expanded="true">Tài khoản</a>
                    </li>
                    <li class="nav-item">
                      <a class="nav-link" id="link-tab2" data-toggle="tab" href="#link2" aria-controls="link2" aria-expanded="false">Tạo yêu cầu hỗ trợ</a>
                    </li>
                    <li class="nav-item">
                      <a class="nav-link" id="link-tab3" data-toggle="tab" href="#link3" aria-controls="link3" aria-expanded="false">Mật khẩu</a>
                    </li>
                    <li class="nav-item">
                      <a class="nav-link" href="/logout" >Đăng xuất</a>
                    </li>
                  </ul>
                  <div class="tab-content px-1 pt-1">
                    <div role="tabpanel" class="tab-pane fade active in" id="active32" aria-labelledby="active-tab32" aria-expanded="true">
                      <form method="POST" action="{{$user->url('update')}}" id="frm_edit" enctype="multipart/form-data">
                        @csrf
                        @method('PUT')
                        <div class="row">
                          <div class="col-sm-12 col-md-6">
                            <input type="text" name="txt-fullname" class="form-control" value="{{$user->name}}">
                          </div>
                        </div>
                        {{-- <hr> --}}
                        <div class="row mt-1">
                          <div class="col-sm-6 "><button type="submit" class="btn rounded">Cập nhật{{$user->name}}</button>
                          </div>
                        </div>
                      </form>
                    </div>
                    <div class="tab-pane fade" id="link2" role="tabpanel" aria-labelledby="link-tab2" aria-expanded="false">
                      <p>2</p>
                    </div>
                    <div class="tab-pane fade" id="link3" role="tabpanel" aria-labelledby="link-tab3" aria-expanded="false">
                      <p>3</p>
                    </div>
                    <div class="tab-pane fade" id="dropdownOpt22" role="tabpanel" aria-labelledby="dropdownOpt2-tab2" aria-expanded="false">
                      <p>Soufflé cake gingerbread apple pie sweet roll pudding. Sweet roll dragée topping cotton candy cake jelly beans. Pie lemon drops sweet pastry candy canes chocolate cake bear claw cotton candy wafer.</p>
                    </div>
                    <div class="tab-pane fade" id="linkOpt2" role="tabpanel" aria-labelledby="linkOpt-tab2" aria-expanded="false">
                      <p>Cookie icing tootsie roll cupcake jelly-o sesame snaps. Gummies cookie dragée cake jelly marzipan donut pie macaroon. Gingerbread powder chocolate cake icing. Cheesecake gummi bears ice cream marzipan.</p>
                    </div>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
@endsection
<x-layout.greenland />
