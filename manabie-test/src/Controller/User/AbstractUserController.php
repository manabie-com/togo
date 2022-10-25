<?php
declare(strict_types=1);

namespace App\Controller\User;

use App\Controller\BaseController;
use App\Service\User\Login;

/**
 * Class AbstractUserController
 * @package App\Controller\User
 */
abstract class AbstractUserController extends BaseController
{
    protected function getLoginUserService(): Login
    {
        return $this->container->get('login_user_service');
    }
}
