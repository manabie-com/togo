<?php
namespace App\Traits;

trait HasGroup
{

    /**
     * 
     * 
     */
    public function viewGroup(string $view, $data = [], $options=[])
    {
        $separator = $options['separator'] ??'.';
        $roleName = $options['role'] ?? $this->getBaseFolder();
        $view = implode($separator, ['groups', $roleName, str_replace('_', '-', $view)]);
        return view($view, $data);
    }
    /**
     * @todo get basename of current folder of reflection class
     * @return string
     */
    private function getBaseFolder():string {
        $reflection = new \ReflectionClass($this);
        return strtolower((string)basename(dirname($reflection->getFileName())));
    }
}
