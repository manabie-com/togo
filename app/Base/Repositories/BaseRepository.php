<?php


namespace App\Base\Repositories;

use Illuminate\Database\Eloquent\Model;

class BaseRepository
{
    protected $model;

    public function __construct(Model $model)
    {
        $this->model = $model;
    }

    public function create($fields) {
        return $this->model::create($fields);
    }

    public function getByField($field) {
        return $this->model::firstOrCreate($field);
    }

    public function updateByField($field, $update){
        $this->model::updateOrCreate($field, $update)->save();
    }

    public function store()
    {
        // TODO: Implement store() method.
    }

    public function delete()
    {
        // TODO: Implement delete() method.
    }

    public function update()
    {
        // TODO: Implement update() method.
    }

    public function edit()
    {
        // TODO: Implement edit() method.
    }
}
