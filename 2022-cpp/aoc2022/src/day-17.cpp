#include <days.h>
#include <algorithm>
#include <iostream>
#include <unordered_map>
#include <tuple>

namespace day_17
{
    using std::string;
    using std::vector;

    struct Point
    {
        int64_t x; // increases "rt"
        int64_t y; // increases "up"

        static Point const lt;
        static Point const rt;
        static Point const dn;
        static Point const up;
    };
    Point const Point::lt{-1, 0};
    Point const Point::rt{1, 0};
    Point const Point::dn{0, -1};
    Point const Point::up{0, 1};
    Point operator+(const Point &lhs, const Point &rhs)
    {
        return Point{lhs.x + rhs.x, lhs.y + rhs.y};
    }

    struct Jet : Point
    {
        Jet() : Point(dn) {}
        explicit Jet(char c) : Point(c == '<' ? lt : rt) {}
    };

    class Jetstream
    {
    public:
        explicit Jetstream(string const &data) : jets_(parse(data)) {}
        explicit Jetstream(vector<Jet> jets) : jets_(jets) {}
        Jet next() const
        {
            Jet j = jets_[current_];
            current_ = (current_ + 1) % repetitionLength();
            ++jets_produced_;
            return j;
        }
        int64_t repetitionLength() const { return static_cast<int64_t>(jets_.size()); }
        int64_t jetsProduced() const { return jets_produced_; }
        static vector<Jet> parse(string const &data)
        {
            vector<Jet> jets(data.length());
            std::transform(data.begin(), data.end(), jets.begin(), [](char c)
                           { return Jet(c); });
            return jets;
        }

