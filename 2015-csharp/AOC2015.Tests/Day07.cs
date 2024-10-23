using Xunit;
using AOC2015.Days;
using System.Runtime.Serialization;

#pragma warning disable CA1861 // Avoid constant arrays as arguments

namespace AOC2015.Tests
{
    public class Day07Test
    {
        [Theory]
        [InlineData(
            new string[] {
                "123 -> x",
                "456 -> y",
                "x AND y -> d",
                "x OR y -> e",
                "x LSHIFT 2 -> f",
                "y RSHIFT 2 -> g",
                "NOT x -> h",
                "NOT y -> i",
                "i -> a"
            },
            "65079"
        )]
        public void PartA(string[] data, string expectation)
        {
            Assert.Equal(expectation, new Day07(data).PartA());
        }


        [Theory]
        [InlineData(
            new string[] {
                "123 -> x",
                "456 -> b",
                "NOT b -> i",
                "i -> a"
            },
            "456"
        )]
        public void PartB(string[] data, string expectation)
        {
            Assert.Equal(expectation, new Day07(data).PartB());
        }
    }
}
