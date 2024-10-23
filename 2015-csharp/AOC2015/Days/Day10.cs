namespace AOC2015.Days;

public class Day10 : Day
{
    private int _iterations;

    public Day10(string[] data, int iterations = 40) : base(data)
    {
        _iterations = iterations;
    }

    private string PlayLookAndSay(string input)
    {
        // break input into runs of repeated digits
        // replace each run with run.length run.digit

        return input;
    }

    public override string PartA()
    {
        string input = Data[0];
        for (int i = 0; i < _iterations; i++)
        {
            input = PlayLookAndSay(input);
        }
        return input.Length.ToString();
    }

    public override string PartB()
    {
        return "";
    }
}
