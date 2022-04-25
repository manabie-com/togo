@php
    $menuData = [];
    $currentURL = url()->current();
@endphp
@role('supplier')
    @php
        $menuData[] = [
            'icon' => '<img src="/images/menu/dashboard.svg">',
            'link' => route('supplier.dashboards.index'),
            'name' => 'Tổng quan',
            'alias'=> 'administrator',
            'sub-menu-alias' => 'dashboards',
            'childrens'=>[],
        ];
        $menuData[] = [
            'icon' => '<img src="/images/menu/warehouse.svg">',
            'link' => route('supplier.warehouses.index'),
            'name' => 'Kho hàng',
            'alias'=> 'warehouses',
            'sub-menu-alias' => 'supplier/warehouses/',
            'childrens'=>[
            ],
        ];
        $menuData[] = [
            'icon' => '<img src="/images/menu/cube.svg">',
            'link' => route('supplier.packages.index'),
            'name' => 'Kiện hàng',
            'alias'=> 'packages',
            'sub-menu-alias' => 'packages',
            'childrens'=>[
                // [
                //     'icon' => '<i class="fa fa-chart"></i>',
                //     'link' => route('supplier.orders.index'),
                //     'name' => 'Quản lý đơn hàng',
                //     'childrens' => []
                // ],
            ],
        ];
        $menuData[] = [
            'icon' => '<img src="/images/Icon/Outline/clipboard-list.svg">',
            'link' => route('supplier.goodsnotes.index'),
            'name' => 'Phiếu kho',
            'alias'=> 'goodsnotes',
            'sub-menu-alias' => 'goodsnotes',
            'childrens'=>[
                // [
                //     'icon' => '<i class="fa fa-chart"></i>',
                //     'link' => route('supplier.orders.index'),
                //     'name' => 'Quản lý đơn hàng',
                //     'childrens' => []
                // ],
            ],
        ];
        $menuData[] = [
            'icon' => '<img src="/images/menu/support.svg">',
            'link' => '#',
            'name' => 'Hỗ trợ',
            'alias'=> '-',
            'sub-menu-alias' => 'setup',
            'childrens'=>[]
        ];
        $menuData[] = [
            'icon' => '<img src="/images/menu/setting.svg">',
            'link' => '#',
            'name' => 'Cài đặt',
            'alias'=> '-',
            'sub-menu-alias' => 'setup',
            'childrens'=>[
                // [
                //     'icon' => '<i class="fa fa-chart"></i>',
                //     'link' => route('supplier.settings.staffs'),
                //     'name' => 'Nhân viên',
                //     'childrens' => []
                // ],
                // [
                //     'icon' => '<i class="fa fa-chart"></i>',
                //     'link' => route('supplier.settings.warehouses'),
                //     'name' => 'Kho hàng',
                //     'childrens' => []
                // ],
                // [
                //     'icon' => '<i class="fa fa-chart"></i>',
                //     'link' => route('supplier.settings.forms'),
                //     'name' => 'Mẫu in',
                //     'childrens' => []
                // ],
            ],
        ];
    @endphp
@endrole
@role('admin')
    @php
        $menuData[] = [
            'icon' => '',
            'link' => '#',
            'name' => trans('titles.administrator'),
            'alias'=> 'administrator',
            'sub-menu-alias' => 'roles,users,warehouse_types',
            'childrens'=>[
                [
                    'icon' => '<i class="fa fa-users"></i>',
                    'link' => url('users'),
                    'name' => trans('Danh sách tài khoản'),
                    'childrens' => []
                ],
                [
                    'icon' => '<i class="fa fa-tasks"></i>',
                    'link' => url('roles'),
                    'name' => trans('Phân quyền'),
                    'childrens' => []
                ],
                // [
                //     'icon' => '<i class="fa fa-users"></i>',
                //     'link' => route('admin.warehouse_types.index'),
                //     'name' => trans('Phân loại kho'),
                //     'childrens' => []
                // ],
                [
                    'icon' => '<i class="fa fa-users"></i>',
                    // 'link' => route('admin.supplier_types.index'),
                    'name' => trans('Loại nhà cung cấp'),
                    'childrens' => []
                ]
            ],
        ];
    @endphp
