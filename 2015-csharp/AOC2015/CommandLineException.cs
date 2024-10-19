namespace AOC2015;

class CommandLineException : System.Exception
{
    override public string Message
    {
        get
        {
            return $"Usage: {AppDomain.CurrentDomain.FriendlyName} [1..25] [A..B] (--test)";
        }
    }
}
