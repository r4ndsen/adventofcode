<?php

use Symfony\Component\Dotenv\Dotenv;
use Werner\Adventofcode\Support\Point;

require_once __DIR__ . '/../composer-vendor/autoload.php';

(new Dotenv())->load(__DIR__ . '/../.env.dist', __DIR__ . '/../.env');

function getInputForDay(int $day): array
{
    $path = sprintf('/tmp/adventofcode/2023/%d/input.txt', $day);
    @mkdir(dirname($path), 0777, true);

    if (!file_exists($path)) {
        $context = stream_context_create([
            'http' => [
                'method' => 'GET',
                'header' => 'Cookie: session=' . $_ENV['SESSION_COOKIE'] . "\r\n",
            ],
        ]);

        $content = file_get_contents(sprintf('https://adventofcode.com/2023/day/%u/input', $day), false, $context);

        file_put_contents($path, $content);
    }

    return file($path, FILE_IGNORE_NEW_LINES | FILE_SKIP_EMPTY_LINES);
}

/** return string[] */
function getPlainInputForDay(int $day): string
{
    $path = sprintf('/tmp/adventofcode/2023/%d/input.txt', $day);
    @mkdir(dirname($path), 0777, true);

    if (!file_exists($path)) {
        $context = stream_context_create([
            'http' => [
                'method' => 'GET',
                'header' => 'Cookie: session=' . $_ENV['SESSION_COOKIE'] . "\r\n",
            ],
        ]);

        $content = file_get_contents(sprintf('https://adventofcode.com/2023/day/%u/input', $day), false, $context);

        file_put_contents($path, $content);

        return $content;
    }

    return file_get_contents($path);
}

function manhattanDistance(Point $a, Point $b): int
{
    return abs($b->x - $a->x) + abs($b->y - $a->y);
}
/** @return string[] */
function toByteArray(string $input): array
{
    $res = [];

    for ($i = 0, $iMax = strlen($input); $i < $iMax; $i++) {
        $res[] = $input[$i];
    }

    return $res;
}

function asciiToInt(string $char): int
{
    return ord($char);
}
