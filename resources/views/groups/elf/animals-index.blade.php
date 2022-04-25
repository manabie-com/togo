@push('css')
<style type="text/css">

</style>
@endpush
@section('content')
<div class="app-view">
    <div class="card p-2 border-info round">
        <div class="card-tittle">
            <h1 class="text-info text-bold-700">Sinh linh huyền bí</h1>
        </div>
        <div class="card-body pt-3">
            <div class="card-block">
                @if(empty($animals))
                    <div class="px-2 text-warning">
                        Có con nào đâu !
                    </div>
                    <div class="mt-2">
                        <a href="{{route('animals.create')}}" class="btn btn-success"> Sáng tạo sinh linh</a>
                    </div>
                @else
                    <div class="h4">
                        Có {{count($animals)}} chúng sinh
                    </div>
                    <div>
                        <table class="table table-striped">
                            <thead>
                                <tr><th colspan="5" class="text-xs-center h3 text-uppercase text-info text-bold-700">Sinh linh bảng</th></tr>
                                <tr>
                                    <th></th>
                                    <th>Sinh linh</th>
                                    <th>Ám hiệu</th>
                                    <th>Giới hạn tập luyện</th>
                                    <th>Thực hiện</th>
                                </tr>
                            </thead>
                            <tbody>
                                @foreach ($animals as $animal)
                                    <tr>
                                        <td></td>
                                        <td><div>{{@$animal->name}}</div></td>
                                        <td><div>{{@$animal->name}}@123</div></td>
                                        <td><div>{{$animal->user->queue_limit??0}}</div></td>
                                        <td><div>{{$animal->user->queue_today??0}}</div></td>
                                    </tr>
                                @endforeach
                            </tbody>
                        </table>
                    </div>
                @endif
            </div>
        </div>
    </div>
</div>
@endsection
<x-layout.greenland />
