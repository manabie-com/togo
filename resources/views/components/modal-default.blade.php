@php
  $titleModal = $titleModal ?? '';
@endphp
<style type="text/css">
  #modal-default .modal-header{border: none}
  #modal-default .modal-header .close {position: absolute;z-index:2000;top: 40px; right: 43px; font-size: 28pt}
  #modal-default .modal-content{min-width: 375px}
  /*fixed*/
  /*.modal-backdrop {position: relative!important;}*/
  @media (min-width: 1920px){}
  @media (min-width: 1366px){
    #modal-default .modal-body{max-width: 1000px;}
  }
  @media (min-width: 800px){}
  @media only screen and (max-width: 450px) {#modal-default .modal-body{}}  
  @media (min-width: 576px){#modal-default .modal-dialog {max-width: fit-content!important;margin: 1.75rem auto;}}
</style>
<div class="utility">
  <div class="modal fade text-xs-left" id="modal-default-">
    <div class="modal-dialog">
      <div class="modal-content">
        <!-- Modal Header -->
        <div class="modal-header p-2" id="utmd-header">
          {{-- <h4 class="modal-title " id="utmd-header-title" title="Thông tin chi tiết"><?=$titleModal?></h4> --}}
          <label class="px-2 text-light" id="utmd-header-title"></label>
          <button type="button" class="close mr-2" data-dismiss="modal">&times;</button>
        </div>
        <!-- Modal body -->
        <div class="modal-body">
        </div>
        <!-- Modal footer -->
        <div class="modal-footer d-none">
          <button type="button" class="btn btn-danger" data-dismiss="modal">Đóng</button>
        </div>
      </div>
    </div>
  </div>
  {{--  --}}
  <div class="modal fade bd-example-modal-lg" tabindex="-1" role="dialog" aria-labelledby="myLargeModalLabel" aria-hidden="true" id="modal-debug">
    <div class="modal-dialog modal-lg">
      <div class="modal-content">
        ...
      </div>
    </div>
  </div>
  <!-- modal-default-->
  <div class="modal fade text-xs-left" id="modal-default" tabindex="-1" role="dialog" aria-labelledby="md-default-3" aria-hidden="true">
    <div class="modal-dialog modal-xs" role="document">
    <div class="modal-content">
      <div class="modal-header">
      <button type="button" class="close" data-dismiss="modal" aria-label="Close">
        <span aria-hidden="true">&times;</span>
      </button>
      <h4 class="modal-title d-none" id="md-default-3">Basic Modal</h4>
      </div>
      <div class="modal-body">
      </div>
      <div class="modal-footer d-none">
        <button type="button" class="btn btn-danger" data-dismiss="modal">Đóng</button>
      </div>
    </div>
  </div>
</div>
<script type="text/javascript">
    $(document).ready(function() {});
        /*initial*/
        var ctrlIsPressed = false;
        /**/
        $(document).keydown(function(e){ if(e.which=="17") ctrlIsPressed = true;});
        $(document).keyup(function(e){ 
          ctrlIsPressed = false;
          if(e.which == "27"){ $('#modal-default').modal('hide');}
        });
        /**/
        /*event listener */
        let btnActive = $('.trigger-utility-modal-default, .btn-utmddf, .btn-get-df' ).not('.disabled');
        btnActive.on("touchstart", function (e) {
          let getURL = $(this).attr('href');
            if (btnActive.hasClass('touch-once')) {getXMLHttpRequest(getURL);} 
            else if (tapHandler(e)) {getXMLHttpRequest(getURL);}
            e.preventDefault();
        });
        btnActive.on('click', function(e){
          e.preventDefault();
          let getURL = $(this).attr('href');
          try {
            if (ctrlIsPressed) {window.open(getURL, '_blank');} 
            else {getXMLHttpRequest(getURL);}
          } catch(e){console.log(e);}
        });
    function getXMLHttpRequest(url) {
        let layout = ((url.search(/\?/) > 0) ? "&" : "?") + "layout=empty"; 
        let progress = swalProgress();
        var xmlhttp = new XMLHttpRequest();
        xmlhttp.onreadystatechange = function() {
                console.log(this.status);
          if (this.readyState == 4) {
            switch(this.status){
                case 200:case 201: case 201: case 301: 
                  $('#modal-default .modal-body').html(this.responseText);
                  setTimeout(() => {
                    $('#modal-default').modal('show'); 
                    Swal.close();
                    $(document).ready(function() { $('.select2').select2();})
                  }, 500);
                  break;
                case 204: case 302:
                    Swal.fire({text: 'Not found data', html: this.responseText})
                    break;
                default: break;
            }
          }
        };
        xmlhttp.open("GET", url + layout, true);
        xmlhttp.send();
    }
</script>