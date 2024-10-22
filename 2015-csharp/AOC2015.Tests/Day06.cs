using Xunit;
using AOC2015.Days;

namespace AOC2015.Tests
{
    public class Day06Test
    {
        [Theory]
        [InlineData(new string[] { "turn on 0,0 through 999,999" }, "1000000")]
        [InlineData(new string[] { "toggle 0,0 through 999,0" }, "1000")]
        [InlineData(new string[] { "turn on 499,499 through 500,500" }, "4")]
        [InlineData(
            new string[] {
                "turn on 0,0 through 999,999",         // +1000000
                "toggle 0,0 through 999,0",            //    -1000
                "turn off 499,499 through 500,500"     //       -4
            },
            "998996"
        )]
        public void PartA(string[] data, string expectation)
        {
            Assert.Equal(expectation, new Day06(data).PartA());
        }


        [Theory]
        [InlineData(new string[] { "turn on 0,0 through 0,0" }, "1")]
        [InlineData(new string[] { "toggle 0,0 through 999,999" }, "2000000")]
        public void PartB(string[] data, string expectation)
        {
            Assert.Equal(expectation, new Day06(data).PartB());
        }
    }
}
