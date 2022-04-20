<?php

namespace Repository;

use Illuminate\Database\Eloquent\Model;
use Illuminate\Pagination\LengthAwarePaginator;
use Illuminate\Support\Collection;
use Illuminate\Support\Facades\DB;
use Repository\IRepositories\IEloquentRepository;
use Spatie\QueryBuilder\QueryBuilder;
use Support\Filter\Filter;
use Support\Sort\Sort;

abstract class EloquentRepository implements IEloquentRepository
{
    /**
     * @var Model
     */
    protected Model $model;

    /**
     * EloquentRepository constructor.
     *
     */
    public function __construct()
    {
        $this->setModel();
    }

    /**
     * get model
     *
     * @return string
     */
    abstract public function getModel(): string;

    /**
     * Set model
     *
     */
    public function setModel()
    {
        $this->model = app()->make(
            $this->getModel()
        );
    }

    /**
     * Get All
     *
     * @param Filter|null $filter
     * @param Sort|null $sort
     * @return Collection
     */
    public function getAll(Filter $filter = null, Sort $sort = null): Collection
    {
        $query = $this->queryBuilder($filter, $sort);

        return $query->get();
    }

    /**
     * Get list
     *
     */
    public function getList(Filter $filter = null, Sort $sort = null): LengthAwarePaginator
    {
        $query = $this->queryBuilder($filter, $sort);

        return $query->paginate(10);
    }

    /**
     * Get one
     *
     * @param $id
     * @return Model|null
     */
    public function find($id): Model
    {
        return $this->model->query()->find($id);
    }

    /**
     * Get more
     *
     * @param array $ids
     * @return Collection
     */
    public function finds(array $ids): Collection
    {
        return $this->model->query()->whereIn('id', $ids)->get();
    }

    /**
     * Create
     *
     * @param array $attributes
     * @return \Illuminate\Database\Eloquent\Builder|Model
     */
    public function create(array $attributes): Model
    {
        $result = null;
        DB::beginTransaction();
        try {
            $result = $this->model->query()->create($attributes);
            DB::commit();
        } catch (\Exception $e) {
            DB::rollBack();
        }

        return $result;
    }

    /**
     * Update
     *
     * @param Model $element
     * @param array $attributes
     * @return Model
     */
    public function update(Model $element, array $attributes = []): Model
    {
        $result = null;
        DB::beginTransaction();
        try {
            $element->update($attributes);
            $result = $element->refresh();
            DB::commit();
        } catch (\Exception $e) {
            DB::rollBack();
        }

        return $result;
    }

    /**
     * Delete
     *
     * @param Model $element
     * @return bool
     */
    public function delete(Model $element): bool
    {
        $result = false;
        DB::beginTransaction();
        try {
            $result = $element->delete();
            DB::commit();
        } catch (\Exception $e) {
            DB::rollBack();
        }

        return $result;
    }

    /**
     * Query Builder
     *
     * @param Filter|null $filter
     * @param Sort|null $sort
     */
    private function queryBuilder(Filter $filter = null, Sort $sort = null): QueryBuilder
    {
        $query = QueryBuilder::for($this->model);
        // Filters
        if ($filter) {
            $query->allowedFilters($filter->apply());
        }
        // Sort
        if ($sort) {
            $query->allowedSorts($sort->allowedSorts());
            // Set sort default
            if (strlen($sort->defaultSort)) {
                $query->defaultSort($sort->defaultSort);
            }
        }

        return $query;
    }
}
