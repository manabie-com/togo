<?php

namespace App\Models;

use Illuminate\Database\Eloquent\Model;

/**
 * Class Task
 *
 * @property int    id
 * @property int    user_id
 * @property string content
 * @property int    created_at
 * @property int    updated_at
 *
 * @package App\Models
 */
class Task extends Model
{
    protected $table = 'tasks';

    protected $fillable = [
        'user_id',
        'content'
    ];
}
