# Snakes and Ladders

A snakes and ladders game.

At the moment the game just autoplays, but it could be interesting as an API
for a front end.


## Install

    go get install github.com/scottjbarr/snakesandladders/cmd/snakesandladders


## Example

    snakesandladders -p1 Bob -p2 Jane
    Bob rolled a 3 and moves from 0 to 3
      Bob landed on a Ladder and finished on 75
    Jane rolled a 12 and moves from 0 to 12
    Bob rolled a 1 and moves from 75 to 76
    Jane rolled a 2 and moves from 12 to 14
    Bob rolled a 5 and moves from 76 to 81
    Jane rolled a 8 and moves from 14 to 22
    Bob rolled a 3 and moves from 81 to 84
    Jane rolled a 5 and moves from 22 to 27
    Bob rolled a 10 and moves from 84 to 94
    Jane rolled a 8 and moves from 27 to 35
    Bob rolled a 3 and moves from 94 to 97
    Jane rolled a 11 and moves from 35 to 46
    Bob rolled a 2 and moves from 97 to 99
    Jane rolled a 8 and moves from 46 to 54
    Bob rolled a 11 and moves from 99 to 110
    Bob is the winner


## Licence

The MIT License (MIT)

Copyright (c) 2016 Scott Barr

See [LICENSE.md](LICENSE.md)
