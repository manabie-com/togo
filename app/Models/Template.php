<?php

namespace App\Models;

use Illuminate\Database\Eloquent\Model;
use Illuminate\Database\Eloquent\SoftDeletes;
use Illuminate\Database\Eloquent\Factories\HasFactory;
use Illuminate\Support\Facades\DB;
use App\Traits\{ListTrait, HasLink};

class Template extends Model
{
    use ListTrait, HasLink;
    use HasFactory;
    use SoftDeletes;

    /**
     * The database table used by the model.
     *
     * @var string
     */
    protected $table = 'templates';

     /**
     * The database table alias
     *
     * @var string
     */
    protected $table_alias = '';

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
        'short_name',
        'prefix',
        'parent_id',
        'description',
        'type_id',
        'status_id',
    ];

    /**
     * The attributes that should be cast to native types.
     *
     * @var array
     */
    protected $casts = [
    ];

    function __construct($attributes = []){
        parent::__construct($attributes);
    }    
    /**
     * 
     * 
     */
    public function status()
    {
        return $this->belongsTo(Status::class);
    }
    /**
     * 
     * 
     */
    public function type()
    {
        return $this->belongsTo(Type::class);
    }
    /**
     * 
     *
     * @var array
     */
    public function user()
    {
        return $this->belongsTo(User::class);
    }
    /**
     * 
     *
     * @var array
     */
    public function hasManyThrough()
    {
        // return $this->hasManyThrough(ClassName1::class, ClassName2::class);
    }
    /**
     * 
     * 
     */
    public function hasMany()
    {
        // return $this->hasMany(ClassName::class);
    }

    /*-----Functions------*/
    /**
     * @todo
     * @return Call SP
     */
    public function callSP()
    {
        // return collect(DB::select("CALL sp()"));
    }

    /**
     * @return boolean
     */
    public function countHasMany()
    {
        // return $this->hasMany(ClassName::class)->count();
    }    
}
