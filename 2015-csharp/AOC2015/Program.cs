namespace AOC2015;

class Program
{
    static void Main(string[] args)
    {
        try
        {
            Runner runner = new(RunOptions.FromArgs(args));
            string result = runner.Run();
            Console.WriteLine(result);
        }
        catch (Exception ex)
        {
            Console.WriteLine(ex.Message);
        }
    }

}
