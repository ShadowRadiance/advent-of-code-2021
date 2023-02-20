#include <days.h>
#include <algorithm>
#include <stack>
#include <unordered_set>
#include <iterator>

namespace day_18
{
    using std::string;
    using std::vector;

    struct Point3D
    {
        union
        {
            struct
            {
                int16_t x;
                int16_t y;
                int16_t z;
            };
            size_t packed;
        };
    };

    bool operator==(Point3D const& lhs, Point3D const& rhs)
    {
        return lhs.x == rhs.x && lhs.y == rhs.y && lhs.z == rhs.z;
    }

    Point3D operator+(Point3D const& lhs, Point3D const& rhs)
    {
        //return { lhs.x + rhs.x, lhs.y + rhs.y, lhs.z + rhs.z };
        return {
            static_cast<int16_t>(lhs.x + rhs.x),
            static_cast<int16_t>(lhs.y + rhs.y),
            static_cast<int16_t>(lhs.z + rhs.z) };
    }

    struct Point3DHasher
    {
        size_t operator()(Point3D const& point) const { return point.packed; }
    };
    using POINTSET = std::unordered_set<Point3D, Point3DHasher>;

    int16_t parseInt(auto& it, auto end)
    {
        size_t size{ 0 };
        int16_t parsed = static_cast<int16_t>(stoi(string{ it, end }, &size));
        it += size;
        return parsed;
    }

    Point3D parseCube(string const& str)
    {
        auto it = str.begin();
        auto end = str.end();
        int16_t x = parseInt(it, end); ++it;
        int16_t y = parseInt(it, end); ++it;
        int16_t z = parseInt(it, end);
        return { x,y,z };
    }

    bool adjacent(Point3D const& lhs, Point3D const& rhs)
    {
        // two cubes are adjacent if they share two coordinates
        // and the third coordinate is exactly one different

        // ie (|x1-x2| == 0 && |y1-y2| == 0 && |z1-z2| == 1)
        // || (|x1-x2| == 0 && |y1-y2| == 1 && |z1-z2| == 0)
        // || (|x1-x2| == 1 && |y1-y2| == 0 && |z1-z2| == 0)

        // ie (|x1-x2| + |y1-y2| + |z1-z2| == 1)

        return 1 ==
            abs(lhs.x - rhs.x) +
            abs(lhs.y - rhs.y) +
            abs(lhs.z - rhs.z);
    }

    vector<Point3D> neighbours(Point3D const& point)
    {
        return {
            point + Point3D{1,0,0},
            point + Point3D{0,1,0},
            point + Point3D{0,0,1},
            point + Point3D{-1,0,0},
            point + Point3D{0,-1,0},
            point + Point3D{0,0,-1},
        };
    }

    int solve(vector<string> const& data)
    {
        // Each 1x1x1 cube has 6 sides.
        // If two cubes are connected by a side, 
        // 2 sides are removed from the total

        vector<Point3D> cubes(data.size());
        std::transform(data.begin(), data.end(), cubes.begin(), parseCube);

        int totalSides = cubes.size() * 6;
        auto end = cubes.end();
        for (auto cube1 = cubes.begin(); cube1 != end - 1; ++cube1) {
            for (auto cube2 = cube1 + 1; cube2 != end; ++cube2) {
                if (adjacent(*cube1, *cube2)) {
                    totalSides -= 2;
                }
            }
        }

        return totalSides;
    }

    void flood_fill(
        Point3D const& minBound,
        Point3D const& maxBound,
        POINTSET const& cubesBound,
        auto fnHitCubeEdge
    )
    {
        // do a six-way flood fill (up down left right front back, no diagonals)
        // from each corner

        // create stack
        // add eight corner points to stack
        // while stack not empty
        //  node = stack.pop
        //  if node is outside minBound/maxBound
        //   continue
        //  if visited
        //   continue
        //  if node is in cubes bound
        //   call fnHitCubeEdge
        //   continue
        //  mark visited
        //  add all six directions from this point to queue
        // return

        POINTSET visited;
        std::stack<Point3D> stack;
        stack.push({ minBound.x, minBound.y, minBound.z }); //000
        stack.push({ minBound.x, minBound.y, maxBound.z }); //001
        stack.push({ minBound.x, maxBound.y, minBound.z }); //010
        stack.push({ minBound.x, maxBound.y, maxBound.z }); //011
        stack.push({ maxBound.x, minBound.y, minBound.z }); //100
        stack.push({ maxBound.x, minBound.y, maxBound.z }); //101
        stack.push({ maxBound.x, minBound.y, maxBound.z }); //110
        stack.push({ maxBound.x, maxBound.y, maxBound.z }); //111

        while (!stack.empty()) {
            Point3D current = stack.top(); stack.pop();
            if (visited.contains(current)) continue;
            if (current.x < minBound.x || current.y < minBound.y || current.z < minBound.z ||
                current.x > maxBound.x || current.y > maxBound.y || current.z > maxBound.z) continue;
            if (cubesBound.contains(current)) {
                fnHitCubeEdge();
                continue;
            }
            visited.insert(current);
            for (Point3D const& neighbour : neighbours(current)) {
                stack.push(neighbour);
            }
        }
    }

    int solve_removing_trapped(vector<string> const& data)
    {
        POINTSET cubes;
        std::transform(data.begin(), data.end(), std::inserter(cubes, cubes.end()), parseCube);

        // create a cube bounding the cubes
        Point3D min{ 
            std::numeric_limits<int16_t>::max(), 
            std::numeric_limits<int16_t>::max(), 
            std::numeric_limits<int16_t>::max() };
        Point3D max{ 
            std::numeric_limits<int16_t>::min(), 
            std::numeric_limits<int16_t>::min(), 
            std::numeric_limits<int16_t>::min() };
        for (auto& cube : cubes) {
            min.x = std::min(min.x, cube.x);
            min.y = std::min(min.y, cube.y);
            min.z = std::min(min.z, cube.z);
            max.x = std::max(max.x, cube.x);
            max.y = std::max(max.y, cube.y);
            max.z = std::max(max.z, cube.z);
        }
        // expand the bounding cube
        min = min + Point3D{ -1, -1, -1 };
        max = max + Point3D{ 1, 1, 1 };

        int surface_area = 0;
        flood_fill(min, max, cubes, [&surface_area]() { surface_area++; });
        return surface_area;
    }

    string answer_a(vector<string> const& input_data)
    {
        return std::to_string(solve(input_data));
    }

    string answer_b(vector<string> const& input_data)
    {
        return std::to_string(solve_removing_trapped(input_data));
    }
}
