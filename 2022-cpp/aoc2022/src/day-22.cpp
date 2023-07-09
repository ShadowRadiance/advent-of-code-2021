#include <days.h>

#include <algorithm>
#include <iostream>
#include <tuple>
#include <variant>
#include <map>

namespace day_22
{
    static bool test_mode = false;
    void set_test_mode(bool testing) { test_mode = testing; }

    using std::string;
    using std::tuple;
    using std::variant;
    using std::vector;
    using std::map;

    class Map
    {
    public:
        const char WALL = '#';

        explicit Map(vector<string> const& data) : map(data)
        {
            auto it = std::max_element(
                data.begin(), data.end(),
                [](string const& lhs, string const& rhs) { return lhs.length() < rhs.length(); });
            uint maxLength = it->length();
            std::for_each(
                map.begin(), map.end(),
                [maxLength](string& str) { str.append(maxLength - str.length(), ' '); });
        }

        [[nodiscard]] bool isWall(uint x, uint y) const { return map[y][x] == WALL; }
        [[nodiscard]] char at(uint x, uint y) const { return map[y][x]; }
        [[nodiscard]] uint firstValidXInRow(uint y) const { return map[y].find_first_not_of(' '); }
        [[nodiscard]] uint lastValidXInRow(uint y) const { return map[y].find_last_not_of(' '); }
        [[nodiscard]] uint firstValidYInCol(uint x) const { return col(x).find_first_not_of(' '); }
        [[nodiscard]] uint lastValidYInCol(uint x) const { return col(x).find_last_not_of(' '); }

        [[nodiscard]] uint height() const { return map.size(); }
        [[nodiscard]] uint width() const { return map[0].size(); }
    protected:
        vector<string> map;
    private:
        [[nodiscard]] string col(uint x) const
        {
            string s;
            std::transform(map.begin(), map.end(), std::back_inserter(s), [x](string const& row) { return row[x]; });
            return s;
        }
    };

    class FlatMap : public Map
    {
    public:
        explicit FlatMap(vector<string> const& data) : Map(data) {}
    private:
    };

    class Face
    {
    public:
        explicit Face(uint size, uint top, uint left)
            : top_(top * size), left_(left * size), bottom_(((top+1)*size)-1), right_(((left+1)*size)-1)
        {}

        [[nodiscard]] bool includesLocation(uint x, uint y) const {
            if (x+1==0 || y+1==0) return false;

            return (x >= left_ && x <= right_ && y >= top_ && y <= bottom_);
        }

        [[nodiscard]] uint top() const { return top_; }
        [[nodiscard]] uint left() const { return left_; }
        [[nodiscard]] uint bottom() const { return bottom_; }
        [[nodiscard]] uint right() const { return right_; }
    private:
        uint top_;
        uint left_;
        uint bottom_;
        uint right_;
    };

    enum class Facing
    {
        right,
        down,
        left,
        up
    };

    Facing nextFacingRight(Facing facing)
    {
        switch (facing) {
            case Facing::right:
                return Facing::down;
            case Facing::down:
                return Facing::left;
            case Facing::left:
                return Facing::up;
            case Facing::up:
                return Facing::right;
            default:
                throw std::invalid_argument{ "Invalid facing" };
        }
    }

    Facing nextFacingLeft(Facing facing)
    {
        switch (facing) {
            case Facing::right:
                return Facing::up;
            case Facing::up:
                return Facing::left;
            case Facing::left:
                return Facing::down;
            case Facing::down:
                return Facing::right;
            default:
                throw std::invalid_argument{ "Invalid facing" };
        }
    }

    class Actor
    {
    public:
        Actor(uint x, uint y, Facing facing) : x_(x), y_(y), facing_(facing) {}
        [[nodiscard]] uint x() const { return x_; }
        [[nodiscard]] uint y() const { return y_; }
        [[nodiscard]] Facing facing() const { return facing_; }
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
        void teleportTo(uint x, uint y)
        {
            x_ = x;
            y_ = y;
        }

        [[nodiscard]] tuple<uint, uint> lookingAt() const
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
        uint x_;
        uint y_;
        Facing facing_;
    };

    class Rule {
    public:
        explicit Rule(uint targetFace, uint turns)
            : targetFace_(targetFace), leftTurns_(turns)
        {}
        [[nodiscard]] uint leftTurns() const { return leftTurns_; }
        [[nodiscard]] uint targetFace() const { return targetFace_; }
    private:
        uint targetFace_{0};
        uint leftTurns_{0};
    };

    class CubeMap : public Map
    {
    public:
        explicit CubeMap(vector<string> const& data) : Map(data)
        {
            determineFaces();
        }

        [[nodiscard]] uint faceIdFrom(uint x, uint y) const {
            auto [id, face] = facePairFrom(x,y);
            return id;
        }

