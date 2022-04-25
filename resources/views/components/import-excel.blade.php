@inject('SupportComponent', 'App\Models\SupportComponent')
@extends('layouts_contents.content-datatable')
@section('content-layout')
  @php
    /**
     * @author toannguyen.dev
     * @todo
     */
    /*intial*/
    $importLink = $importLink ?? '';
    $redirectLink = array_filter($redirectLink ?? ['','']);
    $titleHeading = $titleHeading ?? __('site.import-excel');
    $iconAdd = isset($iconAdd) ? $iconAdd : 'fa fa-upload';
    $iconEdit = isset($iconEdit) ? $iconEdit : 'fas fa-upload';
    $iconRemove = isset($iconRemove) ? $iconRemove : 'fas fa-upload text-red';
    $backLink = $backLink ?? '';
    $noteContent = array_filter($noteContent ?? []);
    $noteContent[] = 'Phải được đánh số thứ tự các dòng thông tin';
    $noteContent[] = 'Không được thay đổi dòng tiêu đề mặc định(dòng 1- màu cam)';
    $noteContent[] = 'Không được thay đổi "SheetNames"';
    $noteContent[] = 'Được đổi tên tệp tin';
  @endphp
  <style type="text/css">
    .file-upload-wrapper .card.card-body input {
        position: absolute;
        top: 0;
        right: 0;
        bottom: 0;
        left: 0;
        height: 100%;
        width: 100%;
        opacity: 0;
        cursor: pointer;
        z-index: 5;
    }
    .file-upload {
      display: -webkit-box;
      display: -webkit-flex;
      display: -ms-flexbox;
      display: flex;
      -webkit-box-pack: center;
      -webkit-justify-content: center;
      -ms-flex-pack: center;
      justify-content: center;
      -webkit-box-align: center;
      -webkit-align-items: center;
      -ms-flex-align: center;
      align-items: center;
      position: relative;
      cursor: pointer;
      overflow: hidden;
      width: 100%;
      max-width: 100%;
      height: 200px;
      padding: 5px 10px;
      font-size: 1rem;
      text-align: center;
      color: #ccc;
    }
    .file-upload-wrapper .card.card-body input {
      position: absolute;
      top: 0;
      right: 0;
      bottom: 0;
      left: 0;
      height: 100%;
      width: 100%;
      opacity: 0;
      cursor: pointer;
      z-index: 5;
    }
    .file-upload-wrapper .card.card-body .file-upload-message {
      position: relative;
      top: 50px;
      -webkit-transform: translateY(-50%);
      -ms-transform: translateY(-50%);
      transform: translateY(-50%);
    }
    .file-upload-wrapper .card.card-body .file-upload-message p.file-upload-error {
      color: #f34141;
      font-weight: 700;
      display: none;
    }
    .file-upload-wrapper .card.card-body .file-upload-message p {
      margin: 5px 0 0;
    }
    .file-upload .mask.rgba-stylish-slight {
      opacity: 0;
      -webkit-transition: all .15s linear;
      -o-transition: all .15s linear;
      transition: all .15s linear;
    }
    .view .mask {
      position: absolute;
      top: 0;
      right: 0;
      bottom: 0;
      left: 0;
      width: 100%;
      height: 100%;
      overflow: hidden;
      background-attachment: fixed;
    }
    .file-upload-wrapper .card.card-body .btn.btn-sm.btn-danger {
      display: none;
      position: absolute;
      opacity: 0;
      z-index: 7;
      top: 10px;
      right: 10px;
      -webkit-transition: all .15s linear;
      -o-transition: all .15s linear;
      transition: all .15s linear;
    }
    .file-upload-wrapper .card.card-body .file-upload-errors-container {
      position: absolute;
      left: 0;
      top: 0;
      right: 0;
      bottom: 0;
      z-index: 3;
      background: rgba(243, 65, 65, .8);
      text-align: left;
      visibility: hidden;
      opacity: 0;
      -webkit-transition: visibility 0s linear .15s, opacity .15s linear;
      -o-transition: visibility 0s linear .15s, opacity .15s linear;
      transition: visibility 0s linear .15s, opacity .15s linear;
    }
    .file-upload-wrapper .card.card-body .file-upload-preview {
      display: none;
      position: absolute;
      z-index: 1;
      background-color: #fff;
      padding: 5px;
      width: 100%;
      height: 100%;
      top: 0;
      right: 0;
      bottom: 0;
      left: 0;
      overflow: hidden;
      text-align: center;
    }
    /*upload*/

    /*.upload-dialog{max-width: 600px!important;}*/
    .upload-header{ position: absolute; z-index: 9999; }
    .upload-header div button:nth-last{float: right;}
    .upload-body {
      position: relative;
      flex: 1 1 auto;
      padding: 0rem; 
    }
    .btn i{display: inline-grid;}
    .b {font-weight: 600}
    .result-import .swal2-popup {max-width: 800px!important; min-width: 375px;min-height: 300px;max-height: 100%;overflow-y: auto;}
    #result-import tr td{min-width: 50px}
    .import-json ul {margin: 0px auto!important}
  </style>
  <template id="my-template">
    <swal-title>
    </swal-title>
    <swal-icon type="" color=""></swal-icon>
    @if(!empty($redirectLink))
    <swal-button type="cancel">
      <a class="btn" href="{{$redirectLink[0] ?? ''}}" target="_self"> {{$redirectLink[1] ?? 'Chuyển hướng'}}</a>
    </swal-button>
    @endif
    <swal-button type="confirm">
      Đồng ý
    </swal-button>
    <swal-button type="deny">
      Đóng
    </swal-button>
    <swal-param name="allowEscapeKey" value="false" />
    <swal-param name="customClass" value='{ "popup": "my-popup" }' />
  </template>
  <div class="import-json">
    <div class="app-main__inner">
      <div class="app-page-title">
        <div class="page-title-wrapper">
          <div class="page-title-heading">
            <div class="page-title-icon bg-success">
              <i class="fa fa-file-excel "></i>
            </div>
            <div class="text-uppercase">{{$titleHeading}}</div>
          </div>
          <div class="page-title-actions">
            <a href="{{$backLink}}" class="btn-back d-flex text-uppercase">
              <img src="{{ asset('images/icon_back.png') }}" width="25px" class="mr-2">
              {!! trans('titles.back-page') !!}
            </a>
          </div>
        </div>
      </div>
      @if (session('failed'))
        <div class="alert alert-danger">
          {!! session('failed') !!}
        </div>
      @endif
      @if (session('success'))
        <div class="alert alert-success">
          {!! session('success') !!}
        </div>
      @endif
      <div class="container-fluid">
        <div class="utility">
          <div class="upload" id="upload-file">
            <div class="upload-header">
              <div class="d-flex justify-content-center">
                
                <!-- <button id="btn-import-edit" api_method="PUT" class="btn-import btn btn-yellow mx-1" type="button" title="edit data"><i class="<?= $iconEdit;?>"></i></button> -->
                <!-- <button id="btn-import-remove" api_method="DELETE" class="btn-import btn btn-dark ml-5 float-right" type="button" title="delete"><i class="<?= $iconRemove; ?>"></i></button> -->
              </div>
            </div>
            <div class="upload-body">
              @if(!empty($noteContent))
                <div class="border rounded bg-light mb-3">
                  <div class="px-2">{{__('site.note')}}</div>
                  @foreach($noteContent as $content)
                    <div class="row">
                      <div class="col">
                        <span class="btn text-danger " style="font-size: 14pt; line-height: 0px">*</span>
                        {!!strip_tags($content,'<a>')!!}</div>
                    </div>
                  @endforeach
                </div>
              @endif
              <form action="{{$importLink}}" class="form-upload my-0" name="form-upload" enctype="multipart/form-data" method="POST" max="1">
               @csrf
                <div class="file-upload-wrapper">
                  <div class="card card-body view file-upload mb-0" style="height: 15vh; min-height: 120px">
                    <div class="card-text file-upload-message">
                      <i class="fas fa-cloud-upload-alt fa-3x"></i>
                      <p>Kéo và thả tệp</p>
                      <p class="file-upload-error">Ooops, something wrong happended.</p>
                    </div>
                    <div class="mask rgba-stylish-slight"></div>
                    <div class="file-upload-errors-container"><ul></ul></div>
                    <input type="file" id="input-excel-to-json" class="file_upload" name="" data-height="500" >
                    <button type="button" class="btn btn-sm btn-danger waves-effect waves-light">Remove<i class="far fa-trash-alt ml-1"></i></button>
                    <div class="file-upload-preview">
                      <span class="file-upload-render"></span>
                      <div class="file-upload-infos">
                        <div class="file-upload-infos-inner">
                          <p class="file-upload-filename"><span class="file-upload-filename-inner"></span></p>
                          <p class="file-upload-infos-message">Drag and drop or click to replace</p>
                        </div>
                      </div>
                    </div>
                  </div>
                </div>
                <hr class="border-0">
                <div class="btn-group w-100 d-none">
                  <textarea class="form-control" rows=5 cols=120 id="xlx_json" name="data-json" style="" placeholder="data json"></textarea>
                </div>
                <div class="row">                  
                  <div class="col rounded bg-light mb-3 mx-3 py-2">
                    <ul class="my-0"></ul>
                  </div>
                </div>
                <div class="row">
                  <div class="col d-flex justify-content-center">
                    <div class="form-group mb-0">
                      <div class="input-group">
                        <input type="text" class="form-control d-none" id="sheetname" name="sheetname" style="max-width:30vh" placeholder="SheetNames" required />
                        <div class="input-group">
                          <div class="input-group-text btn-import btn bg-success " title="Import data">
                            <i class="<?= $iconAdd; ?> px-2 d-none"></i>Import</div>
                        </div>
                      </div>
                    </div>
                  </div>
                </div>
              </form>
            </div>
            <!-- upload footer -->
            <div class="upload-footer">
              
                <div class="row px-2 my-2">
                  <div class="col" style="width: 100vh;overflow-x: auto;">
                    <table id="excel-datatable" class="excel-datatable datatable table table-striped display nowrap" style="width:100%"></table>
                  </div>
                </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
