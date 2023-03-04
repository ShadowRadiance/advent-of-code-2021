#include <days.h>

#include <algorithm>
#include <exception>
#include <format>
#include <iostream>
#include <tuple>
#include <variant>
#include <cassert>
#include <numeric>

namespace day_22
{
    using std::string;
    using std::tuple;
    using std::variant;
    using std::vector;

    class Map
    {
    public:
        const char NOTHING = ' ';
        const char WALL = '#';
        const char SPACE = '.';

        Map(vector<string> const& data) : map(data)
        {
            auto it = std::max_element(
                data.begin(), data.end(),
                [](string const& lhs, string const& rhs) { return lhs.length() < rhs.length(); });
            int maxLength = it->length();
            std::for_each(
                map.begin(), map.end(),
                [maxLength](string& str) { str.append(maxLength - str.length(), ' '); });
        }
        bool isWall(int x, int y) const { return map[y][x] == WALL; }
        char at(int x, int y) const { return map[y][x]; }
        int firstValidXInRow(int y) const { return map[y].find_first_not_of(' '); }
        int lastValidXInRow(int y) const { return map[y].find_last_not_of(' '); }
        int firstValidYInCol(int x) const { return col(x).find_first_not_of(' '); }
        int lastValidYInCol(int x) const { return col(x).find_last_not_of(' '); }

        int height() const { return map.size(); }
        int width() const { return map[0].size(); }
    protected:
        vector<string> map;
    private:
        string col(int x) const
        {
            string s;
            std::transform(map.begin(), map.end(), std::back_inserter(s), [x](string const& row) { return row[x]; });
            return s;
        }
    };

    class FlatMap : public Map
    {
    public:
        FlatMap(vector<string> const& data) : Map(data) {}
    private:
    };

    class Face
    {
    public:
        Face(int top, int left, int bottom, int right)
            : top_(top), left_(left), bottom_(bottom), right_(right)
        {}
    private:
        int top_;
        int left_;
        int bottom_;
        int right_;
    };

    class CubeMap : public Map
    {
    public:
        CubeMap(vector<string> const& data) : Map(data)
        {
            determineFaces();
        }
    private:
        void determineFaces()
        {
            // find the highest common factor of the width and height
            faceSize_ = std::gcd(height(), width());
            int h = height() / faceSize_;
            int w = width() / faceSize_;
            assert(h == 5 && w == 2 || h == 2 && w == 5 || h == 3 && w == 4 || h == 4 && w == 5);

            // for each "face-sized row"

            for (int i = 0; i < h; ++i) {
                int top = i * faceSize_;
                int bottom = top + faceSize_ - 1;
                int left = firstValidXInRow(top);
                int right = left + faceSize_ - 1;
                while (left < lastValidXInRow(top)) {
                    faces_.push_back(Face(top, left, bottom, right));
                    left += faceSize_;
                    right += faceSize_;
                }
            }

            return;
        }
        vector<Face> faces_;
        int faceSize_;
    };

    enum class Facing
    {
        right,
        down,
        left,
        up
    };

    char facingChar(Facing facing)
    {
        switch (facing) {
        case Facing::right:
            return '>';
            break;
        case Facing::down:
            return 'v';
            break;
        case Facing::left:
            return '<';
            break;
        case Facing::up:
            return '^';
            break;
        default:
            return ' ';
            break;
        }
    }

    Facing nextFacingRight(Facing facing)
    {
        switch (facing) {
        case Facing::right:
            return Facing::down;
            break;
        case Facing::down:
            return Facing::left;
            break;
        case Facing::left:
            return Facing::up;
            break;
        case Facing::up:
            return Facing::right;
            break;
        default:
            throw std::invalid_argument{ "Invalid facing" };
        }
    }

    Facing nextFacingLeft(Facing facing)
    {
        switch (facing) {
        case Facing::right:
            return Facing::up;
            break;
        case Facing::up:
            return Facing::left;
            break;
        case Facing::left:
            return Facing::down;
            break;
        case Facing::down:
            return Facing::right;
            break;
        default:
            throw std::invalid_argument{ "Invalid facing" };
        }
    }

