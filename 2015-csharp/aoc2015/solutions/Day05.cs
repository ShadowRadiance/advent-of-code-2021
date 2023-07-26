using System.Text.RegularExpressions;
using aoc.support;

namespace aoc;

public class Day05 : Day
{
    public override string Solve(int part)
    {
        var allStrings = Input.Split("\n");
        if (part == 1)
            return allStrings.Where(StupidRules).Count().ToString();
        return allStrings.Where(BetterRules).Count().ToString();
    }


    private bool StupidRules(string s)
    {
        if (CountVowels(s) < 3) return false;
        if (!HasDoubleLetter(s)) return false;
        if (ContainsBadPair(s)) return false;

        return true;
    }

    private bool BetterRules(string s)
    {
        if (!ContainsMultipleNonOverlappingCopiesOfPair(s)) return false;
        if (!ContainsLtrAnyLtr(s)) return false;

        return true;
    }

    #region Better Methods

    private bool ContainsMultipleNonOverlappingCopiesOfPair(string s)
    {
        var r = new Regex(@"(?<first>\w\w).*(\k<first>)");
        return r.IsMatch(s);
    }

    private bool ContainsLtrAnyLtr(string s)
    {
        var r = new Regex(@"(?<first>\w)\w(\k<first>)");
        return r.IsMatch(s);
    }

    #endregion


    #region Silly Methods

    private int CountVowels(string s)
    {
        return s.Count(c => c == 'a') +
               s.Count(c => c == 'e') +
               s.Count(c => c == 'i') +
               s.Count(c => c == 'o') +
               s.Count(c => c == 'u');
    }

    private bool HasDoubleLetter(string s)
    {
        var r = new Regex(@"(?<first>\w)(\k<first>)");
        return r.IsMatch(s);
    }

    private bool ContainsBadPair(string s)
    {
        var r = new Regex(@"ab|cd|pq|xy");
        return r.IsMatch(s);
    }

    #endregion
}