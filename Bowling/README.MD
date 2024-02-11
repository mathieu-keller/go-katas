# Bowling Game Kata

This repository contains a class for counting bowling games.

## Description

A bowling game consists of rolls to knock down 10 pins at the end of the bowling alley. With each roll, 0 to 10 pins can be knocked down. Per round, a player has a maximum of 2 rolls to knock down the 10 pins.

Rolls are grouped into frames, each assigned a score. A game consists of 10 rounds or frames.

How many rolls per round a player is allowed to make depends on how many pins are knocked down with them:

- If a player's first roll knocks down all 10 pins (strike), it is the only roll in the round.

- Otherwise, the frames always contain two rolls, each with 0 to 10 pins.

- Exception: If the 10th frame contains a strike on the first roll or both rolls together make a spare, a third roll can be made.

There are some rules for calculating the score of a frame:

- A frame whose two rolls together have knocked down a maximum of 9 pins receives the sum of the pins as a score.

- A frame whose two rolls together have knocked down 10 pins (spare) receives a score of 10 + the number of pins of the next roll.

- A frame with a strike receives a score of 10 + the sum of the pins of the next two rolls.

## Running the Kata

To run the Kata run the .go file. For example:

```bash
go run main.go
```