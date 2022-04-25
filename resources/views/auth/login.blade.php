@push('css')
    <style type="text/css">
        #login .card{max-width: 400px; margin: 0px auto; background-color: inherit;}
        #login .card .card-body {background-color: #fff;border-radius: 6px;}
        #login .card .card-body .card-block {padding: 32px 40px;}
        #login .card .card-body .block-1 h4{color: #172E4D;font-size: 18px;}
        #login .card .card-body .block-1 p{font-size: 12px;}
        #login .card .card-body .block-1{color: #7D9AC0;}
        #login .card .card-body .block-2{}
        #login .card .card-body .block-2 input {min-width: 250px;}
        #login .card .card-body .btn-login{color: #fff;background-color: #FAA227; margin: 0px auto;}
        #login .card .card-body label.form-control{border: none}
        #login .card .card-body .input-group .form-control:focus{border-color: #faa227!important;}
        #login .card .card-body .row.r1 >h4{font-weight: 500;font-size: 16px;line-height: 24px;}
        /*input:-webkit-autofill {-webkit-background-clip: text;}*/
    </style>
@endpush
@section('content')
    <div id="login" class="d-flex">
        <div class="container d-flex align-items-center">
            <form id="form_login" method="POST" action="{{ route('login') }}" >
                <div class="row p-2">
                    <div class="col-lg-12 text-xs-center">
                    </div>
                </div>
                <div class="card">
                    @csrf
                    <div class="card-body">
                        <div class="card-block">
                            <div class="row r1 text-xs-center block-1">
                                <h4 class="text-uppercase">Hoang vực</h4>
                                {{-- <p class="text-xs-center">Nhập thông tin để truy cập vào tài khoản của bạn</p> --}}
                                @if($errors->has('notice'))
                                <div class="col-12 p-2">
                                    <span class="alert alert-danger help-block">
                                        <strong>{{ $errors->first('notice')}}</strong>
                                    </span>
                                </div>
                                @endif
                            </div>
                            <div class="row mt-2">
                                <div class="col-xs-12">                                    
                                    <fieldset>
                                        <div class="input-group">
                                            <span class="input-group-addon" id="basic-addon1">Lệnh bài</span>
                                            <input id="email" type="text" class=" form-control {{ $errors->has('email') ? ' is-invalid' : '' }}" name="email" value="{{ old('email') }}" required autofocus placeholder="">
                                            {{-- <input type="text" class="form-control" placeholder="Addon to Left" aria-describedby="basic-addon1"> --}}
                                        </div>
                                            <input id="id" type="text" class="form-control" hidden name="name" value="" >

                                        <span id="hb-username" class="help-block">
                                            <strong>{{ $errors->first('email')}}{{ $errors->first('id') }}</strong>
                                        </span>
                                    </fieldset>
                                </div>
                            </div>                            
                            <div class="row mt-2">
                                <div class="col-xs-12">                                    
                                    <fieldset>
                                        <div class="input-group">
                                            <span class="input-group-addon" id="basic-addon2">
                                                <div class="" style="cursor:pointer;">Ám hiệu</div>
                                            </span>
                                            {{-- <span class="input-group-addon d-xs-block d-none-" id="basic-addon2">Mã xác nhận</span> --}}
                                            {{-- <input id="email" type="text" class=" form-control {{ $errors->has('email') ? ' is-invalid' : '' }}" name="email" value="{{ old('email') }}" required autofocus placeholder=""> --}}
                                            <input id="password" type="password" class="form-control{{ $errors->has('password') ? ' is-invalid' : '' }}" name="password" minlengthp="6" required placeholder="Mã xác nhận" aria-describedby="basic-addon2">
                                            {{-- <input type="text" class="form-control" placeholder="Addon to Left" aria-describedby="basic-addon1"> --}}
                                        </div>

                                        <span id="hb-password" class="help-block">
                                            <strong>{{ $errors->first('password') }}</strong>
                                        </span>
                                    </fieldset>
                                </div>
                            </div>
                            <div class="row form-group d-none">
                                <div class="col-md-6 offset-md-4">
                                    <div class="checkbox">
                                        <label>
                                            <input type="checkbox" name="remember" {{ old('remember') ? 'checked' : '' }}> {!! trans('auth.rememberMe') !!}
                                        </label>
                                    </div>
                                </div>
                            </div>
                            <div class="row mt-2">
                                <div class="col-xs-12 text-xs-center">
                                    <button type="submit" class="btn btn-icon-text btn-login btn-outline-warning" style="width: 100%">
                                        Bắt đầu
                                    </button>
                                </div>
                            </div>
                            <div class="row block-3">
                                <div class="col-xs-12">
                                    <a class="btn btn-link d-none" href="{{ route('password.request') }}">
                                        {{ __('auth.forgot') }}
                                    </a>
                                </div>
                                <div class="col-xs-12">
                                   {{--  <button type="submit" class="btn btn-login">
                                        {!! trans('titles.login') !!}
                                    </button> --}}
                                </div>
                            </div>
                            <div class="row align-text-bottom">
                                <div class="col-12">
                                    {{-- @include('systems.version') --}}
                                </div>
                            </div>
                        </div>
                    </div>
                </div>
            </form>
        </div>
    </div>
@endsection
@push('scripts')
    <script type="text/javascript">
        $(document).ready(function(){
            $('#email').on('change, keyup', function(e){
                $('#id').val($(this).val());
                validate('email');
            })
            $('#password').on('change, keyup', function(e){
                validate('password');
            })
        });
        function validate(inputID) {
          const input = document.getElementById(inputID);
          const validityState = input.validity;
          if (input.value.length < 1) {
            input.classList.add("is-invalid");
          } else{
            input.classList.remove("is-invalid");
            $(input).closest('.col').find('.help-block').html('');
          }
        }
    </script>
@endpush
<x-layout.onlycontent />