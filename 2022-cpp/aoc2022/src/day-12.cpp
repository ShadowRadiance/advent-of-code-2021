#include <days.h>
#include <algorithm>
#include <iterator>
#include <cstdlib>
#include <functional>
#include <optional>
#include <map>
#include <memory>

namespace day_12
{
    using std::string;
    using std::vector;
    using std::to_string;
    using std::function;
    using std::optional;
    using std::map;
    using std::shared_ptr, std::make_shared;
    using std::make_heap, std::pop_heap, std::sort_heap, std::push_heap;

    struct Position {
        int x;
        int y;

        static const Position up() { return { 0,-1 }; }
        static const Position dn() { return { 0,+1 }; }
        static const Position lt() { return { -1,0 }; }
        static const Position rt() { return { +1,0 }; }

        bool operator<(const Position& rhs) const { return x < rhs.x || (x == rhs.x && y < rhs.y); }
        bool operator==(const Position& rhs) const { return x == rhs.x && y == rhs.y; }
        Position operator+(const Position& rhs) const { return { x + rhs.x, y + rhs.y }; }
        Position operator-(const Position& rhs) const { return { x - rhs.x, y - rhs.y }; }
        int manhattanDistanceFrom(const Position& rhs) { Position pos{ *this - rhs }; return abs(pos.x) + abs(pos.y); }
    };

    class CharHeightNode {
    public:
        CharHeightNode(char c) {
            switch (c) {
            case 'E': height_ = 25; break;
            case 'S': height_ = 0; break;
            default: height_ = c - 'a'; break;
            }
        }
        int height() const { return height_; }
    private:
        int height_;
    };

    class Grid {
    public:
        Grid(vector<CharHeightNode> nodes, int w, int h)
            : data_(nodes), width_(w), height_(h)
        {
        }

        int height() const { return height_; }
        int width() const { return width_; }

        CharHeightNode at(Position pos) const { return data_[indexFromPosition(pos)]; }
        void set(Position pos, CharHeightNode node) { data_[indexFromPosition(pos)] = node; }

        vector<Position> neighbouringPositionsOf(Position pos) const {
            vector<Position> neighbours;

            int maxHeight = height_ - 1;
            int maxWidth = width_ - 1;

            if (pos.y - 1 >= 0)
            {
                neighbours.push_back(pos + Position::up());
            }

            if (pos.y + 1 <= maxHeight)
            {
                neighbours.push_back(pos + Position::dn());
            }

            if (pos.x - 1 >= 0)
            {
                neighbours.push_back(pos + Position::lt());
            }

            if (pos.x + 1 <= maxWidth)
            {
                neighbours.push_back(pos + Position::rt());
            }

            if (neighbours.empty()) return {};
            else return neighbours;
        }

        Position positionFromIndex(int index) const { return { index % width_, index / width_ }; }
        int indexFromPosition(Position pos) const { return pos.x + width_ * pos.y; }
    private:
        int height_;
        int width_;
        vector<CharHeightNode> data_;
    };

    class AStar {
        struct Node
        {
            Position value;
            int cheapestCostToNode = INT_MAX;    // gScore
            int bestGuessCostToFinish = INT_MAX; // fScore
            shared_ptr<Node> cameFrom;
        };
    public:
        using EstimateDistanceToTargetFn = function<int(Position)>;
        using IsTargetFn = function<bool(Position)>;
        using GetNeighboursOfFn = function<vector<Position>(Position)>;
        using DistanceToMoveFn = function<int(Position, Position)>;
        using Path = vector<Position>;
    private:
        using NodePtr = shared_ptr<Node>;
        using Lookup = map<Position, NodePtr>;
        using MinHeap = vector<NodePtr>;

    public:
        AStar(
            EstimateDistanceToTargetFn estimateDistanceToTargetFn,
            IsTargetFn isTargetFn,
            GetNeighboursOfFn getNeighboursOfFn,
            DistanceToMoveFn distanceToMoveFn
        )
            : estimateDistanceToTarget(estimateDistanceToTargetFn)
            , isTarget(isTargetFn)
            , getNeighboursOf(getNeighboursOfFn)
            , distanceToMove(distanceToMoveFn)
        {
        }

