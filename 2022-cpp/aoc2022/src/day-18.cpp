#include <days.h>
#include <algorithm>

namespace day_18
{
	using std::string;
	using std::vector;

	struct Cube1x1x1
	{
		int x;
		int y;
		int z;
	};

	int parseInt(auto& it, auto end)
	{
		size_t size{ 0 };
		int parsed = stoi(string{ it, end }, &size);
		it += size;
		return parsed;
	}

	Cube1x1x1 parseCube(string const& str)
	{
		auto it = str.begin();
		auto end = str.end();
		int x = parseInt(it, end); ++it;
		int y = parseInt(it, end); ++it;
		int z = parseInt(it, end);
		return { x,y,z };
	}

	bool adjacent(Cube1x1x1 const& lhs, Cube1x1x1 const& rhs)
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

	int solve(vector<string> const& data)
	{
		// Each 1x1x1 cube has 6 sides.
		// If two cubes are connected by a side, 
		// 2 sides are removed from the total

		vector<Cube1x1x1> cubes(data.size());
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

	string answer_a(vector<string> const& input_data)
	{
		return std::to_string(solve(input_data));
	}

	string answer_b(vector<string> const& input_data)
	{
		return "PENDING";
	}
}
