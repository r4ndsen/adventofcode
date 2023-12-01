<?php

declare(strict_types=1);

require __DIR__ . '/../support.php';

$lines = getInputForDay(1);

$result = 0;
foreach ($lines as $line) {
    [$first, $last] = firstAndLastDigit($line);
    $result += $first * 10 + $last;
}
echo sprintf("result: %u\n", $result);

function firstAndLastDigit(string $row): array
{
    $first = $last = null;

    $row = preg_replace('/[^0-9]/', '', $row);

    foreach (str_split($row) as $char) {
        if ($char > 0 && $char <= 9) {
            if ($first === null) {
                $first = $char;
            }

            $last = $char;
        }
    }

    return [$first, $last];
}