<script src="https://ajax.googleapis.com/ajax/libs/jquery/2.1.1/jquery.min.js"></script>
<script src="https://cdnjs.cloudflare.com/ajax/libs/xlsx/0.8.0/jszip.js"></script>
<!-- <script src="https://cdnjs.cloudflare.com/ajax/libs/xlsx/0.8.0/xlsx.js"></script> -->
<!-- <script src="https://cdnjs.cloudflare.com/ajax/libs/xlsx/0.16.9/xlsx.min.js"></script> -->
<script src="https://cdnjs.cloudflare.com/ajax/libs/xlsx/0.12.9/jszip.js"></script>
<script src="https://cdnjs.cloudflare.com/ajax/libs/xlsx/0.12.9/xlsx.min.js"></script>
<script>
    var ExcelToJSON = function() {
      this.parseExcel = function(file) {
        var reader = new FileReader();
        reader.onload = function(e) {
          var data = e.target.result;
          var workbook = XLSX.read(data, {
            type: 'binary'
          });
            let sheetName = workbook.SheetNames[0];
            var xl_row_object = XLSX.utils.sheet_to_row_object_array(workbook.Sheets[sheetName]);
            var xl_row_object_empty = XLSX.utils.sheet_to_row_object_array(workbook.Sheets[sheetName],{blankrows: false, defval: ''});
            var json_object = JSON.stringify(xl_row_object);
            let columns_object = xl_row_object_empty[0];
            let columns = [];
            let dataSet = [];
            for (const property in columns_object) {
              let title = {title: property.toUpperCase()};
              columns.push(title);
            }
            xl_row_object_empty.forEach(function (ele, rowIndex) {
              let eleArr = [];
              for (const property in ele) {
                value = ele[property]
                eleArr.push(value != '' ? value : '-');
              }
              dataSet.push(eleArr);
            })
            jQuery('#xlx_json' ).val( json_object );
            jQuery('#sheetname').val(sheetName);
            // drawDatatGird(dataSet, columns);
        };
        reader.onerror = function(ex) {
          console.log(ex);
        };
        reader.readAsBinaryString(file);
      };
    };

    function handleFileSelect(evt) {
        jQuery('#xlx_json' ).val('');
        jQuery('#sheetname').val('');
        var files = evt.target.files; // FileList object
        var xl2json = new ExcelToJSON();
        xl2json.parseExcel(files[0]);
        return true;
    } 