@endrole
@role('customer')
    @php
        $menuData[] = [
            'icon' => '<img src="/images/menu/dashboard.svg">',
            'link' => route('dashboards.index'),
            'name' => 'Tổng quan',
            'alias'=> 'administrator',
            'sub-menu-alias' => 'dashboards',
            'childrens'=>[],
        ];
        $menuData[] = [
            'icon' => '<img src="/images/menu/item.svg">',
            'link' => route('items.index'),
            'name' => 'Sản phẫm',
            'alias'=> 'items',
            'sub-menu-alias' => 'items',
            'childrens'=>[],
        ];
        $menuData[] = [
            'icon' => '<img src="/images/menu/report.svg">',
            'link' => '#',
            'name' => 'Báo cáo',
            'alias'=> 'reports',
            'sub-menu-alias' => 'report',
            'childrens'=>[
                [
                    'icon' => '<i class="fa fa-chart"></i>',
                    'link' => route('reports.income'),
                    'name' => 'Doanh thu',
                    'childrens' => []
                ],
                [
                    'icon' => '<i class="fa fa-chart"></i>',
                    'link' => route('reports.orders'),
                    'name' => 'Đơn hàng',
                    'childrens' => []
                ],
                [
                    'icon' => '<i class="fa fa-chart"></i>',
                    'link' => route('reports.items'),
                    'name' => 'Sản phẩm',
                    'childrens' => []
                ],
                [
                    'icon' => '<i class="fa fa-chart"></i>',
                    'link' => '',
                    'name' => 'Tồn kho',
                    'childrens' => []
                ],
                [
                    'icon' => '<i class="fa fa-chart"></i>',
                    'link' => '',
                    'name' => 'Bán sỉ',
                    'childrens' => []
                ],
                [
                    'icon' => '<i class="fa fa-chart"></i>',
                    'link' => '',
                    'name' => 'Bán lẻ',
                    'childrens' => []
                ],
                [
                    'icon' => '<i class="fa fa-chart"></i>',
                    'link' => '',
                    'name' => 'Khách hàng',
                    'childrens' => []
                ],
            ],
        ];        
        $menuData[] = [
            'icon' => '<img src="/images/menu/warehouse.svg">',
            'link' => '#',
            'name' => 'Kho hàng',
            'alias'=> 'warehouses',
            'sub-menu-alias' => ' ',
            'childrens'=>[
                [
                    'icon' => '<i class="fa fa-chart"></i>',
                    'link' => route('orders.import'),
                    'name' => 'Nhập kho',
                    'childrens' => []
                ],
                [
                    'icon' => '<i class="fa fa-chart"></i>',
                    'link' => route('orders.export'),
                    'name' => 'Xuất kho',
                    'childrens' => []
                ],
                [
                    'icon' => '<i class="fa fa-chart"></i>',
                    'link' => route('orders.stock'),
                    'name' => 'Tồn kho',
                    'childrens' => []
                ],
            ],
        ];
        $menuData[] = [
            'icon' => '<img src="/images/menu/shopcart.svg">',
            'link' => route('orders.index'),
            'name' => 'Đơn hàng',
            'alias'=> 'orders',
            'sub-menu-alias' => '-',
            'childrens'=>[
                [
                    'icon' => '<i class="fa fa-chart"></i>',
                    'link' => route('orders.index'),
                    'name' => 'Quản lý đơn hàng',
                    'childrens' => []
                ],
                [
                    'icon' => '<i class="fa fa-chart"></i>',
                    'link' =>  route('orders.index-duplicate'),
                    'name' => 'Đơn hàng trùng',
                    'childrens' => []
                ],
            ],
        ];
        $menuData[] = [
            'icon' => '<img src="/images/menu/delivery.svg">',
            'link' => '#',
            'name' => 'Vận chuyển',
            'alias'=> 'deliveryorders',
            'sub-menu-alias' => 'deliveryorders',
            'childrens'=>[
                [
                    'icon' => '<i class="fa fa-chart"></i>',
                    'link' => route('transports.management'),
                    'name' => 'Quản lý vận chuyển',
                ],
                [
                    'icon' => '<i class="fa fa-chart"></i>',
                    'link' => route('transports.index'),
                    'name' => 'Đơn vị vận chuyển',
                ],
            ],
        ];
        $menuData[] = [
            'icon' => '<img src="/images/menu/customer.svg">',
            'link' => '',
            'name' => 'Khách hàng',
            'alias'=> 'clients',
            'sub-menu-alias' => 'clients',
            'childrens'=>[
               [
                    'icon' => '',
                    'link' => route('clients.supplier'),
                    'name' => 'Nhà cung cấp',
                ],
                [
                    'icon' => '',
                    'link' => route('clients.index'),
                    'name' => 'Khách lẻ/sỉ',
                ],
            ],
        ];
        // $menuData[] = [
        //     'icon' => '<img src="/images/menu/support.svg">',
        //     'link' => route('support.index'),
        //     'name' => 'Hỗ trợ',
        //     'alias'=> '',
        //     'sub-menu-alias' => 'support',
        //     'childrens'=>[],
        // ];
        $menuData[] = [
            'icon' => '<img src="/images/menu/setting.svg">',
            'link' => '#',
            'name' => 'Cài đặt',
            'alias'=> '',
            'sub-menu-alias' => 'setup',
            'childrens'=>[
                
                [
                    'icon' => '<i class="fa fa-chart"></i>',
                    'link' => route('settings.staffs'),
                    'name' => 'Nhân viên',
                    'childrens' => []
                ],
                [
                    'icon' => '<i class="fa fa-chart"></i>',
                    'link' => route('settings.warehouses'),
                    'name' => 'Kho hàng',
                    'childrens' => []
                ],
                [
                    'icon' => '<i class="fa fa-chart"></i>',
                    'link' => route('settings.forms'),
                    'name' => 'Mẫu in',
                    'childrens' => []
                ],

            ],
        ];
    @endphp
