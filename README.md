# godoku
## Sudoku solver with GO
### Solves the given sudoku puzzle

`go get github.com/wviiras/godoku`
`go run github.com/wviiras/godoku path/to/unsolved/file`

It doesn't matter how the file itself is formatted, as long as the total number of digits and dashes total up to a 9x9 grid.

All of these are a valid input
``` 
8-- --6 3-5
-4- --- -7-
--- --- ---

-1- -38 7-4
--- 1-4 ---
3-- -7- 29-

--- --3 ---
-2- --- -4-
5-6 8-- --2
```
```
8----63-5
-4-----7-
---------
-1--387-4
---1-4---
3---7-29-
-----3---
-2-----4-
5-68----2
```

```
8----63-5-4-----7-----------1--387-4---1-4---3---7-29------3----2-----4-5-68----2
```