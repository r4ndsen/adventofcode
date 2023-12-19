<?php

namespace Werner\Adventofcode\Day14;

class Grid extends \Werner\Adventofcode\Support\Grid
{
    public function moveStonesNorth(): self
    {
        foreach ($this->grid as $i => $row) {
            foreach ($row as $j => $cell) {
                if ($cell === 'O') {
                    $offset = $i;
                    while (($this->grid[$offset - 1][$j] ?? null) === '.') {
                        $this->grid[$offset - 1][$j] = 'O';
                        $this->grid[$offset][$j]     = '.';
                        $offset--;
                    }
                }
            }
        }

        return $this;
    }

    public function moveStonesSouth(): self
    {
        for ($i = count($this->grid) - 1; $i >= 0; $i--) {
            foreach ($this->grid[$i] as $j => $cell) {
                if ($cell === 'O') {
                    $offset = $i;
                    while (($this->grid[$offset + 1][$j] ?? null) === '.') {
                        $this->grid[$offset + 1][$j] = 'O';
                        $this->grid[$offset][$j]     = '.';
                        $offset++;
                    }
                }
            }
        }

        return $this;
    }

    public function moveStonesEast(): self
    {
        foreach ($this->grid as $i => $row) {
            for ($j = count($row) - 1; $j >= 0; $j--) {
                $cell = $this->grid[$i][$j];

                if ($cell === 'O') {
                    $offset = $j;
                    while (($this->grid[$i][$offset + 1] ?? null) === '.') {
                        $this->grid[$i][$offset + 1] = 'O';
                        $this->grid[$i][$offset]     = '.';
                        $offset++;
                    }
                }
            }
        }

        return $this;
    }

    public function moveStonesWest(): self
    {
        foreach ($this->grid as $i => $row) {
            foreach ($row as $j => $cell) {
                if ($cell === 'O') {
                    $offset = $j;
                    while (($this->grid[$i][$offset - 1] ?? null) === '.') {
                        $this->grid[$i][$offset - 1] = 'O';
                        $this->grid[$i][$offset]     = '.';
                        $offset--;
                    }
                }
            }
        }

        return $this;
    }
}
