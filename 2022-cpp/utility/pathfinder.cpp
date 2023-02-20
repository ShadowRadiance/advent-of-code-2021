#include <pathfinder.h>

#include <algorithm>
#include <functional>
#include <queue>
#include <unordered_map>
#include <unordered_set>
#include <utility>
#include <vector>
#include <stdexcept>
#include <optional>

namespace advent_of_code
{
    namespace Pathfinder
    {
        using std::greater;
        using std::invalid_argument;
        using std::pair;
        using std::priority_queue;
        using std::queue;
        using std::unordered_map;
        using std::unordered_set;
        using std::vector;

        using std::reverse;
        using std::find;

        using NodeHandle = IGraph::NodeHandle;

        bool SimpleGraph::Node::hasEdge(char id) const
        {
            return edges.end() != find(edges.begin(), edges.end(), id);
        }

        vector<NodeHandle> SimpleGraph::neighbours(NodeHandle handle) const
        {
            Node node = nodes[index(handle)];

            vector<NodeHandle> result(node.edges.size());
            transform(node.edges.begin(),
                      node.edges.end(),
                      result.begin(),
                      [&](char c) { return findNode(c); });
            return result;
        }

        size_t SimpleGraph::numberOfNodes() const
        {
            return nodes.size();
        }

        double SimpleGraph::cost(NodeHandle from, NodeHandle to) const
        {
            // from and to must be adjacent
            const Node& fromNode = nodes[index(from)];
            const Node& toNode = nodes[index(to)];

            if (!fromNode.hasEdge(toNode.id)) throw invalid_argument("from and to must be adjacent");

            return 1;
        }

        vector<NodeHandle> SimpleGraph::allNodes() const
        {
            vector<NodeHandle> result(nodes.size());
            // convert each index into a handle
            for (ptrdiff_t i = 0; i < static_cast<ptrdiff_t>(nodes.size()); i++) {
                result[i] = handle(i);
            }
            return result;
        }

        NodeHandle SimpleGraph::findNode(char id) const
        {
            return handle(index(id));
        }

        NodeHandle SimpleGraph::addNode(char id, std::vector<char> edges)
        {
            return addNode(Node{ id, edges });
        }

        NodeHandle SimpleGraph::addNode(const Node& node)
        {
            if (index(node.id) != nodes.size()) throw invalid_argument("id already exists");

            nodes.push_back(node);
            return handle(index(node.id));
        }

        char SimpleGraph::id(NodeHandle handle) const
        {
            return nodes[index(handle)].id;
        }

        NodeHandle SimpleGraph::handle(ptrdiff_t index) const
        {
            // the 0 handle, nullptr, would equate to index==-1
            // valid handles are > 0

            return reinterpret_cast<NodeHandle>(index + 1);
        }

        ptrdiff_t SimpleGraph::index(NodeHandle handle) const
        {
            // the 0 handle, nullptr, would equate to index==-1
            // valid handles are > 0

            return reinterpret_cast<ptrdiff_t>(handle) - 1;
        }

        ptrdiff_t SimpleGraph::index(char id) const
        {
            auto it = find_if(nodes.begin(), nodes.end(), [id](auto& node) { return node.id == id; });
            return it - nodes.begin();
        }

        CameFromMap breadth_first_search(const IGraph& graph,
                                         NodeHandle start,
                                         NodeHandle goal
        )
        {
            queue<NodeHandle> frontier;
            frontier.push(start);

            CameFromMap came_from;
            came_from[start] = start;

            while (!frontier.empty()) {
                NodeHandle current = frontier.front();
                frontier.pop();
                if (current == goal)
                    break;
                // yield current if yield-fn provided
                for (NodeHandle next : graph.neighbours(current)) {
                    if (came_from.find(next) == came_from.end()) {
                        frontier.push(next);
                        came_from[next] = current;
                    }
                }
            }
            return came_from;
        }

        template <typename T, typename priority_t>
        struct PriorityQueue
        {
            typedef pair<priority_t, T> Element;
            priority_queue<Element, vector<Element>, greater<Element>> elements;

            bool empty() const { return elements.empty(); }

            void put(T item, priority_t priority) { elements.emplace(priority, item); }

            T get()
            {
                T best_item = elements.top().second;
                elements.pop();
                return best_item;
            }
        };

