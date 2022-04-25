@extends('components.layout.greenland')
@section('content')
    <div class=" p-3 text-info h1">
        Chào mừng <span class="text-uppercase text-bold-700">{{ Auth::user()->name}}</span>, đến với thế giới hoang dã
    </div>
@push('css')
    <style>
        .elf_ctrl .card{min-height: 200px;min-width: 200px; max-width: 300px;}
    </style>

@endpush

<div class="elf_ctrl">
    @role('elf')
        <div>
            <div class="row mx-auto" style="max-width: 1200px;">
                <div class="col-xs-12 col-md-4">
                    <a href="{{route('animals.book')}}" style="color:#fff">
                        <div class="card round card-info p-2">
                            <div class="card-title text-xs-center">Thiên thư
                            </div>
                            <div class="card-body">
                                <p>
                                    Chứa đựng thông tin huyền bí
                                </p>
                                <p>
                                    Bí kíp vượt ải
                                </p>
                            </div>
                        </div>
                    </a>
                </div>
                <div class="col-xs-12 col-md-4">
                    <a href="{{route('animals.generic')}}" style="color:#fff">
                        <div class="card round card-success p-2">
                            <div class="card-title text-xs-center">
                                Khai sáng
                            </div>
                            <div class="card-body">
                                <p>
                                    Tạo lập sinh mệnh
                                </p>
                                <p class="h4">
                                    Test case
                                </p>
                            </div>
                        </div>
                    </a>
                </div>
                <div class="col-xs-12 col-md-4">
                    <a href="{{url('elf/animals')}}" style="color:#fff">
                        <div class="card round card-danger p-2">
                            <div class="card-title text-xs-center">Sinh linh bảng
                            </div>
                            <div class="card-body">
                                <p>
                                    Thần quang mục thị chúng sinh
                                </p>
                            </div>
                        </div>
                    </a>
                </div>
            </div>
        </div>
    @endrole
    @role('animal')
        <div>
            <div class="row mx-auto" style="max-width: 1200px;">
                <div class="col-xs-12 col-md-4">
                    <a href="{{route('animal.book')}}" style="color:#fff">
                        <div class="card round card-success p-2">
                            <div class="card-title text-xs-center">Cẩm nan</div>
                            <div class="card-body">
                                <p>
                                    Những thông tin thất quý giá để có thể sinh tồn.
                                </p>
                            </div>
                        </div>
                    </a>
                </div>
                <div class="col-xs-12 col-md-4">
                    <a href="{{route('animal.practice')}}" style="color:#fff">
                        <div class="card round card-warning p-2">
                            <div class="card-title text-xs-center">Luyện tập</div>
                            <div class="card-body">
                                <p>
                                    Bắt đầu hành trình
                                </p>
                            </div>
                        </div>
                    </a>
                </div>
                <div class="col-xs-12 col-md-4">
                    <a href="#" class="disabled" style="filter:">
                        <div class="card round card-dark p-2">
                            <div class="card-title text-xs-center">Mộng cảnh</div>
                            <div class="card-body">
                                <p>
                                    Mộng cảnh huyền bí
                                </p>
                                <p>Chưa mở</p>
                            </div>
                        </div>
                    </a>
                </div>
            </div>
        </div>
    @endrole
</div>

@endsection()
