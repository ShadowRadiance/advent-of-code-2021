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
        Day day = new Day01();
        day.SetInput(input);
        var result = day.Solve(1);
        Assert.AreEqual(expected, result);
    }

    [TestMethod]
    [DataRow(")", "1")]
    [DataRow("()())", "5")]
    [DataRow("(())", "FAIL")]
    [DataRow("()()", "FAIL")]
    [DataRow("(((", "FAIL")]
    [DataRow("(()(()(", "FAIL")]
    [DataRow("))(((((", "1")]
    [DataRow("())", "3")]
    [DataRow("))(", "1")]
    [DataRow(")))", "1")]
    [DataRow(")())())", "1")]
    public void Part2(string input, string expected)
    {
        Day day = new Day01();
        day.SetInput(input);
        var result = day.Solve(2);
        Assert.AreEqual(expected, result);
    }
}