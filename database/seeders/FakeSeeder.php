<?php

namespace Database\Seeders;

use Illuminate\Database\Seeder;
use Illuminate\Support\Str;
use Illuminate\Support\Facades\DB;
use App\Models\{Status, Type};
use App\Models\{Animal,};
use App\Traits\{Random};

class FakeSeeder extends Seeder
{
    /**
     * Run the database seeds.
     *
     * @return void
     */
    public function run()
    {
        // $this->scripts_once();
    }
    /**
     * only run once 
     * 
     */
    public function scripts_once()
    {
        self::fakeAnimalList();
    }
    /**
     * 
     * 
     */
    public function fakeAnimalList()
    {
        
        $role = config('roles.models.role')::where('name', '=', 'animal')->first();
        $list = ['lion', 'cat', 'tiger','dolphin'];
        foreach ($list as $key => $animalName) {
            $code = $animalName . str_pad(random_int(1, 999), 3, 0, STR_PAD_LEFT);
            $codeSlugSpace = Str::slug($code, '');

            $user = config('roles.models.defaultUser')::create([
                'code' => Str::slug($code),
                'name' => $codeSlugSpace,
                'email'=> $codeSlugSpace,
                'password' => bcrypt($codeSlugSpace.'@123'),
                'queue_limit' =>   rand(5, 10) *10,
                'created_at' => '2022-04-01 09:00:00',
            ]);
            $user->attachRole($role);

            /*fake animal*/
            $animal = Animal::create([
                'code'      => $codeSlugSpace,
                'name'      => $code,
                'user_id'   => $user->id,
                'created_at' => '2022-04-01 09:00:00',
            ]);
        }
    }
}
