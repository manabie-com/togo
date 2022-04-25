@php
    $response = explode(':', request()->header('Authorization'));
    $stringToken = request()->header('cookie');
    $tokenArr = explode('=',$stringToken);
    $access_token = trim($tokenArr[2]??'');
    $host = url('/');
@endphp
@push('css')
<style type="text/css">
    .app-view form{max-width: 400px;}
    .app-view form input.form-control{max-width: 300px;}
    
    .help-elf {}
    /*.help-elf fieldset.form-group .form-control{max-width: 300px;}*/
    .help-elf fieldset.form-group .curl-content.form-control{min-height: 100px;max-width: 800px; overflow-x: hidden;}
</style>
@endpush
@section('content')
<div class="app-view">
    <div class="card p-2 border-success round">
        <div class="card-tittle">
            <h1 class="text-success text-bold-700">Sáng tạo sinh linh</h1>
        </div>
        <div class="card-body pt-3">
            <div class="card-block">
                <x-form method="POST" action="{{route('animals.store')}}">
                    <div class="row">
                        <div class="col-xs-12">
                            <fieldset class="form-group">
                                <lable>Tên gọi</lable>
                                <input type="text" class="form-control" name="">
                            </fieldset>
                        </div>
                    </div>
                    <div class="row px-2 pt-3">
                        <button class="btn btn-outline-success">Sáng tạo</button>
                    </div>
                </x-form>
            </div>
        </div>
    </div>
    <div class="help-elf">
        <div class="card px-3 py-1 border-info">
            <div class="card-tittle ">
                <h1>The simple API::CRUD</h1>
            </div>
            <div class="card-body pt-1 px-1">                
                <div class="card-block">
                    <b class="text-danger">May you need:</b>
                    <ul class="">
                        <li>
                            Token_access::
                            <div class="form-control border-warning" style="overflow-x: hidden; max-width: 600px;">
                                {{$access_token;}}
                            </div>
                        </li>
                        <li>
                            <div style="max-width:600px">
                                Host::
                                <div class="form-control">{{$host}}</div>
                            </div>
                        </li>
                    </ul>
                </div>
                <div class="card-block test-case">
                    <ul class="nav">
                        <li class="nav-item">
                            <div class="row">
                                <div class="">
                                    <div class="h3">1. Create:: "Thêm mới 1 sinh linh"</div>
                                    <fieldset class="form-group px-1">
                                        <label for="">CURL::</label>
                                        <div class="form-control curl-content">
                                            <x-content.precode>
                                                curl -X POST {{$host}}/api/elf/animals/generic \
                                                    -H "Accept: application/json" \
                                                    -H "Content-Type: application/json" \
                                                    -d "{\"admin\": \"admin@test.com\", \"password\": \"admin@123\" }"
                                                    {{-- -H "Authorization:Bearer {{$access_token}}" --}}
                                            </x-content.precode>
                                        </div>
                                    </fieldset>
                                </div>
                            </div>                    
                        </li>                        
                        <li class="nav-item">
                            <div class="row">
                                <div class="">
                                    <div class="h3">2. Update:: Cập nhật thông tin</div>
                                    <fieldset class="form-group px-1">
                                        <label for="">CURL::</label>
                                        <div class="form-control curl-content">
                                            <x-content.precode>
                                                curl -X PUT {{$host}}/api/elf/animals/2 \
                                                    -H "Accept: application/json" \
                                                    -H "Content-Type: application/json" \
                                                    -d "{\"email\": \"test2@test.com\", \"password\": \"admin@123456\" }"
                                                    {{-- -H "Authorization:Bearer {{$access_token}}"                                                    --}}
                                            </x-content.precode>
                                        </div>
                                    </fieldset>
                                </div>
                            </div>                    
                        </li>
                        <li class="nav-item">
                            <div class="row">
                                <div class="">
                                    <div class="h3">3. READ:: "danh sách sinh linh"</div>
                                    <fieldset class="form-group px-1">
                                        <label for="">CURL::</label>
                                        <div class="form-control curl-content">
                                            <x-content.precode>
                                                curl -X GET {{$host}}/api/elf/animals \
                                                    -H "Accept: application/json" \
                                                    -H "Content-Type: application/json" \
                                                    {{-- -H "Authorization:Bearer {{$access_token}}"                                                    --}}
                                            </x-content.precode>
                                        </div>
                                    </fieldset>
                                </div>
                            </div>                    
                        </li>                        
                        <li class="nav-item">
                            <div class="row">
                                <div class="">
                                    <div class="h3">4. DELETE:: "Xóa bỏ 1 sinh linh"</div>
                                    <fieldset class="form-group px-1">
                                        <label for="">CURL::</label>
                                        <div class="form-control curl-content">
                                            <x-content.precode>
                                                curl -X DELETE {{$host}}/api/elf/animals/2 \
                                                    -H "Accept: application/json" \
                                                    -H "Content-Type: application/json" \
                                                    {{-- -H "Authorization:Bearer {{$access_token}}"--}}
                                            </x-content.precode>
                                        </div>
                                    </fieldset>
                                </div>
                            </div>                    
                        </li>
                    </ul>
                </div>
            </div>
        </div>
    </div>
</div>
@endsection
<x-layout.greenland />
