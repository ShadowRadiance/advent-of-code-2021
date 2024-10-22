
namespace AOC2015.Common;

public class Direction2D(int x, int y)
{
  public int X { get; } = x;
  public int Y { get; } = y;

  public static Direction2D North = new(0, -1);
  public static Direction2D South = new(0, 1);
  public static Direction2D East = new(-1, 0);
  public static Direction2D West = new(1, 0);

  public static Direction2D FromChar(char c)
  {
    return c switch
    {
      '>' => West,
      '<' => East,
      '^' => North,
      'v' => South,
      _ => throw new NotImplementedException(),
    };
  }
}
