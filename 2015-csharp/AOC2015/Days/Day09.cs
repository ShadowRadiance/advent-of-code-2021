using System.Drawing;
using System.Text;
using System.Text.RegularExpressions;

namespace AOC2015.Days;

public class Day09 : Day
{
    public Day09(string[] data) : base(data) { }

    struct Route(string src, string dst, int distance)
    {
        public string Src { get; } = src;
        public string Dst { get; } = dst;
        public int Distance { get; } = distance;

        static readonly Regex parser = new(@"(\w+) to (\w+) = (\d+)");

        public static Route Parse(string line)
        {
            var matches = parser.Match(line);

            return new(matches.Groups[1].Value, matches.Groups[2].Value, int.Parse(matches.Groups[3].Value));
        }
    }

    class Place(string name)
    {
        public string Name { get; } = name;
    }

    private static int Factorial(int number)
    {
        if (number == 0) return 1;
        if (number == 1) return 1;
        return number * Factorial(number - 1);
    }

    private static void Permutations_Heaps(List<int[]> result, int[] input, int size)
    {
        if (size == 1)
        {
            result.Add((int[])input.Clone());
        }
        else
        {
            for (int i = 0; i < size; i++)
            {
                Permutations_Heaps(result, input, size - 1);
                if (size % 2 == 1)
                {
                    (input[size - 1], input[0]) = (input[0], input[size - 1]); // ooh parallel assignment swap!
                }
                else
                {
                    (input[size - 1], input[i]) = (input[i], input[size - 1]); // ooh parallel assignment swap!
                }
            }
        }
    }

    class Solution(List<Route> routes)
    {
        List<Route> _routes = routes;
        Dictionary<string, Dictionary<string, int>> _distances = [];
        string[] _places = [];
        List<int[]> _indexPermutations = [];
        Tuple<string, int>[] _pathLengths = [];

        public int ShortestPath()
        {
            BuildDistances();
            BuildPlaces();
            BuildIndexPermutations();
            DeterminePathLengths();

            var shortestPath = _pathLengths.MinBy(tuple => tuple.Item2);
            Console.WriteLine($"{shortestPath!.Item1} = {shortestPath!.Item2}");
            return shortestPath!.Item2;
        }

        public int LongestPath()
        {
            BuildDistances();
            BuildPlaces();
            BuildIndexPermutations();
            DeterminePathLengths();

            var longestPath = _pathLengths.MaxBy(tuple => tuple.Item2);
            Console.WriteLine($"{longestPath!.Item1} = {longestPath!.Item2}");
            return longestPath!.Item2;
        }


        private void BuildDistances()
        {
            foreach (Route route in _routes)
            {
                if (!_distances.ContainsKey(route.Src)) _distances.Add(route.Src, new());
                if (!_distances.ContainsKey(route.Dst)) _distances.Add(route.Dst, new());

                _distances[route.Src][route.Dst] = route.Distance;
                _distances[route.Dst][route.Src] = route.Distance;
            }
        }

        private void BuildPlaces()
        {
            _places = _distances.Keys.ToArray();
        }

        private void BuildIndexPermutations()
        {
            int[] indices = new int[_places.Length];
            for (int i = 0; i < indices.Length; i++) indices[i] = i;

            _indexPermutations = new();
            Permutations_Heaps(_indexPermutations, indices, indices.Length);
        }

        private void DeterminePathLengths()
        {
            _pathLengths = _indexPermutations.Select(
                indexPermutation => new Tuple<string, int>(
                    Path(indexPermutation),
                    PathLength(indexPermutation)
                )
            ).ToArray();
        }

        private string Path(int[] indices)
        {
            return String.Join(" -> ", indices.Select(index => _places[index]));
        }

        private int PathLength(int[] indices)
        {
            string src = _places[indices[0]];
            int distance = 0;
            for (int i = 1; i < indices.Length; i++)
            {
                string dst = _places[indices[i]];
                distance += _distances[src][dst];
                src = dst;
            }
            return distance;
        }
    }


    public override string PartA()
    {
        List<Route> routes = Data.Select(line => Route.Parse(line)).ToList();
        Solution solution = new(routes);

        return solution.ShortestPath().ToString();
    }

    public override string PartB()
    {
        List<Route> routes = Data.Select(line => Route.Parse(line)).ToList();
        Solution solution = new(routes);

        return solution.LongestPath().ToString();
    }
}
