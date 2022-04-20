<?php

namespace Api\Controllers;

use Domain\Users\Filters\UserFilter;
use Domain\Users\Sorts\UserSort;
use Domain\Users\Transformers\UserTransformer;
use Illuminate\Http\JsonResponse;
use Repository\IRepositories\IUserRepository;

class UserController extends ApiController
{
    private IUserRepository $userRepository;

    /**
     * UserController constructor.
     *
     * @param IUserRepository $userRepository
     */
    public function __construct(IUserRepository $userRepository)
    {
        $this->userRepository = $userRepository;
    }

    /**
     * Display a listing of the resource.
     *
     * @param UserFilter $filter
     * @param UserSort $sort
     * @return JsonResponse
     */
    public function index(UserFilter $filter, UserSort $sort): JsonResponse
    {
        $user = $this->userRepository->getList($filter, $sort);

        return $this->httpOK($user, UserTransformer::class);
    }
}
