<?php

declare(strict_types=1);

namespace Werner\Adventofcode\Day14;

use Werner\Adventofcode\Support\Grid;

require __DIR__ . '/../support.php';

$input = getPlainInputForDay(14);

/*$input = 'O....#....
O.OO#....#
.....##...
OO.#O....O
.O.....O#.
O.#..O.#.#
..O..#O..O
.......O..
#....###..
#OO..#....';
*/
$grid = [];
foreach (explode("\n", $input) as $row) {
    if ($row === '') {
        continue;
    }
    $grid[] = str_split($row);
}

// echo new Grid($grid) . PHP_EOL;

function moveStonesNorth(array $grid): array
{
    foreach ($grid as $i => $row) {
        foreach ($row as $j => $cell) {
            if ($cell === 'O') {
                $offset = $i;
                while (($grid[$offset - 1][$j] ?? null) === '.') {
                    $grid[$offset - 1][$j] = 'O';
                    $grid[$offset][$j]     = '.';
                    $offset--;
                }
            }
        }
    }

    return $grid;
}

$grid = moveStonesNorth($grid);

echo new Grid($grid);

$height = count($grid);
$total  = 0;
foreach ($grid as $i => $row) {
    $total += ($height - $i) * substr_count(implode('', $row), 'O');
}

echo "Total: $total\n";
