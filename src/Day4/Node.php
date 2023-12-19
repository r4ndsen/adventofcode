<?php

namespace Werner\Adventofcode\Day4;

final class Node
{
    public bool $visited        = false;
    public ?Node $previousNode  = null;
    private static array $nodes = [];

    public function __construct(
        public readonly string $name,
        public ?Node $l = null,
        public ?Node $r = null,
    ) {
        self::$nodes[$name] = $this;
    }

    public static function get(string $name): self
    {
        return self::$nodes[$name] ?? new self($name);
    }

    public function visited(): self
    {
        $this->visited = true;

        return $this;
    }

    public function previous(): ?Node
    {
        return $this->previousNode;
    }
}
