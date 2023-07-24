using System.Reflection;
using aoc.support;

Console.WriteLine("Hello, World!");

foreach (var dayClass in SubTypeFinder.GetSubTypesOf(typeof(IDay), Assembly.GetExecutingAssembly()))
{
    IDay? dayObject = (IDay?)Activator.CreateInstance(dayClass);
    if (dayObject != null)
    {
        for (int part = 1; part <= 2; part++)
        {
            Console.WriteLine($"{dayClass.FullName} / Part {part} / {dayObject.Solve(part)}");
        }
    }
}

;
