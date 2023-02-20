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

    struct Position2I
    {
        int x;
        int y;

        static const Position2I up() { return { 0,-1 }; }
        static const Position2I dn() { return { 0,+1 }; }
        static const Position2I lt() { return { -1,0 }; }
        static const Position2I rt() { return { +1,0 }; }

        bool operator<(const Position2I& rhs) const { return x < rhs.x || (x == rhs.x && y < rhs.y); }
        bool operator==(const Position2I& rhs) const { return x == rhs.x && y == rhs.y; }
        Position2I operator+(const Position2I& rhs) const { return { x + rhs.x, y + rhs.y }; }
        Position2I operator-(const Position2I& rhs) const { return { x - rhs.x, y - rhs.y }; }
        int manhattanDistanceFrom(const Position2I& rhs) { Position2I pos{ *this - rhs }; return abs(pos.x) + abs(pos.y); }
    };

    class CharHeightNode
    {
    public:
        CharHeightNode(char c)
        {
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

    class Grid
    {
    public:
        Grid(vector<CharHeightNode> nodes, int w, int h)
            : data_(nodes), width_(w), height_(h)
        {}

        int height() const { return height_; }
        int width() const { return width_; }

        CharHeightNode at(Position2I pos) const { return data_[indexFromPosition(pos)]; }
        void set(Position2I pos, CharHeightNode node) { data_[indexFromPosition(pos)] = node; }

        vector<Position2I> neighbouringPositionsOf(Position2I pos) const
        {
            vector<Position2I> neighbours;

            int maxHeight = height_ - 1;
            int maxWidth = width_ - 1;

            if (pos.y - 1 >= 0) neighbours.push_back(pos + Position2I::up());
            if (pos.y + 1 <= maxHeight) neighbours.push_back(pos + Position2I::dn());
            if (pos.x - 1 >= 0) neighbours.push_back(pos + Position2I::lt());
            if (pos.x + 1 <= maxWidth) neighbours.push_back(pos + Position2I::rt());

            return neighbours;
        }

        Position2I positionFromIndex(int index) const { return { index % width_, index / width_ }; }
        int indexFromPosition(Position2I pos) const { return pos.x + width_ * pos.y; }
    private:
        int height_;
        int width_;
        vector<CharHeightNode> data_;
    };

    class AStar
    {
        struct Node
        {
            Position2I value;
            int cheapestCostToNode = INT_MAX;    // gScore
            int bestGuessCostToFinish = INT_MAX; // fScore
            shared_ptr<Node> cameFrom;

#if defined(__clang__) && defined(__apple_build_version__)
#if __apple_build_version__ <= 14000029
            Node(
                Position2I value = Position2I{},
                int cheapestCostToNode = INT_MAX,
                int bestGuessCostToFinish = INT_MAX
            )
                : value(value)
                , cheapestCostToNode(cheapestCostToNode)
                , bestGuessCostToFinish(bestGuessCostToFinish)
                , cameFrom(nullptr)
            {}
#endif
#endif
        };
    public:
        using EstimateDistanceToTargetFn = function<int(Position2I)>;
        using IsTargetFn = function<bool(Position2I)>;
        using GetNeighboursOfFn = function<vector<Position2I>(Position2I)>;
        using DistanceToMoveFn = function<int(Position2I, Position2I)>;
        using Path = vector<Position2I>;
    private:
        using NodePtr = shared_ptr<Node>;
        using Lookup = map<Position2I, NodePtr>;
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
        {}

        optional<Path> execute(Position2I startValue)
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

                vector<Position2I> posNeighbours = getNeighboursOf(current->value);
                for (Position2I posNeighbour : posNeighbours) {
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

        Path constructPath(NodePtr end)
        {
            Path path;
            NodePtr n = end;
            do { path.push_back(n->value); } while ((n = n->cameFrom));
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
        GridRunner(const Grid& grid)
            : grid_(grid)
        {}
        optional<vector<Position2I>> findOptimalPath(Position2I startPos, Position2I endPos)
        {
            auto estimateDistanceToTarget = [=](Position2I pos) -> int {
                return 25 - grid_.at(pos).height();
            };
            auto isTarget = [=](Position2I pos) -> bool {
                return endPos == pos;
            };
            auto getValidNeighboursOf = [=](Position2I posCurrent) -> vector<Position2I> {
                vector<Position2I> neighbours = grid_.neighbouringPositionsOf(posCurrent);
                if (neighbours.empty())
                    return neighbours;
                else {
                    neighbours.erase(remove_if(
                        begin(neighbours), end(neighbours),
                        [=](Position2I posNeighbour) {
                            return grid_.at(posNeighbour).height() > grid_.at(posCurrent).height() + 1;
                        }),
                        end(neighbours));
                    return neighbours;
                }
            };
            auto distanceToMove = [=](Position2I a, Position2I b) -> int {
                return 1;
            };

            AStar aStar{
                estimateDistanceToTarget,
                isTarget,
                getValidNeighboursOf,
                distanceToMove
            };

            optional<vector<Position2I>> path = aStar.execute(startPos);
            if (path.has_value()) return path.value();
            return {};
        }
    private:
        const Grid& grid_;
    };

    class ReverseGridRunner
    {
    public:
        ReverseGridRunner(const Grid& grid)
            : grid_(grid)
        {}
        optional<vector<Position2I>> findOptimalPath(Position2I startPos, int targetHeight)
        {
            auto estimateDistanceToTarget = [=](Position2I pos) -> int {
                return grid_.at(pos).height();
            };
            auto isTarget = [=](Position2I pos) -> bool {
                return grid_.at(pos).height() == 0;
            };
            auto getValidNeighboursOf = [=](Position2I posCurrent) -> vector<Position2I> {
                vector<Position2I> neighbours = grid_.neighbouringPositionsOf(posCurrent);
                if (neighbours.empty())
                    return neighbours;
                else {
                    neighbours.erase(remove_if(
                        begin(neighbours), end(neighbours),
                        [=](Position2I posNeighbour) {
                            return grid_.at(posNeighbour).height() < grid_.at(posCurrent).height() - 1;
                    return grid_.at(posNeighbour).height() < grid_.at(posCurrent).height() - 1;
                        }),
                        end(neighbours));
                    return neighbours;
                }
            };
            auto distanceToMove = [=](Position2I a, Position2I b) -> int {
                return 1;
            };

            AStar aStar{
                estimateDistanceToTarget,
                isTarget,
                getValidNeighboursOf,
                distanceToMove
            };

            optional<vector<Position2I>> path = aStar.execute(startPos);
            if (path.has_value()) return path.value();
            return {};
        }
    private:
        const Grid& grid_;
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
        Position2I start = grid.positionFromIndex(indexStart);
        Position2I end = grid.positionFromIndex(indexEnd);

        GridRunner runner{ grid };
        optional<vector<Position2I>> path = runner.findOptimalPath(start, end);

        if (path.has_value()) return to_string(path.value().size() - 1);
        return "no path found from S to E";
    }

    string answer_b(const vector<string>& input_data)
    {
        vector<CharHeightNode> nodes;
        int indexStart;
        int height = input_data.size();
        int width = input_data[0].length();

        for (const string& str : input_data) {
            for (char c : str) {
                nodes.push_back(CharHeightNode{ c });
                if (c == 'E') { indexStart = nodes.size() - 1; }
            }
        }

        Grid grid{ nodes, width, height };
        ReverseGridRunner runner{ grid };

        vector<vector<Position2I>> possiblePaths;

        Position2I startPos = grid.positionFromIndex(indexStart);

        optional<vector<Position2I>> path = runner.findOptimalPath(startPos, 0);

        if (path.has_value()) possiblePaths.push_back(path.value());

        auto it = min_element(begin(possiblePaths), end(possiblePaths), [](auto& pathA, auto& pathB) { return pathA.size() < pathB.size(); });
        return to_string(it->size() - 1);
    }


    string answer_b_1(const vector<string>& input_data)
    {
        vector<CharHeightNode> nodes;
        vector<int> indexesStart;
        int indexEnd = 0;
        int height = input_data.size();
        int width = input_data[0].length();

        for (const string& str : input_data) {
            for (char c : str) {
                nodes.push_back(CharHeightNode{ c });
                if (c == 'S' || c == 'a') { indexesStart.push_back(nodes.size() - 1); }
                if (c == 'E') { indexEnd = nodes.size() - 1; }
            }
        }

        Grid grid{ nodes, width, height };
        Position2I endPos = grid.positionFromIndex(indexEnd);
        GridRunner runner{ grid };

        vector<vector<Position2I>> possiblePaths;
        for (int indexStart : indexesStart) {
            Position2I startPos = grid.positionFromIndex(indexStart);

            optional<vector<Position2I>> path = runner.findOptimalPath(startPos, endPos);

            if (path.has_value()) possiblePaths.push_back(path.value());
        }

        auto it = min_element(begin(possiblePaths), end(possiblePaths), [](auto& pathA, auto& pathB) { return pathA.size() < pathB.size(); });
        return to_string(it->size() - 1);
    }
}
