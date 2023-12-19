<?php

declare(strict_types=1);

namespace Werner\Adventofcode\Day14;

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
#OO..#....';*/

$grid = [];
foreach (explode("\n", $input) as $row) {
    if ($row === '') {
        continue;
    }
    $grid[] = str_split($row);
}

// echo new Grid($grid) . PHP_EOL;

function cycle(Grid $grid)
{
    $grid
        ->moveStonesNorth()
        ->moveStonesWest()
        ->moveStonesSouth()
        ->moveStonesEast()
    ;
}

$configs = [];
$g       = new Grid($grid);

$offset = 0;
$repeat = 0;

for ($i = 0; $i < 1000; $i++) {
    cycle($g);

    $hash = md5((string) $g);

    echo "$i $hash\n";

    if (isset($configs[$hash])) {
        $repeat = $configs[$hash] - $i;
        $offset = $i;
        echo 'Found cycle at ' . $i . ' ' . $configs[$hash] . "\n";
        break;
    }

    $configs[$hash] = $i;
}

// echo $g . PHP_EOL;

/*
$height = count($grid);
$total  = 0;
foreach ($grid as $i => $row) {
    $total += ($height - $i) * substr_count(implode('', $row), 'O');
}

echo "Total: $total\n";
*/
