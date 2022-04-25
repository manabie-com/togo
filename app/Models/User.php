<?php

namespace App\Models;

use Illuminate\Database\Eloquent\Model;
use Illuminate\Database\Eloquent\SoftDeletes;
use Illuminate\Database\Eloquent\Factories\HasFactory;
use Illuminate\Foundation\Auth\User as Authenticatable;
use Illuminate\Notifications\Notifiable;
use jeremykenedy\LaravelRoles\Traits\HasRoleAndPermission;

use App\Traits\{HasLink, ListTrait};

class User extends Authenticatable
{
    use HasLink, ListTrait;
    use HasRoleAndPermission;
    use Notifiable;
    use HasFactory;
    use SoftDeletes;

    /**
     * The database table used by the model.
     *
     * @var string
     */
    protected $table = 'users';

    /**
     * Indicates if the model should be timestamped.
     *
     * @var bool
     */
    public $timestamps = true;

    /**
     * The attributes that are not mass assignable.
     *
     * @var array
     */
    protected $guarded = [
        'id',
    ];

    /**
     * The attributes that are hidden.
     *
     * @var array
     */
    protected $hidden = [
        'password',
        'remember_token',
        'activated',
        'token',
    ];

    /**
     * The attributes that should be mutated to dates.
     *
     * @var array
     */
    protected $dates = [
        'created_at',
        'updated_at',
        'deleted_at',
    ];

    /**
     * The attributes that are mass assignable.
     *
     * @var array
     */
    protected $fillable = [
        'code',
        'name',
        'first_name',
        'last_name',
        'email',
        'password',
        'activated',
        'token',
        'remember_token',
        'remember_login',
        'latest_login_ip',
        'latest_login_time',
        'updated_ip_address',
    ];

    /**
     * The attributes that should be cast to native types.
     *
     * @var array
     */
    protected $casts = [
        'id'                                => 'integer',
        'first_name'                        => 'string',
        'last_name'                         => 'string',
        'email'                             => 'string',
        'password'                          => 'string',
        'activated'                         => 'boolean',
        'token'                             => 'string',
        'updated_ip_address'                => 'string',
    ];
    /**
     * get name and id
     *
     * @param  Profile $profile
     */
    public function getNameAndCode()
    {
        $code = empty($this->code) ? '' : '('.($this->code).')';
        return $this->name . $code;
    }
     /**
     * get name and code in array
     *
     * @param  array
     */   
    public static function getNameAndCodeArray():array
    {
        $nameAndCodeArray = [];
        $data = (array)self::all()->sortBy('name', SORT_NATURAL|SORT_FLAG_CASE)->toArray();
        foreach ($data as $_index => $_model) {
            $code = isset($_model['code']) ? ' ('. $_model['code'].')' : '';
            $nameAndCodeArray[$_model['id']] = $_model['name'] . $code;
        }
        return $nameAndCodeArray;
    }
    /**
     * get fullname
     *
     * @param  array
     */
    public function getFullname()
    {
        return $this->last_name . ' ' . $this->first_name;
    }
    /**
     * 
     * 
     */
    public function customer()
    {
        return $this->hasOne(Customer::class);
    }
    /**
     * @todo  
     * 
     */
    public function supplier()
    {
        return $this->hasOne(Supplier::class);
    }
}
