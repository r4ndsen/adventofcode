<?php

namespace Werner\Adventofcode\Day15;

use Stringable;

class Lens implements Stringable
{
    public function __construct(public readonly string $label, public ?int $focalLength = null)
    {
    }

    public function __toString(): string
    {
        return sprintf('[%s %u]', $this->label, $this->focalLength);
    }
}
