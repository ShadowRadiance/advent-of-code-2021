namespace AOC2015.Days;

public class Day01 : Day
{
    public Day01(string[] data) : base(data) { }

    public override string PartA()
    {
        // To what floor do the instructions take Santa?

        // only expect one line of data
        string instructions = Data[0];

        int floor = 0;
        foreach (char c in instructions)
        {
            switch (c)
            {
                case '(':
                    floor += 1; break;
                case ')':
                    floor -= 1; break;
            }
        }

        return $"{floor}";
    }

    public override string PartB()
    {
        // find the position of the first character that causes him to enter the basement

        // only expect one line of data
        string instructions = Data[0];

        int floor = 0;
        int index = 0;
        foreach (char c in instructions)
        {
            index++;
            switch (c)
            {
                case '(': floor += 1; break;
                case ')': floor -= 1; break;
            }
            if (floor == -1) { return $"{index}"; }
        }

        return "X";
    }
}
