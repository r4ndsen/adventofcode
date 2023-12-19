<?php

use Werner\Adventofcode\Day7\Deck;
use Werner\Adventofcode\Day7\Game;

require __DIR__ . '/../support.php';

$input = getPlainInputForDay(7);

/*
$input = '32T3K 765
T55J5 684
KK677 28
KTJJT 220
QQQJA 483';*/

$decks = [];
foreach (explode("\n", $input) as $line) {
    if ($line === '') {
        continue;
    }

    [$deck, $bid] = explode(' ', $line);

    $decks[] = new Deck($deck, (int) $bid);
}

$game = new Game(...$decks);
echo $game->play() . PHP_EOL;
