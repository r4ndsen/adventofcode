<?php

declare(strict_types=1);

namespace Werner\Adventofcode\Day10;

require __DIR__ . '/../support.php';

$input = getPlainInputForDay(10);

$input = '.....
.F-7.
.|.|.
.L-J.
.....';

$inputs = explode("\n", $input);

$grid = [];

foreach ($inputs as $input) {
    if ($input === '') {
        continue;
    }

    $grid[] = str_split($input);
}
