using aoc.support;

namespace aoc;

// https://adventofcode.com/2015/day/2

public class Day02 : Day
{
    public override string Solve(int part)
    {
        var lines = Input.Split('\n');
        if (part == 1)
            return lines.Sum(WrappingPaperFor).ToString();
        return lines.Sum(RibbonFor).ToString();
    }

    private static int[] SidesFor(string line)
    {
        return line.Split('x').Select(int.Parse).Order().ToArray();
    }

    private static int RibbonFor(string line)
    {
        var sides = SidesFor(line);
        var wrap = 2 * sides[0] + 2 * sides[1];
        var bow = sides[0] * sides[1] * sides[2];

        return wrap + bow;
    }

    private static int WrappingPaperFor(string line)
    {
        var sides = SidesFor(line);

        var slack = sides[0] * sides[1];

        var aSides = sides.ToArray();
        return 2 * aSides[0] * aSides[1]
               + 2 * aSides[0] * aSides[2]
               + 2 * aSides[1] * aSides[2]
               + slack;
    }
}