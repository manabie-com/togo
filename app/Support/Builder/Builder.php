<?php

namespace Support\Builder;

use Illuminate\Pagination\LengthAwarePaginator;
use Illuminate\Support\Str;
use DateTimeInterface;
use Illuminate\Database\Eloquent\Builder as BaseBuilder;
use Illuminate\Database\Query\Builder as QueryBuilder;
use Illuminate\Pagination\Paginator;
use Illuminate\Support\Arr;

class Builder extends BaseBuilder
{
    /**
     * Filter multiple
     *
     * @param $column
     * @param $value
     * @return $this
     */
    public function filterMultiple($column, $value): self
    {
        if (is_string($value)) {
            $value = Str::of($value)->split('/[\s,]+/');
        }

        return $this->whereIn($column, $value);
    }

    /**
     * @param Closure|string|array $column
     * @param string|null $value
     * @return $this
     */
    public function whereStartsWith($column, $value = null): Builder
    {
        $this->where($column, 'like', $value.'%');

        return $this;
    }

    /**
     * @param Closure|string|array $column
     * @param string|null $value
     * @return $this
     */
    public function whereEndsWith($column, $value = null): Builder
    {
        $this->where($column, 'like', '%'.$value);

        return $this;
    }

    /**
     * @param Closure|string|array $column
     * @param string|null $value
     * @return $this
     */
    public function whereLike($column, $value = null): Builder
    {
        $this->where($column, 'like', '%'.$value.'%');

        return $this;
    }

    /**
     * @param Closure|string|array $column
     * @param string|null $value
     * @return $this
     */
    public function whereEqual($column, $value = null): Builder
    {
        $this->where($column, '=', $value);

        return $this;
    }

    /**
     * @param Closure|string|array $column
     * @param string|null $value
     * @return $this
     */
    public function whereNotEqual($column, $value = null): Builder
    {
        $this->where($column, '<>', $value);

        return $this;
    }

    /**
     * @param string $column
     * @param array $value
     * @return Builder
     */
    public function whereDateRange(string $column, array $value = []): Builder
    {
        $from = Arr::get($value, 'from', '');
        $to = Arr::get($value, 'to', '');
        $this->query->where(function (QueryBuilder $query) use ($column, $from, $to) {
            return $query
                ->when($from, function (QueryBuilder $query) use ($column, $from) {
                    return $query->where($column, '>=', $from);
                })
                ->when($to, function (QueryBuilder $query) use ($column, $to) {
                    return $query->where($column, '<=', $to);
                });
        });

        return $this;
    }

    /**
     * @param $column
     * @param null $operator
     * @param null $value
     * @return Builder
     */
    public function whereDateTime($column, $operator = null, $value = null): Builder
    {
        if ($value instanceof DateTimeInterface) {
            $value = $value->format('H:i:s');
        }

        return $this->where($column, $operator, $value);
    }

    /**
     * Paginate the given query.
     *
     * @param int|null $perPage
     * @param array $columns
     * @param string $pageName
     * @param int|null $page
     * @return LengthAwarePaginator
     *
     * @throws \InvalidArgumentException
     */
    public function paginate($perPage = null, $columns = ['*'], $pageName = 'page', $page = null): LengthAwarePaginator
    {
        $page = $page ?: Paginator::resolveCurrentPage($pageName);

        $total = $this->toBase()->getCountForPagination();

        $perPage = $perPage ?: ($total != 0 ? $total : 1);

        $results = $total
            ? $this->forPage($page, $perPage)->get($columns)
            : $this->model->newCollection();

        return $this->paginator($results, $total, $perPage, $page, [
            'path'     => Paginator::resolveCurrentPath(),
            'pageName' => $pageName,
        ]);
    }
}
