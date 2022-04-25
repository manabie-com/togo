@push('css')
<style type="text/css">
    .app-view form{max-width: 400px;}
    .app-view form input.form-control{max-width: 300px;}
</style>
@endpush
@section('content')
<div class="app-view">
    <div class="card p-2 border-success round">
        <div class="card-tittle">
            <h1 class="text-success text-bold-700">Đấng sáng tạo</h1>
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
                        {{-- <div class="col-xs-12">
                            <fieldset class="form-group">
                                <lable>A</lable>
                                <input type="text" class="form-control" name="">
                            </fieldset>
                        </div>
                        <div class="col-xs-12">
                            <fieldset class="form-group">
                                <lable>A</lable>
                                <input type="text" class="form-control" name="">
                            </fieldset>
                        </div> --}}
                    </div>
                    <div class="row px-2 pt-3">
                        <button class="btn btn-outline-success">Sáng tạo</button>
                    </div>
                </x-form>
            </div>
        </div>
    </div>
</div>
@endsection
<x-layout.greenland />
