<?php
namespace App\Traits;
use App\Models\Link;
trait HasParent
{
    public function parent(){
        return $this->belongsTo(static::class,'parent_id');
    }

    public function childs(){
      return static::whereParentId($this->id);
    }

    public function path(){
        return $this->growIds($this,[]);
    }

    public function tree($select = ['id','parent_id']){
      return $this->grow($this,[],$select);
    }


    public function road(){
        return $this->rback($this,[]);
    }

    function rback($model,$arr){
        if(isset($model->parent)){
            $arr[] = $model->parent;
            $this->rback($arr,$model->parent);
        }
        return $arr;
    }

    function grow($node,$arr,$select,$level = 0){
        $level++;
        foreach ($node->childs()->select($select)->get() as $item){
            $item->level  = $level;
            $arr[]        = $item;
            if($item->childs()->count()){
                $arr  = $this->grow($item,$arr,$select,$level);
            }
        }
        return $arr;
    }

    function growIds($node,$arr,$level = 0){
        $level++;
        foreach ($node->childs()->select('id')->get() as $item){
            $arr[] = $item->id;
            if($item->childs()->count()){
                $arr  = $this->grow($item,$arr,$select,$level);
            }
        }
        return $arr;
    }

    public static function root(){
        return static::whereParentId(0);
    }

    public static function branchs($group){
        foreach($group as $item){
            $collect[] = $item;
            foreach($item->tree() as $branch){
                $collect[] = $branch;
            }
        }
        return $collect;
    }

    public function scopeArrange($query){
        $collect = [];
        foreach($query->get() as $item){
            $collect[] = $item;
            foreach($item->tree() as $branch){
                $collect[] = $branch;
            }
        }
        return $collect;
    }
}
