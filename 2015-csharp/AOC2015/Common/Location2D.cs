
namespace AOC2015.Common;

public class Location2D(int x, int y)
{
  public int X { get; } = x;
  public int Y { get; } = y;

  public static Location2D FromString(string xy)
  {
    int[] parts = xy.Split(",")
      .Select(s => int.Parse(s.Trim()))
      .ToArray();
    return new(parts[0], parts[1]);
  }

  public override string ToString()
  {
    return $"{X},{Y}";
  }

  public override int GetHashCode()
  {
    return ToString().GetHashCode();
  }
  public override bool Equals(object? obj)
  {
    return obj is Location2D location && location.X == X && location.Y == Y;
  }

  public static Location2D operator +(Location2D a, Direction2D b)
  {
    return new(a.X + b.X, a.Y + b.Y);
  }
}
