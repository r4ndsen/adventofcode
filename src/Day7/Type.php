<?php

namespace Werner\Adventofcode\Day7;

enum Type: int
{
    case HighCard     = 1;
    case OnePair      = 2;
    case TwoPairs     = 3;
    case ThreeOfAKind = 4;
    case FullHouse    = 5;
    case FourOfAKind  = 6;
    case FifeOfAKind  = 7;
}
