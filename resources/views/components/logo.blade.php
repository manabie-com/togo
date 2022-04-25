@props(['style'=>''])
<a href="{{ url('home') }}">
  <img loading="lazy" {{$attributes}} src="{{ asset("theme/images/logo/logo$style.png") }}" />
</a>
