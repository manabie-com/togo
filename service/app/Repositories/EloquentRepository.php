<?php

namespace App\Repositories;

abstract class EloquentRepository implements RepositoryInterface
{
    protected $redisService;
    /**
     * @var \Illuminate\Database\Eloquent\Model
     */
    protected $_model;

    /**
     * EloquentRepository constructor.
     */
    public function __construct()
    {
        $this->setModel();
    }

    /**
     * get model
     * @return string
     */
    abstract public function getModel();

    /**
     * Set model
     */
    public function setModel()
    {
        $this->_model = app()->make(
            $this->getModel()
        );
    }

    /**
     * Get All
     * @return \Illuminate\Database\Eloquent\Collection|static[]
     */
    public function getAll()
    {
        return $this->_model->all();
    }

    /**
     * Get one
     * @param $id
     * @return mixed
     */
    public function find($id)
    {
        $result = $this->_model->find($id);
        return $result;
    }

    public function findOneField($fieldName, $value)
    {
        $result = $this->_model->where($fieldName, $value)->get()->toArray();
        return $result;
    }

    /**
     * Create
     * @param array $attributes
     * @return mixed
     */
    public function create(array $attributes)
    {
        return $this->_model->create($attributes);
    }

    /**
     * Insert mutiple data
     * @param array $attributes
     * @return mixed
     */

    public function insert(array $attributes)
    {
        return $this->_model->insert($attributes);
    }


    /**
     * Insert mutiple data
     * @param array $attributes
     * @return mixed
     */

    public function insertGetId(array $attributes)
    {
        return $this->_model->insertGetId($attributes);
    }

    /**
     * Save
     * @param array $attributes
     * @return mixed
     */
    public function save(array $attributes)
    {
        return $this->_model->save($attributes);
    }

    /**
     * Update
     * @param $id
     * @param array $attributes
     * @return bool|mixed
     */
    public function update($id, array $attributes)
    {
        $result = $this->find($id);
        if($result) {
            $result->update($attributes);
            return $result;
        }
        return false;
    }

    /**
     * Delete
     *
     * @param $id
     * @return bool
     */

    /**
     * update one field
     *
     * @param $name, $value
     * @return bool
     */
    public function updateByOneField($fieldName, $value, $attributes)
    {
        $result = $this->_model->where($fieldName, $value);
        if($result) {
            $result->update($attributes);
            return $result;
        }

        return false;

    }

    public function delete($id)
    {
        $result = $this->find($id);
        if($result) {
            $result->delete();
            return true;
        }

        return false;
    }

    /**
     * Delete one field
     *
     * @param $name, $value
     * @return bool
     */
    public function deleteByOneField($fieldName, $value)
    {
        $result = $this->_model->where($fieldName, $value);
        if($result) {
            $result->delete();
            return true;
        }

        return false;
    }

    /**
    * Count All
    * @return mixed
    */
   public function countAll()
   {
       return $this->_model->count();
   }
}
