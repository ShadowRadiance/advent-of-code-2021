using aoc.support;

namespace aoc;

// https://adventofcode.com/2015/day/1

public class Day01 : Day
{
    public override string Solve(int part)
    {
        if (part == 1) return Naive().ToString();

        var result = PositionOfCharacterFirstEnteringFloor(-1);

        return result == 0 ? "FAIL" : result.ToString();
    }

    private int PositionOfCharacterFirstEnteringFloor(int targetFloor)
    {
        var currentFloor = 0;
        var position = 0;

        foreach (var c in Input)
        {
            ++position;
            if (c == '(')
                ++currentFloor;
            else
                --currentFloor;

            if (currentFloor == targetFloor)
                return position;
        }

        return 0;
    }

    private int Naive()
    {
        var up = Input.Count(x => x == '(');
        var down = Input.Length - up;
        return up - down;
    }
}