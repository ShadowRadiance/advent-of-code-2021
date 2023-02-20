#include <gtest/gtest.h>

#include <pathfinder.h>

#include <stdexcept>
#include <algorithm>

using namespace advent_of_code::Pathfinder;

SimpleGraph buildSimpleGraph()
{
    SimpleGraph g;
    g.addNode('A', { 'D', 'I', 'B' });
    g.addNode('B', { 'C', 'A' });
    g.addNode('C', { 'D', 'B' });
    g.addNode('D', { 'C', 'A', 'E' });
    g.addNode('E', { 'F', 'D' });
    g.addNode('F', { 'E', 'G' });
    g.addNode('G', { 'F', 'H' });
    g.addNode('H', { 'G' });
    g.addNode('I', { 'A', 'J' });
    g.addNode('J', { 'I' });
    return g;
}

TEST(Pathfinder, SimpleGraph)
{
    SimpleGraph graph = buildSimpleGraph();

    EXPECT_THROW(graph.addNode('H', {}), std::invalid_argument);

    auto a = graph.findNode('A');
    auto d = graph.findNode('D');
    auto j = graph.findNode('J');

    EXPECT_EQ(graph.neighbours(d).size(), 3);

    EXPECT_EQ(graph.cost(d, a), 1);
    EXPECT_THROW(graph.cost(d, j), std::invalid_argument);
}

TEST(Pathfinder, BreadthFirstSearch)
{
    SimpleGraph graph = buildSimpleGraph();
    auto a = graph.findNode('A');
    auto h = graph.findNode('H');
    auto j = graph.findNode('J');

    CameFromMap result = breadth_first_search(graph, a, h);
    auto path = reconstruct_path(a, h, result);
    EXPECT_EQ(path.size(), 6); // ADEFGH

    std::string s;
    std::transform(path.begin(), path.end(), back_inserter(s), [&graph](auto handle) { return graph.id(handle); });
    EXPECT_EQ(s, "ADEFGH");

    result = breadth_first_search(graph, h, j);
    path = reconstruct_path(h, j, result);
    s.clear();
    std::transform(path.begin(), path.end(), back_inserter(s), [&graph](auto handle) { return graph.id(handle); });
    EXPECT_EQ(s, "HGFEDAIJ");
}

TEST(Pathfinder, DijkstraSearch)
{
    SimpleGraph graph = buildSimpleGraph();
    auto a = graph.findNode('A');
    auto h = graph.findNode('H');
    auto j = graph.findNode('J');

    CameFromMap result;
    CostMap costs;
    dijkstra_search(graph, a, h, result, costs);
    auto path = reconstruct_path(a, h, result);
    EXPECT_EQ(path.size(), 6); // ADEFGH

    std::string s;
    std::transform(path.begin(), path.end(), back_inserter(s), [&graph](auto handle) { return graph.id(handle); });
    EXPECT_EQ(s, "ADEFGH");

    result.clear();
    costs.clear();
    dijkstra_search(graph, h, j, result, costs);
    path = reconstruct_path(h, j, result);
    s.clear();
    std::transform(path.begin(), path.end(), back_inserter(s), [&graph](auto handle) { return graph.id(handle); });
    EXPECT_EQ(s, "HGFEDAIJ");
}

TEST(Pathfinder, AStarSearch)
{
    SimpleGraph graph = buildSimpleGraph();
    auto a = graph.findNode('A');
    auto h = graph.findNode('H');
    auto j = graph.findNode('J');

    CameFromMap result;
    CostMap costs;
    auto fn = [](auto _next, auto _goal) { return 1; };
    a_star_search(graph, a, h, fn, result, costs);
    auto path = reconstruct_path(a, h, result);
    EXPECT_EQ(path.size(), 6); // ADEFGH

    std::string s;
    std::transform(path.begin(), path.end(), back_inserter(s), [&graph](auto handle) { return graph.id(handle); });
    EXPECT_EQ(s, "ADEFGH");

    result.clear();
    costs.clear();
    a_star_search(graph, h, j, fn, result, costs);
    path = reconstruct_path(h, j, result);
    s.clear();
    std::transform(path.begin(), path.end(), back_inserter(s), [&graph](auto handle) { return graph.id(handle); });
    EXPECT_EQ(s, "HGFEDAIJ");
}
