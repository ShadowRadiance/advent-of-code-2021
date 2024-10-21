using Xunit;
using AOC2015.Days;

namespace AOC2015.Tests
{
    public class Day04Test
    {
        [Theory]
        [InlineData(new string[] { "" }, "")]
        public void PartA(string[] data, string expectation)
        {
            Assert.Equal(expectation, new Day04(data).PartA());
        }


        [Theory]
        [InlineData(new string[] { "" }, "")]
        public void PartB(string[] data, string expectation)
        {
            Assert.Equal(expectation, new Day04(data).PartB());
        }
    }
}
