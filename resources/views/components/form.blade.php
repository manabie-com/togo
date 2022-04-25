@props(['method'=>'POST'])
<form {{ $attributes }} method="{{ ($method == 'GET') ? $method : 'POST' }}" >
    {{ $slot }}
    <input type="text" name="per-page" hidden value="{{request('per-page')}}">
    <input type="text" name="page" hidden value="{{request('page')}}">
    @csrf
    @method($method)
</form>