@endrole
{{-- <x-block> --}}
    <div data-scroll-to-active="true" class="main-menu menu-fixed menu-light menu-accordion menu-shadow">
        <div class="main-menu-content">
            <ul id="main-menu-navigation" data-menu="menu-navigation" class="navigation navigation-main">
                <li class="navigation-header">
                    <span class=""><a href="/" style="display: contents;">fulfillment</a></span>
                    <i class="icon menu-toggle-ext btn p-0" data-toggle="tooltip" data-placement="right" data-original-title="fulfillment" >F</i>
                    <span class="menu-toggle-ext pl-1"><img src="/images/menu/toggle.svg"></span>
                </li>
                @foreach($menuData as $_index => $_menuItem)
                    @php
                        $childrens = $_menuItem['childrens'] ?? [];
                        $subMenuAlias = explode(',', $_menuItem['sub-menu-alias'] ?? '');

                        $hasSubAlias = function($subMenuAlias){
                            foreach ($subMenuAlias as $alias) {
                                $alias = str_replace('/', '\/', $alias);
                                if (preg_match("/^".$alias."/", request()->path())) {
                                    return true;
                                }
                            }
                            return false;
                        };
                        $subLinks = array_map(function ($k){return $k['link'] ?? '';}, $childrens);
                        $hasOpen = (in_array($currentURL, $_menuItem) || array_key_exists($currentURL, array_flip($subLinks)) || $hasSubAlias($subMenuAlias) ) ? 'open' : '';
                    @endphp
                    <li class="nav-item border-bottom- open- {{$hasOpen}}">
                        @if(!empty($_menuItem['name']))
                            <a href="{{$_menuItem['link']}}">
                                <div class="icon" >{!!$_menuItem['icon']!!}</div>
                                <span data-i18n="" class="menu-title">{{$_menuItem['name']}}</span>
                            </a>
                        @endif
                        @if(!empty($_menuItem['childrens']) && !empty($_menuItem['name']))
                            <ul class="menu-content">
                                @foreach($_menuItem['childrens'] as $__index => $__menuItem)
                                @php 
                                    $active = ($currentURL === ($__menuItem['link'] ?? '')) ? 'active' : '';
                                    $hr = (($__menuItem['icon'] ?? '') === '<hr>') ? 'hr' : '';
                                @endphp
                                    <li class=" nav-item border-bottom- ">
                                        <a class="menu-item {{$active}} {{$hr}}" href="{{$__menuItem['link'] ?? ''}}">
                                            <label class=" m-0 icon pr-">{!!$__menuItem['icon']!!}</label>
                                            <span class="">{{$__menuItem['name']}}</span>
                                        </a>
                                    </li>
                                @endforeach
                            </ul>
                        @else
                            @foreach($_menuItem['childrens'] as $__index => $__menuItem)
                                @php 
                                    $active = ($currentURL === ($__menuItem['link'] ?? '')) ? 'active' : '';
                                    $hr = (($__menuItem['icon'] ?? '') === '<hr>') ? 'hr' : '';
                                @endphp
                                    <li class=" nav-item border-bottom-r {{$active}}">
                                        <a class="menu-item {{$active}} {{$hr}}" href="{{$__menuItem['link'] ?? ''}}">
                                            <label class=" m-0 icon- pr-">{!!$__menuItem['icon'] ?? ''!!}</label>
                                            <span class="">{{$__menuItem['name']}}</span>
                                        </a>
                                    </li>
                                @endforeach
                        @endif
                    </li>
                @endforeach
                <li class="nav-item" style="position: absolute; bottom: 39px;border-left: inherit;">
                    <a href="#" id="menu-toggle-ext" class="">
                        <svg width="18" height="18" viewBox="0 0 18 18" fill="none" xmlns="http://www.w3.org/2000/svg">
                            <path d="M17 1V8C17 9.06087 16.5786 10.0783 15.8284 10.8284C15.0783 11.5786 14.0609 12 13 12H1L6 7M4 15L6 17" stroke="#4D6E99" stroke-linecap="round" stroke-linejoin="round"/>
                        </svg>
                        <span class="menu-title- ml-2">Ẩn bảng điều khiển</span>
                    </a>
                </li>
                {{-- <li style="position: absolute; bottom: 0px;border-left: inherit; width: 100%;background-color: #fff; padding: 5px 10px 20px 10px;">
                    <a href="{{Auth::user()->route('edit')}}" target="_self" data-toggle="dropdown-" class="" aria-expanded="false" style="padding:0px!important">
                        <span class="avatar avatar-online-">
                            <img src="/images/avatars/avatar_default.png" alt="avatar">
                        </span>
                        <span class="user-name " style="color:#000000;">{{ucfirst(Auth::user()->name ?? '')}}</span>
                    </a>
                </li>  --}}
            </ul>
            <div class="profiles" style="position: absolute; bottom: 0px;border-left: inherit; width: 100%;background-color: #fff; padding: 5px 10px 20px 10px;">
                <a href="{{Auth::user()->route('edit')}}" target="_self" data-toggle="dropdown-" class="" aria-expanded="false" style="padding:0px!important">
                    <span class="avatar avatar-online-">
                        <img src="/images/avatars/avatar_default.png" alt="avatar">
                    </span>
                    <span class="user-name " style="color:#000000;">{{ucfirst(Auth::user()->name ?? '')}}</span>
                </a>
            </div>
        </div>
    </div>
{{-- </x-block> --}}
@push('scripts')
    <script type="text/javascript">
        $('#menu-toggle-ext, .menu-toggle-ext').on('click', function (e) {
            $('#menu-toggle').trigger('click');
        })
    </script>
@endpush
