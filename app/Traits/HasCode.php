<?php

namespace App\Traits;

use Illuminate\Support\Str;

trait HasCode
{
  public static function findCode($code)
  {
    return static::whereCode($code)->first();
  }

  public function codeNotNull()
  {
    if (empty($this->code)) {
      $this->code  = uniqid();
    }
  }

  public function codeReal()
  {
    if ($this->isDirty('code') && static::findCode($this->code)) {
      $this->code .= uniqid();
    }
  }
}
