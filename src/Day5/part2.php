<?php

use Werner\Adventofcode\Day5\Boundary;

require __DIR__ . '/../support.php';

$input = getPlainInputForDay(5);

/*
$input = 'seeds: 79 14 55 13

seed-to-soil map:
50 98 2
52 50 48

soil-to-fertilizer map:
0 15 37
37 52 2
39 0 15

fertilizer-to-water map:
49 53 8
0 11 42
42 0 7
57 7 4

water-to-light map:
88 18 7
18 25 70

light-to-temperature map:
45 77 23
81 45 19
68 64 13

temperature-to-humidity map:
0 69 1
1 0 69

humidity-to-location map:
60 56 37
56 93 4';
*/
$mapdata = explode("\n\n", $input);
preg_match('/seeds: (.*)/', $mapdata[0], $matches);
$seedInput = explode(' ', $matches[1]);

$seeds = [];
foreach ($seedInput as $idx => $seed) {
    if ($idx % 2 === 0) {
        $seeds[$seed] = $seedInput[$idx + 1] - 1;
    }
}

$fn = [];

for ($i = 1; $i < count($mapdata); ++$i) {
    $data    = explode("\n", $mapdata[$i]);
    $mapname = substr($data[0], 0, -5);

    /** @var Boundary[] $boundaries */
    $boundaries = [];
    foreach ($data as $idx => $datum) {
        if ($idx === 0) {
            continue;
        }

        if (!$datum) {
            continue;
        }

        [$dest, $source, $range] = explode(' ', $datum);

        $boundaries[] = new Boundary(...explode(' ', $datum));
    }

    $fn[$mapname] = function ($in) use ($boundaries) {
        foreach ($boundaries as $boundary) {
            if ($boundary->hit($in)) {
                return $boundary->convert($in);
            }
        }

        return $in;
    };
}

$min = PHP_INT_MAX;

$iteration = 0;
foreach ($seeds as $seed => $amount) {
    for ($i = 0; $i < $amount; ++$i) {
        ++$iteration;

        if ($iteration % 100000 === 0) {
            echo $iteration . PHP_EOL;
        }
        $test = $seed + $amount;

        $soil        = $fn['seed-to-soil']($test);
        $fertilizer  = $fn['soil-to-fertilizer']($soil);
        $water       = $fn['fertilizer-to-water']($fertilizer);
        $light       = $fn['water-to-light']($water);
        $temperature = $fn['light-to-temperature']($light);
        $humidity    = $fn['temperature-to-humidity']($temperature);
        $location    = $fn['humidity-to-location']($humidity);

        if ($location < $min) {
            $min = $location;
        }
    }
}

print_r($min);
