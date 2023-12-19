<?php

namespace Werner\Adventofcode\Day11;

use Werner\Adventofcode\Support\Point;

class Grid extends \Werner\Adventofcode\Support\Grid
{
    /** @return Point[] */
    public function getGalaxies(): array
    {
        $res = [];

        foreach ($this->grid as $rIdx => $row) {
            foreach ($row as $cIdx => $column) {
                if ($column === '#') {
                    $res[] = new Point($cIdx, $rIdx);
                }
            }
        }

        return $res;
    }
}
