## Conway's Game Of Life

The **Game Of Life** also known simply as **Life**, is a cellular automation devised by the British mathematician Jhon Horton
Conway in 1979. It is a zero player game, meaning that its evolution is determined by initial state, requiring no further 
input. One interaction with the Game Of Life by creating an initial configuration and observing how it evolves.

### Rules

The universe of **Game Of Life** is an infinite, two-dimensional orthogonal grid of square cells, each of wich is in one of
two possible state, _live_ or _dead_. Every cell interacts with eight neighbor, which are the cells that are 
horizontally, vertically or diagonally adjacent. At each step in time, the following transitions occur:

* Any live cell with fewer than two live neighbors dies, as if by underpopulation.
* Any live cell with two or three live neighbors lives on to the next generation.
* Any live cell with more than three live neighbors dies, as if by overpopulation.
* Any dead cell with exactly three live neighbors becomes a live cell, as if by reproduction.

* ![#1589F0] Build entirely  in Go technology.
