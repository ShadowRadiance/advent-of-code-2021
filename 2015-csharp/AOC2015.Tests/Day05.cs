using Xunit;
using AOC2015.Days;

namespace AOC2015.Tests
{
    public class Day05Test
    {
        [Theory]
        [InlineData(new string[] { "ugknbfddgicrmopn" }, "1")]
        [InlineData(new string[] { "aaa" }, "1")]
        [InlineData(new string[] { "jchzalrnumimnmhp" }, "0")]
        [InlineData(new string[] { "haegwjzuvuyypxyu" }, "0")]
        [InlineData(new string[] { "dvszwmarrgswjxmb" }, "0")]
        public void PartA(string[] data, string expectation)
        {
            Assert.Equal(expectation, new Day05(data).PartA());
        }


        [Theory]
        [InlineData(new string[] { "qjhvhtzxzqqjkmpb" }, "1")]
        [InlineData(new string[] { "xxyxx" }, "1")]
        [InlineData(new string[] { "uurcxstgmygtbstg" }, "0")]
        [InlineData(new string[] { "ieodomkazucvgmuy" }, "0")]
        public void PartB(string[] data, string expectation)
        {
            Assert.Equal(expectation, new Day05(data).PartB());
        }
    }
}
