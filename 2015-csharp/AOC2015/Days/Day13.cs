using System.Collections.ObjectModel;
using System.Dynamic;
using System.Text.RegularExpressions;

namespace AOC2015.Days;

public class Day13 : Day
{
    // Alice would gain 54 happiness units by sitting next to Bob.
    // Alice would lose 79 happiness units by sitting next to Carol.
    // Alice would lose 2 happiness units by sitting next to David.
    // Bob would gain 83 happiness units by sitting next to Alice.
    // Bob would lose 7 happiness units by sitting next to Carol.
    // Bob would lose 63 happiness units by sitting next to David.
    // Carol would lose 62 happiness units by sitting next to Alice.
    // Carol would gain 60 happiness units by sitting next to Bob.
    // Carol would gain 55 happiness units by sitting next to David.
    // David would gain 46 happiness units by sitting next to Alice.
    // David would lose 7 happiness units by sitting next to Bob.
    // David would gain 41 happiness units by sitting next to Carol.

    static readonly Regex parser = new(@"(\w)\w* would (gain|lose) (\d+) happiness units by sitting next to (\w)\w*.");

    readonly struct HappinessModifier
    {
        public readonly char person1;
        public readonly char person2;
        public readonly int amount;
        public readonly string peopleAlpha;
        public readonly string peopleAlphaReversed;

        public HappinessModifier(char person1, char person2, int amount)
        {
            this.person1 = person1;
            this.person2 = person2;
            this.amount = amount;
            this.peopleAlpha = DeterminePeopleAlpha();
            this.peopleAlphaReversed = new(peopleAlpha.Reverse().ToArray());
        }

        private readonly string DeterminePeopleAlpha()
        {
            char[] people = [person1, person2];
            Array.Sort(people);
            return String.Join("", people);
        }

        public static HappinessModifier Parse(string line)
        {
            var matches = parser.Match(line);
            char person1 = matches.Groups[1].Value[0];
            char person2 = matches.Groups[4].Value[0];
            bool lose = matches.Groups[2].Value == "lose";
            int amount = int.Parse(matches.Groups[3].Value);
            if (lose) amount = -amount;

            return new(person1, person2, amount);
        }
    }

    public Day13(string[] data) : base(data) { }

    public static string Solve(Dictionary<string, int> happinessMap, char[] people)
    {
        // start with A... do a DFS to determine order to add others?
        // or just run each permutation of ABC…?

        int[] indices = new int[people.Length];
        for (int i = 0; i < people.Length; i++) { indices[i] = i; }
        List<int[]> permutations = new PermutationBuilder().Permutations(indices);
        // permutations = { [0,1,2,…], [2,1,0,…], etc }

        Dictionary<string, int> deducedHappinesses = [];
        foreach (int[] permutation in permutations)
        {
            string sPermutation = new(permutation.Select(i => people[i]).ToArray());
            if (!deducedHappinesses.ContainsKey(sPermutation))
            {
                deducedHappinesses.Add(sPermutation, HappinessOf(sPermutation, happinessMap));
            }
        }
        int maxHappiness = 0;
        string maxHappinessOrdering = "";
        foreach (var kv in deducedHappinesses)
        {
            if (kv.Value > maxHappiness)
            {
                maxHappiness = kv.Value;
                maxHappinessOrdering = kv.Key;
            }
        }

        Console.WriteLine($"Max Happiness from {maxHappinessOrdering} is {maxHappiness}");

        return maxHappiness.ToString();
    }

    private static int HappinessOf(string sPermutation, Dictionary<string, int> happinessMap)
    {
        int totalHappiness = 0;
        for (int i = 1; i < sPermutation.Length; i++)
        {
            char first = sPermutation[i - 1];
            char second = sPermutation[i];
            string peopleAlpha = $"{first}{second}";
            totalHappiness += happinessMap[peopleAlpha];
        }
        string peopleAlphaCycle = $"{sPermutation[^1]}{sPermutation[0]}";
        totalHappiness += happinessMap[peopleAlphaCycle];
        return totalHappiness;
    }

    private static Dictionary<string, int> BuildHappinessMap(List<HappinessModifier> happinessModifiers)
    {
        Dictionary<string, int> happinessMap = [];
        foreach (var happinessModifier in happinessModifiers)
        {
            if (!happinessMap.ContainsKey(happinessModifier.peopleAlpha))
            {
                happinessMap.Add(happinessModifier.peopleAlpha, happinessModifier.amount);
                happinessMap.Add(happinessModifier.peopleAlphaReversed, happinessModifier.amount);
            }
            else
            {
                happinessMap[happinessModifier.peopleAlpha] += happinessModifier.amount;
                happinessMap[happinessModifier.peopleAlphaReversed] += happinessModifier.amount;
            }
        }
        return happinessMap;
    }

    public override string PartA()
    {
        List<HappinessModifier> happinessModifiers = Data.Select(HappinessModifier.Parse).ToList();
        HashSet<char> people = happinessModifiers.Select(hm => hm.person1).ToHashSet();  // A, B, C, …
        Dictionary<string, int> happinessMap = BuildHappinessMap(happinessModifiers);    // AB: -2, AC: 13, BC: 44, …
        return Solve(happinessMap, [.. people]);
    }

    public override string PartB()
    {
        List<HappinessModifier> happinessModifiers = Data.Select(HappinessModifier.Parse).ToList();
        HashSet<char> people = happinessModifiers.Select(hm => hm.person1).ToHashSet();  // A, B, C, …
        Dictionary<string, int> happinessMap = BuildHappinessMap(happinessModifiers);    // AB: -2, AC: 13, BC: 44, …

        foreach (char person in people)
        {
            happinessMap.Add($"{person}X", 0);
            happinessMap.Add($"X{person}", 0);
        }
        people.Add('X');

        return Solve(happinessMap, [.. people]);
    }
}

// Brute Force (fast enough for this set!)
//   4 People = 4! = 4•3•2•1 = 24 permutations
//   8 People = 8! = 8•7•6•5•4•3•2•1 = 40320 permutations
// Check the happiness gained on each permutation

// Slightly smarter - notice that a circle of 1-2-3-4 == 2-3-4-1 == 3-4-1-2 == 4-1-2-3
// and  1-2-3-4-5-6-7-8 == 2-3-4-5-6-7-8-1 == 3-4-5-6-7-8-1-2 == etc

// Or:
//   Create nodes for each person with undirected edges for each other person weighted by the combination of their gain/loss
//   Nodes: Alice, Bob, Carol, David
//   Edge: Alice-Bob   (weight = +54 +83 = +137)
//   Edge: Alice-Carol (weight = -79 -62 = -141)
//   Edge: Alice-David (weight = -2  +46 =  +44)
//   etc
// Find the path that visits all nodes once starting and ending at Alice, maximising happiness