        [[nodiscard]] const Face& faceFrom(uint x, uint y) const {
            auto [id, face] = facePairFrom(x,y);
            return face;
        }

        [[nodiscard]] tuple<uint, uint, uint> relativeLocationOnFace(
            Actor& actor
        ) const {
            uint currentFaceId = faceIdFrom(actor.x(), actor.y());
            const Face& currFace = faces_.at(currentFaceId);
            const std::map<Facing, Rule>& rules = rules_.at(currentFaceId);
            const Rule& rule = rules.at(actor.facing());
            const Face& newFace = faces_.at(rule.targetFace());

            uint newX, newY;

            uint distance_from_left = actor.x() - currFace.left();
            uint distance_from_top = actor.y() - currFace.top();
            switch (actor.facing()) {
                case Facing::up:
                    switch(rule.leftTurns()) {
                        case 0:
                            newX = newFace.left() + distance_from_left;
                            newY = newFace.bottom();
                            break;
                        case 1:
                            newX = newFace.right();
                            newY = newFace.bottom() - distance_from_left;
                            break;
                        case 2:
                            newX = newFace.right() - distance_from_left;
                            newY = newFace.top();
                            break;
                        case 3:
                            newX = newFace.left();
                            newY = newFace.top() + distance_from_left;
                            break;
                    }
                    break;
                case Facing::down:
                    switch(rule.leftTurns()) {
                        case 0:
                            newX = newFace.left() + distance_from_left;
                            newY = newFace.top();
                            break;
                        case 1:
                            newX = newFace.left();
                            newY = newFace.bottom() - distance_from_left;
                            break;
                        case 2:
                            newX = newFace.right() - distance_from_left;
                            newY = newFace.bottom();
                            break;
                        case 3:
                            newX = newFace.right();
                            newY = newFace.top() + distance_from_left;
                            break;
                    }
                    break;
                case Facing::left:
                    switch(rule.leftTurns()) {
                        case 0:
                            newX = newFace.right();
                            newY = newFace.top() + distance_from_top;
                            break;
                        case 1:
                            newX = newFace.left() + distance_from_top;
                            newY = newFace.top();
                            break;
                        case 2:
                            newX = newFace.left();
                            newY = newFace.bottom() - distance_from_top;
                            break;
                        case 3:
                            newX = newFace.right() - distance_from_top;
                            newY = newFace.bottom();
                            break;
                    }
                    break;
                case Facing::right:
                    switch(rule.leftTurns()) {
                        case 0:
                            newX = newFace.left();
                            newY = newFace.top() + distance_from_top;
                            break;
                        case 1:
                            newX = newFace.left() + distance_from_top;
                            newY = newFace.bottom();
                            break;
                        case 2:
                            newX = newFace.right();
                            newY = newFace.bottom() - distance_from_top;
                            break;
                        case 3:
                            newX = newFace.right() - distance_from_top;
                            newY = newFace.top();
                            break;
                    }
                    break;
            }

            return {newX, newY, rule.leftTurns()};
        }
    private:
        [[nodiscard]] tuple<uint, const Face&> facePairFrom(uint x, uint y) const {
            auto it = std::find_if(
                faces_.begin(),
                faces_.end(),
                [x,y](auto &pair) {return pair.second.includesLocation(x,y);}
            );
            if (it==faces_.end()) throw std::invalid_argument("no faces with x, y");
            return {it->first, it->second};
        }

