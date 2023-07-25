using aoc.support;

namespace aoc;

[TestClass]
public class DayTest
{
    private static readonly Dictionary<(int, int), List<Expectation>> ExpectationsSpecs = new();

    public DayTest()
    {
        ExpectationsSpecs[(1, 1)] = new List<Expectation>
        {
            new("(())", "0"),
            new("()()", "0"),
            new("(((", "3"),
            new("(()(()(", "3"),
            new("))(((((", "3"),
            new("())", "-1"),
            new("))(", "-1"),
            new(")))", "-3"),
            new(")())())", "-3")
        };
        ExpectationsSpecs[(1, 2)] = new List<Expectation>
        {
            new(")", "1"),
            new("()())", "5"),

            new("(())", "FAIL"),
            new("()()", "FAIL"),
            new("(((", "FAIL"),
            new("(()(()(", "FAIL"),
            new("))(((((", "1"),
            new("())", "3"),
            new("))(", "1"),
            new(")))", "1"),
            new(")())())", "1")
        };
    }


    [TestMethod]
    [DataRow(1, 1)]
    [DataRow(1, 2)]
    [DataRow(2, 1)]
    [DataRow(2, 2)]
    [DataRow(3, 1)]
    [DataRow(3, 2)]
    [DataRow(4, 1)]
    [DataRow(4, 2)]
    [DataRow(5, 1)]
    [DataRow(5, 2)]
    [DataRow(6, 1)]
    [DataRow(6, 2)]
    [DataRow(7, 1)]
    [DataRow(7, 2)]
    [DataRow(8, 1)]
    [DataRow(8, 2)]
    [DataRow(9, 1)]
    [DataRow(9, 2)]
    [DataRow(10, 1)]
    [DataRow(10, 2)]
    [DataRow(11, 1)]
    [DataRow(11, 2)]
    [DataRow(12, 1)]
    [DataRow(12, 2)]
    [DataRow(13, 1)]
    [DataRow(13, 2)]
    [DataRow(14, 1)]
    [DataRow(14, 2)]
    [DataRow(15, 1)]
    [DataRow(15, 2)]
    [DataRow(16, 1)]
    [DataRow(16, 2)]
    [DataRow(17, 1)]
    [DataRow(17, 2)]
    [DataRow(18, 1)]
    [DataRow(18, 2)]
    [DataRow(19, 1)]
    [DataRow(19, 2)]
    [DataRow(20, 1)]
    [DataRow(20, 2)]
    [DataRow(21, 1)]
    [DataRow(21, 2)]
    [DataRow(22, 1)]
    [DataRow(22, 2)]
    [DataRow(23, 1)]
    [DataRow(23, 2)]
    [DataRow(24, 1)]
    [DataRow(24, 2)]
    [DataRow(25, 1)]
    [DataRow(25, 2)]
    public void CanSolve(int day, int part)
    {
        var t = Type.GetType($"aoc.Day{day:D2},solutions")!;

        ExpectationsSpecs.TryGetValue((day, part), out var expectations);

        if (expectations == null)
            Assert.Inconclusive("PENDING");
        else
            foreach (var expectation in expectations)
            {
                var dayObject = (Day)Activator.CreateInstance(t)!;
                dayObject.SetInput(expectation.Input);
                Assert.AreEqual(expectation.Expected, dayObject.Solve(part));
            }
    }

    private record Expectation(string Input, string Expected);
}