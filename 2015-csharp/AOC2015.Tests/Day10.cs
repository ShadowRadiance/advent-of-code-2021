using AOC2015.Days;

namespace AOC2015.Tests
{
    public class Day10Test
    {
        [Theory]
        [InlineData("1", "11")]
        [InlineData("11", "21")]
        [InlineData("21", "1211")]
        [InlineData("1211", "111221")]
        [InlineData("111221", "312211")]
        public void PlayLookAndSay_Test(string input, string expectation)
        {
            Assert.Equal(expectation, Day10.PlayLookAndSay(input));
        }

        [Theory]
        [InlineData(new string[] { "1" }, 1, "2")]
        [InlineData(new string[] { "11" }, 1, "2")]
        [InlineData(new string[] { "21" }, 1, "4")]
        [InlineData(new string[] { "1211" }, 1, "6")]
        [InlineData(new string[] { "111221" }, 1, "6")]
        public void PartA(string[] data, int iterations, string expectation)
        {
            Assert.Equal(expectation, new Day10(data, iterations).PartA());
        }
    }
}