        void determineFaces()
        {
            if (test_mode) {
                faceSize_ = 4;

                faces_.insert_or_assign(1, Face(faceSize_, 0, 2));
                faces_.insert_or_assign(2, Face(faceSize_, 1, 0));
                faces_.insert_or_assign(3, Face(faceSize_, 1, 1));
                faces_.insert_or_assign(4, Face(faceSize_, 1, 2));
                faces_.insert_or_assign(5, Face(faceSize_, 2, 2));
                faces_.insert_or_assign(6, Face(faceSize_, 2, 3));

                rules_.insert_or_assign(1, std::map<Facing, Rule>{
                    {Facing::left,  Rule(3,1)},
                    {Facing::right, Rule(6,2)},
                    {Facing::up,    Rule(2,2)},
                    {Facing::down,  Rule(4,0)}
                });
                rules_.insert_or_assign(2, std::map<Facing, Rule>{
                    {Facing::left,  Rule(6,3)},
                    {Facing::right, Rule(3,0)},
                    {Facing::up,    Rule(1,2)},
                    {Facing::down,  Rule(5,2)}
                });
                rules_.insert_or_assign(3, std::map<Facing, Rule>{
                    {Facing::left,  Rule(2,0)},
                    {Facing::right, Rule(4,0)},
                    {Facing::up,    Rule(1,3)},
                    {Facing::down,  Rule(5,1)}
                });
                rules_.insert_or_assign(4, std::map<Facing, Rule>{
                    {Facing::left,  Rule(3,0)},
                    {Facing::right, Rule(6,3)},
                    {Facing::up,    Rule(1,0)},
                    {Facing::down,  Rule(5,0)}
                });
                rules_.insert_or_assign(5, std::map<Facing, Rule>{
                    {Facing::left,  Rule(3,3)},
                    {Facing::right, Rule(6,0)},
                    {Facing::up,    Rule(4,0)},
                    {Facing::down,  Rule(2,2)}
                });
                rules_.insert_or_assign(6, std::map<Facing, Rule>{
                    {Facing::left,  Rule(5,0)},
                    {Facing::right, Rule(1,2)},
                    {Facing::up,    Rule(4,1)},
                    {Facing::down,  Rule(2,1)}
                });
            } else {
                faceSize_ = 50;

                faces_.insert_or_assign(1, Face{faceSize_, 0, 1});
                faces_.insert_or_assign(2, Face{faceSize_, 0, 2});
                faces_.insert_or_assign(3, Face{faceSize_, 1, 1});
                faces_.insert_or_assign(4, Face{faceSize_, 2, 0});
                faces_.insert_or_assign(5, Face{faceSize_, 2, 1});
                faces_.insert_or_assign(6, Face{faceSize_, 3, 0});

                rules_.insert_or_assign(1, std::map<Facing, Rule>{
                    {Facing::left,  Rule(4, 2)},
                    {Facing::right, Rule(2, 0)},
                    {Facing::up,    Rule(6, 3)},
                    {Facing::down,  Rule(3, 0)}
                });
                rules_.insert_or_assign(2, std::map<Facing, Rule>{
                    {Facing::left,  Rule(1, 0)},
                    {Facing::right, Rule(5, 3)},
                    {Facing::up,    Rule(6, 0)},
                    {Facing::down,  Rule(3, 1)}
                });
                rules_.insert_or_assign(3, std::map<Facing, Rule>{
                    {Facing::left,  Rule(4, 3)},
                    {Facing::right, Rule(2, 3)},
                    {Facing::up,    Rule(1, 0)},
                    {Facing::down,  Rule(5, 0)}
                });
                rules_.insert_or_assign(4, std::map<Facing, Rule>{
                    {Facing::left,  Rule(1, 2)},
                    {Facing::right, Rule(5, 0)},
                    {Facing::up,    Rule(3, 1)},
                    {Facing::down,  Rule(6, 0)}
                });
                rules_.insert_or_assign(5, std::map<Facing, Rule>{
                    {Facing::left,  Rule(4, 0)},
                    {Facing::right, Rule(2, 2)},
                    {Facing::up,    Rule(3, 0)},
                    {Facing::down,  Rule(6, 3)}
                });
                rules_.insert_or_assign(6, std::map<Facing, Rule>{
                    {Facing::left,  Rule(1, 3)},
                    {Facing::right, Rule(5, 1)},
                    {Facing::up,    Rule(4, 0)},
                    {Facing::down,  Rule(2, 0)}
                });
            }
        }
        std::map<uint, Face> faces_;
        uint faceSize_{};
        std::map<uint, std::map<Facing, Rule>> rules_;
    };

    class Instruction
    {
    public:
        explicit Instruction(char c) : data(c) {}
        explicit Instruction(int distance) : data(distance) {}

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
                    } else if (actor.facing() == Facing::left && x < map.firstValidXInRow(y) || x+1==0) {
                        x = map.lastValidXInRow(y);
                        teleport = true;
                    } else if (actor.facing() == Facing::down && y > map.lastValidYInCol(x)) {
                        y = map.firstValidYInCol(x);
                        teleport = true;
                    } else if (actor.facing() == Facing::up && y < map.firstValidYInCol(x) || y+1==0) {
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

                    const Face& currentFace = map.faceFrom(actor.x(), actor.y());
                    // if facing the edge of a cube, determine target "across the edge"
                    if (currentFace.includesLocation(x,y)) {
                        if (map.isWall(x, y)) {
                            break;
                        }
                        actor.moveForward(1);
                   } else {
                        auto [newX, newY, leftTurns] = map.relativeLocationOnFace(actor);

                        if (map.isWall(newX, newY)) {
                            break;
                        }

                        // teleport!
                        actor.teleportTo(newX, newY);
                        // remember this will change facing (relative to the overhead map)
                        for(int i = 0; i<leftTurns; i++) actor.turnLeft();
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
                result.emplace_back( num );
            } else {
                char c = *it;
                result.emplace_back( c );
                ++it;
            }
        }

