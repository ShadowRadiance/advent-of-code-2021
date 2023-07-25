using System.Text;
using aoc.support;

namespace aoc;

public class Day04 : Day
{
    public override string Solve(int part)
    {
        return BruteForce(part == 1 ? "00000" : "000000").ToString();
    }

    private int BruteForce(string startsWith)
    {
        var key = Input;
        var n = 0;
        while (true)
            if (MD5($"{Input}{++n}").StartsWith(startsWith))
                return n;
    }

    private string MD5(string input)
    {
        using var md5 = System.Security.Cryptography.MD5.Create();
        var inputBytes = Encoding.ASCII.GetBytes(input);
        var hashBytes = md5.ComputeHash(inputBytes);
        return Convert.ToHexString(hashBytes);
    }
}