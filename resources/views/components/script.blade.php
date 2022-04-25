@php
global $x_item;
@endphp
{{-- <script src="{{asset('theme/js/vendors.min.js')}}"></script> --}}
{{-- <script src="{{asset('theme/js/theme.js')}}"></script> --}}
<script>
  @if($x_item)
  $.each(@json($x_item), function(key, value) {
    $('[xitem-id=' + key + ']').html(value);
  });
  @endif


  $('.lazy').each(function(i) {
    var $this = $('.lazy').eq(i);
    $this.html($this.data('value'));

    var bg = $this.data('bg-img');
    if (typeof bg != 'undefined') {
      $this.css('background-image', 'url(' + bg + ')');
    }

  });

  @stack('x-script')
</script>
