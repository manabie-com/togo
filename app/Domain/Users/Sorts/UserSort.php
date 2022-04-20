<?php

namespace Domain\Users\Sorts;

use Spatie\QueryBuilder\AllowedSort;
use Support\Sort\Sort;

class UserSort extends Sort
{
    public string $defaultSort = '-updated_at';

    function allowedSorts(): array
    {
        return [
            'updated_at',
            AllowedSort::custom('name', new NameSort),
        ];
    }
}
