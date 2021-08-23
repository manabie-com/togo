<?php
declare(strict_types=1);

namespace App\Entity;

/**
 * Class User
 * @package App\Entity
 */
class User
{
    /**
     * @var string
     */
    private string $id;

    /**
     * @var string
     */
    private string $password;

    /**
     * @var int
     */
    private int $maxTodo;

    /**
     * @return object
     */
    public function toJson(): object
    {
        return json_decode((string) json_encode(get_object_vars($this)), false);
    }

    /**
     * @return string
     */
    public function getId(): string
    {
        return $this->id;
    }

    /**
     * @return string
     */
    public function getPassword(): string
    {
        return $this->password;
    }

    /**
     * @param string $password
     * @return $this
     */
    public function setPassword(string $password): self
    {
        $this->password = $password;

        return $this;
    }

    /**
     * @return int
     */
    public function getMaxToDo(): int
    {
        return $this->maxTodo;
    }

    /**
     * @param int $maxTodo
     * @return $this
     */
    public function setMaxToDo(int $maxTodo): self
    {
        $this->maxTodo = $maxTodo;

        return $this;
    }

    /**
     * @param string $id
     */
    public function setId(string $id): void
    {
        $this->id = $id;
    }
}
