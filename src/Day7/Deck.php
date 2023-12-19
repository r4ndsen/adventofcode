<?php

namespace Werner\Adventofcode\Day7;

class Deck
{
    public const CARDS = [
        '2' => 2,
        '3' => 3,
        '4' => 4,
        '5' => 5,
        '6' => 6,
        '7' => 7,
        '8' => 8,
        '9' => 9,
        'T' => 10,
        'J' => 1,
        'Q' => 12,
        'K' => 13,
        'A' => 14,
    ];

    public Type $type;

    /** @var Card[] */
    private array $cards = [];

    private ?array $ocurrences = null;

    public function __construct(public readonly string $given, public readonly int $bid)
    {
        for ($i = 0, $iMax = strlen($given); $i < $iMax; $i++) {
            $this->cards[] = new Card($given[$i], $this->given[$i] === 'J');
        }

        $this->determineType();

        foreach ($this->cards as $i => $card) {
            if ($card->isJoker) {
                $this->cards[$i] = $this->switchJokerCard();
            }
        }

        $this->determineType();
    }

    public function __toString(): string
    {
        return $this->type->name . ': ' . $this->given;
    }

    public function cards(): array
    {
        return $this->cards;
    }

    private function determineType(): void
    {
        $this->ocurrences = null;

        $this->type = match (true) {
            $this->fiveOfAKind()  => Type::FifeOfAKind,
            $this->fourOfAKind()  => Type::FourOfAKind,
            $this->fullHouse()    => Type::FullHouse,
            $this->threeOfAKind() => Type::ThreeOfAKind,
            $this->twoPairs()     => Type::TwoPairs,
            $this->onePair()      => Type::OnePair,
            default               => Type::HighCard,
        };
    }

    private function switchJokerCard(): Card
    {
        $o = $this->occurrences();
        unset($o[1]); // no jokers

        $value = current(array_keys($o));
        $ident = array_search($value, self::CARDS, true);

        return $ident ? new Card($ident) : Card::Ace(true);
    }

    private function fiveOfAKind(): bool
    {
        return array_values($this->occurrences()) === [5];
    }

    private function fourOfAKind(): bool
    {
        return array_values($this->occurrences()) === [4, 1];
    }

    private function onePair(): bool
    {
        return array_values($this->occurrences()) === [2, 1, 1, 1];
    }

    private function fullHouse(): bool
    {
        return array_values($this->occurrences()) === [3, 2];
    }

    private function threeOfAKind(): bool
    {
        return array_values($this->occurrences()) === [3, 1, 1];
    }

    private function twoPairs(): bool
    {
        return array_values($this->occurrences()) === [2, 2, 1];
    }

    private function occurrences(): array
    {
        if ($this->ocurrences === null) {
            $this->ocurrences = [];
            foreach ($this->cards as $card) {
                @$this->ocurrences[$card->value]++;
            }

            arsort($this->ocurrences);
        }

        return $this->ocurrences;
    }
}
