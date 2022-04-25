@inject('Str', 'Illuminate\Support\Str')
@php
  /**/
  $dataRows = $dataRows ?? [];
  $colTitle = $colTitle ?? [];
  $pageTitle = $pageTitle ?? '';
  $importLink = $importLink ?? '';
  $exportLink = $exportLink ?? '';
  $addLink = $addLink ?? '';
@endphp
@push('csss')
  <style type="text/css">
    .app-view.app-view-index {}
    /*.app-view table tbody tr td {padding: .5rem}*/
    .app-view.app-view-index table thead tr th.th-no,
    .app-view.app-view-index table tbody tr td.td-no {max-width: 30px!important;padding-right:7px!important;padding-left:7px!important;}
    .app-view.app-view-index table tbody tr td >div.lbl-description{max-height: 100px;overflow-y: auto;}
    .app-view.app-view-index table tbody tr td.td-description{min-width: 300px}
    /*customize*/
  </style>
@endpush
@section('content')
  <x-block>
    <div class="app-view app-view-index">
      <div class="app-view-title">
        <form method="GET" action="{{url()->current()}}">
            <div class="row">
              <div class="mb-h1 col col-xs-12 col-md-6 float-xs-left">
               {{--  <div class="btn-">
                  @if(!empty($filterData))
                    @include('components.form.filter')
                  @else
                    @include('components.form.search')
                  @endif
                </div> --}}
                <x-form.search />
              </div>
              <div class="mb-h1 col col-xs float-xs-left float-md-right">
                @if(!empty($exportLink))
                  <a class="btn btn-export-excel export-list-" href="{{$exportLink}}" target="_blank">
                    <span class="" >{!! trans('app.export-excel') !!}</span>
                    <i class="fa fa-file-excel-o fa-1x mr-1" aria-hidden="true" title="{!! trans('app.export-excel') !!}"></i>
                  </a>
                @endif
                @if(!empty($addLink))
                <a class="btn btn-add " href="{{$addLink}}">Thêm mới<i class="fa fa-plus ml-1"></i></a>
                @endif
              </div>
            </div>
        </form>
      </div>
      <div class="container-fluid">
        <div class="row">
          <div class="col-xs-12 p-0">
            @if(!empty($dataRows))
              <div class="card card-upos mb-0">
                <div class="card-body">
                  <div class="table-responsive form-group mb-0">
                    <table class="datatable align-middle nowarp mb-0 table table-borderless table-striped- table-hover ">
                      <thead class="border-bottom text-uppercase">
                        @foreach($colTitle as $_tt_k => $_tt_v)
                          <th class="{{'th-'.$_tt_k}}">{!!$_tt_v!!}</th>
                        @endforeach
                      </thead>
                      <tbody>
                        @foreach($dataRows as $_key => $_item)
                          @php 
                            $id = $_item['id'] ?? '';
                          @endphp
                          <tr id="tr-{{$id}}" style="border-radius: 25px">
                            @foreach($colTitle as $_tt_k => $_tt_v)
                              @php
                                $value = $_item[$_tt_k] ?? '';
                                $slugAtt = str_replace('_', '-', $_tt_k);
                                $tdClass = 'lbl-'.$slugAtt;
                              @endphp
                              <td class="td-{{$slugAtt}}">
                                @switch(1)
                                  {{-- Default --}}
                                  @case(!!preg_match('/^code-link$/', $_tt_k))
                                    <div class="{{$tdClass}}">
                                      <a class=" text-info btn-utmddf-" href="{{$value['href'] ?? ''}}">{{$value['label']??''}}</a>
                                    </div>
                                    @break
                                  @case(!!preg_match('/-link$/', $_tt_k))
                                    <div class="{{$tdClass}}">
                                      <a class=" text-info btn-utmddf-" href="{{$value['href'] ?? ''}}">{{$value['label']??''}}</a>
                                    </div>
                                    @break
                                  @case(!!preg_match('/smalltool/', $_tt_k))
                                    <div class="{{$tdClass}} row">
                                      @if(is_array($value)) 
                                        @foreach ($value as $t_key => $tool)
                                         {!!$tool!!}
                                        @endforeach
                                      @endif
                                    </div>
                                    @break
                                   @case(!!preg_match('/^_action$/', $_tt_k))
                                    <div class="{{$tdClass}} row">
                                        @foreach ($value as $t_key => $url)
                                         <a href="{{$url}}" class="{{$_tt_k.$t_key}}" title="{{__('app.'.$t_key)}}">
                                            @switch($t_key)
                                              @case('show')
                                                <span class="fonticon-wrap pr-1 text-info"><i class="fa fa-eye"></i></span>
                                                @break
                                              @case('edit')
                                                <span class="fonticon-wrap pr-1 text-warning"><i class="fa fa-edit"></i></span>
                                                @break
                                              @case('destroy')
                                                <span class="fonticon-wrap pr-0 text-danger"><i class="fa fa-trash"></i></span>
                                                @break
                                              @default
                                            @endswitch
                                         </a>
                                        @endforeach
                                    </div>
                                    @break
                                  @case(!!preg_match('/(_date|datetime|created_at|updated_at|deleted_at)$/', $_tt_k))
                                    @if(empty($value) || !date_create($value)) @break @endif
                                    <div class="{{$tdClass}}">{{date('d/m/Y H:i',strtotime($value))}}</div>
                                    @break
                                  @case(!!preg_match('/^type$/', $_tt_k))
                                    @if(is_string($value))
                                      <div class=" {{$tdClass}} {{$Str::slug($value)}}">{{$value}}</div>
                                    @elseif(is_array($value))
                                      <div class=" {{$tdClass}} {{$Str::slug($value['code']??'')}}">{{$value['name']??''}}</div>
                                    @endif
                                    @break
                                  @case(!!preg_match('/^status$/', $_tt_k))
                                    @if(is_string($value))
                                      <div class="tag {{$tdClass}} {{$Str::slug($value)}}">{{$value}}</div>
                                    @elseif(is_array($value))
                                      <div class="tag {{$tdClass}} {{$Str::slug($value['code']??'')}}">{{$value['name']??''}}</div>
                                    @endif
                                    @break
                                  {{-- Customize --}}
                                  @case(!!preg_match('/^roles$/', $_tt_k))
                                    @foreach( $value as $_index => $role)
                                      {{ucfirst(__('app.'.$role['name'] ?? ''))}} <br>
                                    @endforeach
                                    @break
                                  @default
                                    @if(!is_array($value) && !is_object($value))
                                      <div class="{{$tdClass}}">{{$value}}</div>
                                    @endif
                                @endswitch
                              </td>
                            @endforeach
                          </tr>
                        @endforeach
                      </tbody>
                    </table>
                    @if(isset($pagination))
                      <div class="px-3">
                        {{ $pagination->appends(request()->input())->render('partials.pagination') }}
                      </div>
                    @endif
                  </div>
                </div>
              </div>
            @endif
          </div>
        </div>
      </div>
    </div>
  </x-block>
