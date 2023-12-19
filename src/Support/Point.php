<?php

namespace Werner\Adventofcode\Support;

use Stringable;

class Point implements Stringable
{
    public function __construct(public int $x, public int $y)
    {
    }

    public function __toString(): string
    {
        return sprintf('(%d, %d)', $this->x, $this->y);
    }

    public function distanceTo(Point $other): Point
    {
        return new Point($other->x - $this->x, $other->y - $this->y);
    }
}
