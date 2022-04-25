<?php

namespace App\Http\Controllers;
use Auth;
use \Exception;

use Maatwebsite\Excel\Concerns\FromCollection;
use Maatwebsite\Excel\Concerns\FromArray;
use Maatwebsite\Excel\Concerns\ShouldAutoSize;
use Maatwebsite\Excel\Concerns\WithEvents;
use Maatwebsite\Excel\Events\AfterSheet;
use Maatwebsite\Excel\Concerns\{WithHeadings, WithMapping, WithTitle};

class ExportExcelController implements FromCollection, WithHeadings, WithMapping, WithTitle, ShouldAutoSize, WithEvents
{
    /**
     * @author toannguyen.dev
     * @todo export excel from contract
     * @version 1.0
     * @return Excel file
     * #date 2021.08.07
     **/
    public $title;
    protected $collection;
    protected $customHeading;
    protected $countRow;

    public function __construct($collection, $customHeading = []){
        $this->collection = $collection;
        $this->customHeading = $customHeading;
        $this->countRow = 1;
    }
    /**
     * @todo set data
     * @return 
     */
    public function collection()
    {
        return $this->collection;
    }
    /**
     * @todo set title
     * @return string
     */
    public function title(): string
    {
        return $this->title;
    }
    /**
     * @todo mapping data
     * @return array
     */
    public function map($object): array
    {
        $countRow = $this->countRow++;
        return array_map(
            function($attribute) use($object, $countRow){
                $value = $object->{$attribute} ?? $object[$attribute] ?? null;
                if (preg_match('/^(no|stt)$/', strtolower($attribute))) return $countRow;
                if (preg_match('/^(type)$/', strtolower($attribute)) && is_array((array)$value)) return $value['name'] ?? '';
                if (preg_match('/^(status)$/', strtolower($attribute)) && is_array((array)$value)) return $value['name'] ?? '';
                if (preg_match('/(-link)$/', strtolower($attribute))) return $value['label'] ?? '';
                return $value;
            },
            array_flip($this->customHeading)
        );
    }
    /**
     * @todo set heading
     * @return array
     */
    public function headings(): array
    {
        return ($this->customHeading);
    }
    /**
     * @todo formating
     * @return array
     */
    public function registerEvents(): array
    {
        return [
            AfterSheet::class => function(AfterSheet $event) {
                $cellRange = 'A1:AZ1';
                $font = $event->sheet->getDelegate()->getStyle($cellRange)->getFont();
                $font->setSize(14)->setBold(1);
            },
        ];
    }
}
