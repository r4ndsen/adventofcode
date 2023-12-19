<?php

namespace Werner\Adventofcode\Day4;

class Map
{
    private array $map = [];
    private Node $lastNode;
    private Node $currentNode;

    public function __construct(
        /** @var Node[] */
        public readonly array $nodes,
        /** @var Direction[] */
        public readonly array $directions,
    ) {
        foreach ($this->nodes as $node) {
            $this->map[$node->name] = $node;
        }

        $this->currentNode = $this->map['AAA'];
        $this->lastNode    = $this->currentNode;
    }

    public function navigate(): int
    {
        $visited = [];

        foreach ($this->directions as $step => $leftOrRight) {
            $this->move($leftOrRight);

            if ($this->current->name === 'ZZZ') {
                break;
            }
        }

        echo "Ran out of instructions, alternating R/L\n";

        $lastDirection = Direction::L;
        $lastNode      = $current;

        while ($current->name !== 'ZZZ') {
            $nextNodeName = $this->map[$current]->r;
            if ($lastDirection === Direction::L) {
                $nextNodeName = $this->map[$current]->l;

                $current       = $this->move($current, Direction::R);
                $lastDirection = Direction::R;
            } else {
                $current       = $this->move($current, Direction::L);
                $lastDirection = Direction::L;
            }

            if (isset($visited[$nextNodeName])) {
                echo "Already visited $nextNodeName, going back\n";
                $current = $this->move($current, $lastDirection);
                continue;
            }

            $step++;
        }

        return $step + 1;
    }

    private function move(Node $n, Direction $d): Node
    {
        if ($d === Direction::L) {
            $newNodeName = $this->map[$n->name]->l;
        } else {
            $newNodeName = $this->map[$n->name]->r;
        }

        return $this->map[$newNodeName]->visited();
    }
}
