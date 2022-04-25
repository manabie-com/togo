<x-modal class="modal-sm" id="modal-login-email" >
  <div class="modal-content">
      <div class="modal-header">
        <button type="button" class="close" onClick="$('#modal-login-email').modal('toggle')">
          <span aria-hidden="true">×</span>
        </button>
        <h4 class="modal-title text-xs-center" >
        <i class="fa fa-envelope-o mr-1"></i> {{ __('Sign in with email') }}
        </h4>
      </div>
      <div class="modal-body">

        <div class="card">
          <div class="card-block">
            <div class="card-body">
              <livewire:auth.login />
            </div>
          </div>
        </div>
      </div>

  </div>
</x-modal>
