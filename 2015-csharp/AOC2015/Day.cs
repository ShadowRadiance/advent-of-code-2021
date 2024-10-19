namespace AOC2015;

public abstract class Day(string[] data)
{
    public string[] Data { get; } = data;

    public abstract string PartA();
    public abstract string PartB();
}
