<?php

namespace App\Models;

use Illuminate\Database\Eloquent\Model;

class Task extends Model
{
    const TABLE = 'task';
    protected $table      = self::TABLE;
    protected $connection = 'edu_test';
    public    $timestamps = false;

    protected $fillable = [
        'id',
        'user_id',
        'name',
        'date',
        'time_created',
    ];
}
