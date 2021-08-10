<?php

namespace App\Services\Contracts;

interface UserInterface
{
    public function register(string $username,  string $password);
}
