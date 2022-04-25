<?php

namespace App\Models;

use Illuminate\Database\Eloquent\Model;
use Illuminate\Database\Eloquent\SoftDeletes;
use Illuminate\Database\Eloquent\Factories\HasFactory;
use App\Traits\ListTrait;

class Animal extends Model
{
    use ListTrait;
    use HasFactory;
    use SoftDeletes;

    /**
     * The database table used by the model.
     *
     * @var string
     */
    protected $table = 'animals';

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
        'status_id',
        'type_id',
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
}
