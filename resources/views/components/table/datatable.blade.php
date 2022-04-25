@props(['colTitle'=> [], 'dataRows' => [], 'class'=>'datatable', 'id'=>''])
@inject('Str', 'Illuminate\Support\Str')
@push('css')
<style type="text/css"> 
</style>
@endpush
<div class="table-responsive form-group mb-0 bg-danger-">
    <table id="{{$id}}" class="{{$class}} align-middle nowarp mb-0 table table-borderless table-striped- table-hover ">
        <thead class="border-bottom text-uppercase">
            @foreach($colTitle as $_tt_k => $_tt_v)
            <th class="{{'th-'.$_tt_k}}">{!!$_tt_v!!}</th>
            @endforeach
        </thead>
        <tbody>
            @foreach($dataRows as $_key => $_item)
            <tr id="tr-{{$_item['id'] ?? ''}}" style="border-radius: 25px">
                @foreach($colTitle as $_tt_k => $_tt_v)
                @php
                $value = $_item[$_tt_k] ?? '';
                $slugAtt = str_replace('_', '-', $_tt_k);
                $tdClass = 'lbl-'.$slugAtt;
                @endphp
                <td class="td-{{$slugAtt}}">
                    @switch(1)
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
                    @case(!!preg_match('/smalltool/', $_tt_k))
                    <div class="{{$tdClass}} row">
                        @if(is_array($value)) 
                        @foreach ($value as $t_key => $tool)
                        {!!$tool!!}
                        @endforeach
                        @endif
                    </div>
                    @break
                    @case(!!preg_match('/^_action$/', $_tt_k) && is_array($value))
                    <div class="{{$tdClass}} row">
                        @foreach ($value as $t_key => $url)
                        @switch($t_key)
                        @case('show')
                        <a href="{{$url}}" class="{{$_tt_k.$t_key}}" title="{{__('app.'.$t_key)}}">
                            <span class="fonticon-wrap pr-1 text-info"><i class="fa fa-eye"></i></span>
                        </a>
                        @break
                        @case('edit')
                        <a href="{{$url}}" class="{{$_tt_k.$t_key}}" title="{{__('app.'.$t_key)}}">
                            <span class="fonticon-wrap pr-1 text-warning"><i class="fa fa-edit"></i></span>
                        </a>
                        @break
                        @case('destroy')
                        <a href="{{$url}}" label="{{$_item['code'] ?? ''}}" class="btn-delete {{$_tt_k.$t_key}}" title="{{__('app.'.$t_key)}}">
                            <span class="fonticon-wrap pr-0 text-danger"><i class="fa fa-trash"></i></span>
                        </a>
                        @break
                        @default
                        @endswitch
                        @endforeach
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
            @endforeach
        </tbody>
    </table>
    @if(isset($pagination))
    <div class="px-3">
        {{ $pagination->appends(request()->input())->render('partials.pagination') }}
    </div>
    @endif
</div>
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
        });
    </script>
    @endpush
