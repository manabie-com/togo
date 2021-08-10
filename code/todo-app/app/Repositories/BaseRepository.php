<?php

namespace App\Repositories;

use Illuminate\Database\Eloquent\Model;
use Illuminate\Database\Eloquent\Builder;
use App\Repositories\Contracts\RepositoryInterface;
use Prettus\Repository\Exceptions\RepositoryException;
use Prettus\Repository\Eloquent\BaseRepository as L5Repository;

/**
 * Class BaseRepository
 * @property Model|Builder $model
 * @package App\Repositories
 */
abstract class BaseRepository extends L5Repository implements RepositoryInterface
{
	/**
	 * @param array $conditions
	 * @param array $columns
	 * @return mixed
	 * @throws RepositoryException
	 */
	public function findWhereForUpdate(array $conditions, $columns = ['*'])
    {
        $this->applyConditions($conditions);

        $results = $this->model->lockForUpdate()->first($columns);

        $this->resetModel();

        return $this->parserResult($results);
    }

    public function findWhereFirst(array $conditions, $columns = ['*'])
    {
	    $this->applyConditions($conditions);

	    return $this->first($columns);
    }

	/**
	 * @param array $where
	 */
	protected function applyConditions(array $where)
	{
		foreach ($where as $field => $value) {
			if (is_array($value)) {
				list($field, $condition, $val) = $value;
				if (strtoupper($condition) == 'IN') {
					$this->model = $this->model->whereIn($field, $val);
				} else if (strtoupper($condition) == 'NOT_IN') {
					$this->model = $this->model->whereNotIn($field, $val);
				} else {
					$this->model = $this->model->where($field, $condition, $val);
				}
			} else {
				$this->model = $this->model->where($field, '=', $value);
			}
		}
	}
}
