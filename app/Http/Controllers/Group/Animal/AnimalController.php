<?php
namespace App\Http\Controllers\Group\Animal;
use Auth;
use App;
use \Throwable;

use Illuminate\Http\Request;
use Illuminate\Support\Facades\{Validator, Redirect};
use Illuminate\Validation\Rule;
use Illuminate\Pagination\LengthAwarePaginator;
use Illuminate\Support\Facades\DB;
use Illuminate\Support\Str;
use Carbon\Carbon;

use App\Models\{Animal};

class AnimalController extends App\Http\Controllers\Controller
{
    // use CastRequestTrait;
    
    /**
     * Display a listing of the resource.
     *
     * @return \Illuminate\Http\Response
     */
    public function index(Request $request, Animal $animal)
    {
        try {
            $animals = Animal::all();
            return view('groups.animal.animals-index')->with(compact('animals'));
        } catch (Throwable $th) {
            logger($th);
        }
    }

    /**
     * Show the form for creating a new resource.
     *
     * @return \Illuminate\Http\Response
     */
    public function create(Request $request, Animal $animal)
    {
        try {
            return view('groups.animal.animals-create')->with(compact('animal'));
        } catch (Throwable $th) {logger($th);}
    }
    /**
     * Store a newly created resource in storage.
     *
     * @param  \Illuminate\Http\Request  $request
     * @return \Illuminate\Http\Response
     */
    public function store(Request $request, Animal $animal)
    {
        try {
            dd($request->toArray());
           
        } catch (Throwable $th) {
            logger($th);
            return back()->withInput()->with('failed', trans('site.created_failed'));
        }
    }

    /**
     * Display the specified resource.
     *
     * @param  \App\Models\Animal  $animal
     * @return \Illuminate\Http\Response
     */
    public function show(Animal $animal, Request $request)
    {
        try {
            $user = Auth::user();
            return view('groups.animal.cage')->with(compact('animal', 'user'));
        } catch (Throwable $th) {
            logger($th);
        }
    }

    /**
     * Show the form for editing the specified resource.
     *
     * @param  \App\Models\Animal  $animal
     * @return \Illuminate\Http\Response
     */
    public function edit(Animal $animal)
    {
        try {
            // dd(Auth::user());
            return view('groups.animal.animals-edit')->with(compact('Animal'));
        } catch (Throwable $th) {
            logger($th);
        }
    }
    /**
     * Update the specified resource in storage.
     *
     * @param  \Illuminate\Http\Request  $request
     * @param  \App\Models\Animal  $animal
     * @return \Illuminate\Http\Response
     */
    public function update(Request $request, Animal $animal)
    {
        try {
            $validator = Validator::make($input,
                [
                    'code' => ['required','alpha_dash','max:16',Rule::unique($animal->getTable())->ignore($animal->id)],
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
            if ($animal->fill($input)->save()) return back()->with('success', trans('site.updated_success')); 
            throw new Throwable("Can't save");
        } catch (Throwable $th) {
            logger($th);
            return back()->withInput()->with('failed', trans('site.created_failed'));
        }
    }

    /**
     * Remove the specified resource from storage.
     *
     * @param  \App\Models\Animal  $animal
     * @return \Illuminate\Http\Response
     */
    public function destroy(Animal $animal)
    {
        try {
            if ($animal->delete()) return Response(__('titles.deleted_success'), 301);
            return Response(__('titles.deleted_success'), 302);
        } catch (Throwable $th) {logger($th);}
    }

    /**
     * Display the specified resource.
     *
     * @param  \App\Models\Animal  $animal
     * @return \Illuminate\Http\Response
     */
    public function generic(Animal $animal, Request $request)
    {
        try {

            $user = Auth::user();
            return view('groups.animal.generic')->with(compact('animal', 'user'));
        } catch (Throwable $th) {
            logger($th);
        }
    }
    /**
     * Display the specified resource.
     *
     * @param  \App\Models\Animal  $animal
     * @return \Illuminate\Http\Response
     */
    public function book(Request $request)
    {
        try {
            $user = Auth::user();
            $animal = Animal::where('user_id', $user->id??'')->first();
            // dd($animal->user->name??'');
            return view('groups.animal.book')->with(compact('animal'));
        } catch (Throwable $th) {
            logger($th);
        }
    }
    /**
     * Display the specified resource.
     *
     * @param  \App\Models\Animal  $animal
     * @return \Illuminate\Http\Response
     */
    public function practice(Request $request)
    {
        try {
            $user = Auth::user();
            $animal = Animal::where('user_id', $user->id??'')->first();
            $task_today = $user->queue_today??0;
            $task_limit = $user->queue_limit??0;
            if($task_today >= $task_limit) return view('groups.animal.npc_talk')->with(['message'=> "Luyện tập hôm nay đã đủ"]);
            $user->queue_today ++;
            $user->save();
            return view('groups.animal.practice')->with(compact('animal'));
        } catch (Throwable $th) {
            logger($th);
        }
    }
}
