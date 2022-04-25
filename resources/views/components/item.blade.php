@props(['tag'=>'div','out'=>''])
@php
  global $x_item;
  $id = isset($x_item) ? 'x'.count($x_item) :'x0';
  $x_item[$id] = $out;
@endphp

<{{$tag}}
  {{ $attributes->merge(['class' => 'x-item']) }}
  xitem-id='{{ $id }}'></{{$tag}}>
