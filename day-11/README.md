# Day 11: Dumbo Octopus

You enter a large cavern full of rare bioluminescent dumbo octopuses! They seem to not like the Christmas lights on your submarine, so you turn them off for now.

There are 100 octopuses arranged neatly in a 10 by 10 grid. Each octopus slowly gains energy over time and flashes brightly for a moment when its energy is full. Although your lights are off, maybe you could navigate through the cave without disturbing the octopuses if you could predict when the flashes of light will happen.

Each octopus has an energy level - your submarine can remotely measure the energy level of each octopus (your puzzle input). For example:

```
5483143223
2745854711
5264556173
6141336146
6357385478
4167524645
2176841721
6882881134
4846848554
5283751526
```

The energy level of each octopus is a value between 0 and 9. Here, the top-left octopus has an energy level of 5, the bottom-right one has an energy level of 6, and so on.

You can model the energy levels and flashes of light in steps. During a single step, the following occurs:

- First, the energy level of each octopus increases by 1.
- Then, any octopus with an energy level greater than 9 flashes. This increases the energy level of all adjacent octopuses by 1, including octopuses that are diagonally adjacent. If this causes an octopus to have an energy level greater than 9, it also flashes. This process continues as long as new octopuses keep having their energy level increased beyond 9. (An octopus can only flash at most once per step.)
- Finally, any octopus that flashed during this step has its energy level set to 0, as it used all of its energy to flash.
- Adjacent flashes can cause an octopus to flash on a step even if it begins that step with very little energy. Consider the middle octopus with 1 energy in this situation:

Before any steps:

```
11111
19991
19191
19991
11111
```

After step 1:

```
34543
4•••4
5•••5
4•••4
34543
```

After step 2:

```
45654
51115
61116
51115
45654
```

An octopus is highlighted when it flashed during the given step.

Here is how the larger example above progresses:

Before any steps:

```
5483143223
2745854711
5264556173
6141336146
6357385478
4167524645
2176841721
6882881134
4846848554
5283751526
```

After step 1:

```
6594254334
3856965822
6375667284
7252447257
7468496589
5278635756
3287952832
7993992245
5957959665
6394862637
```

After step 2:

```
88•7476555
5•89•87•54
85978896•8
84857696••
87••9•88••
66•••88989
68••••5943
••••••7456
9••••••876
87••••6848
```

After step 3:

```
••5•9••866
85••8••575
99••••••39
97••••••41
9935•8••63
77123•••••
791125•••9
221113••••
•421125•••
••21119•••
```

After step 4:

```
2263•31977
•923•31697
••3222115•
••41111163
••76191174
••53411122
••4236112•
5532241122
1532247211
113223•211
```

After step 5:

```
4484144•••
2•44144•••
2253333493
1152333274
11873•3285
1164633233
1153472231
6643352233
2643358322
2243341322
```

After step 6:

```
5595255111
3155255222
33644446•5
2263444496
2298414396
2275744344
2264583342
7754463344
3754469433
3354452433
```

After step 7:

```
67•7366222
4377366333
4475555827
34966557•9
35••6256•9
35•9955566
3486694453
8865585555
486558•644
4465574644
```

After step 8:

```
7818477333
5488477444
5697666949
46•876683•
473494673•
474••97688
69••••7564
••••••9666
8•••••4755
68••••7755
```

After step 9:

```
9•6••••644
78•••••976
69••••••8•
584•••••82
5858••••93
69624•••••
8•2125•••9
222113•••9
9111128•97
7911119976
```

After step 10:

```
•481112976
••31112••9
••411125•4
••811114•6
••991113•6
••93511233
•44236113•
553225235•
•53225•6••
••3224••••
```

After step 10, there have been a total of 204 flashes. Fast forwarding, here is the same configuration every 10 steps:

After step 20:

```
3936556452
56865568•6
449655569•
444865558•
445686557•
568••86577
7•••••9896
•••••••344
6••••••364
46••••9543
```

After step 30:

```
•643334118
4253334611
3374333458
2225333337
2229333338
2276733333
2754574565
5544458511
9444447111
7944446119
```

After step 40:

```
6211111981
•421111119
••42111115
•••3111115
•••3111116
••65611111
•532351111
3322234597
2222222976
2222222762
```

After step 50:

```
9655556447
48655568•5
448655569•
445865558•
457486557•
57•••86566
6•••••9887
8••••••533
68•••••633
568••••538
```

After step 60:

```
25333342••
274333464•
2264333458
2225333337
2225333338
2287833333
3854573455
1854458611
1175447111
1115446111
```

After step 70:

```
8211111164
•421111166
••42111114
•••4211115
••••211116
••65611111
•532351111
7322235117
5722223475
4572222754
```

After step 80:

```
1755555697
59655556•9
448655568•
445865558•
457•86557•
57•••86566
7•••••8666
•••••••99•
•••••••8••
••••••••••
```

After step 90:

```
7433333522
2643333522
2264333458
2226433337
2222433338
2287833333
2854573333
4854458333
3387779333
3333333333
```

After step 100:

```
•397666866
•749766918
••53976933
•••4297822
•••4229892
••53222877
•532222966
9322228966
7922286866
6789998766
```

After 100 steps, there have been a total of 1656 flashes.

Given the starting energy levels of the dumbo octopuses in your cavern, simulate 100 steps. **How many total flashes are there after 100 steps?**

To begin, [get your puzzle input](https://adventofcode.com/2021/day/11/input).

Post the answer to https://adventofcode.com/2021/day/11

## Part Two

---

Post the answer to https://adventofcode.com/2021/day/11
