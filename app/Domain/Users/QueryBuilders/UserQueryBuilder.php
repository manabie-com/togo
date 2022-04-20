<?php

namespace Domain\Users\QueryBuilders;

use Support\Builder\Builder;

class UserQueryBuilder extends Builder
{
    /**
     * Login from user_name
     *
     * @param $userName
     * @return $this
     */
    public function login($userName): self
    {
        return $this->where('user_name', $userName);
    }
}
