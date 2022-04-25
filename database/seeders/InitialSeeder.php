<?php

namespace Database\Seeders;

use Illuminate\Database\Seeder;
use Illuminate\Support\Str;
use App\Models\{Warehouse, Status, Type, WarehouseType, SupplierType, GoodsNote, GoodsNoteType};
use App\Models\{Package, PackageMaterialStandard, PackageDimensionStandard};
use App\Models\{Category, Item, Brand};
use App\Models\{Customer, Transport};

class InitialSeeder extends Seeder
{
    /**
     * Run the database seeds.
     *
     * @return void
     */
    public function run()
    {
        $this->statuses();
        $this->types();
        $this->roles();

        $this->users();
    }
    /**
     * @todo only for the first initial and setup. To create some accounts by role
     * 
     * */
    public function users()
    {
        $adminRole = config('roles.models.role')::where('name', '=', 'elf')->first();
        $newUser = config('roles.models.defaultUser')::create([
            'code'      => 'groot',
            'name'      => 'groot',
            'email'     => 'groot@wildlife.com',
            'queue_limit' =>   rand(5, 10) *10,
            'password'  => bcrypt('groot@123'),
        ]);
        
        $newUser->attachRole($adminRole);
    }
    /**
     * [roles description]
     * @return [type] [description]
     */
    public function roles()
    {
        /*
         * Role Types
         *
         */
        $RoleItems = [
            [
                'name'          => 'elf',
                'slug'          => 'elf',
                'description'   => 'elf',
                'level'         => 5,
            ],
            [
                'name'          => 'animal',
                'slug'          => 'animal',
                'description'   => 'animal',
                'level'         => 1,
            ]
        ];

        /*
         * Add Role Items
         *
         */
        foreach ($RoleItems as $RoleItem) {
            $newRoleItem = config('roles.models.role')::where('slug', '=', $RoleItem['slug'])->first();
            if ($newRoleItem === null) {
                $newRoleItem = config('roles.models.role')::create([
                    'name'         => $RoleItem['name'],
                    'slug'         => $RoleItem['slug'],
                    'description'  => $RoleItem['description'],
                    'level'        => $RoleItem['level'],
                ]);
            }
        }
    }
    /**
     * @author photrucco
     * @todo
     * 
     * */
    public function statuses()
    {
        /*for animal*/
        $prefix = 'animal';
        Status::create(['code' => 'living', 'name' =>'đang sống', 'prefix' => $prefix]);
        Status::create(['code' => 'dead', 'name' =>'đã chết', 'prefix' => $prefix]);

        /*for food*/
        $prefix = 'food';
    }
    /**
     * @author photrucco
     * @todo
     * 
     * */
    public function types()
    {
        /*for elf*/
        $prefix = 'elf';
        Type::create(['code' => 'caption','name' =>'Đội trưởng', 'prefix'=>$prefix]);
        Type::create(['code' => 'guard','name' =>'Thủ vệ', 'prefix'=>$prefix]);

        /*for animal*/
        $prefix = 'animal';
        Type::create(['code' => 'wild','name' =>'Hoang dã', 'prefix'=>$prefix]);
        Type::create(['code' => 'pet','name' =>'Thuần hóa', 'prefix'=>$prefix]);        

        /*for food*/
        $prefix = 'food';
        Type::create(['code' => 'meet','name' =>'Thịt sống', 'prefix'=>$prefix]);
        Type::create(['code' => 'vegetable','name' =>'Thực vật', 'prefix'=>$prefix]);
        Type::create(['code' => 'gobble','name' =>'Hỗn tạp', 'prefix'=>$prefix]);
    }    
}
