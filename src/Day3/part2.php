<?php

declare(strict_types=1);

require __DIR__ . '/../support.php';

$lines = getInputForDay(3);

$input = '467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..';

$grid = [];
foreach ($lines as $idx => $line) {
    $len = strlen($line);
    for ($columnIdx = 0; $columnIdx < $len; ++$columnIdx) {
        $grid[$idx][] = $line[$columnIdx];
    }
}

function getSurroundingParts(array $grid, int $lineIdx, int $columnIdx): array
{
    $rows    = count($grid);
    $columns = count($grid[0]);

    $numbersFound = [];

    for ($col = max(0, $columnIdx - 1); $col <= min($columns, $columnIdx + 1); ++$col) {
        for ($row = max(0, $lineIdx - 1); $row <= min($rows, $lineIdx + 1); ++$row) {
            if (!isNumber($grid, $row, $col)) {
                continue;
            }

            $found          = getNumberAt($grid, $row, $col);
            $numbersFound[] = $found;
        }
    }

    return $numbersFound;
}

function getNumberAt(array &$grid, mixed $row, mixed $col): int
{
    $colIdx = $col;

    while (isNumber($grid, $row, $colIdx - 1)) {
        --$colIdx;
    }

    $num = $grid[$row][$colIdx];

    $grid[$row][$colIdx] = '.';

    while (isNumber($grid, $row, $colIdx + 1)) {
        $num .= $grid[$row][$colIdx + 1];
        $grid[$row][$colIdx + 1] = '.';
        ++$colIdx;
    }

    return (int) $num;
}

function isNumber(array $grid, mixed $row, mixed $col): bool
{
    return isset($grid[$row][$col]) && $grid[$row][$col] >= 0 && $grid[$row][$col] <= 9;
}

function isGear(array $grid, mixed $row, mixed $col): bool
{
    return ($grid[$row][$col] ?? '') === '*';
}

$result = 0;

foreach ($grid as $lineIdx => $line) {
    foreach ($line as $columnIdx => $value) {
        if (!preg_match('#[^.\d]#', $value, $symbols)) {
            continue;
        }

        $parts = getSurroundingParts($grid, $lineIdx, $columnIdx);

        if (isGear($grid, $lineIdx, $columnIdx) && count($parts) === 2) {
            $result += array_product($parts);
        }
    }
}

echo "result: $result\n";
