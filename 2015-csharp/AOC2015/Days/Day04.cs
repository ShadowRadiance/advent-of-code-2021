using System.Security.Cryptography;
using System.Text;

namespace AOC2015.Days;

public class Day04 : Day
{
    public Day04(string[] data) : base(data) { }

    private static IEnumerable<Tuple<long, string>> MD5Generator(string secret_key)
    {
        long index = 1;
        while (index < long.MaxValue)
        {
            string input = $"{secret_key}{index}";
            yield return new Tuple<long, string>(
                index,
                Convert.ToHexString(MD5.HashData(Encoding.UTF8.GetBytes(input)))
            );
            index += 1;
        }
    }

    public override string PartA()
    {
        // find MD5 hashes which, in hexadecimal, start with at least five zeroes.
        // The input to the MD5 hash is some secret key (your puzzle input, given below) followed by a number in decimal.
        // find Santa the lowest positive number (no leading zeroes: 1, 2, 3, ...) that produces such a hash.

        string secret_key = Data[0];

        foreach (var result in MD5Generator(secret_key))
        {
            if (result.Item2.StartsWith("00000"))
            {
                return result.Item1.ToString();
            }
        }

        return "NOT FOUND";
    }

    public override string PartB()
    {
        string secret_key = Data[0];

        foreach (var result in MD5Generator(secret_key))
        {
            if (result.Item2.StartsWith("000000"))
            {
                return result.Item1.ToString();
            }
        }

        return "NOT FOUND";
    }
}