        optional<Path> execute(Position startValue)
        {
            Lookup allNodes;

            MinHeap frontier;
            auto orderHeap = [](NodePtr a, NodePtr b) {
                return a->bestGuessCostToFinish > b->bestGuessCostToFinish;
            };
            // Usually make_heap and friends order in largest to smallest order, using a function that
            // determines less_than ordering, that is the HIGHEST value is at the "top" of the heap
            // Since we want the LOWEST values on the top of the heap we need to give it the reverse function
            // using greater_than

            NodePtr start = make_shared<Node>(startValue, 0, estimateDistanceToTarget(startValue));

            allNodes.insert_or_assign(startValue, start);
            frontier.push_back(start);
            push_heap(begin(frontier), end(frontier), orderHeap); // unnecessary but whatever

            while (!frontier.empty()) {
                // move node in frontier with lowest bestGuessCostToFinish to end
                pop_heap(begin(frontier), end(frontier), orderHeap);
                NodePtr current = frontier.back();
                // and remove from frontier
                frontier.pop_back();

                // are we done?
                if (isTarget(current->value)) return constructPath(current);

                vector<Position> posNeighbours = getNeighboursOf(current->value);
                for (Position posNeighbour : posNeighbours) {
                    if (!allNodes.contains(posNeighbour)) {
                        allNodes.insert({ posNeighbour, make_shared<Node>(posNeighbour) });
                    }
                    NodePtr neighbour = allNodes[posNeighbour];

                    int score = current->cheapestCostToNode + distanceToMove(current->value, posNeighbour);
                    if (score < neighbour->cheapestCostToNode) {
                        neighbour->cameFrom = current;
                        neighbour->cheapestCostToNode = score;
                        neighbour->bestGuessCostToFinish = score + estimateDistanceToTarget(neighbour->value);
                        auto it = std::find(begin(frontier), end(frontier), neighbour);
                        bool frontierContainsNeighbour = (it != end(frontier));
                        if (!frontierContainsNeighbour) {
                            frontier.push_back(neighbour);
                            push_heap(begin(frontier), end(frontier), orderHeap);
                        }
                    }
                }
            }

            return {};
        }

        Path constructPath(NodePtr end) {
            Path path;
            NodePtr n = end;
            do { path.push_back(n->value); } while (n = n->cameFrom);
            return path;
        }
    private:
        EstimateDistanceToTargetFn estimateDistanceToTarget;
        IsTargetFn isTarget;
        GetNeighboursOfFn getNeighboursOf;
        DistanceToMoveFn distanceToMove;
    };

    class GridRunner
    {
    public:
        GridRunner(const Grid& grid, Position start, Position end)
            : grid_(grid), start_(start), end_(end)
        {
        }
        optional<vector<Position>> findOptimalPath() {
            auto estimateDistanceToTarget = [=](Position pos) -> int {
                return 25 - grid_.at(pos).height();
            };
            auto isTarget = [=](Position pos) -> bool {
                return end_ == pos;
            };
            auto getValidNeighboursOf = [=](Position posCurrent) -> vector<Position> {
                vector<Position> neighbours = grid_.neighbouringPositionsOf(posCurrent);
                if (neighbours.empty())
                    return neighbours;
                else {
                    neighbours.erase(remove_if(
                        begin(neighbours), end(neighbours),
                        [=](Position posNeighbour)
                        {
                            return grid_.at(posNeighbour).height() > grid_.at(posCurrent).height() + 1;
                        }),
                        end(neighbours));
                    return neighbours;
                }
            };
            auto distanceToMove = [=](Position a, Position b) -> int {
                return 1;
            };

            AStar aStar{
                estimateDistanceToTarget,
                isTarget,
                getValidNeighboursOf,
                distanceToMove
            };

            optional<vector<Position>> path = aStar.execute(start_);
            if (path.has_value()) return path.value();
            return {};
        }
    private:
        const Grid& grid_;
        Position start_{ 0,0 };
        Position end_{ 0,0 };
    };

    string answer_a(const vector<string>& input_data)
    {
        vector<CharHeightNode> nodes;
        int indexStart = 0;
        int indexEnd = 0;
        int height = input_data.size();
        int width = input_data[0].length();

        for (const string& str : input_data) {
            for (char c : str) {
                nodes.push_back(CharHeightNode{ c });
                if (c == 'S') { indexStart = nodes.size() - 1; }
                if (c == 'E') { indexEnd = nodes.size() - 1; }
            }
        }

        Grid grid{ nodes, width, height };
        Position start = grid.positionFromIndex(indexStart);
        Position end = grid.positionFromIndex(indexEnd);

        GridRunner runner{ grid, start, end };
        optional<vector<Position>> path = runner.findOptimalPath();

        if (path.has_value()) return to_string(path.value().size()-1);
        return "no path found from S to E";
    }

    string answer_b(const vector<string>& input_data)
    {
        return "PENDING";
    }
}
