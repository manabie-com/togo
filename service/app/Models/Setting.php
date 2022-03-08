<?php

namespace App\Models;

use Illuminate\Database\Eloquent\Model;

class Setting extends Model
{
    const TABLE = 'setting';
    protected $table      = self::TABLE;
    protected $connection = 'edu_test';
    public    $timestamps = false;

    protected $fillable = [
        'id',
        'user_id',
        'limit',
        'time_created',
        'time_updated',
    ];
}
