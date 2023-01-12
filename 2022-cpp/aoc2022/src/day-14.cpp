#include <days.h>

#include <unordered_map>
#include <algorithm>
#include <iterator>
#include <cassert>
#include <numeric>
#include <optional>

namespace day_14
{
	using std::string;
	using std::vector;

	struct Vector2D
	{
		double x;
		double y;

		static const Vector2D up() { return { 0.0,-1.0 }; }
		static const Vector2D dn() { return { 0.0,+1.0 }; }
		static const Vector2D lt() { return { -1.0,0.0 }; }
		static const Vector2D rt() { return { +1.0,0.0 }; }

		bool operator<(const Vector2D& rhs) const { return x < rhs.x || (x == rhs.x && y < rhs.y); }
		bool operator==(const Vector2D& rhs) const { return x == rhs.x && y == rhs.y; }
		Vector2D operator+(const Vector2D& rhs) const { return { x + rhs.x, y + rhs.y }; }
		Vector2D operator-(const Vector2D& rhs) const { return { x - rhs.x, y - rhs.y }; }
		Vector2D unit() const { return { x / magnitude(), y / magnitude() }; }
		double manhattanDistanceFrom(const Vector2D& rhs) {
			Vector2D pos{ *this - rhs };
			return abs(pos.x) + abs(pos.y);
		}
		double magnitude() const {
			return sqrt(x * x + y * y);
		}
	};

	struct Vector2D_Hash {
		size_t operator()(const Vector2D& vec2d) const noexcept {
			size_t x = std::hash<int>()(vec2d.x);
			size_t y = std::hash<int>()(vec2d.y);
			return x ^ (y << 1);
		}
	};

	class Cave
	{
		using Map = std::unordered_map<Vector2D, std::string, Vector2D_Hash>;
		using iterator = Map::iterator;
		using const_iterator = Map::const_iterator;
	public:
		int count(std::string what)
		{
			return std::count_if(
				begin(), end(), [what](auto& pair) { return pair.second == what; }
			);
		}

		std::string at(Vector2D where) {
			if (point_contents.contains(where)) {
				return point_contents.at(where);
			}
			return "";
		}

		void add(std::string what, Vector2D where) {
			point_contents[where] = what;
		}

		std::string remove(Vector2D where) {
			if (point_contents.contains(where)) {
				std::string old = point_contents.at(where);
				point_contents.erase(where);
				return old;
			}
			return "";
		}

		iterator begin() { return point_contents.begin(); }
		const_iterator begin() const { return point_contents.begin(); }
		iterator end() { return point_contents.end(); }
		const_iterator end() const { return point_contents.end(); }

	private:
		Map point_contents;
	};

	class SandInjector
	{
	public:
		SandInjector(Cave& cave, Vector2D injectionPoint)
			: cave(cave), injectionPoint(injectionPoint)
		{
			auto lesserY = [](auto& lPair, auto& rPair) {
				return lPair.first.y < rPair.first.y;
			};
			falloffPoint = {
				0.0,
				std::max_element(
					cave.begin(), cave.end(), lesserY
				)->first.y
			};
		}
		std::optional<Vector2D> injectSand() {
			cave.add("sand", injectionPoint);
			return settle(injectionPoint);
		}
	private:
		std::optional<Vector2D> settle(Vector2D position) {
			return settle_iterative(position);
		}

		std::optional<Vector2D> settle_recursive(Vector2D position) {
			assert(cave.at(position) == "sand");
			if (position.y > falloffPoint.y) {
				cave.remove(position);
				return {};
			}

			auto dn = position + Vector2D::dn();
			auto dn_lt = dn + Vector2D::lt();
			auto dn_rt = dn + Vector2D::rt();
			if (cave.at(dn) == "") {
				cave.add(cave.remove(position), dn);
				return settle(dn);
			}
			if (cave.at(dn_lt) == "") {
				cave.add(cave.remove(position), dn_lt);
				return settle(dn_lt);
			}
			if (cave.at(dn_rt) == "") {
				cave.add(cave.remove(position), dn_rt);
				return settle(dn_rt);
			}
			return position;
		}

		std::optional<Vector2D> settle_iterative(Vector2D position) {
			Vector2D targetPos = position;
			assert(cave.at(position) == "sand");

			while (true) {
				if (targetPos.y > falloffPoint.y) {
					cave.remove(position);
					return {};
				}

				Vector2D newPos = targetPos + Vector2D::dn();
				if (cave.at(newPos) == "") {
					targetPos = newPos;
					continue;
				}
				newPos = targetPos + Vector2D::dn() + Vector2D::lt();
				if (cave.at(newPos) == "") {
					targetPos = newPos;
					continue;
				}
				newPos = targetPos + Vector2D::dn() + Vector2D::rt();
				if (cave.at(newPos) == "") {
					targetPos = newPos;
					continue;
				}
				cave.add(cave.remove(position), targetPos);
				return targetPos;
			}
		}
	private:
		Cave& cave;
		Vector2D injectionPoint;
		Vector2D falloffPoint;
	};

	class Parser
	{
	public:
		Parser(const std::vector<std::string>& data)
			: paths(parse(data))
		{
		}

		void populate(Cave& cave)
		{
			for (auto& path : paths) {
				bool penDown = false;
				Vector2D current;
				for (auto& position : path) {
					if (penDown) {
						Vector2D unitDirection = (position - current).unit();
						assert(unitDirection.x == 0.0 || unitDirection.y == 0.0);
						while (position != current) {
							current = current + unitDirection;
							cave.add("rock", current);
						}
					}
					else {
						current = position;
						cave.add("rock", current);
						penDown = true;
					}
				}
			}
		}
	private:
		std::vector<std::vector<Vector2D>> parse(const std::vector<std::string>& data)
		{
			auto parseInt = [](auto& it, auto end) -> int {
				size_t size{ 0 };
				int parsed = std::stoi(string{ it, end }, &size);
				it += size;
				return parsed;
			};

			auto parseVector2D = [&parseInt](auto& it, auto end) -> Vector2D {
				int x = parseInt(it, end);
				assert(*it == ',');
				++it;
				int y = parseInt(it, end);
				return { double(x), double(y) };
			};

			auto parsePath = [&parseVector2D](std::string str) -> std::vector<Vector2D> {
				// 491,131 -> 502,131 -> 502,130
				std::vector<Vector2D> path;

				auto it = str.begin();
				auto end = str.end();

				while (it != end) {
					if (std::isdigit(*it)) {
						path.push_back(parseVector2D(it, end)); // advances current
					}
					else if (*it == ' ' || *it == '-' || *it == '>') {
						++it;
					}
				}
				return path;
			};

			std::vector<std::vector<Vector2D>> result;
			std::transform(data.begin(), data.end(), std::back_inserter(result), parsePath);
			return result;
		}
		const std::vector<std::vector<Vector2D>> paths;
	};

	string answer_a(const vector<string>& input_data)
	{
		Cave cave;
		Parser parser(input_data);
		parser.populate(cave);
		SandInjector sandInjector{ cave, {500, 0} };
		while (sandInjector.injectSand()) {}
		return std::to_string(cave.count("sand"));
	}

	string answer_b(const vector<string>& input_data)
	{
		return "PENDING";
	}
}
