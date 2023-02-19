#include <days.h>
#include <algorithm>
#include <iostream>

namespace day_17
{
	using std::string;
	using std::vector;

	struct Point
	{
		int x; // increases "rt"
		int y; // increases "up"

		static Point const lt;
		static Point const rt;
		static Point const dn;
		static Point const up;
	};
	Point const Point::lt{ -1, 0 };
	Point const Point::rt{ 1, 0 };
	Point const Point::dn{ 0, -1 };
	Point const Point::up{ 0, 1 };
	Point operator+(const Point& lhs, const Point& rhs)
	{
		return Point{ lhs.x + rhs.x, lhs.y + rhs.y };
	}

	struct Jet : Point
	{
		Jet() : Point(dn) {}
		explicit Jet(char c) : Point(c == '<' ? lt : rt) {}
	};

	class Jetstream
	{
	public:
		explicit Jetstream(string const& data) : jets_(parse(data)) {}
		explicit Jetstream(vector<Jet> jets) : jets_(jets) {}
		Jet next() const
		{
			Jet j = jets_[current_];
			current_ = (current_ + 1) % jets_.size();
			++jets_produced_;
			return j;
		}
		size_t jets_produced() const { return jets_produced_; }
		static vector<Jet> parse(string const& data)
		{
			vector<Jet> jets(data.length());
			std::transform(data.begin(), data.end(), jets.begin(), [](char c) { return Jet(c); });
			return jets;
		}
	private:
		vector<Jet> const jets_;
		size_t mutable jets_produced_ = 0;
		size_t mutable current_ = 0;
	};

	struct RockShape
	{
		vector<Point> const points;

		static RockShape const flat;
		static RockShape const cross;
		static RockShape const backl;
		static RockShape const beam;
		static RockShape const box;
	};
	RockShape const RockShape::flat{
		{
			Point{0,0}, Point{1,0}, Point{2,0}, Point{3,0},
		}
	};
	RockShape const RockShape::cross{
		{
			/*       */ Point{1,2} /*       */,
			Point{0,1}, Point{1,1}, Point{2,1},
			/*       */ Point{1,0} /*       */,
		}
	};
	RockShape const RockShape::backl{
		{
			/*                   */ Point{2,2},
			/*                   */ Point{2,1},
			Point{0,0}, Point{1,0}, Point{2,0},
		}
	};
	RockShape const RockShape::beam{
		{
			Point{0,0},
			Point{0,1},
			Point{0,2},
			Point{0,3},
		}
	};
	RockShape const RockShape::box{
		{
			Point(0,0), Point(1,0),
			Point(0,1), Point(1,1),
		}
	};

	struct Rock
	{
		RockShape shape;
		Point position;
	};

	class RockShapeGenerator
	{
	public:
		RockShapeGenerator() : rock_shapes_({ RockShape::flat, RockShape::cross, RockShape::backl, RockShape::beam, RockShape::box }) {}
		RockShape next() const
		{
			RockShape rs = rock_shapes_[current_];
			current_ = (current_ + 1) % rock_shapes_.size();
			return rs;
		}
	private:
		vector<RockShape> const rock_shapes_;
		size_t mutable current_ = 0;
	};

	class RockGenerator
	{
	public:
		RockGenerator(const RockShapeGenerator& rsg) : rock_shape_generator_(rsg) {}
		Rock next(Point const& position) const
		{
			++rocks_produced_;
			return Rock{ rock_shape_generator_.next(), position };
		}
		size_t rocks_produced() const { return rocks_produced_; }
	private:
		RockShapeGenerator const& rock_shape_generator_;
		size_t mutable rocks_produced_ = 0;
	};

	class Board
	{
	public:
		explicit Board() {}
		int top() const { return static_cast<int>(map.size()); }
		bool isValidLocation(const Rock& rock)
		{
			for (Point const& point : rock.shape.points) {
				Point offset = rock.position + point;
				if (offset.x < 0) return false;
				if (offset.x >= width) return false;
				if (offset.y < 0) return false;
				if (offset.y < map.size() && map[offset.y][offset.x] == '#') return false;
			}
			return true;
		}
		void addRock(const Rock& rock)
		{
			// add the rock to the map
			for (Point const& point : rock.shape.points) {
				Point offset = rock.position + point;
				while (map.size() <= offset.y) { map.push_back(string(width, ' ')); }
				map[offset.y][offset.x] = '#';
			}
		}
		void print(std::ostream& os) const
		{
			if (map.size() == 0) {
				os << "\n";
				os << "+" << string(width, '=') << "+\n";
				return;
			}

			os << "\n";
			std::for_each(
				map.rbegin(),
				map.rend(),
				[&os](string const& s) { os << "|" << s << "|\n"; }
			);
			os << "\\" << string(width, '=') << "/" << std::endl;
		}
	private:
		vector<string> map;
		static const int width = 7;
	};

	string answer_a(vector<string> const& input_data)
	{
		Jetstream jetstream(input_data[0]);
		RockShapeGenerator rsg;
		RockGenerator rg(rsg);
		Board b;

		while (rg.rocks_produced() < 2022) {
			Rock rock = rg.next({ 2, b.top() + 3 });
			Rock rockCopy(rock);
			while (true) {
				Jet j = jetstream.next();
				rockCopy.position = rockCopy.position + j;
				if (b.isValidLocation(rockCopy)) {
					rock.position = rockCopy.position; // accept
				} else {
					rockCopy.position = rock.position; // undo
				}

				rockCopy.position = rockCopy.position + Point::dn;
				if (b.isValidLocation(rockCopy)) {
					rock.position = rockCopy.position; // accept
				} else {
					rockCopy.position = rock.position; // undo
					// could not move down
					break;
				}
			}
			b.addRock(rock);
			//if (rg.rocks_produced() < 11) {
			//	b.print(std::cout);
			//}
		}

		return std::to_string(b.top());
	}

	string answer_b(vector<string> const& input_data)
	{
		return "PENDING";
	}
}
