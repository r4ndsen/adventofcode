<?php

declare(strict_types=1);

namespace Werner\Adventofcode\Day4;

require __DIR__ . '/../support.php';

ini_set('memory_limit', '2G');

$input = getPlainInputForDay(8);

/*$input = 'LLR

AAA = (BBB, BBB)
BBB = (AAA, ZZZ)
ZZZ = (ZZZ, ZZZ)';
*/
$directions = null;
$nodes      = [];

foreach (explode("\n", $input) as $line) {
    if ($line === '') {
        continue;
    }

    if ($directions === null) {
        $directions = toByteArray($line);
        foreach ($directions as $key => $decision) {
            $directions[$key] = match ($decision) {
                'L' => Direction::L,
                'R' => Direction::R,
            };
        }
        continue;
    }

    preg_match('/^([A-Z]{3}) = \(([A-Z]{3}), ([A-Z]{3})\)$/', $line, $matches);

    $n    = Node::get($matches[1]);
    $n->l = Node::get($matches[2]);
    $n->r = Node::get($matches[3]);

    $nodes[] = $n;
}

print_r($nodes);
exit;

$map = new Map($nodes, $directions);
echo $map->navigate() . PHP_EOL;
