using System.Text.RegularExpressions;

namespace AOC2015.Days;

public class Day11 : Day
{
    public Day11(string[] data) : base(data) { }

    private static Regex _nonOverlappingPairs = new(@"(\w)\1.*?(\w)\2");

    private static bool ContainsILO(string input) => (input.Contains('i') || input.Contains('l') || input.Contains('o'));
    private static bool ContainsIncreasingStraight(string input)
    {
        for (int i = 2; i < input.Length; i++)
        {
            if (input[i - 2] == input[i - 1] - (char)1 && input[i - 1] == input[i] - (char)1) return true;
        }
        return false;
    }
    private static bool ContainsTwoNonOverlappingPairs(string input) => _nonOverlappingPairs.IsMatch(input);
    public static bool PasswordValid(string input)
    {
        if (ContainsILO(input)) return false;
        if (!ContainsIncreasingStraight(input)) return false;
        if (!ContainsTwoNonOverlappingPairs(input)) return false;
        return true;
    }

    public static string Increment(string input)
    {
        // increment the last letter (skip I L O)
        char last = (char)(input.Last() + 1);

        if (last == 'i' || last == 'l' || last == 'o') last = (char)(last + 1);

        // "carry the one" i.e. z -> a but the next letter back gets incremented (carry the one again!)
        if (last == 'z' + 1)
        {
            return Increment(input[0..^1]) + 'a';
        }

        return input[0..^1] + last;
    }

    public override string PartA()
    {
        string oldPassword = Data[0];
        string newPassword = Increment(oldPassword);
        while (!PasswordValid(newPassword)) newPassword = Increment(newPassword);
        return newPassword;
    }

    public override string PartB()
    {
        string oldPassword = PartA();
        string newPassword = Increment(oldPassword);
        while (!PasswordValid(newPassword)) newPassword = Increment(newPassword);
        return newPassword;
    }
}
