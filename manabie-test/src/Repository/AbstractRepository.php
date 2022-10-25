<?php
declare(strict_types=1);

namespace App\Repository;

/**
 * Class AbstractRepository
 * @package App\Repository
 */
abstract class AbstractRepository
{
    /**
     * @var \PDO
     */
    protected \PDO $database;

    /**
     * @param \PDO $database
     */
    public function __construct(\PDO $database)
    {
        $this->database = $database;
    }

    /**
     * @return \PDO
     */
    protected function getDb(): \PDO
    {
        return $this->database;
    }

    /**
     * @param string $query
     * @param int $page
     * @param int $perPage
     * @param array $params
     * @param int $total
     * @return array
     */
    protected function getResultsWithPagination(string $query, int $page, int $perPage, array $params, int $total): array
    {
        return [
            'pagination' => [
                'totalRows' => $total,
                'totalPages' => ceil($total / $perPage),
                'currentPage' => $page,
                'perPage' => $perPage,
            ],
            'data' => $this->getResultByPage($query, $page, $perPage, $params),
        ];
    }

    /**
     * @param string $query
     * @param int $page
     * @param int $perPage
     * @param array $params
     * @return array
     */
    protected function getResultByPage(string $query, int $page, int $perPage, array $params): array
    {
        $offset = ($page - 1) * $perPage;
        $query .= " LIMIT ${perPage} OFFSET ${offset}";
        $statement = $this->database->prepare($query);

        $statement->execute($params);

        return (array)$statement->fetchAll();
    }
}
