@props(['message'=>null, 'options' => 0])

@php
	xdebug_print_function_stack($message, $options);
@endphp
