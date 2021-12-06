<?php

namespace App\Models;

use App\Models\BaseModel;
use App\Models\User;

class Task extends BaseModel
{
    protected $fillable = [
        'name',
        'user_id'
    ];

    public function user()
    {
        return $this->belongsTo(User::class);
    }
}
