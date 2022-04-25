<?php

namespace App\Traits;

use Illuminate\Support\Str;
use App\Models\Link;
use App\Traits\Slug;

trait HasLink
{
    use Slug;
    public function link()
    {
        return $this->morphOne(Link::class, 'belong');
    }

    function upLink()
    {
        $slug   =   $this->slug;
        
        return Link::create([
            'pretty'      =>  static::slug(empty($this->$slug) ? $this->code : $this->$slug),
            'belong_type' =>  get_class($this),
            'belong_id'   =>  $this->id
        ]);
    }
    function dropLink()
    {
        $this->link->delete();
    }
    /**
     * 
     * @return URL
     */
    function route($action = null, $options = [])
    {
        $tables     = $options['tables']?? $this->table?? $this->table_alias?? '';
        $glue       = $options['glue']?? '.';
        $prefix     = $options['prefix']?? '';
        $prefixArr  = preg_split('/(,|\.|\/|-)/', $prefix, -1, PREG_SPLIT_NO_EMPTY);
        $id         = is_numeric($options)? $options: $options['id']?? $this->id?? null;
        $actionList = [
            'index'     => [$tables, 'index'],
            'create'    => [$tables, 'create'],
            'store'     => [$tables, 'store'],
            'show'      => [$tables, 'show'],
            'edit'      => [$tables, 'edit'],
            'update'    => [$tables, 'update'],
            'destroy'   => [$tables, 'destroy'],
            'export'    => [$tables, 'export'],
            'import'    => [$tables, 'import'],
            'back'      => [],
        ];
        $linkTemp = implode($glue, array_filter(array_merge($prefixArr, $actionList[$action]??[])));
        switch ($action) {
            case 'show': case 'edit': case 'update': case 'destroy':
                return route($linkTemp, $id);
            default: 
                return route($linkTemp);
        }
        return false;
    }
    /**
    * 
    * @return URL
    */
   function url($action = null, $options = null)
    {   
        $tables     = $options['tables']?? $this->table?? $this->table_alias?? '';
        $glue       = $options['glue']?? '/';
        $prefix     = $options['prefix']?? '';
        $prefixArr  = preg_split('/(,|\.|\/|-)/', $prefix, -1, PREG_SPLIT_NO_EMPTY);
        $id         = is_numeric($options)? $options: $options['id']?? $this->id?? null;
        $actionList = [
            'index'     => implode($glue, array_filter(array_merge($prefixArr, [$tables]))),
            'create'    => implode($glue, array_filter(array_merge($prefixArr, [$tables, 'create']))),
            'store'     => implode($glue, array_filter(array_merge($prefixArr, [$tables]))),
            'show'      => implode($glue, array_filter(array_merge($prefixArr, [$tables, $id]))),
            'edit'      => implode($glue, array_filter(array_merge($prefixArr, [$tables, $id, 'edit']))),
            'update'    => implode($glue, array_filter(array_merge($prefixArr, [$tables, $id]))),
            'destroy'   => implode($glue, array_filter(array_merge($prefixArr, [$tables, $id]))),
            'export'    => implode($glue, array_filter(array_merge($prefixArr, [$tables, 'export']))),
            'import'    => implode($glue, array_filter(array_merge($prefixArr, [$tables, 'import']))),
            'back'      => back()->getTargetUrl(),
        ];
        return $actionList[$action] ? url($actionList[$action]) : '';
    }
}