@endsection
@push('scripts')
  <script>
    let dataTable = null;
    $(document).ready(function() {
      @if(session('success'))
        toastFlash({title:"{!! session('success') !!}",timer:2000});
      @elseif(session('failed'))
        toastFlash({title:"{!! session('failed') !!}",timer:2000,icon:'warning'});
      @endif
      /*datatable*/
      let config = {
          sDom: 'tb',
          paging: false,
          oLanguage: {
            sLengthMenu: "_MENU_",
            sInfo: "_START_ - _END_ of _TOTAL_ ",
            oPaginate: {
              "sFirst": '<i class="fas fa-step-backward"></i>',"sLast": '<i class="fas fa-step-forward"></i>',
              "sNext": '<i class="fas fa-chevron-right"></i>',
              "sPrevious": '<i class="fas fa-chevron-left"></i>'},
            sSearch: "",
          },
          language:{
            infoFiltered:" / _MAX_ ",
            select: { rows: '<span class="px-1"> Đã chọn %d </span>'},
          },
          colReorder: true,
        }
      dataTable = $('.datatable').DataTable(config);
      dataTable.on( 'order.dt search.dt', function () {
        dataTable.column(0, {search:'applied', order:'applied'}).nodes().each( function (cell, i) {cell.innerHTML = '<div class="lbl-no text-center px-1">' + (i+1) + '</div>';});
        }).draw(false);
      /**/
      $('input[type=reset]').on('click', function(e){
        $('#search_range_date').val("");
      })
      /**/
      $('#btn-clear-filter').on('click', function(e){
        $(this).closest('form').find('input[type=text], select').val("");
        $(this).closest('form').submit();
      })
      /*delete*/
      $('.btn-delete').on('click', function(e){
        e.preventDefault();
        let this_btn = this;
        let label = $(this_btn).attr('label');
        let url = $(this_btn).attr('href');
        if(label === undefined) {swalAlert('Không xác định đối tượng !', 'error');return false;}
        if(url === '') {swalAlert('Không tìm thấy URL !', 'error');return false;}
        swalConfirm("Xóa " + label,"Bạn đã chắc chắn ?", function(r){
          if (r) {
            $.ajax({
              url: url,
              type:'POST',
              data: {_method:'DELETE', _token:'{{csrf_token()}}'},
              dataType : "json",
              beforeSend: function(request) {
                request.setRequestHeader("X-CSRF-TOKEN",'{{csrf_token()}}');
              },
              success: function(response) {},
              statusCode: {
                301: function() {
                  let row = $(this_btn).closest('tr');
                  console.log(row);
                  dataTable.row(row).remove().draw()
                  toastFlash({title:"{{trans('titles.deleted_success')}}",timer:2000});
                },
                302: function(){
                  toastFlash({title:"{{trans('titles.deleted_failure')}}",timer:2000,icon:"warning"});
                }
              }
            });
          }
        },{
          confirmButtonText: "Xóa",
          cancelButtonText:"Đóng",
          icon:"warning",
        });
      })
    });
  </script>
  @yield('scripts-more')
@endpush
<x-layout.greenland />