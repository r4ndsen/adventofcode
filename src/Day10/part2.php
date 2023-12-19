<?php

declare(strict_types=1);

namespace Werner\Adventofcode\Day9;

require __DIR__ . '/../support.php';

$input  = getPlainInputForDay(9);
$inputs = explode("\n", $input);

$result = 0;

foreach ($inputs as $input) {
    if ($input === '') {
        continue;
    }
    $reading = array_map('intval', explode(' ', $input));

    $sequence = [$reading];
    while (!isAllZeroes($reading)) {
        $reading    = diffs($reading);
        $sequence[] = $reading;
    }

    $extrapolated = extrapolate($sequence);

    echo sprintf('%s => %d', $input, $extrapolated) . PHP_EOL;

    $result += $extrapolated;
}

echo $result . PHP_EOL;

function extrapolate(array $sequence): int
{
    array_unshift($sequence[count($sequence) - 1], 0);
    for ($i = count($sequence) - 1; $i > 0; $i--) {
        array_unshift($sequence[$i - 1], $sequence[$i - 1][0] - $sequence[$i][0]);
    }

    return $sequence[0][0];
}

function isAllZeroes(array $in): bool
{
    foreach ($in as $item) {
        if ($item !== 0) {
            return false;
        }
    }

    return true;
}

function diffs(array $in): array
{
    $res = [];
    for ($i = 1, $iMax = count($in); $i < $iMax; $i++) {
        $res[] = $in[$i] - $in[$i - 1];
    }

    return $res;
}
