@props(['row','action'])
<x-modal id="modal-delete-{{$row->id}}">
  <div class="modal-content">
    <x-form method='DELETE' action="{{ route($action,$row) }}">
      <div class="modal-header">
        <button type="button" class="close" data-dismiss="modal" aria-label="Close">
          <span aria-hidden="true">×</span>
        </button>
        <h4 class="modal-title" id="myModalLabel2"><i class="fa fa-road2"></i> {{ $row->name }}</h4>
      </div>
      <div class="modal-body">
        <p>
          This data will be deleted
        </p>
      </div>
      <div class="modal-footer">
        <button type="button" class="btn grey btn-outline-secondary" data-dismiss="modal">Close</button>
        <button type="submit" class="btn btn-outline-primary">Accept</button>
      </div>
    </x-form>
  </div>
</x-modal>