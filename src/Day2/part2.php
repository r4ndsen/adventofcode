<?php

declare(strict_types=1);

require __DIR__ . '/../support.php';

$lines = getInputForDay(2);

$input = 'Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green';

#$lines = explode("\n", $input);

$drawRe = '#(\d+) (red|blue|green)#';
$gameRe = '#^Game (\d+)#';

$result = 0;

foreach ($lines as $line) {
    preg_match($gameRe, $line, $m);
    $game = $m[1];

    $draws = explode(';', $line);

    $max = [
        'red' => 0,
        'green' => 0,
        'blue' => 0,
    ];

    foreach ($draws as $draw) {
        preg_match_all($drawRe, $draw, $m);
        $amount = $m[1];
        $colors = $m[2];

        foreach ($colors as $idx => $color) {
            if ($amount[$idx] > $max[$color]) {
                $max[$color] = $amount[$idx];
            }
        }
    }

    $result += array_product($max);
}

echo "result: $result\n";

