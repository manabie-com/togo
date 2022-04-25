@section('content')
@inject('Str', 'Illuminate\Support\Str')
@php
    $statusList = $statusList ?? [];
    $exportLink = $exportLink ?? '';
    $addLink    = $addLink ?? '';
    array_unshift($statusList, ['code'=> '', 'name'=>'Tất cả', 'count' => null]);
@endphp

@push('css')
<style type="text/css">

    .app-view .container-fluid .row.r1 .col-xs-12{display: inline-flex;}
    .app-view .container-fluid .row.r1 .form-group{margin-right: 17px;min-width: 200px;}
    .app-view .container-fluid .row.r1 .form-group >label{min-height: 24px;font-weight: 400;}

    .app-view .container-fluid .filter-content{display: inline-flex;user-select: none;overflow-x: auto;width: 100%;}
    .app-view .container-fluid .filter-content input.filter{display: none;}
    .app-view .container-fluid .filter-content div.form-control{width: auto;max-width: 250px; margin: 4px 4px;border-radius: 16px;min-width: 120px!important; text-align: center;border: 0px solid #faa227;background-color: #F2F3F5;white-space:nowrap; cursor: pointer; }
    .app-view .container-fluid .filter-content label{font-weight: 400;}
    .app-view .container-fluid .filter-content div.form-control .tag{background-color: #faa227;}
    .app-view .container-fluid .filter-content div.form-control.checked{background-color: #faa227;color: white;}
    .app-view .container-fluid .filter-content div.form-control.checked .tag{background-color:#F2F3F5;color: #faa227;}
    .app-view .container-fluid table thead tr th{padding: 18px 20px;text-align: left;}
    .app-view table tbody tr:hover{cursor: pointer;user-select: none;}
    .app-view table tbody tr td .tag{color: black;min-width: 140px;font-size: 100%;border-radius: 26px; padding: 8px 20px;font-size: 16px;line-height: 24px;}
    .app-view table tbody tr td .tag.dang-van-chuyen{color: #0052FF;background-color: #E7EEFF;}
    .app-view table tbody tr td .tag.dang-cho-duyet{color: #AE9615;background-color: #FBF6CF;}
    .app-view table tbody tr td .tag.hoan-thanh{color: #1EAE71;background-color:  #E6FAF1;}
    .app-view table tbody tr td .tag.huy{color: #000000;background-color: #E7EEFF;}
</style>
@endpush
<x-block>
    <div class="app-view">
        <div class="row">
            <div class="col-xs-12 col-md-5">
                <div class="app-view-title">vận chuyển</div>
            </div>
            <div class="col-xs-12 col-md-7 text-md-right d-inline">
              <div class="btn btn-outline-icon-text btn-export-excel disabled" href="{{$exportLink}}"><i class="fa fa-file-excel-o d-inline"></i>Xuất file Excel</div>
            </div>
        </div>
        <div class="container-fluid">
            <x-form method="GET" action="{{url()->current()}}">
                <div class="row">
                    <div class="col-xs-12">
                        <div class="filter-content">
                            @foreach($statusList as $status)
                                <label for="chb-{{$status['code']}}">
                                    <div class="form-control {{request('filter') == $status['code'] ? 'checked':''}}">{{$status['name']}}
                                        @if(!empty($status['count']))
                                            <span class="tag tag-pill tag-warning">{{$status['count']}}</span>
                                        @endif
                                    </div>
                                </label>
                                <input type="radio" id="chb-{{$status['code']}}" class="filter" name="filter" value="{{$status['code']}}" {{request('filter') == $status['code'] ? 'checked':''}} />
                            @endforeach
                        </div>
                    </div>
                </div>
                <div class="row">
                    <div class="col-xs-12" style="display: inline-flex;">
                        <x-form.search placeholder="Tìm mã đơn hàng" show-btn-clear="false" />
                    </div>
                </div>
                <div class="row r1">
                    <div class="col-xs-12">
                        <fieldset class="form-group">
                            <label class="" style="display:block">Loại vận chuyển</label>
                            <select class="select2 form-control block" name="report-by">
                                <option>Oto</option>
                                <option>Xe khách</option>
                                <option>Máy bay</option>
                                <option>Tàu hỏa</option>
                            </select>
                        </fieldset>
                        <fieldset class="form-group">
                            <label>Chọn thời gian </label>
                            <div class='input-group' >
                               <input type='text' name="date-range" class="form-control dateranges" value="{{old('date-range-items') ?? $dateRangeValue ?? ''}}" />
                                <span class="input-group-addon">
                                    <svg width="26" height="24"></svg>
                                </span>
                            </div>
                            <small class="text-muted"></small>
                        </fieldset>
                    </div>
                </div>
            </x-form>
            <div class="card card-upos">
                <div class="card-body">
                    <div class="card-block p-0">
                        @if(!empty($dataRows))
                        <div class="table-responsive table-striped form-group mb-0 bg-danger-">
                            <table id="datatable" class="datatable align-middle nowarp mb-0 table table-borderless table-striped- table-hover ">
                                <thead class="border-bottom text-uppercase">
                                    @foreach($colTitle as $_tt_k => $_tt_v)
                                        <th class="{{'th-'.$_tt_k}}">{!!$_tt_v!!}</th>
                                    @endforeach
                                </thead>
                                <tbody>
                                    @php $countRow = 0; @endphp
                                    @foreach($dataRows as $_key => $_elements)
                                        @php $countRow++ @endphp
                                        <tr id="tr-{{$countRow}}" style="border-radius: 25px">
                                            @foreach($colTitle as $_tt_k => $_tt_v)
                                                @php
                                                    $value = $_elements[$_tt_k] ?? '';
                                                    $slugAtt = str_replace('_', '-', $_tt_k);
                                                    $tdClass = 'lbl-'.$slugAtt;
                                                @endphp
                                                <td class="td-{{$slugAtt}}">
                                                    @switch(1)
                                                        {{-- @case(!!preg_match('/^items$/', $_tt_k))
                                                            @foreach($_elements['items'] as $_index =>$_item)
                                                                <div class="row">
                                                                    <label class="" style="min-width: 250px">{{$_item['item']??''}}</label>
                                                                    <label class="">x{{$_item['quantity']??''}}</label>
                                                                </div>
                                                            @endforeach
                                                            @break --}}
                                                        @case(!!preg_match('/^ship_code$/', $_tt_k))
                                                            <div class="{{$tdClass}}">
                                                                <div class="text-bold-700 text-blue">{{$value}}</div>
                                                                <div>15:40 29/11/2021</div>
                                                            </div>
                                                            @break
                                                        @case(!!preg_match('/^_action2$/', $_tt_k) && is_array($value))
                                                            <div class="{{$tdClass}} row">
                                                                <a class=" text-info" href="" ><i class="fa fa-file-o"></i></a>
                                                                </div>
                                                            </div>
                                                            @break
                                                        @case(!!preg_match('/^sum_amount$/', $_tt_k))
                                                            <div class="{{$slugAtt}} text-xs-right">{{number_format($value)}}<span class="unit-currency"> đ</span></div>
                                                            @break
                                                        @case(!!preg_match('/^transport$/', $_tt_k))
                                                            <div style="display: inline-flex;">
                                                                <div class="mr-1">
                                                                    <img src="/images/avatars/ghn.svg" width="36" height="36">
                                                                </div>
                                                                <div class="">
                                                                    <div class=" text-bold-700 text-nowarp">{{$value}}</div>
                                                                    <div class=""> 5 - 6 ngày</div>
                                                                </div>
                                                            </div>
                                                            @break
                                                        @case(!!preg_match('/^code-link$/', $_tt_k))
                                                            <div class="{{$tdClass}}">
                                                                <a class=" text-info btn-utmddf-" href="{{$value['href'] ?? ''}}">{{$value['label']??''}}</a>
                                                            </div>
                                                            @break
                                                        @case(!!preg_match('/-link$/', $_tt_k))
                                                            {{s($_tt_k)}}
                                                            <div class="{{$tdClass}}">
                                                                <a class=" text-info btn-utmddf-" href="{{$value['href'] ?? ''}}">{{$value['label']??''}}</a>
                                                            </div>
                                                            @break
                                                        @case(!!preg_match('/(_date|datetime|created_at)$/', $_tt_k))
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
                                                        @default
                                                            @if(!is_array($value) && !is_object($value))
                                                                <div class="{{$tdClass}}">{{$value}}</div>
                                                            @endif
                                                    @endswitch
                                                </td>
                                            @endforeach
                                        </tr>
                                        <ul id="sub-tr-{{$countRow}}" class="sub-row sub-tr-{{$countRow}} d-none">
                                            <li style="">
                                                @foreach($_elements['items'] as $_index =>$_item)
                                                    <div class="row" style="display:block;">
                                                        <label class="m-1" style="min-width: 250px">{{$_item['item']??''}}</label>
                                                        <label class="m-1">x{{$_item['quantity']??''}}</label>
                                                    </div>
                                                @endforeach
                                            </li>
                                        </ul>
                                    @endforeach
                                </tbody>
                            </table>
                            @if(isset($pagination))
                                <div class="px-3">
                                    {{ $pagination->appends(request()->input())->render('partials.pagination') }}
                                </div>
                            @endif
                        </div>
                        @else
                            <hr>
                            <div class="alert alert-info">Không có kết quả phù hợp !</div>
                        @endif
                    </div>
                </div>
            </div>
        </div>
    </div>
</x-block>
@endsection
@push('scripts')
    <script type="text/javascript">
        $('input.filter').on('change', function(){
            $(this).closest('form').submit();
        })
    </script>
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
                    ordering: false,
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
                if($('table.datatable th.th-no').length){
                    dataTable.column(0, {search:'applied', order:'applied'}).nodes().each( function (cell, i) {cell.innerHTML = '<div class="lbl-no text-center px-1">' + (i+1) + '</div>';});
                }
            }).draw(false);
            /**/
            $('input[type=reset]').on('click', function(e){
                $('#search_range_date').val("");
            })
            /**/
            
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
            /*toggle row show detail items */
            $('table.datatable tbody tr').on('click', function(){
                let trID = $(this)[0].id;
                let colspan = $('table.datatable thead tr th').length;
                let subHTML = $('ul#sub-'+trID).find('li').html();
                console.log('tr#show-'+trID, subHTML);
                if ($('tr#show-'+trID)[0] != undefined) {
                    console.log($('tr#show-'+trID));
                    $('tr#show-'+trID)[0].remove();
                }else{
                    $(this).after('<tr id="show-'+trID+'"><td colspan="'+colspan+'" style="padding: 0rem 2rem"><div style="padding: 0rem 2rem">'+subHTML+'</div></td></tr>');
                }
            })
        });
    </script>
    <script type="text/javascript">
        let daterange1 = $('.dateranges').daterangepicker({
            "showCustomRangeLabel": false,
            "alwaysShowCalendars": false,
            ranges: {
              // 'Hôm nay': [moment(), moment()],
              // 'Hôm qua': [moment().subtract(1, 'days'), moment().subtract(1, 'days')],
              // '7 Ngày trước': [moment().subtract(6, 'days'), moment()],
              // '30 Ngày trước': [moment().subtract(29, 'days'), moment()],
              'Tháng này': [moment().startOf('month'), moment().endOf('month')],
              'Tháng trước': [moment().subtract(1, 'month').startOf('month'), moment().subtract(1, 'month').endOf('month')]
            },
            // drops: "down",
            buttonClasses: "btn",
            applyClass: "btn-info mx-3",
            cancelClass: "btn-danger d-none",
            locale: {
              format: 'DD/MM/YYYY',
              applyLabel: "Áp dụng",
              cancelLabel: 'Hủy',
              startLabel: 'Bắt đầu',
              customRangeLabel: 'Tùy chọn',
            }
        });
        $('input.dateranges').on('apply.daterangepicker', function(ev, picker) {$(this).closest('form').submit();});
        $('.input-group-addon').on('click', function(e){$(this).parent().find('.dateranges').trigger('click');})
        /*submit form*/
        $('form#order-report select[name=report-by]').on('change', function(e){
            $(this).closest('form').submit();
        });
    </script>
@endpush
<x-layout.default/>

