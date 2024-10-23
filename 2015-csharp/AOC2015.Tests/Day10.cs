using Xunit;
using AOC2015.Days;

namespace AOC2015.Tests
{
    public class Day10Test
    {
        [Theory]
        [InlineData(new string[] { "1" }, 1, "11")]
        [InlineData(new string[] { "11" }, 1, "21")]
        [InlineData(new string[] { "21" }, 1, "1211")]
        [InlineData(new string[] { "1211" }, 1, "111221")]
        [InlineData(new string[] { "111221" }, 1, "312211")]
        public void PartA(string[] data, int iterations, string expectation)
        {
            Assert.Equal(expectation, new Day10(data, iterations).PartA());
        }


        [Theory]
        [InlineData(new string[] { "" }, 1, "")]
        public void PartB(string[] data, int iterations, string expectation)
        {
            Assert.Equal(expectation, new Day10(data, iterations).PartB());
        }
    }
}
