using System.Globalization;

namespace AOC2015.Days;

public class Day05 : Day
{
    public Day05(string[] data) : base(data) { }

    private static bool IsVowel(char input) => "aeiou".Contains(input);
    private static bool HasThreeVowels(string input) => input.Count(IsVowel) >= 3;
    private static bool HasDoubleLetter(string input)
    {
        if (input.Length < 2) return false;

        for (int i = 1; i < input.Length; i++)
        {
            if (input[i] == input[i - 1]) return true;
        }
        return false;
    }
    private static bool NoBadPairs(string input)
    {
        if (input.Contains("ab")) return false;
        if (input.Contains("cd")) return false;
        if (input.Contains("pq")) return false;
        if (input.Contains("xy")) return false;
        return true;
    }
    private static bool DuplicatedNonOverlappingPair(string input)
    {
        if (input.Length < 4) return false;

        for (int outer = 0; outer <= input.Length - 4; outer++)
        {
            for (int inner = outer + 2; inner <= input.Length - 2; inner++)
            {
                if (input.Substring(outer, 2) == input.Substring(inner, 2)) return true;
            }
        }
        return false;
    }
    private static bool RepeatedLetterOffByOne(string input)
    {
        if (input.Length < 3) return false;

        for (int i = 2; i < input.Length; i++)
        {
            if (input[i] == input[i - 2]) return true;
        }
        return false;

    }

    private static bool IsNice(string input)
    {
        return HasThreeVowels(input) && HasDoubleLetter(input) && NoBadPairs(input);
    }
    private static bool IsNicer(string input)
    {
        return DuplicatedNonOverlappingPair(input) && RepeatedLetterOffByOne(input);
    }

    public override string PartA()
    {
        // nice:
        // - at least three vowels
        // - at least one letter that appears twice in a row
        // - does NOT contain 'ab', 'cd', 'pq', or 'xy'
        // naughty: !nice

        // how many string are nice?
        return Data.Select(IsNice).Count(el => el == true).ToString();
    }

    public override string PartB()
    {
        // nice:
        // - It contains a pair of any two letters that appears at least twice in the string without overlapping
        //   like xyxy (xy) or aabcdefgaa (aa), but not like aaa (aa, but it overlaps)
        // - It contains at least one letter which repeats with exactly one letter between them
        //   like xyx, abcdefeghi (efe), or even aaa.

        return Data.Select(IsNicer).Count(el => el == true).ToString();
    }
}
