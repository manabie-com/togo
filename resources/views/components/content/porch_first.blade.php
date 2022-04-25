@php
    $linkData = [];
    $currentURL = url()->current();
@endphp
@role('elf')
    @php
        $linkData = [
            'icon' => '',
            'link_home' => url('/elf/home'),
            'link_practice' => url('/elf/practice'),
            'link_discovery' => url('/elf/discovery'),
        ];
    @endphp
@endrole
@role('animal')
    @php
        $linkData = [
            'icon' => '',
            'link_home' => url('/animal/home'),
            'link_practice' => url('/animal/practice'),
            'link_discovery' => url('/animal/discovery'),
        ];
    @endphp
@endrole

<div class="row mx-auto" style="max-width: 600px;">
    <div class="col-xs-12 col-md-4">
        <div class="card round card-success p-2">
            <div class="card-title text-xs-center">
                <a href="{{url($linkData['link_home']??'')}}" style="color:#fff">Về chuồng</a>
            </div>
            <div class="card-body">
                
            </div>
        </div>
    </div>
    <div class="col-xs-12 col-md-4">
        <div class="card round card-warning p-2">
            <div class="card-title text-xs-center">
                <a href="{{url($linkData['link_practice']??'')}}" style="color:#fff">Tập luyện</a>
            </div>
            <div class="card-body"></div>
        </div>
    </div>
    <div class="col-xs-12 col-md-4">
        <div class="card round card-danger p-2">
            <div class="card-title text-xs-center">
                <a href="{{url($linkData['link_discovery']??'')}}" style="color:#fff">Khám phá</a>
            </div>
            <div class="card-body">
            </div>
        </div>
    </div>
</div>
