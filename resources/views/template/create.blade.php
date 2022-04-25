@php
    /**
     * @author toannguyen.dev
     * @todo
     */
    /*intial*/
    $model = $model ?? die('404');
    $trans = $trans ?? [];
@endphp
@push('css')
  <link rel="stylesheet" type="text/css" href="{{asset('app-assets/vendors/css/forms/icheck/icheck.css')}}">
  <link rel="stylesheet" type="text/css" href="{{asset('app-assets/vendors/css/forms/icheck/custom.css')}}">
  <link rel="stylesheet" type="text/css" href="{{asset('app-assets/css/plugins/forms/checkboxes-radios.css')}}">
  <link rel="stylesheet" type="text/css" href="{{asset('app-assets/vendors/css/pickers/datetime/bootstrap-datetimepicker.css')}}">
  <style>
    .app-view {}
    .app-view.app-view-create .card {max-width: 1141px;margin: 0px;}
    .app-view.app-view-create input.form-control,
    .app-view.app-view-create select.form-control,
    .app-view.app-view-create textarea.form-control,
    .app-view.app-view-create .singledate-picker{min-height: 43px; height: 43px}
    .app-view.app-view-create .control-field > .form-control, 
    .app-view.app-view-create .select2,
    .app-view.app-view-create .control-field > .form-control, 
    .app-view.app-view-create .form-group{max-width: 540px;min-height: 43px;}
    /*.app-view.app-view-create .form-actions{max-width:1100px;}*/
    .app-view.app-view-create .form-actions{max-width:550px;}
    .app-view.app-view-create .form-actions .btn {min-width: 140px;font-size: 18px;}
    .app-view.app-view-create .select2-container .select2-selection--multiple{display: flex;}
    .app-view.app-view-create .select2-container .select2-selection--multiple .select2-selection__rendered {display: inline!important;}
  </style>
@endpush
@section('content')
  <div class="app-view app-view-create bg-info" >
    <div class="app-view-title"></div>
    <div class="container-fluid">
      <form class="form" action="{{$model->route('store')}}" id="frm_create" enctype="multipart/form-data" method="POST">
        @csrf
        @method('POST')
        <div class="card m-0">
          @if(!empty($createTitle))
          <div class="card-header p-0">
            <div class="heading-title bg-upos p-1 rounded">{{trans($createTitle)}}</div>
          </div>
          @endif
          <div class="card-body collapse in">
            <div class="card-block pt-0 pb-1">
              <div class="row d-non">
                @foreach($fillable as $_attr => $_label)
                  @php
                    $value = old($_attr) ?? $_attrValue['value'] ?? null;
                    $has_error = $errors->has($_attr) ? ' has-error ' : '';
                  @endphp
                  <div class="">
                    <div class="{{$has_error}} as-row form-group row">
                      <label for="" class="col control-label">
                        {{$trans[$_attr] ?? $_label}}
                      </label>
                      <div class="col control-field d-non">
                        @switch(1)
                          @case(!!preg_match('/_date|created_at|updated_at/', $_attr))
                            @php
                              $value = date_create($value) && !empty($value) ? date_create($value)->format('d/m/Y') : '';
                              $value = old($_attr) ?? $value;
                                $txtInputId = 'txt-input-date'.$_attr;
                            @endphp
                            <div class="form-group singledate-picker m-0">
                              <div class='input-group date' id='{{$txtInputId}}'>
                                <input type='text' class="form-control" value="{{$value}}"/>
                                <span class="input-group-addon">
                                  <span class="fa fa-calendar-o"></span>
                                </span>
                              </div>
                            </div>
                            <script type="text/javascript">
                            $(document).ready(function() {
                              $('#{{$txtInputId}}').datetimepicker({format: 'DD/MM/YYYY'});
                            })
                            </script>
                            @break
                          @case(isset($attrArray['type']) && !!preg_match('/radio/', $attrArray['type']))
                            <div class="form-control">
                              <div class="row skin skin-line ">
                                @foreach($option as $_index => $_tagArr)
                                  @php
                                    $id_rio = 'rio-' . $_attr . $_index;
                                  @endphp
                                  <div class="col-md-12 col-sm-12">
                                    <fieldset>
                                      <input type="{{$attrArray['type']}}" name="{{$_attr}}" id="{{$id_rio}}" {{$_tagArr['checked']}} value="{{$_tagArr['value']}}" {{$attrString}}>
                                      <label for="{{$id_rio}}">{{$_tagArr['label']}}</label>
                                    </fieldset>
                                  </div>
                                @endforeach
                              </div>
                            </div>
                            @break
                          @default
                            <input type="text" class="form-control" name="{{$_attr}}"  value="{{$value}}" />
                              
                        @endswitch
                      </div>
                      @error($_attr)
                        <div class="has-error-message">
                          @foreach ($errors->get($_attr) as $message)
                            {{$message}} <br>
                          @endforeach
                        </div>
                      @enderror
                    </div>
                  </div>
                @endforeach
              </div>
              @hasSection('form-actions')
                @yield('form-actions')
              @else
                <div class="form-actions right row">
                  <button type="button" class="btn" style="background-color: inherit; color: rgb(125, 154, 192);">
                    <a class="" href="{{$model->route('index')}}" style="color:inherit;">
                      <svg width="18" height="18" viewBox="0 0 18 18" fill="none" xmlns="http://www.w3.org/2000/svg">
                        <path d="M19 12H5L12 19M12 5L8 9" stroke="#7D9AC0" stroke-linecap="round" stroke-linejoin="round"/>
                      </svg>
                      <span >Danh sách</span>
                    </a>
                  </button>
                  <button type="submit" tabindex="0" class="btn btn-add mx-auto" style="max-width: 200px">Thêm mới</button>
              </div>
              @endif
            </div>
          </div>
        </div>
      </form>
    </div>
  </div>
@endsection
@push('scripts')
  <script src="{{asset('app-assets/vendors/js/forms/icheck/icheck.min.js')}}" type="text/javascript"></script>
  <script src="{{asset('app-assets/js/scripts/forms/checkbox-radio.js')}}" type="text/javascript"></script>
  <script type="text/javascript">
    $(document).on('keyup', function (e) {
      if(e.keyCode === 13) $('#btn-create').trigger('click');
    })
    $(document).ready(function() {
      @if (session('failed'))
        swalAlert("{!! session('failed') !!}", "error");
      @elseif(session('success'))
        let title ="{{session('success')}}", text = '{{session('text')}}';
        if(text === '' || text == undefined){
          text = "Bạn có muốn đến trang danh sách ?";
        }
        swalConfirm(title, text, function(r){
          if (r) {location.replace('{{$listLink}}')}
        },{
          imageUrl:'/images/site/accepted.png',
          confirmButtonText: "Có",
          cancelButtonText:" Không",
        });
      @endif
      /**/
      $('.timepicker').each(function(i, ele){
        $(ele).datetimepicker({format: 'YYYY-MM-DD HH:mm:ss'})
      })
      $('.txt-effect-date').daterangepicker({
        singleDatePicker: true,
        locale: {format: 'DD/MM/YYYY'}
      });
    });
    /*prevent set validation default*/
    $('input').on('invalid', function (e) {
      e.preventDefault();
      e.target.required = false;
      $(this).closest('form').submit();
    });
    /*set autofocus the first input has error*/
    $('.has-error').first().find('input')[0].select();
  </script>
@endpush
<x-layout.default />
