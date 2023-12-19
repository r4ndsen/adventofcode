<?php

namespace Werner\Adventofcode\Day15;

use Stringable;

class Box implements Stringable
{
    private array $lenses = [];

    public function __toString(): string
    {
        if ($this->lenses === []) {
            return '[ ]';
        }

        return implode(' ', $this->lenses);
    }

    public function setLens(Lens $lens): void
    {
        $this->lenses[$lens->label] = $lens;
    }

    public function hasLens(Lens $lens): bool
    {
        return isset($this->lenses[$lens->label]);
    }

    public function removeLens(Lens $lens): void
    {
        unset($this->lenses[$lens->label]);
    }

    public function power(int $idx): int
    {
        if ($this->lenses === []) {
            return 0;
        }

        $res = 0;
        foreach (array_values($this->lenses) as $i => $lens) {
            $res += $idx * $lens->focalLength * ($i + 1);
        }

        return $res;
    }
}
