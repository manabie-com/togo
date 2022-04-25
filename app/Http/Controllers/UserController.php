<?php

namespace App\Http\Controllers;

use Auth;
use App;
use \Exception;
use \Throwable;

use App\Http\Controllers\Controller;
use Illuminate\Http\{Request,Response};
use Illuminate\Support\Facades\{Validator, Redirect, Hash};
use Illuminate\Validation\Rule;
use Illuminate\Pagination\LengthAwarePaginator;
use Illuminate\Support\Str;
use Maatwebsite\Excel\Facades\Excel;
use Carbon\Carbon;

use App\View\Components\{ViewModel, ActionModel, FilterRaws};
use App\Traits\{CaptureIpTrait, CastRequestTrait, Paginator};
use App\Models\{User, Category};
use jeremykenedy\LaravelRoles\Models\Role;

class UserController extends Controller
{
    use CastRequestTrait;
    /**
     * Create a new controller instance.
     *
     * @return void
     */
    public function __construct()
    {
        /*double check role*/
        $this->middleware('auth');
    }

    /**
     * Display a listing of the resource.
     *
     * @return \Illuminate\Http\Response
     */
    public function index(Request $request, User $user)
    {
        try {
            $action = $request->segment(2);
            $filterKeys = [];
            $filterData = [];
            // $perPage = (int)$request->input('per-page');
            /*setup*/
            $colTitle = [
                'no'            => 'STT',
                'code-link'     => 'Tài khoản',
                'email'         => 'Email',
                'roles'         => 'Vai trò',
                'customer'      => 'Khách hàng',
                'created_at'    => 'Ngày tạo',
                'updated_at'    => 'Thay đổi',
                'description'   => 'Ghi chú',
                '_action'       => '<i class="fa fa-cog" aria-hidden="true"></i>',
            ];
            $filterModel = new FilterRaws;
            $filterKeys = [
                'keyword' => $request->input('keyword'),
                'category_id' => $request->input('category_id'),
                'brand_id' => $request->input('brand_id'),
            ];
            // $filterModel->with('category_id', Category::getListKVByPrefix(), $filterKeys['category_id'], 'Tất cả');
            /**/
            $subQuery = User::query();
            if (!empty($filterKeys['keyword'])) {
                $subQuery->where('users.code', $filterKeys['keyword'])
                ->orwhere('users.name', 'LIKE', '%'. $filterKeys['keyword'].'%');
            }
            $newQuery = $user->fromSub($subQuery, 'users')->orderBy('id', 'asc');
            /*filter*/
            /*get pagination*/
            $pagination = (new Paginator(($newQuery->get()), $request->segment(2) === 'export' ? -1 : null));
            $collection = $pagination->getCollection();
            $collection->map(function ($item, $key) {
                if (is_object($item->roles)) {}
                $item->__set('code-link', ['href' => $item->route('show'), 'label' => $item->name ?? '']);
                /*set action*/
                $item->__set('_action', ['show' => $item->route('show'), 'edit' => $item->route('edit'), 'destroy' => $item->route('destroy')]);
            });
            /*export list*/
            if ($action === 'export') {
                unset($colTitle['_action']);
                $collection->map(function ($item, $key) {});
                return $this->exportList($collection, $colTitle);
            }
            return view('users.index')->with([
                'filterKeys'    => $filterKeys,
                'filterData'    =>$filterModel->toData(),
                'pageTitle'     => 'Danh sách người dùng',
                'colTitle'      => $colTitle,
                'pagination'    => $pagination,
                'dataRows'      => $collection->toArray(),
                'addLink'       => $user->route('create'),
                'exportLink'    => $user->route('export'),
            ]);
        } catch (Exception $e) {
            logger($e);
        }
    }

    /**
     * Show the form for creating a new resource.
     *
     * @return \Illuminate\Http\Response
     */
    public function create()
    {
        $roles = Role::all();

        return view('users.create', compact('roles'));
    }

    /**
     * Store a newly created resource in storage.
     *
     * @param \Illuminate\Http\Request $request
     *
     * @return \Illuminate\Http\Response
     */
    public function store(Request $request)
    {
        $validator = Validator::make(
            $request->all(),
            [
                'name'                  => 'required|max:255|unique:users|alpha_dash',
                'first_name'            => 'required|max:64',
                'last_name'             => 'required|max:64',
                'email'                 => 'required|email|max:255|unique:users',
                'password'              => 'required|min:6|max:20|confirmed',
                'password_confirmation' => 'required|same:password',
                'role'                  => 'required',
            ],
            [
                'name.unique'         => trans('auth.userNameTaken'),
                'name.required'       => trans('auth.userNameRequired'),
                'first_name.required' => trans('auth.fNameRequired'),
                'last_name.required'  => trans('auth.lNameRequired'),
                'email.required'      => trans('auth.emailRequired'),
                'email.email'         => trans('auth.emailInvalid'),
                'password.required'   => trans('auth.passwordRequired'),
                'password.min'        => trans('auth.PasswordMin'),
                'password.max'        => trans('auth.PasswordMax'),
                'role.required'       => trans('auth.roleRequired'),
            ]
        );

        if ($validator->fails()) {
            return back()->withErrors($validator)->withInput();
        }

        $ipAddress = new CaptureIpTrait();
        $profile = new Profile();

        $user = User::create([
            'name'             => strip_tags($request->input('name')),
            'first_name'       => strip_tags($request->input('first_name')),
            'last_name'        => strip_tags($request->input('last_name')),
            'email'            => $request->input('email'),
            'password'         => Hash::make($request->input('password')),
            'token'            => str_random(64),
            // 'admin_ip_address' => $ipAddress->getClientIp(),
            'activated'        => 1,
        ]);

        $user->profile()->save($profile);
        $user->attachRole($request->input('role'));
        $user->save();

        return redirect('users')->with('success', trans('usersmanagement.createSuccess'));
    }

