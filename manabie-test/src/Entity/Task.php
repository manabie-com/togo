<?php
declare(strict_types=1);

namespace App\Entity;

/**
 * Class Task
 * @package App\Entity
 */
class Task
{
    /**
     * @var string
     */
    private string $id;

    /**
     * @var string
     */
    private string $content;

    /**
     * @var string
     */
    private string $createdDate;

    /**
     * @var string
     */
    private string $userId;

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
    public function getContent(): string
    {
        return $this->content;
    }

    /**
     * @param string $content
     */
    public function setContent(string $content): void
    {
        $this->content = $content;
    }

    /**
     * @return string
     */
    public function getCreatedDate(): string
    {
        return $this->createdDate;
    }

    /**
     * @param string $createdDate
     * @return $this
     */
    public function setCreatedDate(string $createdDate): self
    {
        $this->createdDate = $createdDate;

        return $this;
    }

    /**
     * @return string
     */
    public function getUserId(): string
    {
        return $this->userId;
    }

    /**
     * @param string $userId
     * @return $this
     */
    public function setUserId(string $userId): self
    {
        $this->userId = $userId;

        return $this;
    }

    /**
     * @return object
     */
    public function toJson(): object
    {
        return json_decode((string) json_encode(get_object_vars($this)), false);
    }

    /**
     * @param string $id
     */
    public function setId(string $id): void
    {
        $this->id = $id;
    }
}
