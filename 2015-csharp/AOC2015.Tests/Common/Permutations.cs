using Xunit;

namespace AOC2015.Tests.Common
{
  public class PermutationBuilderTest
  {
    public void TestPermutations()
    {
      int[] data = [1, 2, 3];
      List<int[]> expectation = [
        [1, 2, 3],
        [1, 3, 2],
        [2, 1, 3],
        [2, 3, 1],
        [3, 1, 2],
        [3, 2, 1],
      ];

      Assert.Equal(expectation, new PermutationBuilder().Permutations(data));
    }
  }
}
