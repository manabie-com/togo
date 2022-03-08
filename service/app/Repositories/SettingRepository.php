<?php

namespace App\Repositories;

use App\Models\Setting;
use App\Repositories\EloquentRepository;

class SettingRepository extends EloquentRepository
{
    /**
     * get model
     * @return string
     */
    public function getModel()
    {
        return Setting::class;
    }

    public function getLimit($userId)
    {
        $query = $this->_model
            ->select('limit');
        if ($userId) {
            $query = $query->where('user_id', '=', $userId);
        }
        return $query->first();
    }

}
