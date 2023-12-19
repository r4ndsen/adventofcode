<?php

declare(strict_types=1);

namespace Werner\Adventofcode\Day11;

require __DIR__ . '/../support.php';

$input = getPlainInputForDay(13);

/*$input = '#.##..##.
..#.##.#.
##......#
##......#
..#.##.#.
..##..##.
#.#.##.#.

#...##..#
#....#..#
..##..###
#####.##.
#####.##.
..##..###
#....#..#';
*/

/*
$input = '...##...#.#####
.######...#####
#......##...#.#
..#..#..#######
#..###.#.....##
...##..........
...##...####..#
########....##.
##.##.##..##..#
.##..##.##....#
#.#..#.##.#..#.
.#.##.#...####.
.#.##.#...####.';
*/
$grids = explode("\n\n", $input);

$res = 0;

foreach ($grids as $inputs) {
    $grid = [];
    foreach (explode("\n", $inputs) as $input) {
        if ($input === '') {
            continue;
        }

        $grid[] = str_split($input);
    }

    $h = horizontalMirrorValue($grid);

    if ($h) {
        $ng           = $grid;
        $ng[$h - 1][] = 'v';
        $ng[$h][]     = '^';
    // echo "$h:\n" . new Grid($ng) . PHP_EOL . PHP_EOL;
    } else {
        // echo "no horizontal mirror found:\n" . new Grid($ng) . PHP_EOL . PHP_EOL;
    }

    $flip = flipRight($grid);
    $v    = horizontalMirrorValue($flip);

    if (!$v && !$h) {
        echo "no mirrors found:\n" . new Grid($grid) . "\n\n" . new Grid($flip) . PHP_EOL . PHP_EOL;
    }

    $res += ($v + 100 * $h);
}

echo "res: $res\n";

function horizontalMirrorValue(array $grid): int
{
    for ($i = 1; $i < count($grid) - 1; $i++) {
        for ($offset = 0; $offset < count($grid) / 2; $offset++) {
            $lineup   = $grid[$i - $offset - 1] ?? null;
            $linedown = $grid[$i + $offset] ?? null;

            if ($lineup === null || $linedown === null) {
                return $i;
            }

            if ($lineup !== $linedown) {
                continue 2;
            }
        }

        return $i;
    }

    return 0;
}

function flipLeft(array $grid): array
{
    $rows    = count($grid);
    $columns = count($grid[0]);

    $rotatedArray = array_fill(0, $columns, array_fill(0, $rows, null));

    for ($i = 0; $i < $rows; $i++) {
        for ($j = 0; $j < $columns; $j++) {
            $rotatedArray[$columns - $j - 1][$i] = $grid[$i][$j];
        }
    }

    return $rotatedArray;
}

function flipRight(array $grid): array
{
    // Get the number of rows and columns in the array
    $rows    = count($grid);
    $columns = count($grid[0]);

    // Create a new array with the dimensions swapped
    $rotatedArray = array_fill(0, $columns, array_fill(0, $rows, null));

    // Fill the new array with rotated values
    for ($i = 0; $i < $rows; $i++) {
        for ($j = 0; $j < $columns; $j++) {
            $rotatedArray[$j][$rows - $i - 1] = $grid[$i][$j];
        }
    }

    return $rotatedArray;
}

// echo getVerticalMirrorValue($grid) . PHP_EOL;
