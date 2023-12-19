<?php

declare(strict_types=1);

namespace Werner\Adventofcode\Day7;

class Card
{
    public readonly int $value;

    public function __construct(string $card, public bool $isJoker = false)
    {
        $this->value = Deck::CARDS[$card];
    }

    public static function Ace(bool $isJoker): self
    {
        return new self('A', $isJoker);
    }
}
