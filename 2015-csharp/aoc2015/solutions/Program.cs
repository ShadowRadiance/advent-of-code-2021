using System.Reflection;
using aoc.support;

Console.WriteLine("Hello, World!");

foreach (var dayClass in SubTypeFinder.GetSubTypesOf(typeof(Day), Assembly.GetExecutingAssembly()))
{
    Day? dayObject = (Day?)Activator.CreateInstance(dayClass);
    if (dayObject != null)
    {
        string fileName = $"data/{dayClass.Name}.txt";
        string input = (File.Exists(fileName)) ? File.ReadAllText(fileName) : "";

        dayObject.SetInput(input);
        for (int part = 1; part <= 2; part++)
        {
            Console.WriteLine($"{dayClass.FullName} / Part {part} / {dayObject.Solve(part)}");
        }
    }
}