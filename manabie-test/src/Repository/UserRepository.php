<?php
declare(strict_types=1);

namespace App\Repository;

use App\Entity\User;
use App\Exception\UserException;

/**
 * Class UserRepository
 * @package App\Entity\User
 */
class UserRepository extends AbstractRepository implements UserRepositoryInterface
{
    /**
     * @param string $userId
     * @param string $password
     * @return User
     * @throws UserException
     */
    public function loginUser(string $userId, string $password): User
    {
        $query = '
            SELECT *
            FROM `users`
            WHERE `id` = :id AND `password` = :password
            ORDER BY `id`
        ';
        $statement = $this->database->prepare($query);
        $statement->bindParam('id', $userId);
        $statement->bindParam('password', $password);

        $statement->execute();

        $user = $statement->fetch();

        if (!$user) {
            throw new UserException('Login failed: Id or password incorrect.', 400);
        }

        $user = new User();
        $user->setId($userId);

        return $user;
    }
}
