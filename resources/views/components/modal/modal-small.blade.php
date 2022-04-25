@props(['id', 'title'=>'', 'lang'=>[], 'config'=>['footer_hide'=>true, 'header_hide'=>true]])
@push('outer')
  <x-modal id="{{$id}}" class="modal-dialog modal-md" >
      <div class="modal-content">
        <div class="modal-header {{empty($config['header_hide']) ?:'d-sm-none d-none'}}">
          <button type="button" class="close" data-dismiss="modal" aria-label="Close">
            <span aria-hidden="true">&times;</span>
          </button>
          <h4 class="modal-title" id="">{{$title ?? ''}}</h4>
          </div>
          <div class="modal-body">
            {{$slot}}
          </div>
          <div class="modal-footer {{empty($config['footer_hide']) ?:'d-sm-none d-none'}}">
          <button type="button" class="btn grey btn-outline-secondary" data-dismiss="modal">{{$lang['close']??'Đóng'}}</button>
          <button type="button" id="{{$lang['btn_save_id'] ?? ''}}" class="{{$lang['btn_save_class'] ?? 'btn btn-outline-primary'}}">{{$lang['btn_save_label'] ?? 'Lưu'}}</button>
        </div>
      </div>
  </x-modal>
@endpush
