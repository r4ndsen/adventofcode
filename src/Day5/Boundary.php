<?php

namespace Werner\Adventofcode\Day5;

class Boundary
{
    public function __construct(public int $dest, public int $source, public int $range)
    {
    }

    public static function fromStrings(string $dest, string $source, string $range): self
    {
        return new self(
            dest: (int) $dest,
            source: (int) $source,
            range: (int) $range
        );
    }

    public function hit(int $value): bool
    {
        return $value >= $this->source && $value < $this->source + $this->range;
    }

    public function convert(int $value): int
    {
        if (!$this->hit($value)) {
            return $value;
        }

        return $this->dest - $this->source + $value;
    }
}
