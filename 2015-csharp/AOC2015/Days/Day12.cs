using System.Collections;
using System.Globalization;
using System.Runtime.Serialization;
using System.Text.Json;
using System.Text.RegularExpressions;

namespace AOC2015.Days;

public class Day12 : Day
{
    public Day12(string[] data) : base(data) { }

    public override string PartA()
    {
        Regex numbers = new(@"(-?\d+)");
        var matches = numbers.Matches(Data[0]);
        return matches.Select(match => int.Parse(match.Groups[1].Value)).Sum().ToString();
    }

    private bool ObjectHasRedProperty(JsonElement element) => element.EnumerateObject().Any(child => child.Value.ToString() == "red");

    private int SumOfNumbersIn(JsonElement element) => element.ValueKind switch
    {
        JsonValueKind.Number => element.GetInt32(),
        JsonValueKind.Undefined => 0,
        JsonValueKind.Object => ObjectHasRedProperty(element) ? 0 : element.EnumerateObject().Sum(child => SumOfNumbersIn(child.Value)),
        JsonValueKind.Array => element.EnumerateArray().Sum(child => SumOfNumbersIn(child)),
        JsonValueKind.String => 0,
        JsonValueKind.True => 0,
        JsonValueKind.False => 0,
        JsonValueKind.Null => 0,
        _ => throw new NotImplementedException(),
    };

    public override string PartB()
    {
        string json = Data[0];

        JsonElement element = JsonSerializer.Deserialize<JsonElement>(json);

        return SumOfNumbersIn(element).ToString();
    }
}


/*
obj
    <JsonElement ValueKind = Array : "[[[3,"red",4]]]">
((JsonElement)obj).ValueKind
    Array
((JsonElement)obj)[0]
    <JsonElement ValueKind = Array : "[[3,"red",4]]">
((JsonElement)obj)[0].ValueKind
    Array
((JsonElement)obj)[0][0][1]
    <JsonElement ValueKind = String : "red">
((JsonElement)obj)[0][0][2]
    <JsonElement ValueKind = Number : "4">
((JsonElement)obj)[0][0][2].GetType()
    {System.Text.Json.JsonElement}
((JsonElement)obj)[0][0][2].ValueKind
    Number
((JsonElement)obj)[0][0][2].GetInt32()
    4
*/
