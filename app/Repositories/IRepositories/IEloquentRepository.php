<?php

namespace Repository\IRepositories;

use Illuminate\Database\Eloquent\Model;
use Illuminate\Pagination\LengthAwarePaginator;
use Illuminate\Support\Collection;
use Support\Filter\Filter;
use Support\Sort\Sort;

interface IEloquentRepository
{
    public function find($id): Model;

    public function finds(array $ids): Collection;

    public function getAll(Filter $filter = null, Sort $sort = null): Collection;

    public function getList(Filter $filter = null, Sort $sort = null): LengthAwarePaginator;

    public function create(array $attributes): Model;

    public function update(Model $element, array $attributes = []): Model;

    public function delete(Model $element);
}
