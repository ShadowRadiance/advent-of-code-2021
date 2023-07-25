using System.Drawing;
using aoc.support;

namespace aoc;

public class Day03 : Day
{
    private readonly Size _unitX = new(1, 0);
    private readonly Size _unitY = new(0, 1);

    public override string Solve(int part)
    {
        var instructions = Input.ToCharArray();
        var visitedLocations = new HashSet<Point> { new(0, 0) };

        if (part == 1)
            VisitHouses(1, ref visitedLocations, instructions);
        else
            VisitHouses(2, ref visitedLocations, instructions);
        return visitedLocations.Count.ToString();
    }

    private void VisitHouses(int visitors, ref HashSet<Point> visitedLocations, char[] instructions)
    {
        var currentLocations = new Point[visitors];
        var who = 0;
        foreach (var instruction in instructions)
        {
            switch (instruction)
            {
                case '>':
                    currentLocations[who] += _unitX;
                    break;
                case '<':
                    currentLocations[who] -= _unitX;
                    break;
                case 'v':
                    currentLocations[who] += _unitY;
                    break;
                case '^':
                    currentLocations[who] -= _unitY;
                    break;
            }

            visitedLocations.Add(currentLocations[who]);
            who = (who + 1) % visitors;
        }
    }
}