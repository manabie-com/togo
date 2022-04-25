@props([
'id'=>uniqid('editor_'),
'mode'=>'simple',
'name'=>'content'
])

@push('script')
<script>
  $('#{{$id}}').summernote({
    height: 280, // set editor height
    minHeight: null, // set minimum height of editor
    maxHeight: null, // set maximum height of editor

    toolbar: [
      // [groupName, [list of button]]
      @if($mode == 'editor')['insert', ['link', 'picture', 'video']],
      @endif['style', ['bold', 'italic', 'underline', 'clear']],
      ['font', ['strikethrough', 'superscript', 'subscript']],
      ['fontsize', ['fontsize']],
      ['color', ['color']],
      ['para', ['ul', 'ol', 'paragraph']],
      ['height', ['height']]
    ],
    popover: {
      image: [
        ['image', ['resizeFull', 'resizeHalf', 'resizeQuarter', 'resizeNone']],
        ['float', ['floatLeft', 'floatRight', 'floatNone']],
        ['remove', ['removeMedia']]
      ],
    },
    callbacks: {
      onImageUpload: function(files, editor, welEditable) {
        data = new FormData();
        for (var x = 0; x < files.length; x++) {
          data.append("images[]", files[x]);
        }
        data.append("_token", "{{ csrf_token() }}");
        $.ajax({
          data: data,
          type: "POST",
          url: "{{ route('editor.imgupload') }}",
          cache: false,
          contentType: false,
          processData: false,
          success: function(urls) {
            if (urls.length > 0) {
              urls.filter(el => Object.keys(el).length).map((url) => {
                $('#{{$id}}').summernote('editor.insertImage', '{{asset("media")}}/'+ url);
              })
            }
          },
          error: function(xhr, status, error) {
            console.log(status + ': ' + error);
          }
        });
      }
    }
  });
</script>
@endpush
<textarea id="{{$id}}" name="{{ $name }}">{{ $slot }}</textarea>