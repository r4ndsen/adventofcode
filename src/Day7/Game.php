<?php

namespace Werner\Adventofcode\Day7;

class Game
{
    /** @var Deck[] */
    private array $decks;

    public function __construct(Deck ...$decks)
    {
        $this->decks = $decks;
    }

    public function play(): int
    {
        usort($this->decks, static function (Deck $a, Deck $b): int {
            $cmpResult = $a->type->value <=> $b->type->value;
            if ($cmpResult !== 0) {
                return $cmpResult;
            }

            $bCards = $b->cards();

            foreach ($a->cards() as $i => $aCard) {
                if ($aCard->isJoker) {
                    if ($bCards[$i]->isJoker) {
                        continue;
                    }

                    return -1;
                }

                if ($bCards[$i]->isJoker) {
                    return 1;
                }

                if (($res = ($aCard <=> $bCards[$i])) !== 0) {
                    return $res;
                }
            }

            return 0;
        });

        $sum = 0;
        foreach ($this->decks as $idx => $deck) {
            $score = $deck->bid * ($idx + 1);

            echo sprintf('deck %s score = %u * %u = %u', $deck, $idx + 1, $deck->bid, $score) . PHP_EOL;
            $sum += $score;
        }

        return $sum;
    }
}
