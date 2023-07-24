using System.Reflection;
using aoc.support;

namespace aoc;

[TestClass]
public class DayTest
{
    [TestMethod]
    [DataRow(2,1, "PENDING")]
    [DataRow(2,2, "PENDING")]
    [DataRow(3,1, "PENDING")]
    [DataRow(3,2, "PENDING")]
    [DataRow(4,1, "PENDING")]
    [DataRow(4,2, "PENDING")]
    [DataRow(5,1, "PENDING")]
    [DataRow(5,2, "PENDING")]
    [DataRow(6,1, "PENDING")]
    [DataRow(6,2, "PENDING")]
    [DataRow(7,1, "PENDING")]
    [DataRow(7,2, "PENDING")]
    [DataRow(8,1, "PENDING")]
    [DataRow(8,2, "PENDING")]
    [DataRow(9,1, "PENDING")]
    [DataRow(9,2, "PENDING")]
    [DataRow(10,1, "PENDING")]
    [DataRow(10,2, "PENDING")]
    [DataRow(11,1, "PENDING")]
    [DataRow(11,2, "PENDING")]
    [DataRow(12,1, "PENDING")]
    [DataRow(12,2, "PENDING")]
    [DataRow(13,1, "PENDING")]
    [DataRow(13,2, "PENDING")]
    [DataRow(14,1, "PENDING")]
    [DataRow(14,2, "PENDING")]
    [DataRow(15,1, "PENDING")]
    [DataRow(15,2, "PENDING")]
    [DataRow(16,1, "PENDING")]
    [DataRow(16,2, "PENDING")]
    [DataRow(17,1, "PENDING")]
    [DataRow(17,2, "PENDING")]
    [DataRow(18,1, "PENDING")]
    [DataRow(18,2, "PENDING")]
    [DataRow(19,1, "PENDING")]
    [DataRow(19,2, "PENDING")]
    [DataRow(20,1, "PENDING")]
    [DataRow(20,2, "PENDING")]
    [DataRow(21,1, "PENDING")]
    [DataRow(21,2, "PENDING")]
    [DataRow(22,1, "PENDING")]
    [DataRow(22,2, "PENDING")]
    [DataRow(23,1, "PENDING")]
    [DataRow(23,2, "PENDING")]
    [DataRow(24,1, "PENDING")]
    [DataRow(24,2, "PENDING")]
    [DataRow(25,1, "PENDING")]
    [DataRow(25,2, "PENDING")]
    public void CanSolve(int day, int part, string expected)
    {
        Type t = Type.GetType($"aoc.Day{day:D2},solutions")!;
        IDay dayObject = (IDay)Activator.CreateInstance(t)!;
        Assert.AreEqual(expected, dayObject.Solve(part));
    }

}