</script>
<script>
</script>
{{-- customize --}}
<script type="text/javascript">
    $(document).ready(function() {
      var datatable = null;
      var action = null;
      $('#input-excel-to-json').on('change', function(evt){
        let input_excel = this;
        let max = parseInt($(input_excel).attr('max'));
        let maxSize = $(input_excel).attr('max-size') != undefined ? $(input_excel).attr('max-size') : 2;
            maxSize = Math.max(3, maxSize) * 1024;
        let fileList = this.files;
        let aceptFile = ['xls', 'xlsx'];
        /*check only one file*/
        let ext = getExt(fileList[0].name);
        if (!aceptFile.includes(ext)) {
          toastFlash({icon:'warning',title:"Định dạng file ["+ext+"] không hỗ trợ"});
          return false;
        }
        if (byteToKB(fileList[0].size) >= maxSize) {
          toastFlash({icon:'warning',title:"Dung lượng tệp quá lớn ("+byteToKB(fileList[0].size) +'/'+ maxSize+") KB!"});
          return false;
        }
        if (this.files.length > max) {
          toastFlash({icon:'warning',title:"Chọn quá số tệp !("+max+")"});
          return false;
        }
        handleFileSelect(evt);
        $('.upload ul').html('');
        $.each(fileList, function(i,file){
          $('.upload ul').append($("<li>").text(file.name +' Size: '+ byteToKB(file.size) + ' KB'));
        });
      });
      $('form').on('submit', function (e) {e.preventDefault(); $('.btn-import').trigger('click')})
      $('#btn-upload-file, .btn-import ').on('click',function(e){
        btn_import = this;
        file_upload = $('input.file_upload');
        var confirmed = false;
        if ($('#xlx_json').val() == '') {
          $('#xlx_json').select().focus();
          toastFlash({icon:'warning',title:"Không có dữ liệu !"}); 
          return false;
        }
        if ($('#sheetname').val() == '') {$('#sheetname').select().focus(); return false;}
        if ($(file_upload).val() != '' || $('#xlx_json').val() != '') {
          swalConfirm("{{trans('site.are_you_sure')}}",'', function(r){
            if (r) {
              let form_upload = $('.form-upload');
              let url = $(form_upload).attr('action');
              let formData = $('.form-upload').serialize();
              swalProgress(90000);
              var xmlhttp = new XMLHttpRequest();
              xmlhttp.onreadystatechange = function() {
              if (this.readyState == 4) {
                  console.log(this.status);
                  switch(this.status){
                    case 201:
                      object = JSON.parse(this.responseText);
                      var txt_res = '<div class="text-left mx-3">';
                      $(object.import).each(function(index1, ele1){

                        txt_res += '<table id="result-import" class="table table-striped">';
                          $.each(ele1, function(line, ele2){
                            var imported = '';
                            var cla_failed = '';
                            if (ele2.imported != undefined && ele2.imported === 'failed') {
                              imported = 'Chưa thêm';
                              cla_failed = 'bg-danger';
                              delete ele2.imported;
                            }
                          txt_res += '<tr class="'+cla_failed+'">';
                              txt_res += '<td class="" style="white-space: nowrap;">'+ line + '</td>';
                              txt_res += '<td class="" style="white-space: nowrap;"><ul class="my-0">';
                                $.each(ele2, function(key, value){
                                  txt_res += '<li><b>' + key + "</b> : " + value + '</li>';
                                });
                              txt_res += '<ul></td>';
                              txt_res += '<td class="" style="white-space: nowrap;">'+ imported + '</td>';
                          txt_res += '</tr>';
                          });
                        txt_res += '</table>';
                        // txt_res += '<hr>Total ' + Object.entries(ele1).length;
                        // txt_res += '</div>';
                      });
                       txt_res += '<hr><h4>' + object.total+'</h4>';
                      Swal.fire({
                        template:'#my-template',
                        title:'Kết quả thực hiện:',
                        html: txt_res,
                        icon:'success',
                        confirmButtonColor: '#2cc194',
                        cancelButtonColor:'#2cc194',
                        showConfirmButton:false,
                        width:'100%',
                        height:'1550px',
                        customClass:{container:"result-import"}
                      });

                      // $('#modal-default').find('.modal-body').html(txt_res);
                      // $('#modal-default').modal();
                      // Swal.close();
                      break;
                    case 202:
                      object = JSON.parse(this.responseText);
                      // Swal.fire({
                      //   icon:'danger',
                      //   title:"Thao tác không thành công !",
                      //   html:JSON.stringify(object.data) ,
                      //   icon:'success'
                      // })
                      swalAlert('Không thành công !','error',{
                        html:JSON.stringify(object.data)
                      });
                      break;
                    default:
                      // showResponse(this);

                      swalAlert('Không thành công !','error',{
                        html:this.responseText,
                        confirmButtonText: 'Đóng',
                      });
                      //  Swal.fire({
                      //   icon:'warning',
                      //   title:"Thao tác không thành công !",
                      //   html:(this.responseText) ,
                      //   confirmButtonColor: '#2cc194',
                      //   confirmButtonText: 'Đóng',
                      // })
                      break;
                  }
                }
              };
              xmlhttp.open("POST", url, true);
              xmlhttp.setRequestHeader("Content-Type", "application/x-www-form-urlencoded");
              xmlhttp.setRequestHeader('X-CSRF-TOKEN', $('meta[name="csrf-token"]').attr('content'));
              xmlhttp.send(formData);
            }
          },{});
        } else {
          $('.upload ul').html('');
          $('.upload ul').append($("<li>").append($('<lable class="text-danger">').text('Choose a file to upload')));
        }
      });
    });
    function drawDatatGird(dataSet, columns) {
      let config = {
          data: dataSet,
          columns: columns,
          sDom: 'tlipb',
          searching: false,
          ordering: false,
          info: true,
          "autoWidth": true,
          "scrollCollapse": true,
          "scrollY": "100%",
          scrollX: "700px",
          // responsive: true,
          stateSave: true,
          paging: true,
          "pagingType": 'full',
          "lengthChange": true,
          "lengthMenu": [ [10, 20, 50, 100, -1], [10, 20, 50, 100, "All"] ],
          "pageLength": 10,
          orderCellsTop: true,
          // fixedHeader: true,
          // fixedFooter: true,
          // "select": {"style": "os"},
          // "order": [[ 1, 'desc' ], [0, 'asc' ]],
          oLanguage: {
            sLengthMenu: "_MENU_",
            sInfo: "_START_ - _END_ of _TOTAL_ ",
            oPaginate: {
              "sFirst": '<i class="fas fa-step-backward"></i>',"sLast": '<i class="fas fa-step-forward"></i>',
              "sNext": '<i class="fas fa-chevron-right"></i>',
              "sPrevious": '<i class="fas fa-chevron-left"></i>'},
            sSearch: "",
          },
          language:{
            infoFiltered:" / _MAX_ ",
            select: { rows: '<span class="px-1"> Đã chọn %d </span>'},
          },
          colReorder: true,
        }
        try{
          if ( $.fn.dataTable.isDataTable( '#excel-datatable' ) ) {
              datatable.destroy();
              $('#excel-datatable').DataTable(config).draw();
          } else {
            datatable = $('#excel-datatable').DataTable(config).draw();
          }
        } catch(e){
          console.log(e);
        }
       
    }
    function showResponse(reponse){
      $('#modal-default').find('.modal-body').html(reponse.responseText);
      $('#modal-default').modal();
      Swal.close();
    }
    function getExt(filename)
    {
        var ext = filename.split('.').pop();
        if(ext == filename) return "";
        return ext;
    }
    function byteToMB(byte){
      return Math.round(byte/ (1024 *1024));
    }
    function byteToKB(byte){return Math.round(byte/(1024));}
</script>
@endsection
