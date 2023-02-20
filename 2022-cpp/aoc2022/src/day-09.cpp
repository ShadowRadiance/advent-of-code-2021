#include <days.h>

#include <sstream>
#include <algorithm>
#include <cassert>
#include <iostream>

namespace day_09
{
    using std::string;
    using std::vector;
    using std::stringstream;
    using std::unique;
    using std::to_string;
    using std::cout; using std::endl; using std::ostream;

    using strings = vector<string>;

    struct vec2d
    {
        int x;
        int y;

        bool operator==(const vec2d& other) const
        {
            return x == other.x && y == other.y;
        }

        vec2d operator+(const vec2d& other) const
        {
            return { x + other.x, y + other.y };
        }

        vec2d operator-(const vec2d& other) const
        {
            return { x - other.x, y - other.y };
        }

        vec2d unit() const
        {
            return {
                (x == 0) ? 0 : x / abs(x),
                (y == 0) ? 0 : y / abs(y)
            };
        }
    };

    ostream& operator<<(ostream& os, const vec2d& v)
    {
        return os << v.x << "," << v.y;
    }

    class command
    {
    public:
        command(const string& data)
        {
            stringstream ss{ data };
            char c;
            ss >> c >> distance;
            direction = vec2dFromChar(c);
        }
        vec2d direction;
        int distance;

    private:
        vec2d vec2dFromChar(char c)
        {
            switch (c) {
            case 'U': return { 0, -1 };
            case 'D': return { 0, 1 };
            case 'L': return { -1, 0 };
            case 'R': return { 1, 0 };
            default: return { 0, 0 };
            }
        }
    };

    bool locationExists(const vec2d& location, const vector<vec2d>& vecs)
    {
        auto it = find(vecs.begin(), vecs.end(), location);
        return it != vecs.end();
    }

    bool tailOkay(const vec2d& head, const vec2d& tail)
    {
        return (abs(head.x - tail.x) <= 1 && abs(head.y - tail.y) <= 1); // adjacent/overlapping
    }

    vec2d tailMove(const vec2d& head, const vec2d& tail)
    {
        return (head - tail).unit();
    }

    string answer_a(const strings& input_data)
    {
        vector<command> commands{ input_data.begin(), input_data.end() };

        vec2d head{ 0,0 };
        vec2d tail{ 0,0 };
        vector<vec2d> tailHistory{ tail };

        for (auto& command : commands) {
            for (size_t i = 0; i < command.distance; i++) {
                head = head + command.direction;
                if (!tailOkay(head, tail)) {
                    vec2d move = tailMove(head, tail);
                    // cout << "move tail(" << tail << ") toward head(" << head << ") by ("<<move<<")" << endl;
                    tail = tail + move;
                    if (!locationExists(tail, tailHistory))
                        tailHistory.push_back(tail);
                }
            }
        }

        return to_string(tailHistory.size());
    }

    string answer_b(const strings& input_data)
    {
        vector<command> commands{ input_data.begin(), input_data.end() };

        vector<vec2d> knots{ 10, { 0, 0 } };
        vector<vec2d> tailHistory{ { 0, 0 } };

        for (auto& command : commands) {
            for (size_t i = 0; i < command.distance; i++) {
                knots[0] = knots[0] + command.direction;

                for (size_t knotIndex = 1; knotIndex < knots.size(); knotIndex++) {
                    if (!tailOkay(knots[knotIndex - 1], knots[knotIndex])) {
                        vec2d move = tailMove(knots[knotIndex - 1], knots[knotIndex]);
                        knots[knotIndex] = knots[knotIndex] + move;

                        if (knotIndex == 9) {
                            if (!locationExists(knots[knotIndex], tailHistory))
                                tailHistory.push_back(knots[knotIndex]);
                        }
                    }
                }
            }
        }

        return to_string(tailHistory.size());
    }
}







