<?php
namespace App\Traits;
use Illuminate\Support\Collection;
use Illuminate\Pagination\LengthAwarePaginator;

class Paginator extends LengthAwarePaginator
{
    /**
     * Create a new paginator instance.
     *
     * @param  Collection  $items
     * @param  int  $perPage
     * @param  int|null  $currentPage
     * @param  array  $options  (path, query, fragment, pageName)
     * @return void
     */
    function __construct(Collection $items, $perPage = null, $currentPage = null, array $options = [])
    {
        $total = $items->count();
        $perPage = max(5, $perPage ==-1 ? $total : $perPage ?: request('per_page') ?: request('per-page') ?:15);
        $currentPage = min((int)ceil($total/$perPage), $currentPage ?: request('current_page') ?: request('page') ?:1);
        parent::__construct($items, $total, $perPage, $currentPage, $options);
        $this->items = $this->items->forPage($currentPage, $perPage);
        $this->setPath();
    }

    public function setPath($path = null)
    {
        $this->path = $path ?:url()->current();
    }
    /**
     * Render the paginator using the given view.
     *
     * @param  string|null  $view
     * @param  array  $data
     * @return \Illuminate\Contracts\Support\Htmlable
     */
    public function render($view = 'partials.pagination', $data = [])
    {
        return parent::render($view, $data);
    }
    
}
