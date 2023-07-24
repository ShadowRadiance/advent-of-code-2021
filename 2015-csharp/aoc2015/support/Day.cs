namespace aoc.support;

public class Day
{
    protected string Input = "";

    public virtual string Solve(int part)
    {
        return "PENDING";
    }

    public void SetInput(string input)
    {
        Input = input;
    }
}