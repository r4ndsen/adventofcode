<?php

declare(strict_types=1);

namespace Werner\Adventofcode\Day4;

require __DIR__ . '/../support.php';

$input = 'LLR

AAA = (BBB, BBB)
BBB = (ZZZ, ZZZ)
ZZZ = (ZZZ, ZZZ)';

/*$input = 'RL

AAA = (BBB, CCC)
BBB = (DDD, EEE)
CCC = (ZZZ, GGG)
DDD = (DDD, DDD)
EEE = (EEE, EEE)
GGG = (GGG, GGG)
ZZZ = (ZZZ, ZZZ)';*/

$input = getPlainInputForDay(8);

$inputs = explode("\n", $input);

$directions = toByteArray($inputs[0]);

$nodes = [];

for ($i = 2; $i < count($inputs); $i++) {
    preg_match('/^([A-Z]{3}) = \(([A-Z]{3}), ([A-Z]{3})\)$/', $inputs[$i], $matches);

    if (!$matches) {
        continue;
    }

    $nodes[$matches[1]] = [
        'L' => $matches[2],
        'R' => $matches[3],
    ];
}

$current = 'AAA';

$steps = 0;

while (true) {
    if ($current === 'ZZZ') {
        break;
    }

    $current = $nodes[$current][$directions[$steps % count($directions)]];

    $steps++;
}

echo $steps . PHP_EOL;
