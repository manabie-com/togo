<?php


namespace App\Http\Repositories;


use App\Base\Repositories\BaseRepository;
use App\Models\Todo;

class TodoRepository extends BaseRepository
{
    public function __construct(Todo $model)
    {
        parent::__construct($model);
    }
}
