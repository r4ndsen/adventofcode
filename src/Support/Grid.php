<?php

namespace Werner\Adventofcode\Support;

use Stringable;

class Grid implements Stringable
{
    public function __construct(public array $grid)
    {
    }

    public function __toString(): string
    {
        $buffer = '';

        foreach ($this->grid as $row) {
            $buffer .= implode('', $row) . "\n";
        }

        return $buffer;
    }

    public function rotateRight(): self
    {
        $grid = $this->grid;

        // Get the number of rows and columns in the array
        $rows    = count($grid);
        $columns = count($grid[0]);

        // Create a new array with the dimensions swapped
        $rotatedArray = array_fill(0, $columns, array_fill(0, $rows, null));

        // Fill the new array with rotated values
        for ($i = 0; $i < $rows; $i++) {
            for ($j = 0; $j < $columns; $j++) {
                $rotatedArray[$j][$rows - $i - 1] = $grid[$i][$j];
            }
        }

        return new self($rotatedArray);
    }

    public function rotateLeft(): self
    {
        $grid = $this->grid;

        $rows    = count($grid);
        $columns = count($grid[0]);

        $rotatedArray = array_fill(0, $columns, array_fill(0, $rows, null));

        for ($i = 0; $i < $rows; $i++) {
            for ($j = 0; $j < $columns; $j++) {
                $rotatedArray[$columns - $j - 1][$i] = $grid[$i][$j];
            }
        }

        return new self($rotatedArray);
    }
}
