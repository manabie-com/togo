@php
    $config = $config ?? [];
@endphp
@push('css')
    <style>
        .fa, .fas, .far, .fal, .fab {
            line-height: 1.25;
        }
        .paginator .pagination {
            margin: 0;
            display: flex;
            padding-left: 0;
            list-style: none;
            border-radius: .25rem;
        }
        .paginator .pagination li.active {
            position: relative;
            display: block;
            padding: 9px;
            /*margin-left: -1px;*/
            line-height: 1.25;
            color: #fff;
            background-color: #FAA227;
            min-width: 36px;
            min-height: 36px;
            padding: 9px;
            text-align: center;
            font-weight: 700;
            /*border: 1px solid #2cc194;*/
        }
        .paginator .pagination .disabled {
            position: relative;
            display: block;
            padding: .5rem .75rem;
            margin-left: -1px;
            line-height: 1.25;
            color: #2cc194;
            background-color: #ffffff;
            border: 1px solid #e9ecef;
        }
        .paginator .pagination li{height: 40px;width: 40px; border-radius: 3px;font-size: 16px;}
        .paginator .pagination li a {
            position: relative;
            display: block;
            line-height: 1.25;
            color: #aaa;
            background-color: #fff;
            min-width: 40px;
            min-height: 40px;
            padding: 9px;
            text-align: center;
        }
        .paginator .pagination li a:hover {
            color: #FAA227;
            font-weight: 700;
            /*border: 1px solid #24C4A4;*/
        }
        /*.pagination li i{display: none;}*/
        .paginator .pagination li.page-first a > div > span,
        .paginator .pagination li.page-last a > div > span,
        .paginator .pagination li.page-previous a div span,
        .paginator .pagination li.page-next a div span{ display: none; }
        .paginator .per-page {font-weight:600;display: inline-flex;min-width: 40px;}
        .paginator .per-page select {
            -webkit-appearance: none;
            -moz-appearance: none;
            appearance: none;
            padding: 10px 27px;
            border: 1px solid #DBDBDB;
            box-sizing: border-box;
            border-radius: 40px;
            width: 132px;
            height: 66px;
        }
        .paginator .per-page .select2-selection.select2-selection--single{min-width: 90px;}
        .paginator .pull-left, .paginator .pull-right {padding: 9px;}
        .paginator ._info{font-size: 12px;color: #999999}
    </style>
@endpush
    <div class="form-inline paginator">
        <input type="number" name="page" value="{{ $paginator->currentPage() ?? request('page') ?? 1 }}" hidden>
        <div class="pull-left" style="display:inline-flex;">
            <div class="_counter" style="padding:10px 0px">
                @if(!!($config['counter'] ??true))
                <span style="color: #4D6E99">{{__('app.sum')}}: <span style="font-weight:700;">{{$paginator->total()}}</span></span>
                @endif
            </div>
            <div class="_length px-1 ">
                @php
                    $all = $paginator->total();
                    $perPage = $paginator->perPage();
                    $perPageList = $perPageList ?? [10 => 10, 20 => 20, 50 => 50, 100 => 100, $perPage => $perPage, $all => 'Tất cả'];
                    $perPage = $perPage > $all ? $all : $perPage;
                    foreach ($perPageList as $key => $label) {
                        if($key > $all) {unset($perPageList[$key]); continue;}
                        $perPageList[$key] = ['value' => $key, 'label' => $label, 'selected' => ''];
                    }
                    $perPageList[$perPage]['selected'] = 'selected';
                    asort($perPageList);
                @endphp
                <span class="">{{$lang['display'] ?? ''}}</span class="">
                <div class="per-page form-group h5 m-0" style="" title="Số dòng mỗi trang">
                    <label for="perPage"></label>
                    <select class="select form-control border-bottom" id="perPage"  name="per-page">
                        @foreach($perPageList as $key => $value)
                            <option value="{{$key}}" {{$value['selected'] ?? ''}} > {{$value['label'] ?? ''}}</option>
                        @endforeach
                    </select>
                </div>
            </div>
            <div class="_info" style="padding:10px">
                @if(!!($config['info'] ??true))
                    <span>
                        Từ {{min($paginator->currentPage() * $paginator->PerPage() - $paginator->PerPage() + 1, $paginator->total())}}
                        đến {{min($paginator->currentPage() * $paginator->PerPage(), $paginator->total())}}
                        trên {{$paginator->total()}}
                        {{$lang['_OF_'] ?? ''}}
                    </span>
                @endif
            </div>
        </div>
        <div class="pull-right" style="display:inline-flex;">
            <div class="_paging">
                <ul class="pagination">
                    {{-- Previous Page Link --}}
                    @if ($paginator->onFirstPage())

                    @else
                        <li class="page-first">
                            <a class="paginator-link" data-page="{{ 1 }}" href="{{ $paginator->url(1) . '&per-page='.$paginator->perPage() }}">
                                <div ><i class="fa fa-angle-double-left" title="Trang đầu"></i> <span>{{$lang['page-first'] ?? '<<'}}</span></div>
                            </a>
                        </li>
                        <li class="page-previous">
                            <a class="paginator-link" data-page="{{ $paginator->currentPage() -1 }}" href="{{ $paginator->previousPageUrl() . '&per-page='.$paginator->perPage() }}">
                                <div><i class="fa fa-angle-left" title="Quay lại"></i><span>{{$lang['page-previous'] ?? '<'}}</span></div>
                            </a>
                        </li>
                    @endif

                    {{-- Pagination Elements --}}
                    @foreach ($elements as $element)
                        {{-- Array Of Links --}}
                        @if (is_array($element))
                            @foreach ($element as $page => $url)
                                @if ($page == $paginator->currentPage())
                                    <li class="active"><span>{{ $page }}</span></li>
                                @elseif ($page >= ($paginator->currentPage() - 3) && $page <= ($paginator->currentPage() + 3) )
                                    <li><a class="paginator-link" data-page="{{ $page}}" href="{{ $url . '&per-page='.$paginator->perPage() }}">{{ $page }}</a></li>
                                @endif
                            @endforeach
                        @endif
                    @endforeach
                    {{-- Next Page Link --}}
                    @if ($paginator->hasMorePages())
                        <li class="page-next">
                            <a class="paginator-link" data-page="{{ $paginator->currentPage()+1}}" href="{{ $paginator->nextPageUrl() . '&per-page='.$paginator->perPage()}}">
                                <div><i class="fa fa-angle-right" title="Trang sau"></i><span>{{$lang['page-next'] ?? '>'}}</span></div>
                            </a>
                        </li>
                        <li class="page-last">
                            <a class="paginator-link" data-page="{{ $paginator->lastPage()}}" href="{{ $paginator->url($paginator->lastPage()) . '&per-page='.$paginator->perPage() }}">
                                <div><i class="fa fa-angle-double-right" title="Trang cuối"></i><span>{{$lang['page-last'] ?? '<<'}}</span></div>
                            </a>
                        </li>
                    @endif
                </ul>
            </div>
        </div>
    </div>
@push('scripts')
<script type="text/javascript">
    $(document).on('click', '.paginator-link', function(e){
        $('input[name=page]').val(parseInt($(this).attr('data-page')));
        let this_form = $(this).closest('form');
        if(isset(this_form)) {
            e.preventDefault();
            $(this).closest('form').submit();
        }
    })
    $(document).on('change', 'select[name=per-page]', function(e){
        let total = parseInt({{ $paginator->total() }});
        let perPage = parseInt($(this).val());
        let page = parseInt($('input[name=page]').val());
        page = Math.min(parseInt(total/perPage), page);

        $('input[name=per-page]').val(perPage);
        $('input[name=page]').val(page);

        let this_form = $(this).closest('form');
        if(isset(this_form)) {
            e.preventDefault();
            $(this_form).submit();
            return true;
        }
        location.search = '&per-page=' + $(this).val() +'&page='+ page;
    })
</script>
@endpush
