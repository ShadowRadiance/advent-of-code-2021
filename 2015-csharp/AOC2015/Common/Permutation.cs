public class PermutationBuilder
{
  private List<int[]> _indexPermutations { get; set; } = [];

  public List<int[]> Permutations(int[] indices)
  {
    _indexPermutations = [];
    Permutations_Heaps(indices, indices.Length);
    return _indexPermutations;
  }

  private void Permutations_Heaps(int[] input, int size)
  {
    if (size == 1)
    {
      _indexPermutations.Add((int[])input.Clone());
      return;
    }

    int lastIndex = size - 1;
    for (int i = 0; i < size; i++)
    {
      Permutations_Heaps(input, lastIndex);
      if (size % 2 == 1)
      {
        // swap last with first
        (input[lastIndex], input[0]) = (input[0], input[lastIndex]); // ooh parallel assignment swap!
      }
      else
      {
        // swap last with i'th
        (input[lastIndex], input[i]) = (input[i], input[lastIndex]); // ooh parallel assignment swap!
      }
    }
  }

}
