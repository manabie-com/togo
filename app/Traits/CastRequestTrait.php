<?php
namespace App\Traits;
use Illuminate\Support\Str;
use Carbon\Carbon;

trait CastRequestTrait
{
    /**
     * 
     * 
     * 
     */
    static function castNameRequest(array &$input, ?string $attributes = null, $from_glue = '-', $to_glue = '_')
    {
        $attributes = array_filter(explode(',', $attributes));
        if (empty($attributes) || is_null($attributes)) {
            foreach ($input as $name => $value) {
                unset($input[$name]);
                $newName = str_replace($from_glue, $to_glue, $name);
                $input[$newName] = request($name);
            }
        } else{
            foreach ($input as $name => $value) {
                if (!array_key_exists($name, $attributes)) continue;
                unset($input[$name]);
                $newName = str_replace($from_glue, $to_glue, $name);
                $input[$newName] = request($name);
            }
        }
        return true;
    }
    /**
    * cast attributes to datetime format
    *
    * @param array &$input
    * @param string $attributes
    * @param string $fromFormat = 'd-m-Y'
    * @return boolean
    * @throws 
    */
    static function castDatetime(array &$input, string $attributes, string $fromFormat = 'd-m-Y')
    {
        try {
            if (array_key_exists($attributes, $input)) {
                $_value = str_replace('/', '-', $input[$attributes]);
                if(empty($_value)) return false;
                $input[$attributes] = Carbon::createFromFormat($fromFormat, $_value);
            }
        } catch (Exception $e) {
            logger($e);
        }
        return true;
    }
    /**
     * @param array &$input
     * @param string $attributes
     * @param string $fromFormat = 'd-m-Y H:i:s'
     * @param array $onFormat = ['00:00:01','23:59:59']
     * @return boolean
     */
    public static function castRangeDatetime(array &$input, string $attributes, string $fromFormat = 'd-m-Y H:i:s', $onFormat = ['00:00:01','23:59:59'])
    {
        try {
            if (!array_key_exists($attributes, $input) || empty($input[$attributes])) return false; 
                $_value = explode('-', $input[$attributes]);
                $onFormatFrom = ' '.trim($onFormat[0]);
                $onFormatTo = ' '.trim($onFormat[1]);
                $_value[0] = Carbon::createFromFormat($fromFormat, str_replace('/', '-', trim($_value[0]??'')).$onFormatFrom);
                $_value[1] = Carbon::createFromFormat($fromFormat, str_replace('/', '-', trim($_value[1]??'')).$onFormatTo);
                if(empty($_value)) return false;
                $input[$attributes] = $_value;
        } catch (Exception $e) {
            logger($e);
        }
        return true;
    }
    /**
    * cast:: set default value
    *
    * @param array &$input,
    * @param string $attributes,
    * @param mix $value,
    * @return boolean
    * @throws 
    */ 
    public function castDefault(array &$input, string $attributes, $value)
    {
        try {
            $input[$attributes] = $value;
        } catch (Exception $e) {
            logger($e);
        }
        return true;
    }
    /**
    * cast:: set default value if null
    *
    * @param array &$input,
    * @param string $attributes,
    * @param mix $value,
    * @return boolean
    * @throws 
    */ 
    public static function castDefaultIfNull(array &$input, string $attributes, $value)
    {
        try {
            if (!array_key_exists($attributes, $input) || is_null($input[$attributes])) {
                $input[$attributes] = $value;
            } else return false;
        } catch (Exception $e) {
            logger($e);
        }
        return true;
    }
    /**
    * cast:: implode 
    *
    * @param array &$input
    * @param string $attributes
    * @param string $glue
    * @return boolean
    *
    * @throws 
    */ 
    public static function castImplode(array &$input, string $attributes, string $glue = ',')
    {
        try {
            if (array_key_exists($attributes, $input)) {
                $_value = (array) $input[$attributes];
                $input[$attributes] = implode($glue, $_value);
            }
        } catch (Exception $e) {
            logger($e);
        }
        return true;
    }
    /**
    * cast attribute to JSON
    * @param array &$input
    * @param string $attribute
    * @return boolean
    *
    * @throws 
    */ 
    public static function castJSON(array &$input, string $attribute)
    {
        try {
            if (array_key_exists($attribute, $input)) {
                $_value = (array) $input[$attribute];
                $input[$attribute] = json_encode($_value);
            }
        } catch (Exception $e) {
            logger($e);
        }
        return true;
    }
    /**
    * cast attributes to string 
    * @param  array  &$attributes
    * @param string $attributes
    * @param string $pattern = '/\D/i'
    * @param string $replacement
    * @return boolean
    *
    * @throws 
    */ 
    public static function castTrim(array &$input, string $attributes, string $pattern = '/\D/i', string $replacement = '')
    {
        try {
            if (array_key_exists($attributes, $input)) {
                $subject = (string) $input[$attributes];
                $input[$attributes] = preg_replace($pattern, $replacement, $subject);
            }
        } catch (Exception $e) {
            logger($e);
        }
        return true;
    }
    /**
    * cast attributes to decimal money 
    * @param  array  &$attributes
    * @param string $attributes
    * @param string $pattern = '/\D/i'
    * @return boolean
    *
    * @throws 
    */ 
    public static function castMoney(array &$input, string $attributes, string $pattern = '/\D/i')
    {
        try {
            if (array_key_exists($attributes, $input)) {
                $subject = (string) $input[$attributes];
                $subject = preg_replace($pattern, '', $subject);
                $input[$attributes] = (int)$subject;
            }
        } catch (Exception $e) {
            logger($e);
        }
        return true;
    }
    /**
    * cast attributes to decimal money 
    * @param array &$input
    * @param string $attributes
    * @param string $slugSign
    * @return boolean
    * @throws 
    */ 
    public static function castSlug(array &$input, string $attributes, string $slugSign = '-')
    {
        try {
            if (array_key_exists($attributes, $input)) {
                $subject = (string) $input[$attributes];
                $subject = Str::of($subject)->slug($slugSign);
                $input[$attributes] = strtolower((string)$subject);
            }
        } catch (Exception $e) {
            logger($e);
        }
        return true;
    }
    /**
     * @param array &$input
     * @param string $attributes
     * @param string $formatRex = 'd-m-Y'
     * @param string $glue = ' - '
     * @return boolean
     */
    public static function rangeDatetimeToString(array &$input, string $attributes, string $formatRex = 'd-m-Y', string $glue = ' - ')
    {
        try {
            if (array_key_exists($attributes, $input)) {
                $_value = (array)($input[$attributes] ?? []);
                foreach ($_value as $_index => $_datetime) {
                    $_value[$_index] = $_datetime->format($formatRex);
                }
                return implode($glue, $_value ?? []);
            }
        } catch (Exception $e) {
            logger($e);
        }
        return '';
    }
    /**
     * [expect description]
     * @param  array  &$input [description]
     * @param  array  $keys   [description]
     * @return [type]         [description]
     */
    public function expect(array &$input, array $keys = [])
    {        
        return $input = array_intersect_key($input, array_flip($keys));
    }
    /**
     * [except description]
     * @param  array  &$input [description]
     * @param  array  $keys   [description]
     * @return [type]         [description]
     */
    public function except(array &$input, array $keys = [])
    {        
        return $input = array_diff_key($input, array_flip($keys));
    }
}
