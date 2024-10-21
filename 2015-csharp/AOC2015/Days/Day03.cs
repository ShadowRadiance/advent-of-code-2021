
namespace AOC2015.Days;

public class Day03 : Day
{
    public Day03(string[] data) : base(data) { }

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

    public class Location2D(int x, int y)
    {
        public int X { get; } = x;
        public int Y { get; } = y;

        public override int GetHashCode()
        {
            return $"{X},{Y}".GetHashCode();
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

    public class Actor(string name, Location2D initial)
    {
        public Location2D Location { get; private set; } = initial;

        public string Name { get; } = name;

        public void Move(Direction2D direction)
        {
            Location += direction;
        }
    }

    class Neighbourhood
    {
        readonly Dictionary<Location2D, int> visitedLocations = [];

        public void Visit(Location2D location)
        {
            if (!visitedLocations.ContainsKey(location))
            {
                visitedLocations.Add(location, 1);
            }
            else
            {
                visitedLocations[location] += 1;
            }
        }

        public int VisitedLocationCount()
        {
            return visitedLocations.Keys.Count;
        }
    }

    public override string PartA()
    {
        // Santa is delivering presents to an infinite two-dimensional grid of houses.
        // He begins by delivering a present to the house at his starting location,
        // and then an elf at the North Pole calls him via radio and tells him where to
        // move next. Moves are always exactly one house to the north (^), south (v),
        // east (>), or west (<). After each move, he delivers another present to the
        // house at his new location.
        // How many houses receive at least one present?
        string instructions = Data[0];
        Neighbourhood neighbourhood = new();

        Location2D location = new(0, 0);
        neighbourhood.Visit(location);

        foreach (char inst in instructions)
        {
            location += Direction2D.FromChar(inst);
            neighbourhood.Visit(location);
        }

        return neighbourhood.VisitedLocationCount().ToString();
    }

    public override string PartB()
    {
        // Santa creates a robot version of himself, Robo-Santa, to deliver presents with him.
        // Santa and Robo-Santa start at the same location (delivering two presents to the
        // same starting house), then take turns moving based on instructions from the elf,
        // who is eggnoggedly reading from the same script as the previous year.
        // This year, how many houses receive at least one present?
        string instructions = Data[0];
        Neighbourhood neighbourhood = new();

        Actor santa = new("Santa", new(0, 0));
        Actor robot = new("Robot", new(0, 0));

        neighbourhood.Visit(santa.Location);
        neighbourhood.Visit(robot.Location);

        Actor activeActor = santa;
        foreach (char inst in instructions)
        {
            activeActor.Move(Direction2D.FromChar(inst));

            neighbourhood.Visit(activeActor.Location);

            activeActor = (activeActor == santa) ? robot : santa;
        }

        return neighbourhood.VisitedLocationCount().ToString();
    }
}
