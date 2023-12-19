<?php

declare(strict_types=1);

namespace Werner\Adventofcode\Day11;

use Werner\Adventofcode\Support\Point;

require __DIR__ . '/../support.php';

$input = getPlainInputForDay(11);

/*$input = '...#......
.......#..
#.........
..........
......#...
.#........
.........#
..........
.......#..
#...#.....';
*/
$inputs = explode("\n", $input);

$grid = [];
foreach ($inputs as $input) {
    if ($input === '') {
        continue;
    }

    $grid[] = str_split($input);
}

function findEmptyColumns(array $grid): array
{
    $columns = array_fill(0, count($grid[0]), true);

    foreach ($grid as $row) {
        foreach (array_keys($columns) as $columnIndex) {
            if ($row[$columnIndex] === '#') {
                unset($columns[$columnIndex]);
            }
        }
    }

    return array_keys($columns);
}

function findEmptyRows(array $grid): array
{
    foreach ($grid as $idx => $row) {
        if (str_contains(implode('', $row), '#')) {
            unset($grid[$idx]);
        }
    }

    return array_keys($grid);
}

$xGaps = findEmptyColumns($grid);
$yGaps = findEmptyRows($grid);

function gapsCrossed(Point $a, Point $b, array $xGaps, array $yGaps): int
{
    $crossed = 0;

    $xmin = min($a->x, $b->x);
    $ymin = min($a->y, $b->y);
    $xmax = max($a->x, $b->x);
    $ymax = max($a->y, $b->y);

    foreach ($xGaps as $xGap) {
        if ($xmin < $xGap && $xmax > $xGap) {
            $crossed++;
        }
    }

    foreach ($yGaps as $yGap) {
        if ($ymin < $yGap && $ymax > $yGap) {
            $crossed++;
        }
    }

    return $crossed;
}

$g        = new Grid($grid);
$galaxies = $g->getGalaxies();

const SCALING = 1_000_000;

$distances = [];
foreach ($galaxies as $idx => $galaxy) {
    foreach ($galaxies as $oidx => $otherGalaxy) {
        if ($oidx <= $idx) {
            continue;
        }

        $md   = manhattanDistance($galaxy, $otherGalaxy);
        $gaps = gapsCrossed($galaxy, $otherGalaxy, $xGaps, $yGaps);

        $distances[] = $md - $gaps + SCALING * $gaps;
    }
}

print_r(array_sum($distances));