    class Actor
    {
    public:
        Actor(int x, int y, Facing facing) : x_(x), y_(y), facing_(facing) {}
        int x() const { return x_; }
        int y() const { return y_; }
        Facing facing() const { return facing_; }
        void turnLeft() { facing_ = nextFacingLeft(facing_); }
        void turnRight() { facing_ = nextFacingRight(facing_); }
        void moveForward(int n)
        {
            switch (facing_) {
            case Facing::right:
                x_ += n;
                break;
            case Facing::down:
                y_ += n;
                break;
            case Facing::left:
                x_ -= n;
                break;
            case Facing::up:
                y_ -= n;
                break;
            default:
                throw std::invalid_argument{ "Invalid facing" };
            }
        }
        void teleportTo(int x, int y)
        {
            x_ = x;
            y_ = y;
        }

        tuple<int, int> lookingAt() const
        {
            switch (facing_) {
            case Facing::right:
                return { x_ + 1, y_ };
            case Facing::down:
                return { x_, y_ + 1 };
            case Facing::left:
                return { x_ - 1, y_ };
            case Facing::up:
                return { x_, y_ - 1 };
            default:
                throw std::invalid_argument{ "facing_ is invalid" };
            }
        }

    private:
        int x_;
        int y_;
        Facing facing_;
    };

    class Instruction
    {
    public:
        Instruction(char c) : data(c) {}
        Instruction(int distance) : data(distance) {}

        void apply(Actor& actor, FlatMap const& map)
        {
            if (std::holds_alternative<char>(data)) {
                // turn actor
                switch (std::get<char>(data)) {
                case 'R':
                    actor.turnRight();
                    break;
                case 'L':
                    actor.turnLeft();
                    break;
                default:
                    break;
                }
            } else {
                // move actor forward, jumping over sides, stopped by rocks
                int moves = std::get<int>(data);
                while (moves > 0) {
                    auto [x, y] = actor.lookingAt();
                    bool teleport = false;
                    // if facing the edge of available space, determine target "on the other side"
                    if (actor.facing() == Facing::right && x > map.lastValidXInRow(y)) {
                        x = map.firstValidXInRow(y);
                        teleport = true;
                    } else if (actor.facing() == Facing::left && x < map.firstValidXInRow(y)) {
                        x = map.lastValidXInRow(y);
                        teleport = true;
                    } else if (actor.facing() == Facing::down && y > map.lastValidYInCol(x)) {
                        y = map.firstValidYInCol(x);
                        teleport = true;
                    } else if (actor.facing() == Facing::up && y < map.firstValidYInCol(x)) {
                        y = map.lastValidYInCol(x);
                        teleport = true;
                    }

                    if (map.isWall(x, y)) {
                        break;
                    }
                    if (teleport) {
                        actor.teleportTo(x, y);
                    } else {
                        actor.moveForward(1);
                    }

                    moves--;
                }
            }
        }

        void apply(Actor& actor, CubeMap const& map)
        {
            if (std::holds_alternative<char>(data)) {
                // turn actor
                switch (std::get<char>(data)) {
                case 'R':
                    actor.turnRight();
                    break;
                case 'L':
                    actor.turnLeft();
                    break;
                default:
                    break;
                }
            } else {
                // move actor forward, jumping over sides, stopped by rocks
                int moves = std::get<int>(data);
                while (moves > 0) {
                    auto [x, y] = actor.lookingAt();
                    bool teleport = false;

                    // if facing the edge of a cube, determine target "across the edge"
                    // remember to set teleport to true
                    // remember this will change facing (relative to the overhead map)



                    if (map.isWall(x, y)) {
                        break;
                    }
                    if (teleport) {
                        actor.teleportTo(x, y);
                    } else {
                        actor.moveForward(1);
                    }

                    moves--;
                }
            }
        }
    private:
        variant<char, int> data;
    };

    int parseInt(auto& it, auto end)
    {
        size_t size{ 0 };
        int parsed = std::stoi(string{ it, end }, &size);
        it += size;
        return parsed;
    }

