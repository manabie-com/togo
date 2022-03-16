<?php


namespace App\Http\Repositories;

use App\Base\Repositories\BaseRepository;
use App\Models\TodoCounter;

class TodoCounterRepository extends BaseRepository
{

    public function __construct(TodoCounter $model)
    {
        parent::__construct($model);
    }
}
