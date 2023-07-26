using System.Reflection;
using aoc.support;

Console.WriteLine("Hello, World!");

foreach (var dayClass in SubTypeFinder.GetSubTypesOf(typeof(Day), Assembly.GetExecutingAssembly()))
{
    var dayObject = (Day?)Activator.CreateInstance(dayClass);
    if (dayObject != null)
    {
        var fileName = $"data/{dayClass.Name}.txt";
        var input = File.Exists(fileName) ? File.ReadAllText(fileName) : "";

        dayObject.SetInput(input);
        for (var part = 1; part <= 2; part++)
        {
            if (dayClass.Name == "Day04" && part == 2)
            {
                Console.WriteLine("aoc.Day04 / Part 2 / 9962624 (slow)");
                continue;
            }

            var result = dayObject.Solve(part);
            if (result != "PENDING")
                Console.WriteLine($"{dayClass.FullName} / Part {part} / {result}");
        }
    }
}