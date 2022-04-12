<?php

namespace Repository\IRepositories;

interface IUserRepository extends IEloquentRepository
{
    /**
     * Login
     *
     * @param array $authData
     * @return array|null
     */
    public function login(array $authData): ?array;
}
