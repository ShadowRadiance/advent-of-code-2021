namespace AOC2015;

public class RunOptions(int day, char part, bool test = false)
{
    public int Day { get; } = day;
    public char Part { get; } = part;
    public bool Test { get; } = test;

    public static RunOptions FromArgs(string[] args)
    {
        if (args.Length != 2) throw new CommandLineException();

        return new RunOptions(ExtractDay(args), ExtractPart(args));
    }

    private static char ExtractPart(string[] args)
    {
        char part = args[1].ToUpper()[0];
        if (part != 'A' && part != 'B')
        {
            throw new CommandLineException();
        }

        return part;
    }

    private static int ExtractDay(string[] args)
    {
        int day;
        try
        {
            day = int.Parse(args[0]);
        }
        catch (Exception)
        {
            throw new CommandLineException();
        }

        if (day < 1 || day > 25)
        {
            throw new CommandLineException();
        }

        return day;
    }
}
