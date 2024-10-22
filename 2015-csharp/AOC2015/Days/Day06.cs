using System.Text.RegularExpressions;
using AOC2015.Common;

namespace AOC2015.Days;

public struct Instruction(string command, Location2D from, Location2D to)
{
    private static Regex instructionSplitter = new(@"(turn on|turn off|toggle) (\d+,\d+) through (\d+,\d+)");

    public string Command { get; } = command;
    public Location2D FromLocation { get; } = from;
    public Location2D ToLocation { get; } = to;

    public static Instruction FromString(string instruction)
    {
        // split instruction ("turn on 766,112 through 792,868") into
        // command: (turn on / turn off / toggle)
        // from_location: (766,112)
        // to_location: (792,868)

        var matches = instructionSplitter.Matches(instruction);

        return new(
            matches[0].Groups[1].Value,
            Location2D.FromString(matches[0].Groups[2].Value),
            Location2D.FromString(matches[0].Groups[3].Value)
        );
    }
}

public class LightGrid
{
    private bool[,] lights = new bool[1000, 1000];

    public void TurnOn(int x, int y) => lights[x, y] = true;
    public void TurnOff(int x, int y) => lights[x, y] = false;
    public void Toggle(int x, int y) => lights[x, y] = !lights[x, y];

    public int Lit()
    {
        int count = 0;
        foreach (var light in lights)
        {
            if (light == true) count++;
        }
        return count;
    }
    public int Unlit()
    {
        return lights.Length - Lit();
    }

    public void Process(string instructionString)
    {
        // loop through locations calling command for each location

        Instruction instruction = Instruction.FromString(instructionString);

        for (int y = instruction.FromLocation.Y; y <= instruction.ToLocation.Y; y++)
            for (int x = instruction.FromLocation.X; x <= instruction.ToLocation.X; x++)
                switch (instruction.Command)
                {
                    case "turn on": TurnOn(x, y); break;
                    case "turn off": TurnOff(x, y); break;
                    case "toggle": Toggle(x, y); break;
                    default: break;
                }
    }
}

public class LightGrid_AncientNordicElvish
{
    private int[,] lights = new int[1000, 1000];
    private static Regex instructionSplitter = new(@"(turn on|turn off|toggle) (\d+,\d+) through (\d+,\d+)");

    public void TurnOn(int x, int y) => lights[x, y] += 1;
    public void TurnOff(int x, int y) => lights[x, y] = Math.Clamp(lights[x, y] - 1, 0, int.MaxValue);
    public void Toggle(int x, int y) => lights[x, y] += 2;

    public int Brightness()
    {
        int total = 0;
        foreach (var light in lights)
        {
            total += light;
        }
        return total;
    }

    public void Process(string instructionString)
    {
        // loop through locations calling command for each location

        Instruction instruction = Instruction.FromString(instructionString);

        for (int y = instruction.FromLocation.Y; y <= instruction.ToLocation.Y; y++)
            for (int x = instruction.FromLocation.X; x <= instruction.ToLocation.X; x++)
                switch (instruction.Command)
                {
                    case "turn on": TurnOn(x, y); break;
                    case "turn off": TurnOff(x, y); break;
                    case "toggle": Toggle(x, y); break;
                    default: break;
                }
    }
}

public class Day06 : Day
{
    public Day06(string[] data) : base(data) { }


    public override string PartA()
    {
        LightGrid lightGrid = new();

        foreach (string instruction in Data)
        {
            lightGrid.Process(instruction);
        }

        return lightGrid.Lit().ToString();
    }

    public override string PartB()
    {
        LightGrid_AncientNordicElvish lightGrid = new();

        foreach (string instruction in Data)
        {
            lightGrid.Process(instruction);
        }

        return lightGrid.Brightness().ToString();
    }
}
