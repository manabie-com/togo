<?php

namespace Support\Filter;

use Illuminate\Support\Arr;
use Illuminate\Support\Carbon;
use Illuminate\Support\Str;
use Spatie\QueryBuilder\AllowedFilter;
use Support\Builder\Builder;

/*
 * Use filter
 * https://spatie.be/docs/laravel-query-builder/v3/features/filtering
 */

abstract class Filter
{
    protected array $allowedFilters = [];

    /**
     * Apply filter
     *
     * @return array
     */
    public function apply(): array
    {
        $filters = [];
        $request = request(config('query-builder.parameters.filter'), []);
        if (is_array($request)) {
            foreach ($request as $key => $filter) {
                $method = Str::camel($key);
                if (! method_exists($this, $method)) {
                    continue;
                }
                if ((is_string($filter) && strlen($filter)) || (is_array($filter) && ! empty($filter))) {
                    $filters[] = AllowedFilter::callback($key,
                        function (Builder $builder) use ($method, $filter) {
                            $this->{$method}($builder, $filter);
                        });
                }
            }
        }

        return Arr::collapse([$this->allowedFilters, $filters]);
    }

    /**
     * Filter by created_at
     *
     * @param Builder $builder
     * @param $date
     * @return Builder
     */
    public function createdAt(Builder $builder, $date): Builder
    {
        $date = Str::of($date)->split('/[\s,]+/')->map(function ($date) {
            return Carbon::parse(trim($date));
        });
        $date = [
            'from' => $date->count() ? $date[0]->startOfDay() : null,
            'to'   => $date->count() > 1 ? $date[1]->endOfDay() : null,
        ];

        return $builder->whereDateRange('created_at', $date);
    }
}
