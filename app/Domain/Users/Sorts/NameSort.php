<?php

namespace Domain\Users\Sorts;

use Illuminate\Database\Eloquent\Builder;
use Spatie\QueryBuilder\Sorts\Sort;

class NameSort implements Sort
{
    /**
     * @param Builder $query
     * @param bool $descending
     * @param string $property
     * @return Builder
     */
    public function __invoke(Builder $query, bool $descending, string $property): Builder
    {
        $direction = $descending ? 'DESC' : 'ASC';

        return $query->orderBy($property, $direction);
    }
}
