<div class="modal fade" id="setting_column" tabindex="-1" role="dialog" aria-hidden="true">
    <div class="modal-dialog" role="document">
        <div class="modal-content">
            <div class="modal-header">
                <h5 class="modal-title" id="exampleModalLongTitle">{!! trans('site.custom_column_display') !!}</h5>
                <button type="button" class="close" data-dismiss="modal" aria-label="Close">
                    <span aria-hidden="true">&times;</span>
                </button>
            </div>
            <form method="POST" action="{{ url()->current() }}" id="frm_setting_column">
                @csrf
            <div class="modal-body" id="sortable">
                <div class="form-check form-check-column">
                    <label class="form-check-label">
                        <input type="checkbox" class="form-check-input" name="restore_default" id="restore_default">{!! trans('site.restore_default') !!}
                    </label>
                    <i class="fas fa-arrows-alt"></i>
                </div>
                @foreach ($colTitle as $key => $item_column)
                <div class="form-check form-check-column">
                    <label class="form-check-label">
                        <input type="checkbox" class="form-check-input" name="{{ $key }}" id="{{ $key }}" {{ ($item_column == 'true') ? 'checked' : '' }}><?php echo trans('site.'.$key) ?>
                    </label>
                    <i class="fas fa-arrows-alt"></i>
                </div>
                @endforeach
            </div>
            <div class="modal-footer">
                <button type="button" class="btn btn-secondary" data-dismiss="modal">{!! trans('site.btn_exit') !!}</button>
                <button type="button" class="btn btn-primary btn_save">{!! trans('site.btn_save') !!}</button>
            </div>
            </form>
        </div>
    </div>
</div>