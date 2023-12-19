<?php

declare(strict_types=1);

namespace Werner\Adventofcode\Day4;

require __DIR__ . '/../support.php';

$input = 'LR

11A = (11B, XXX)
11B = (XXX, 11Z)
11Z = (11B, XXX)
22A = (22B, XXX)
22B = (22C, 22C)
22C = (22Z, 22Z)
22Z = (22B, 22B)
XXX = (XXX, XXX)';

$input = getPlainInputForDay(8);

$inputs = explode("\n", $input);

$directions = toByteArray($inputs[0]);

$nodes          = [];
$startingNodes  = [];
$cycleDurations = [];

for ($i = 2, $iMax = count($inputs); $i < $iMax; $i++) {
    preg_match('/^([A-Z\d]{3}) = \(([A-Z\d]{3}), ([A-Z\d]{3})\)$/', $inputs[$i], $matches);

    if (!$matches) {
        continue;
    }

    $nodes[$matches[1]] = [
        'L'         => $matches[2],
        'R'         => $matches[3],
        'isEndNode' => str_ends_with($matches[1], 'Z'),
    ];

    if (str_ends_with($matches[1], 'A')) {
        $startingNodes[] = $matches[1];
    }
}

function getCycleDuration(string $startNode, array $nodes, array $directions): int
{
    $current = $startNode;

    $steps = 0;
    while (!$nodes[$current]['isEndNode']) {
        $current = $nodes[$current][$directions[$steps % count($directions)]];
        $steps++;
    }

    return $steps;
}

foreach ($startingNodes as $node) {
    $cycleDurations[$node] = getCycleDuration($node, $nodes, $directions);
}

function biggestCommonDenominator(int $a, int $b)
{
    if ($a === 0) {
        return $b;
    }

    while ($b !== 0) {
        if ($a > $b) {
            $a = $a - $b;
        } else {
            $b = $b - $a;
        }
    }

    return $a;
}

function lowestCommonMultiple(int $a, int $b)
{
    return ($a * $b) / biggestCommonDenominator($a, $b);
}

$lcm = 1;
foreach ($cycleDurations as $cycleDuration) {
    $lcm = lowestCommonMultiple($lcm, $cycleDuration);
}

echo $lcm;
