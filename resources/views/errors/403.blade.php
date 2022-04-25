{{-- @extends('layouts.app') --}}
@section('page-title', __('Forbidden'))
@section('code', '403')
@section('content')
	<div class="px-1">
		<div class="alert alert-info" style="min-height: 120px;">
			<span>Chúng tôi có thể giúp gì cho bạn ? </span>
		</div>
		{{-- <hr> --}}
		<a class="btn btn-primary round " href="/">Trở về trang chính</a>
	</div>
@endsection
<x-layout.simple />
