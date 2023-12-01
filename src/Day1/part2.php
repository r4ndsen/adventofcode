<?php

declare(strict_types=1);

require __DIR__ . '/../support.php';

$lines = getInputForDay(1);

$searchMap = [
    'one'   => 1,
    'two'   => 2,
    'three' => 3,
    'four'  => 4,
    'five'  => 5,
    'six'   => 6,
    'seven' => 7,
    'eight' => 8,
    'nine'  => 9,
];

foreach ($searchMap as $v) {
    $searchMap[$v] = $v;
}

$result = 0;
foreach ($lines as $line) {
    [$first, $last] = firstAndLastDigit($line, $searchMap);
    $result += $first * 10 + $last;
}
echo sprintf("result: %u\n", $result);

function firstAndLastDigit(string $row, array $searchMap): array
{
    $firstIdx = PHP_INT_MAX;
    $lastIdx  = -1;

    $first = $last = 0;

    foreach ($searchMap as $term => $value) {
        if (($idx = strpos($row, (string) $term)) !== false) {
            if ($idx < $firstIdx) {
                $firstIdx = $idx;
                $first    = $value;
            }

            if ($idx > $lastIdx) {
                $lastIdx = $idx;
                $last    = $value;
            }
        }

        if (($idx = strrpos($row, (string) $term)) !== false) {
            if ($idx > $lastIdx) {
                $lastIdx = $idx;
                $last    = $value;
            }
        }
    }

    return [$first, $last];
}
