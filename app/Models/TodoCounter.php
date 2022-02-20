<?php


namespace App\Models;


use Illuminate\Database\Eloquent\Model;

class TodoCounter extends Model
{
    protected $table = 'todo_counters';

    /**
     * The attributes that are mass assignable.
     *
     * @var array<int, string>
     */
    protected $fillable = [
        'user_id',
        'max_count',
        'count_today'
    ];

}
