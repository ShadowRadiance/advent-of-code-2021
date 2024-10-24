namespace AOC2015.Days;

public class Day10 : Day
{
    private int? _iterations;

    // standard constructor
    public Day10(string[] data) : base(data) => _iterations = null;
    // constructor for tests
    public Day10(string[] data, int iterations) : base(data) => _iterations = iterations;

    private struct DigitRun(char digit, int count)
    {
        public char digit = digit;
        public int count = count;
    }

    public static string PlayLookAndSay(string input)
    {
        if (input == "") return "";

        // break input into runs of repeated digits
        List<DigitRun> runs = [];
        DigitRun currentRun = new(digit: input[0], count: 1);
        for (int index = 1; index < input.Length; index++)
        {
            if (input[index] == currentRun.digit)
            {
                currentRun.count += 1;
            }
            else
            {
                runs.Add(currentRun);
                currentRun = new(digit: input[index], count: 1);
            }
        }
        runs.Add(currentRun);

        // replace each run with run.length run.digit
        return string.Join("", runs.Select(run => $"{run.count}{run.digit}"));
    }

    private string Solution()
    {
        string input = Data[0];
        for (int i = 0; i < _iterations; i++)
        {
            input = PlayLookAndSay(input);
        }
        return input.Length.ToString();
    }

    public override string PartA()
    {
        if (_iterations == null) _iterations = 40;
        return Solution();
    }

    public override string PartB()
    {
        if (_iterations == null) _iterations = 50;

        return Solution();
    }
}
