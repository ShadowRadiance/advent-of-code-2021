using System.Text;
using System.Text.RegularExpressions;

namespace AOC2015.Days;

public class Day08 : Day
{
    public Day08(string[] data) : base(data) { }

    private const string bslash = "\\";
    private const string dquote = "\"";

    Regex _hexFinder = new("""
        \\                  # literal backslash
        X                   # literal X
        (                   # capture group 1 {
            [0-9A-F]        #   hexdigit
            [0-9A-F]        #   hexdigit
        )                   # } capture group 1
        """, RegexOptions.IgnoreCase | RegexOptions.IgnorePatternWhitespace);
    private string Actual(string representation)
    {
        // strip the wrapping quotes
        // replace every `\\` with `\`
        // replace every `\"` with `"`
        // replace every `\xXX` with (char)Uint16.Parse("XX", NumberStyles.AllowHexSpecifier)
        //                        or System.Convert.ToChar(System.Convert.ToUInt16("XX", 16));

        string actual = representation[1..^1]
            .Replace($"{bslash}{bslash}", @"\")
            .Replace($"{bslash}{dquote}", dquote);

        Match match = _hexFinder.Match(actual);
        while (match.Success)
        {
            Capture capture = match.Groups[1].Captures[0];

            actual =
                actual[0..(capture.Index - 2)]
                + Convert.ToChar(Convert.ToUInt16(capture.Value, 16))
                + actual[(capture.Index + 2)..];

            match = _hexFinder.Match(actual);
        }
        return actual;
    }

    private string Encode(string original)
    {
        StringBuilder stringBuilder = new();
        stringBuilder.Append(dquote);
        foreach (char c in original)
        {
            switch (c)
            {
                case '\\': stringBuilder.Append($"{bslash}{bslash}"); break;
                case '\"': stringBuilder.Append($"{bslash}{dquote}"); break;
                default: stringBuilder.Append(c); break;
            }
        }
        stringBuilder.Append(dquote);
        return stringBuilder.ToString();
    }

    public override string PartA()
    {
        // determine the sum of differences between
        //  the number of characters in the literals `"a\"a"` (6)
        //  the number of characters in the represented strings  `a"a` (3)
        // - e.g. "" is the empty string (2-0)
        // - e.g. "abc" is the string abc (5-3)
        // - e.g. "\xDE\xA7" is two characters in hex notation (10-2)
        // The only escape sequences used are:
        //      \\ (which represents a single backslash),
        //      \" (which represents a lone double-quote character), and
        //      \x plus two hexadecimal characters (which represents a single character with that ASCII code)

        return Data
            .Select(representation => representation.Length - Actual(representation).Length)
            .Sum()
            .ToString();
    }

    public override string PartB()
    {
        // let's go the other way.
        // In addition to finding the number of characters of code,
        //  you should now encode each code representation as a new string and find the number of characters of the new encoded representation,
        //  including the surrounding double quotes.
        // - e.g. ""            => "\"\""
        // - e.g. "abc"         => "\"abc\""
        // - e.g. "\xDE\xA7"    => "\"\\xDE\\xA7\""

        return Data
            .Select(text => Encode(text).Length - text.Length)
            .Sum()
            .ToString();
    }
}