    vector<Instruction> parseInstructions(string const& line)
    {
        vector<Instruction> result;

        // 10R5L5R10L4R5L5 // alternating number/R/L
        auto it = line.begin();
        auto end = line.end();
        while (it != end) {
            if (std::isdigit(*it)) {
                int num = parseInt(it, end);
                result.push_back(Instruction{ num });
            } else {
                char c = *it;
                result.push_back(Instruction{ c });
                ++it;
            }
        }

        return result;
    }

    void log(std::ostream& out, Map const& map, Actor const& actor)
    {
        return;
        for (int y = 0; y < map.height(); y++) {
            for (int x = 0; x < map.width(); x++) {
                if (actor.x() == x && actor.y() == y) {
                    out << facingChar(actor.facing());
                } else {
                    out << map.at(x, y);
                }
            }
            out << "\n";
        }
        out << "\n\n";
    }

    string answer_a(vector<string> const& input_data)
    {
        auto itBlankLine = std::find(input_data.begin(), input_data.end(), string{});
        FlatMap map(vector<string>(input_data.begin(), itBlankLine));
        string lastLine = input_data[input_data.size() - 1];
        vector<Instruction> instructions = parseInstructions(lastLine);

        Actor actor{ map.firstValidXInRow(0), 0, Facing::right };
        for (auto& instruction : instructions) {
            instruction.apply(actor, map);
            log(std::cout, map, actor);
        }

        return std::to_string(
            4 * (actor.x() + 1) +
            1000 * (actor.y() + 1) +
            static_cast<int>(actor.facing()));
    }

    string answer_b(const vector<string>& input_data)
    {
        auto itBlankLine = std::find(input_data.begin(), input_data.end(), string{});
        CubeMap map(vector<string>(input_data.begin(), itBlankLine));
        string lastLine = input_data[input_data.size() - 1];
        vector<Instruction> instructions = parseInstructions(lastLine);

        Actor actor{ map.firstValidXInRow(0), 0, Facing::right };
        for (auto& instruction : instructions) {
            instruction.apply(actor, map);
            log(std::cout, map, actor);
        }

        return std::to_string(
            4 * (actor.x() + 1) +
            1000 * (actor.y() + 1) +
            static_cast<int>(actor.facing()));
    }
}




/*
# TEST PATTERN
#
#       +-----+             1-/-3-/-5---6-X-1                                             #
#       |     |             1-X-2-X-5---4---1                                             #
#       |  +--1-----+       2---3---4-/-6-\-2                                             #
#       |  |  |     |                                                                     #
#    +--2--3--4--+  |       1X2 (180 rotation either way)  2-3 (no rotations)             #
#    |  |  |  |  |  |       1/3 (1->3 rot-lt 3->1 rot-rt)  2X5 (180 rotation either way)  #
#    |  |  +--5--6--+       1-4 (no rotations)             2/6 (2->6 rot-rt 6->2 rot-lt)  #
#    |  |     |  |          1X6 (180 rotation either way)  3-4 (no rotations)             #
#    |  +-----+  |                                                                        #
#    |           |                                                                        #
#    +-----------+                                                                        #
#
#
# REAL PATTERN
#
#     +-------+  +----+      1---3---5-\-6-/-1                                            #
#     |       |  |    |      1---2-X-5---4-X-1                                            #
#     | +-----1--2--+ |      2-/-3-\-4---6---2                                            #
#     | |     |  |  | |                                                                   #
#     | |  +--3--+  | |      1-2 (no rotations)             2/3 (2->3 rot-lt 3->2 rot-rt) #
#     | |  |  |     | |      1-3 (no rotations)             2X5 (180 rotation either way) #
#     | +--4--5-----+ |      1X4 (180 rotation either way)  2-6 (no rotations)            #
#     |    |  |       |      1\6 (1->6 rot-rt 6->1 rot-lt)  3\4 (3->4 rot-rt 4->3 rot-lt) #
#     +----6--+       |                                                                   #
#          |          |                                                                   #
#          +----------+                                                                   #
#
*/
