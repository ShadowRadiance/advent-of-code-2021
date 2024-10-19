using Xunit;
using AOC2015.Days;

namespace AOC2015.Tests
{
    public class Day03Test
    {
        [Theory]
        [InlineData(new string[] { ">" }, "2")]
        [InlineData(new string[] { "^>v<" }, "4")]
        [InlineData(new string[] { "^v^v^v^v^v" }, "2")]
        public void PartA(string[] data, string expectation)
        {
            Assert.Equal(expectation, new Day03(data).PartA());
        }

        [Theory]
        [InlineData(new string[] { "^v" }, "3")]
        [InlineData(new string[] { "^>v<" }, "3")]
        [InlineData(new string[] { "^v^v^v^v^v" }, "11")]
        public void PartB(string[] data, string expectation)
        {
            Assert.Equal(expectation, new Day03(data).PartB());
        }
    }
}
