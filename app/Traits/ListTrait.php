<?php
namespace App\Traits;
use Illuminate\Support\Str;
use Carbon\Carbon;



trait ListTrait
{
    /**
     * The attributes that should be cast to native types.
     * @todo the list content key by option attribute
     * @var array
     * @return array
     */
    public static function getListKVByPrefix($prefix = null, $attribute = 'name', string $id = 'id')
    {
        $list = [];
        try {
            $collection = is_null($prefix) ? self::All() : self::where('prefix', $prefix)->get();
            foreach ($collection as $key => $_item) {
                $list[$_item->__get($id)] = $_item->__get($attribute);
            }
        } catch (Exception $e) {
            logger($e);
        }
        return $list;
    }
    /**
     * @author toannguyen.dev
     * @todo array normalization
     * @param array &$list
     * @param $selected = null
     * @return boolean 
     */
    public static function toSelectTag(array &$list, $key_selected = null, $labelDefault = '-')
    {
        $lenght = 5;
        $chars = '0';
        foreach ($list as $value => $label) {
            $element = ['value'=>'', 'selected'=>'', 'label'=>'', 'count'=>0];
            unset($list[$value]);
            $newValue = str_pad($label->value??$label['value']??$value, $lenght, $chars, STR_PAD_LEFT);
            if (is_array($label) || is_object($label)) {
                $element = array_merge($element, (array)$label);
            } else{
                $element['value'] = $value;
                $element['label'] = $label;
            }
            $list[$newValue] = $element;
        }
        $list['_default'] = ['value'=>'', 'selected'=>'selected', 'label'=> $labelDefault, 'count'=>null];
        ksort($list);
        $key_selected = is_array($key_selected) ? $key_selected : explode(',',$key_selected);
        foreach ($key_selected as $key => $_key) {
            $_key = str_pad($_key, $lenght, $chars, STR_PAD_LEFT);
            if (array_key_exists($_key, $list)) {
                $list[$_key]['selected'] = 'selected';
                $list['_default']['selected'] = '';
            }
        }
        if ($labelDefault === 'hide' || is_null($labelDefault)) unset($list['_default']);
        return true;
    }
    /**
    *| @author toannguyen.dev
    *| @todo sort by key of node child
    *| @param array 
    *| @param string keyword,..
    *| @return pass value of array
    *| @version 1
    */
    public static function toMSort(array &$array, array $sort = []) 
    {
        try{
            $ka = array_map('strtolower', array_keys($sort));
            foreach ($sort as $key => $value) {
                if (is_numeric($key)) {
                    $ik = array_search(strtolower((string)$key), $ka);
                    $sort = array_slice($sort, 0, $ik, true) + [$value => 1] + array_slice($sort, $ik+1, count($sort), true);
                }
            }
            usort($array, function($a, $b) use($sort) {
                $a = (array)$a;
                $b = (array)$b;
                $i = 0;
                $c = count($sort);
                $v = array_keys($sort);
                $s = array_values($sort);
                $cmp = 0;
                while($cmp == 0 && $i < $c){
                    $sa = $s[$i] === 1 ? 1: -1;
                    $cmp = $sa * strcmp((string) ($a[$v[$i]] ?? ''), (string) ($b[$v[$i]]??''));
                    $i++;
                }
                return $cmp;
            });
        } catch(Exception $e){}
        return $array;
    }
}
