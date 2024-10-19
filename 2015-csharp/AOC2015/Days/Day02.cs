namespace AOC2015.Days;

public class Day02 : Day
{
    public Day02(string[] data) : base(data) { }

    private struct Dimensions
    {
        int length;
        int width;
        int height;

        public static Dimensions FromString(string dim)
        {
            string[] dimsStr = dim.Split('x');
            int[] dimsInt = dimsStr.Select(x => int.Parse(x)).ToArray();
            return new Dimensions
            {
                length = dimsInt[0],
                width = dimsInt[1],
                height = dimsInt[2],
            };
        }

        public readonly int SurfaceArea()
        {
            return 2 * length * width + 2 * width * height + 2 * height * length;
        }

        public readonly int WrappingSlack()
        {
            return Math.Min(Math.Min(length * width, width * height), height * length);
        }

        public readonly int WrappingRibbon()
        {
            // shortest distance around its sides, or the smallest perimeter of any one face
            return Math.Min(
                Math.Min(
                    2 * (length + width),
                    2 * (length + height)
                ),
                2 * (width + height)
            );
        }

        public readonly int BowRibbon()
        {
            // equal to cubic feet of volume of the present
            return length * width * height;
        }
    }

    public override string PartA()
    {
        var result = Data.Select(Dimensions.FromString)
                         .Sum(d => d.SurfaceArea() + d.WrappingSlack());
        return result.ToString();
    }

    public override string PartB()
    {
        // A present with dimensions 2x3x4 requires 2+2+3+3 = 10 feet of ribbon to wrap the present plus 2*3*4 = 24 feet of ribbon for the bow, for a total of 34 feet.
        // A present with dimensions 1x1x10 requires 1+1+1+1 = 4 feet of ribbon to wrap the present plus 1*1*10 = 10 feet of ribbon for the bow, for a total of 14 feet.

        var result = Data
            .Select(Dimensions.FromString)
            .Sum(d => d.WrappingRibbon() + d.BowRibbon());

        return result.ToString();
    }
}