    private:
        vector<Jet> const jets_;
        int64_t mutable jets_produced_ = 0;
        int64_t mutable current_ = 0;
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
            Point{0, 0},
            Point{1, 0},
            Point{2, 0},
            Point{3, 0},
        }};
    RockShape const RockShape::cross{
        {
            /*       */ Point{1, 2} /*       */,
            Point{0, 1},
            Point{1, 1},
            Point{2, 1},
            /*       */ Point{1, 0} /*       */,
        }};
    RockShape const RockShape::backl{
        {
            /*                   */ Point{2, 2},
            /*                   */ Point{2, 1},
            Point{0, 0},
            Point{1, 0},
            Point{2, 0},
        }};
    RockShape const RockShape::beam{
        {
            Point{0, 0},
            Point{0, 1},
            Point{0, 2},
            Point{0, 3},
        }};
    RockShape const RockShape::box{
        {
            Point{0, 0},
            Point{1, 0},
            Point{0, 1},
            Point{1, 1},
        }};

    struct Rock
    {
        RockShape shape;
        Point position;
    };

    class RockShapeGenerator
    {
    public:
        RockShapeGenerator() : rock_shapes_({RockShape::flat, RockShape::cross, RockShape::backl, RockShape::beam, RockShape::box}) {}
        RockShape next() const
        {
            RockShape rs = rock_shapes_[current_];
            current_ = (current_ + 1) % repetitionLength();
            return rs;
        }
        int64_t repetitionLength() const { return static_cast<int64_t>(rock_shapes_.size()); }

    private:
        vector<RockShape> const rock_shapes_;
        int64_t mutable current_ = 0;
    };

    class RockGenerator
    {
    public:
        RockGenerator(const RockShapeGenerator &rsg) : rock_shape_generator_(rsg) {}
        Rock next(Point const &position) const
        {
            ++rocks_produced_;
            return Rock{rock_shape_generator_.next(), position};
        }
        int64_t rocksProduced() const { return rocks_produced_; }
        void pretend_rocks_produced(int64_t fake_rocks) { rocks_produced_ += fake_rocks; }

    private:
        RockShapeGenerator const &rock_shape_generator_;
        int64_t mutable rocks_produced_ = 0;
    };

    class Board
    {
    public:
        explicit Board() {}
        int64_t top() const { return static_cast<int64_t>(map.size()); }
        bool isValidLocation(const Rock &rock)
        {
            for (Point const &point : rock.shape.points)
            {
                Point offset = rock.position + point;
                if (offset.x < 0)
                    return false;
                if (offset.x >= width)
                    return false;
                if (offset.y < 0)
                    return false;
                if (offset.y < top() && map[offset.y][offset.x] == '#')
                    return false;
            }
            return true;
        }
        void addRock(const Rock &rock)
        {
            // add the rock to the map
            for (Point const &point : rock.shape.points)
            {
                Point offset = rock.position + point;
                while (top() <= offset.y)
                {
                    map.push_back(string(width, ' '));
                }
                map[offset.y][offset.x] = '#';
            }
        }
        void print(std::ostream &os) const
        {
            if (map.size() == 0)
            {
                os << "\n";
                os << "+" << string(width, '=') << "+\n";
                return;
            }

            os << "\n";
            std::for_each(
                map.rbegin(),
                map.rend(),
                [&os](string const &s)
                { os << "|" << s << "|\n"; });
            os << "\\" << string(width, '=') << "/" << std::endl;
        }

    private:
        vector<string> map;
        static const int width = 7;
    };

    int64_t solve_with_simulation(vector<string> const &input_data, int64_t requiredRocks)
    {
        Jetstream jetstream(input_data[0]);
        RockShapeGenerator rsg;
        RockGenerator rg(rsg);
        Board b;

        using seen_key = std::tuple<int64_t, int64_t>; // rocks index, jet index
        struct seen_key_hasher
        {
            size_t operator()(const seen_key &key) const
            {
                std::hash<int64_t> hasher;
                auto h1 = hasher(std::get<0>(key));
                auto h2 = hasher(std::get<1>(key));
                return h1 ^ (h2 << 1);
            }
        };
        using seen_value = std::tuple<int64_t, int64_t, int64_t>; // timesSeen, rocksProduced, top
        struct
        {
            std::unordered_map<seen_key, seen_value, seen_key_hasher> seen;
            int64_t addedByRepetition = 0;
        } state;

        while (rg.rocksProduced() < requiredRocks)
        {
            Rock rock = rg.next({2, b.top() + 3});
            Rock rockCopy(rock);
            while (true)
            {
                Jet j = jetstream.next();
                rockCopy.position = rockCopy.position + j;
                if (b.isValidLocation(rockCopy))
                {
                    rock.position = rockCopy.position; // accept
                }
                else
                {
                    rockCopy.position = rock.position; // undo
                }

                rockCopy.position = rockCopy.position + Point::dn;
                if (b.isValidLocation(rockCopy))
                {
                    rock.position = rockCopy.position; // accept
                }
                else
                {
                    rockCopy.position = rock.position; // undo
                    // could not move down
                    break;
                }
            }
            b.addRock(rock);

            // look for a cycle
            if (state.addedByRepetition == 0)
            {
                seen_key key{
                    rg.rocksProduced() % rsg.repetitionLength(),
                    jetstream.jetsProduced() % jetstream.repetitionLength()};

                // at third occurrence of key, the values in the seen-map repeat
                // add as many of them as possible without hitting the goal piece_count
                if (state.seen.contains(key))
                {
                    auto &[timesSeen, oldRocksProduced, oldTop] = state.seen[key];
                    if (timesSeen == 2)
                    {
                        auto deltaTop = b.top() - oldTop;
                        auto deltaRocks = rg.rocksProduced() - oldRocksProduced;
                        auto repeats = (requiredRocks - b.top()) / deltaRocks;
                        state.addedByRepetition += repeats * deltaTop;
                        rg.pretend_rocks_produced(repeats * deltaRocks);
                    }

                    // update seen map with key: (timesSeen+1, rocksProduced, top)
                    state.seen[key] = std::make_tuple(timesSeen + 1, rg.rocksProduced(), b.top());
                }
                else
                {
                    // update seen map with key: (1, rocksProduced, top)
                    state.seen[key] = std::make_tuple(int64_t{1}, rg.rocksProduced(), b.top());
                }
            }
        }
        return b.top() + state.addedByRepetition;
    }

    string answer_a(vector<string> const &input_data)
    {
        return std::to_string(solve_with_simulation(input_data, 2022));
    }

    string answer_b(vector<string> const &input_data)
    {
        int64_t one_trillion = 1000000000000;
        return std::to_string(solve_with_simulation(input_data, one_trillion));
    }
}