        return result;
    }

    char facingChar(Facing facing) {
        switch (facing) {
            case Facing::right: return '>';
            case Facing::down: return 'v';
            case Facing::left: return '<';
            case Facing::up: return '^';
            default: return ' ';
        }
    }

    void log(std::ostream& out, Map const& map, Actor const& actor)
    {
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

    void tinylog(std::ostream& out, CubeMap const& map, Actor const& actor)
    {
        const Face& currFace = map.faceFrom(actor.x(), actor.y());
        uint currFaceId = map.faceIdFrom(actor.x(), actor.y());

        out << "[" << actor.x() << "," << actor.y() << ",f:" << facingChar(actor.facing()) << "]";
        out << " (" << actor.x() - currFace.left() <<  "," << actor.y() - currFace.top() << ") of face " << currFaceId << "\n";
    }

    string answer_a(vector<string> const& input_data)
    {
        auto itBlankLine = std::find(input_data.begin(), input_data.end(), string{});
        FlatMap map(vector<string>(input_data.begin(), itBlankLine));
        const string& lastLine = input_data[input_data.size() - 1];
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
        const string& lastLine = input_data[input_data.size() - 1];
        vector<Instruction> instructions = parseInstructions(lastLine);

        Actor actor{ map.firstValidXInRow(0), 0, Facing::right };
        tinylog(std::cout, map, actor);
        for (auto& instruction : instructions) {
            instruction.apply(actor, map);
            // log(std::cout, map, actor);
            tinylog(std::cout, map, actor);
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
#       +-----+             1-/-3-/-5---6-X-1              1-5, 2-4, 3-6 NOT CONNECTED    #
#       |     |             1-X-2-X-5---4---1                                             #
#       |  +--1-----+       2---3---4-/-6-\-2                                             #
#       |  |  |     |                                                                     #
#    +--2--3--4--+  |       1X2 (180 rotation either way)  2-3 (no rotations)             #
#    |  |  |  |  |  |       1/3 (1->3 rot-lt 3->1 rot-rt)  2X5 (180 rotation either way)  #
#    |  |  +--5--6--+       1-4 (no rotations)             2/6 (2->6 rot-rt 6->2 rot-lt)  #
#    |  |     |  |          1X6 (180 rotation either way)                                 #
#    |  +-----+  |                                         4-5 (no rotations)             #
#    |           |          3-4 (no rotations)             4\6 (4->6 rot-rt 6->4 rot-lt)  #
#    +-----------+          3/5 (3->5 rot-lt 5->3 rot-rt)  5-6 (no rotations)             #
#
#

 Face(1, left: (3,1), right: (6,2), up: (2,2), down(4,0)),
 Face(2, left: (6,3), right: (3,0), up: (1,2), down(5,2)),
 Face(3, left: (2,0), right: (4,0), up: (1,3), down(5,1)),
 Face(4, left: (3,0), right: (6,3), up: (1,0), down(5,0)),
 Face(5, left: (3,3), right: (6,0), up: (4,0), down(2,2)),
 Face(6, left: (5,0), right: (1,2), up: (4,1), down(2,1)),

# REAL PATTERN
#
#     +-------+  +----+      1---3---5-\-6-/-1              1-5, 2-4, 3-6 NOT CONNECTED   #
#     |       |  |    |      1---2-X-5---4-X-1                                            #
#     | +-----1--2--+ |      2-/-3-\-4---6---2                                            #
#     | |     |  |  | |                                                                   #
#     | |  +--3--+  | |      1-2 (no rotations)             2/3 (2->3 rot-lt 3->2 rot-rt) #
#     | |  |  |     | |      1-3 (no rotations)             2X5 (180 rotation either way) #
#     | +--4--5-----+ |      1X4 (180 rotation either way)  2-6 (no rotations)            #
#     |    |  |       |      1\6 (1->6 rot-rt 6->1 rot-lt)                                #
#     +----6--+       |                                     4-5 (no rotations)            #
#          |          |      3\4 (3->4 rot-lt 4->3 rot-rt)  4-6 (no rotations)            #
#          +----------+      3-5 (no rotations)             5\6 (5->6 rot-rt 6->5 rot-lt) #
#

 Face(1, left: (4,2), right: (2,0), up: (6,3), down(3,0)),
 Face(2, left: (1,0), right: (5,2), up: (6,0), down(3,1)),
 Face(3, left: (4,3), right: (2,3), up: (1,0), down(5,0)),
 Face(4, left: (1,2), right: (5,0), up: (3,1), down(6,0)),
 Face(5, left: (4,0), right: (2,2), up: (3,0), down(6,3)),
 Face(6, left: (1,3), right: (5,1), up: (4,0), down(2,0)),

 */
