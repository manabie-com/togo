`<?php
namespace App\Http\Controllers\Group\Cat;
use Auth;
use App;
use \Throwable;

use App\Http\Controllers\ExportExcelController;
use Illuminate\Http\Request;
use Illuminate\Support\Facades\{Validator, Redirect};
use Illuminate\Validation\Rule;
use Illuminate\Pagination\LengthAwarePaginator;
use Illuminate\Support\Facades\DB;
use Illuminate\Support\Str;
use Maatwebsite\Excel\Facades\Excel;
use Symfony\Component\HttpFoundation\File\UploadedFile;
use Carbon\Carbon;

use App\Models\{Cage};

class CageController extends App\Http\Controllers\Controller
{
    // use CastRequestTrait;
    
    /**
     * Display a listing of the resource.
     *
     * @return \Illuminate\Http\Response
     */
    public function index(Request $request, Cage $Cage)
    {
        try {
            
        } catch (Throwable $th) {
            logger($th);
        }
    }

    /**
     * Show the form for creating a new resource.
     *
     * @return \Illuminate\Http\Response
     */
    public function create(Request $request, Cage $Cage)
    {
        try {
            return view('groups.customer.Cages-create')->with(compact('Cage'));
        } catch (Throwable $th) {logger($th);}
    }
    /**
     * Store a newly created resource in storage.
     *
     * @param  \Illuminate\Http\Request  $request
     * @return \Illuminate\Http\Response
     */
    public function store(Request $request, Cage $Cage)
    {
        try {

           
        } catch (Throwable $th) {
            logger($th);
            return back()->withInput()->with('failed', trans('site.created_failed'));
        }
    }

    /**
     * Display the specified resource.
     *
     * @param  \App\Models\Cage  $Cage
     * @return \Illuminate\Http\Response
     */
    public function show(Cage $Cage)
    {
        try {
            dd('ok');
            return view('Cages.show')->with(compact('Cage'));
        } catch (Throwable $th) {
            logger($th);
        }
    }

    /**
     * Show the form for editing the specified resource.
     *
     * @param  \App\Models\Cage  $Cage
     * @return \Illuminate\Http\Response
     */
    public function edit(Cage $Cage)
    {
        try {
            // dd(Auth::user());
            return view('groups.customer.Cages-edit')->with(compact('Cage'));
        } catch (Throwable $th) {
            logger($th);
        }
    }
    /**
     * Update the specified resource in storage.
     *
     * @param  \Illuminate\Http\Request  $request
     * @param  \App\Models\Cage  $Cage
     * @return \Illuminate\Http\Response
     */
    public function update(Request $request, Cage $Cage)
    {
        try {
            $input = $request->all();
            return back()->with('success', 'Đã gửi cập nhật');
            $validator = Validator::make($input,
                [
                    'code' => ['required','alpha_dash','max:16',Rule::unique($Cage->getTable())->ignore($Cage->id)],
                    'name' => 'required|max:100',
                ], 
                [
                    'required'  => 'Trường :attribute là bắt buộc',
                    'unique'    => 'Trường :attribute đã tồn tại',
                    'max'       => 'Độ dài tối đa là :max',
                    'alpha_dash'=> 'Ký tự phải là chữ/số/"-"/"_"'
                ]
            );
            if ($validator->fails()) return back()->withErrors($validator)->withInput();
            /**/
            if ($Cage->fill($input)->save()) return back()->with('success', trans('site.updated_success')); 
            throw new Throwable("Can't save");
        } catch (Throwable $th) {
            logger($th);
            return back()->withInput()->with('failed', trans('site.created_failed'));
        }
    }

    /**
     * Remove the specified resource from storage.
     *
     * @param  \App\Models\Cage  $Cage
     * @return \Illuminate\Http\Response
     */
    public function destroy(Cage $Cage)
    {
        try {
            if ($Cage->delete()) return Response(__('titles.deleted_success'), 301);
            return Response(__('titles.deleted_success'), 302);
        } catch (Throwable $th) {logger($th);}
    }
}