    /**
     * Display the specified resource.
     *
     * @param User $user
     *
     * @return \Illuminate\Http\Response
     */
    public function show(User $user)
    {
        return view('users.show', compact('user'));
    }

    /**
     * Show the form for editing the specified resource.
     *
     * @param User $user
     *
     * @return \Illuminate\Http\Response
     */
    public function edit(User $user)
    {
        $roles = Role::all();

        foreach ($user->roles as $userRole) {
            $currentRole = $userRole;
        }
        
        $data = [
            'user'        => $user,
            'roles'       => $roles,
            'currentRole' => $currentRole,
            'listLink'    => $user->route('index'),
        ];

        return view('users.edit-combo')->with($data);
    }

    /**
     * Update the specified resource in storage.
     *
     * @param \Illuminate\Http\Request $request
     * @param User $user
     *
     * @return \Illuminate\Http\Response
     */
    public function update(Request $request, User $user)
    {
        try {
            $input = $request->all();
            $this->castDefault($input, 'role', $user->role);
            $validator = Validator::make($input,
                [
                    // 'code' => ['required','alpha_dash','max:16',Rule::unique($user->getTable())->ignore($user->id)],
                    'first_name' => 'required|max:100',
                    'full_name'     => 'required|max:10',
                    'email'      => ['required','email',Rule::unique($user->getTable())->ignore($user->id)],
                ], 
                [
                    'required'  => ':attribute là bắt buộc',
                    'first_name.required' => 'Tên là bắt buộc',
                    'unique'    => ':attribute đã tồn tại',
                    'max'       => 'Độ dài tối đa là :max',
                    'alpha_dash'=> 'Ký tự phải là chữ/số/"-"/"_"',
                    'email.email' => 'Địa chỉ email không hợp lệ',
                ]
            );
            if ($validator->fails()) return back()->withErrors($validator)->withInput();
            
            /**/
            if ($user->fill($input)->save()) return back()->with('success', trans('titles.updated_success')); 
            throw new Throwable("Can't save");
        } catch (Throwable $e) {
            logger($e);
            return back()->withInput()->with('failed', trans('titles.updated_failed'));
        }
    }

    /**
     * Remove the specified resource from storage.
     *
     * @param User $user
     *
     * @return \Illuminate\Http\Response
     */
    public function destroy(User $user)
    {
        $currentUser = Auth::user();
        $ipAddress = new CaptureIpTrait();

        if ($user->id !== $currentUser->id) {
            $user->deleted_ip_address = $ipAddress->getClientIp();
            $user->save();
            $user->delete();

            return redirect('users')->with('success', trans('usersmanagement.deleteSuccess'));
        }

        return back()->with('error', trans('usersmanagement.deleteSelfError'));
    }

    /**
     * Method to search the users.
     *
     * @param Request $request
     *
     * @return \Illuminate\Http\Response
     */
    public function search(Request $request)
    {
        $searchTerm = $request->input('user_search_box');
        $searchRules = [
            'user_search_box' => 'required|string|max:255',
        ];
        $searchMessages = [
            'user_search_box.required' => 'Search term is required',
            'user_search_box.string'   => 'Search term has invalid characters',
            'user_search_box.max'      => 'Search term has too many characters - 255 allowed',
        ];

        $validator = Validator::make($request->all(), $searchRules, $searchMessages);

        if ($validator->fails()) {
            return response()->json([
                json_encode($validator),
            ], Response::HTTP_UNPROCESSABLE_ENTITY);
        }

        $results = User::where('id', 'like', $searchTerm.'%')
                            ->orWhere('name', 'like', $searchTerm.'%')
                            ->orWhere('email', 'like', $searchTerm.'%')->get();

        // Attach roles to results
        foreach ($results as $result) {
            $roles = [
                'roles' => $result->roles,
            ];
            $result->push($roles);
        }

        return response()->json([
            json_encode($results),
        ], Response::HTTP_OK);
    }

    /**
     * @author toannguyen.dev
     * @todo export the list to excel
     * @param Collection $collection
     * @param Array $customHeading
     * @return xlsl file
     * @var 
    */
    public function exportList($collection = [], $customHeading = [])
    {
        try {
            $customHeadingDefault = [];
            if (empty($customHeading)) $customHeading = $customHeadingDefault;
            $nowString = Carbon::now()->format('YmdHi');
            $fileName = 'Product-list-'. $nowString . '.xlsx';
            $exportExcel = new ExportExcelController($collection, $customHeading);
            $exportExcel->title = 'Danh sách';
            return Excel::download($exportExcel, $fileName);
        } catch (Exception $e) {logger($e);}
    }
    /**
     * show default index
     *
     * @param
     * @return
     */
    public function dashboard(Request $request, User $User)
    {
        try {
            return view('dashboard');
        } catch (Exception $e) {
            logger($e);
        }
    }
        /**
     * show default index
     *
     * @param
     * @return
     */
    public function welcome(Request $request, User $User)
    {
        try {
            return view('welcome');
        } catch (Exception $e) {
            logger($e);
        }
    }
}
