# Game-of-Life

This is an implementation of Conway’s Game of Life, which I wrote to familiarize myself with Go.

Rules:
1) A living cell with 2 or 3 neighboring living cells survives
2) A living cell with fewer than 2 living neighbors dies of loneliness
3) A living cell with 3 more neighbors dies from overcrowding
4) A dead cell with 3 neighbors comes to life

Input:
This program gets its initial state from an input file where dead cells are indicated by dashes and live cells are indicated by asterisks.

Output:
This program returns the initial state and ten generations

Run:
This program requires filename as a command line argument

To run this program via the command line:
./gol /filename
