<?php

declare(strict_types=1);

namespace Werner\Adventofcode\Day11;

use Werner\Adventofcode\Day15\Box;
use Werner\Adventofcode\Day15\Lens;

require __DIR__ . '/../support.php';

$input = getPlainInputForDay(15);

// $input = 'rn=1,cm-,qp=3,cm=2,qp-,pc=4,ot=9,ab=5,pc-,pc=6,ot=7';

/** @var Box[] $boxes */
$boxes = [];

foreach (range(0, 255) as $i) {
    $boxes[$i] = new Box();
}

foreach (explode(',', trim($input)) as $line) {
    if (trim($line) === '') {
        continue;
    }

    if (str_contains($line, '=')) {
        [$name, $val] = explode('=', $line);
        $l            = new Lens($name, (int) $val);

        $hash = hash($name);

        $boxes[$hash]->setLens($l);
    } else {
        [$name] = explode('-', $line);
        $l      = new Lens($name);
        $hash   = hash($name);

        $boxes[$hash]->removeLens($l);
    }
}

function hash(string $input): int
{
    $res = 0;
    foreach (str_split($input) as $char) {
        $res = ($res + asciiToInt($char)) * 17 % 256;
    }

    return $res;
}

$sum = 0;
foreach ($boxes as $i => $box) {
    $sum += $box->power($i + 1);
}

echo "sum $sum\n";
