using System.Reflection;

namespace aoc.support;

public static class SubTypeFinder
{
    public static IEnumerable<Type> GetSubTypesOf(Type baseType, Assembly assembly)
    {
        return assembly.GetTypes()
            .Where(baseType.IsAssignableFrom)
            .Where(t => baseType != t);
    }
}