        // if you pass nullptr for goal, all paths from start are determined
        void dijkstra_search(
            const IWeightedGraph& graph,
            NodeHandle start,
            NodeHandle goal,
            CameFromMap& came_from,
            CostMap& cost_so_far)
        {
            PriorityQueue<NodeHandle, double> frontier;
            frontier.put(start, 0);
            came_from[start] = start;
            cost_so_far[start] = 0;

            while (!frontier.empty()) {
                NodeHandle current = frontier.get();
                if (current == goal)
                    break;
                for (NodeHandle next : graph.neighbours(current)) {
                    double new_cost = cost_so_far[current] + graph.cost(current, next);
                    if (cost_so_far.find(next) == cost_so_far.end() || new_cost < cost_so_far[next]) {
                        cost_so_far[next] = new_cost;
                        came_from[next] = current;
                        frontier.put(next, new_cost);
                    }
                }
            }
        }

        vector<NodeHandle> reconstruct_path(NodeHandle start,
                                            NodeHandle goal,
                                            CameFromMap came_from)
        {
            vector<NodeHandle> path;
            NodeHandle current = goal;
            if (came_from.find(goal) == came_from.end())
                return path;
            while (current != start) {
                path.push_back(current);
                current = came_from[current];
            }
            path.push_back(start);
            reverse(path.begin(), path.end());
            return path;
        }

        void a_star_search(const IWeightedGraph& graph,
                           NodeHandle start,
                           NodeHandle goal,
                           HeuristicFunction heuristic_fn,
                           CameFromMap& came_from,
                           CostMap& cost_so_far)
        {
            PriorityQueue<NodeHandle, double> frontier;
            frontier.put(start, 0);
            came_from[start] = start;
            cost_so_far[start] = 0;

            while (!frontier.empty()) {
                NodeHandle current = frontier.get();
                if (current == goal)
                    break;

                for (NodeHandle next : graph.neighbours(current)) {
                    double new_cost = cost_so_far[current] + graph.cost(current, next);
                    if (cost_so_far.find(next) == cost_so_far.end() || new_cost < cost_so_far[next]) {
                        cost_so_far[next] = new_cost;
                        double priority = new_cost + heuristic_fn(next, goal);
                        frontier.put(next, priority);
                        came_from[next] = current;
                    }
                }
            }
        }

        void floyd_warshall(const IWeightedGraph& graph, FloydWarshallDistances& distances, FloydWarshallNexts& nexts)
        {
            vector<NodeHandle> allNodes = graph.allNodes();
            size_t size = allNodes.size();

            for (NodeHandle v : allNodes) {
                distances[v] = FloydWarshallDistances::mapped_type{}; // std::unordered_map<NodeHandle, double>{};
                nexts[v] = FloydWarshallNexts::mapped_type{}; // std::unordered_map<NodeHandle, NodeHandle>{};
                for (NodeHandle u : allNodes) {
                    distances[v][u] = (std::numeric_limits<double>::max)();
                    nexts[v][u] = nullptr;
                }
            }

            // HANDLE VERTICES
            for (NodeHandle v : allNodes) {
                distances[v][v] = 0;
                nexts[v][v] = v;
            }

            // HANDLE DIRECT EDGES
            for (NodeHandle u : allNodes) {
                for (NodeHandle v : graph.neighbours(u)) {
                    distances[u][v] = graph.cost(u, v);
                    nexts[u][v] = v;
                }
            }

            // DETERMINE SHORTEST PATHS
            for (NodeHandle k : allNodes) {
                for (NodeHandle u : allNodes) {
                    for (NodeHandle v : allNodes) {
                        double newDistance = distances[u][k] + distances[k][v];
                        if (distances[u][v] > newDistance) {
                            distances[u][v] = newDistance;
                            nexts[u][v] = nexts[u][k];
                        }
                    }
                }
            }
        }

        std::vector<NodeHandle> reconstruct_path(
            NodeHandle from,
            NodeHandle to,
            const FloydWarshallNexts& nexts)
        {
            vector<NodeHandle> path;

            if (nexts.at(from).at(to) == nullptr) return path;

            path.push_back(from);
            while (from != to) {
                from = nexts.at(from).at(to);
                path.push_back(from);
            }

            return path;
        }
    }
}