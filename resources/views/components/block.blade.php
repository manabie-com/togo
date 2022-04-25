@props(['tag'=>'div'])
@if(request('_form2') || request()->ajax())
    <{{$tag}} {{ $attributes }}>
        {{ $slot }}
    </{{$tag}}>
@else
    <x-item {{ $attributes }} out='{{ $slot }}' tag='{{$tag}}'/>
@endif
