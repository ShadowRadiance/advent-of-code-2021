using System.Data;
using aoc.support;

namespace aoc;

[TestClass]
public class Day01Test
{
    [TestMethod]
    [DataRow("(())", "0")] 
    [DataRow("()()", "0")]
    [DataRow("(((", "3")]
    [DataRow("(()(()(", "3")]
    [DataRow("))(((((", "3")]
    [DataRow("())", "-1")]
    [DataRow("))(", "-1")]
    [DataRow(")))", "-3")]
    [DataRow(")())())", "-3")]
    public void Part1(string input, string expected)
    {
        IDay day = new Day01(input);
        string result = day.Solve(1);
        Assert.AreEqual(expected, result);
    }
    
    [TestMethod]
    public void CanSolvePart2()
    {
        IDay day = new Day01("");
        string result = day.Solve(2);
        Assert.AreEqual("PENDING", result);
    }
}