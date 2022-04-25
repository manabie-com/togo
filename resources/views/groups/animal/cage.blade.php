
@section('template_title', trans('usersmanagement.showing-user', ['name' => @$Animal->name]))
<style type="text/css">
</style>
@section('navbar-more')
@endsection
@section('content')
<div class="app-view">
    <div class="form-control">
        <div>
            Đây là nhà của <span class="text-uppercase">{{@$user->name}}</span>
        </div>
        <hr>
        <div class="card p-3">
            <div class="pb-2 text-bold-700">The simple APIs by CRUD</div>
            <div class="card-body">
                <ul class="nav">
                    <li class="nav-item">
                        1. Add new "Cat"
                        <p class="px-1">
                            more info  
                        </p>
                        <div class="row px-1">
                            <div class="col-xs-12 col-md-4">
                                <span>CURL </span>
                                <span>
                                    <textarea class="form-control" style="min-height: 110px"></textarea>
                                </span>
                            </div>
                            <div class="col-xs-12 col-md-4">
                                <span>By </span>
                                <div class="pb-1">
                                    <button class="btn btn-success">Passed</button>
                                </div>
                                <div class="pb-1">
                                    <button class="btn btn-danger">Failed</button>
                                </div>
                            </div>
                        </div>
                    </li>
                    <li class="nav-item">
                        2. Update :: "There is a maximum limit of N tasks per user that can be added per day"
                        <p class="px-1">
                            more info  
                        </p>
                        <div class="row px-1">
                            <div class="col-xs-12 col-md-4">
                                <span>CURL </span>
                                <span>
                                    <textarea class="form-control" style="min-height: 110px"></textarea>
                                </span>
                            </div>
                            <div class="col-xs-12 col-md-2">
                                <span>By </span>
                                <div class="pb-1">
                                    <button class="btn btn-success">Passed</button>
                                </div>
                                <div class="pb-1">
                                    <button class="btn btn-danger">Failed</button>
                                </div>
                            </div>

                            <div class="col-xs-12 col-md-4">
                                <label>Maximum limit of N tasks:</label>
                                <fieldset class="form-group">
                                    <input type="number" class="form-control" name="limit-qps">
                                    {{-- <input type="number" class="form-control" name="limit-qps"> --}}
                                </fieldset>
                            </div>
                        </div>
                    </li>

                    <li class="nav-item">
                        
                    </li>
                    <li class="nav-item">
                        
                    </li>
                    <li class="nav-item">
                        
                    </li>
                    <li class="nav-item">
                        
                    </li>
                    <li class="nav-item">
                        
                    </li>
                </ul>
            </div>
        </div>
            <div class="form-control border-warning">
                <div class="text-bold-700 pb-2">Result</div>
                <div class="content-result">
                    
                </div>
            </div>
  </div>
</div>
@endsection
<x-layout.greenland />
