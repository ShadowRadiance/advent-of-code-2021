# --- Day 22: Monkey Map ---

The monkeys take you on a surprisingly easy trail through the jungle. They're even going in roughly the right direction according to your handheld device's Grove Positioning System.

As you walk, the monkeys explain that the grove is protected by a force field. To pass through the force field, you have to enter a password; doing so involves tracing a specific path on a strangely-shaped board.

At least, you're pretty sure that's what you have to do; the elephants aren't exactly fluent in monkey.

The monkeys give you notes that they took when they last saw the password entered (your puzzle input).

For example:

```
        ...#
        .#..
        #...
        ....
...#.......#
........#...
..#....#....
..........#.
        ...#....
        .....#..
        .#......
        ......#.

10R5L5R10L4R5L5
```

The first half of the monkeys' notes is a map of the board. It is comprised of a set of _open tiles _(on which you can move, drawn `.`) and _solid walls_ (tiles which you cannot enter, drawn `#`).

The second half is a description of the path you must follow. It consists of alternating numbers and letters:

A _number_ indicates the number of tiles to move in the direction you are facing. If you run into a wall, you stop moving forward and continue with the next instruction.
A _letter_ indicates whether to turn 90 degrees clockwise (R) or counterclockwise (L). Turning happens in-place; it does not change your current tile.
So, a path like 10R5 means "go forward 10 tiles, then turn clockwise 90 degrees, then go forward 5 tiles".

You begin the path in the _leftmost open tile of the top row of tiles_. Initially, you are _facing to the right_ (from the perspective of how the map is drawn).

If a movement instruction would take you off of the map, _you wrap around to the other side of the board_. In other words, if your next tile is off of the board, you should instead look in the direction opposite of your current facing as far as you can until you find the opposite edge of the board, then reappear there.

For example, if you are at A and facing to the right, the tile in front of you is marked B; if you are at C and facing down, the tile in front of you is marked D:

```
        ...#
        .#..
        #...
        ....
...#.D.....#
........#...
B.#....#...A
.....C....#.
        ...#....
        .....#..
        .#......
        ......#.
```

It is possible for the next tile (after wrapping around) to be a wall; this still counts as there being a wall in front of you, and so movement stops before you actually wrap to the other side of the board.

By drawing the last facing you had with an arrow on each tile you visit, the full path taken by the above example looks like this:

```
        >>v#    
        .#v.    
        #.v.    
        ..v.    
...#...v..v#    
>>>v...>#.>>    
..#v...#....    
...>>>>v..#.    
        ...#....
        .....#..
        .#......
        ......#.
```

To finish providing the password to this strange input device, you need to determine numbers for your final _row, column, and facing_ as your final position appears from the perspective of the original map. **Rows** start from 1 at the top and count downward; **columns** start from 1 at the left and count rightward. (In the above example, row 1, column 1 refers to the empty space with no tile on it in the top-left corner.) **Facing** is `0` for right (`>`), `1` for down (`v`), `2` for left (`<`), and `3` for up (`^`). The final password is the _sum of 1000 times the row, 4 times the column, and the facing_.

In the above example, the final row is 6, the final column is 8, and the final facing is 0. So, the final password is `1000 * 6 + 4 * 8 + 0`: `6032`.

Follow the path given in the monkeys' notes. What is the final password?

To begin, get your puzzle input.

Answer: 106094

Day 22 Answer A: (   51675us) 106094

# --- Part Two ---

As you reach the force field, you think you hear some Elves in the distance. Perhaps they've already arrived?

You approach the strange input device, but it isn't quite what the monkeys drew in their notes. Instead, you are met with a large cube; each of its six faces is a square of 50x50 tiles.

To be fair, the monkeys' map does have six 50x50 regions on it. If you were to carefully fold the map, you should be able to shape it into a cube!

In the example above, the six (smaller, 4x4) faces of the cube are:

```
        1111
        1111
        1111
        1111
222233334444
222233334444
222233334444
222233334444
        55556666
        55556666
        55556666
        55556666
```

You still start in the same position and with the same facing as before, but the wrapping rules are different. Now, if you would walk off the board, you instead proceed around the cube. From the perspective of the map, this can look a little strange. In the above example, if you are at A and move to the right, you would arrive at B facing down; if you are at C and move down, you would arrive at D facing up:

```
        ...#
        .#..
        #...
        ....
...#.......#
........#..A
..#....#....
.D........#.
        ...#..B.
        .....#..
        .#......
        ..C...#.
```

Walls still block your path, even if they are on a different face of the cube. If you are at E facing up, your movement is blocked by the wall marked by the arrow:

```
        ...#
        .#..
     -->#...
        ....
...#..E....#
........#...
..#....#....
..........#.
        ...#....
        .....#..
        .#......
        ......#.
```

Using the same method of drawing the last facing you had with an arrow on each tile you visit, the full path taken by the above example now looks like this:

```
        >>v#    
        .#v.    
        #.v.    
        ..v.    
...#..^...v#    
.>>>>>^.#.>>    
.^#....#....    
.^........#.    
        ...#..v.
        .....#v.
        .#v<<<<.
        ..v...#.
```

The final password is still calculated from your final position and facing from the perspective of the map. In this example, the final row is 5, the final column is 7, and the final facing is 3, so the final password is 1000 * 5 + 4 * 7 + 3 = `5031`.

Fold the map into a cube, then follow the path given in the monkeys' notes. What is the final password?

Answer:

- Day 22 Answer B: (  61886µs) 19304 WRONG

Note to self: running from the command line shows:
```
FACES:
 .12
 .3. 
 45.
 6..

7L36R16L35R32R7R49L24R37R5L48L43L40...R6R8L30

[50,0,f:>] (0,0) of face 1          (after start at 50,0 facing >)
[53,0,f:>] (3,0) of face 1          (after 7 -- ie there's a rock at F1:4,0 -- there is)
[53,0,f:^] (3,0) of face 1          (after L)
[53,0,f:^] (3,0) of face 1          (after 36 -- ie there's a rock in the way up from face 1 -> F6:0,3 -- there is)
[53,0,f:>] (3,0) of face 1          (after R)
[53,0,f:>] (3,0) of face 1          (after 16 -- ie there's a rock at F1:4,0 -- there is)
[53,0,f:^] (3,0) of face 1          (after L)
[53,0,f:^] (3,0) of face 1          (after 35 -- ie there's a rock in the way up from face 1 -> F6:0,3 -- there is)
[53,0,f:>] (3,0) of face 1          (after R)
[53,0,f:>] (3,0) of face 1          (after 32 -- ie there's a rock at F1:4,0 -- there is)
[53,0,f:v] (3,0) of face 1          (after R)
[53,5,f:v] (3,5) of face 1          (after 7 -- ie there's a rock at F1:3,6 -- there is)
[53,5,f:<] (3,5) of face 1          (after R)
[15,144,f:>] (15,44) of face 4      (after 49 -- 3 steps to F1:0,0,<, 1 step to F4:0,44,>, 15 steps to F4:15,44 -- rock at F4:16,44 -- there is)
[15,144,f:^] (15,44) of face 4      (after L)
[15,132,f:^] (15,32) of face 4      (after 24 -- ie there is a rock at F4:15,31 -- there is)
...
So far so good...

After running all the instructions we seem to end up at:
...
[71,18,f:v] (21,18) of face 1       (after 8)
[71,18,f:>] (21,18) of face 1       (after L)
[75,18,f:>] (25,18) of face 1       (after 30 -- ie there is a rock at F1:26,18 -- there is)

Based on the rules, the final row is 18+1, col is 75+1, the final facing is 0 (>)
so the password should be 1000 * 19 + 4 * 76 + 0 == (19000+304+0) == 19304 
```

