<?php

declare(strict_types=1);

namespace Werner\Adventofcode\Day11;

require __DIR__ . '/../support.php';

$input = getPlainInputForDay(15);

// $input = 'rn=1,cm-,qp=3,cm=2,qp-,pc=4,ot=9,ab=5,pc-,pc=6,ot=7';

$sum = 0;

foreach (explode(',', trim($input)) as $line) {
    if ($line === '') {
        continue;
    }

    echo sprintf('%s = %u', $line, hash($line)) . PHP_EOL;

    $sum += hash($line);
}

echo $sum;

function hash(string $input): int
{
    $res = 0;
    foreach (str_split($input) as $char) {
        $res = ($res + asciiToInt($char)) * 17 % 256;
    }

    return $res;
}
