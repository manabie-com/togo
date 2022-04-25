@push('css')
<style type="text/css">
    .app-view form{max-width: 400px;}
    .app-view form input.form-control{max-width: 300px;}

    .help-elf {}
    .help-elf fieldset.form-group .form-control{max-width: 300px;}
    .help-elf fieldset.form-group .curl-content.form-control{min-height: 100px;}
</style>
@endpush
@section('content')
<div class="app-view">
    <div class="card p-2 border-info round">
        <div class="card-tittle">
            <h1 class="text-info text-bold-700">Thiên Thư Các</h1>
            <p class="h3 px-2">Tầng 1</p>
        </div>
        <div class="card-body pt-1">
            <div class="card-block">
                <p>
                    Chỉ giới ::
                </p>
                <ul>
                    <li>
                        <p>
                            Nhóm "ELF" chỉ được thấu thị vào {{env('APP_NAME')}} <b>100</b> lần mỗi ngày
                        </p>
                    </li>
                    <li>
                        <p>
                            Sinh linh("animal") được vào <b>200</b> lần mỗi ngày
                        </p>
                    </li>
                    <li></li>
                    <li></li>
                    <li></li>
                    <li></li>
                    <li></li>
                    <li></li>
                    <li></li>
                    <li>
                    </li>
                </ul>
            </div>
            <div class="card-block">
                <b>Requirements:</b>
                <pre class="h5">

                Implement one single API which accepts a todo task and records it

                There is a maximum limit of N tasks per user that can be added per day.

                Different users can have different maximum daily limit.

                Write integration (functional) tests

                Write unit tests

                Choose a suitable architecture to make your code simple, organizable, and maintainable

                Write a concise README

                How to run your code locally?

                A sample “curl” command to call your API

                How to run your unit tests locally?

                What do you love about your solution?

                What else do you want us to know about however you do not have enough time to complete?
                                    
                </pre>
            </div>

            <div class="card-block">
                <p> Đã năm rõ đại đạo</p>
                <a href="{{route('animals.generic')}}" class="btn btn-warning">Chiến ngay</a>
            </div>
        </div>
    </div>
</div>
@endsection
<x-layout.greenland />
