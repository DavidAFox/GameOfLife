GameOfLife is a go implementation of Conway's game of life.

#### Rules
1. Any live cell with fewer than two live neighbors dies, as if caused by under-population.
2. Any live cell with two or three live neighbors lives on to the next generation.
3. Any live cell with more than three live neighbors dies, as if by over-population.
4. Any dead cell with exactly three live neighbors becomes a live cell, as if by reproduction.

[more info](https://en.wikipedia.org/wiki/Conway%27s_Game_of_Life)

#### Use
Use the runner to run the package.  It will look for a file called testfile.life in its directory.  The file should be a rectangular grid of .'s and 0's where .'s represent dead cells and zeros represent live cells.

The program will the continually output subsequent states.


For a more user friendly browser implementation go [here](https://github.com/DavidAFox/GameOfLifeAngularMod).