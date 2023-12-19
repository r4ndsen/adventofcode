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
    $add = 0;

    while (count($sequence) > 0) {
        $last = array_pop($sequence);
        $add += end($last);
    }

    return $add;
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
