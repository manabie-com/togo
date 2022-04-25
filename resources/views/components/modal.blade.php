@props(['id'])
<x-block class='modal fade' id="{{ $id }}"
  tabindex='-1' role='dialog' >
  <div {{ $attributes->merge( ['class' => 'modal-dialog'] ) }}
     role="document">
    {{ $slot }}
  </div>
</x-block>
