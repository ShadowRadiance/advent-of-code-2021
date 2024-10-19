namespace AOC2015;

class Runner
{
    private readonly RunOptions _runOptions;
    private readonly string _dayName;
    private readonly string _klassName;
    private readonly string[] _data;

    public Runner(RunOptions runOptions)
    {
        _runOptions = runOptions;
        _dayName = $"Day{_runOptions.Day:00}";
        _klassName = $"AOC2015.Days.{_dayName}";
        _data = File.ReadAllLines($"Data/{_dayName}.txt");
    }

    public string Run()
    {
        Day day = MakeDay();
        return _runOptions.Part switch
        {
            'A' => day.PartA(),
            'B' => day.PartB(),
            _ => throw new CommandLineException(),
        };
    }

    private Day MakeDay()
    {
        Type? t = Type.GetType(_klassName) ?? throw new CommandLineException();
        object? o = Activator.CreateInstance(t, [_data]) ?? throw new CommandLineException();
        return (Day)o;
    }
}
