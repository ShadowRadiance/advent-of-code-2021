#include <days.h>

#include <iostream>
#include <iomanip>
#include <optional>
#include <algorithm>

namespace day_08
{
    using std::optional;
    using std::string;
    using std::vector;
    using std::ostream;
    using std::cout;
    using std::endl;
    using std::setw;
    using std::for_each;
    using std::max;
    using std::to_string;

    class Forest;

    struct Tree
    {
        int height = 0;
        bool isVisible;

        int maxNorth{ 0 };
        int maxEast{ 0 };
        int maxSouth{ 0 };
        int maxWest{ 0 };

        int scenicScore{ 0 };
    };

    ostream& operator<<(ostream& os, const Tree& tree)
    {
        os
            << "{"
            << tree.height
            << "{" << tree.maxNorth << tree.maxEast << tree.maxSouth << tree.maxWest << "}"
            << (tree.isVisible ? "!" : "_")
            << setw(3) << tree.scenicScore
            << "}";
        return os;
    }

    class Forest
    {
    public:
        void addTree(int height = 0, bool visible = false)
        {
            trees_.push_back({ height, visible });
        }

        void endRow()
        {
            if (width_ == 0) width_ = trees_.size();
            height_++;
        }

        void determineRelationships()
        {
            // each grid location stores: maxNorth, maxEast, etc. and isVisible
            //	as well as seeNorth, seeEast, etc
            // pass 1 (forward - or "moving W->E and N->S"
            //   look north
            //     record max height in that direction (if current > max, is visible)
            //     record number of trees visible in that direction (
            //		 examine the tree in that direction
            //		 - if it is shorter, use its number + 1
            //		 - if it is larger or the same, just use 1
            //   look west (& do the same)
            for (size_t idx = 0; idx < trees_.size(); idx++) {
                size_t x = idx % width_;
                size_t y = idx / width_;
                Tree& currentTree = trees_[idx];
                if (y > 0) { // unless top row
                    Tree& northTree = trees_[indexFromXY(x, y - 1)];
                    currentTree.maxNorth = max(northTree.maxNorth, northTree.height);
                    if (currentTree.height > currentTree.maxNorth) currentTree.isVisible = true;
                }
                if (x > 0) {
                    Tree& westTree = trees_[indexFromXY(x - 1, y)];
                    currentTree.maxWest = max(westTree.maxWest, westTree.height);
                    if (currentTree.height > currentTree.maxWest) currentTree.isVisible = true;
                }
            }

            // pass 2 (backward - or "moving W<-E and N<-S"
            //   look south and east as above 
            for (int idx = trees_.size() - 1; idx >= 0; --idx) {
                size_t x = idx % width_;
                size_t y = idx / width_;
                Tree& currentTree = trees_[idx];
                if (y < height_ - 1) {
                    Tree& southTree = trees_[indexFromXY(x, y + 1)];
                    currentTree.maxSouth = max(southTree.maxSouth, southTree.height);
                    if (currentTree.height > currentTree.maxSouth) currentTree.isVisible = true;
                }
                if (x < width_ - 1) {
                    Tree& eastTree = trees_[indexFromXY(x + 1, y)];
                    currentTree.maxEast = max(eastTree.maxEast, eastTree.height);
                    if (currentTree.height > currentTree.maxEast) currentTree.isVisible = true;
                }
            }
        }

        void determineScenicScores()
        {
            for (int idx = 0; idx < trees_.size(); idx++) {
                int x = idx % width_;			// 2
                int y = idx / width_;			// 1
                bool outer = y == 0 || x == 0 || y == width_ - 1 || x == height_ - 1;
                if (outer) { continue; }

                Tree& currentTree = trees_[idx];	// 5h
                int n{ 0 }, e{ 0 }, w{ 0 }, s{ 0 };

                // look north
                while (y - n > 0) {
                    n++;
                    if (trees_[indexFromXY(x, y - n)].height >= currentTree.height) break;
                }
                // look east
                while (x + e < width_ - 1) {
                    e++;
                    if (trees_[indexFromXY(x + e, y)].height >= currentTree.height) break;
                }
                // look south
                while (y + s < height_ - 1) {
                    s++;
                    if (trees_[indexFromXY(x, y + s)].height >= currentTree.height) break;
                }
                // look west
                while (x - w > 0) {
                    w++;
                    if (trees_[indexFromXY(x - w, y)].height >= currentTree.height) break;
                }

                currentTree.scenicScore = n * e * w * s;
            }
        }

        size_t numberVisible()
        {
            // summarize: count is visible
            return count_if(trees_.begin(), trees_.end(), [](Tree& tree) { return tree.isVisible; });
        }

        size_t highestScenicScore()
        {
            auto it = max_element(trees_.begin(), trees_.end(),
                                  [](Tree& a, Tree& b) {return a.scenicScore < b.scenicScore; });
            Tree& t = *it;
            return t.scenicScore;
        }

        friend ostream& operator<<(ostream& os, const Forest& forest);
    private:
        int width_ = 0;
        int height_ = 0;
        vector<Tree> trees_;

        size_t indexFromXY(size_t x, size_t y)
        {
            return x + y * width_;
        }
    };

    ostream& operator<<(ostream& os, const Forest& forest)
    {
        for (size_t idx = 0; idx < forest.trees_.size(); idx++) {
            os << forest.trees_[idx];
            if (idx % forest.width_ == forest.width_ - 1) os << endl;
        }
        return os;
    }

    using strings = vector<string>;

    Forest generateForest(const strings& input_data)
    {
        Forest forest;
        for (size_t y{ 0 }; y < input_data.size(); y++) {
            const string& row = input_data[y];
            vector<Tree> treeRow;
            for (size_t x{ 0 }; x < row.length(); x++) {
                char cHeight = row[x];
                int nHeight = cHeight - '0';
                bool outer = y == 0 || x == 0 || y == input_data.size() - 1 || x == row.size() - 1;
                forest.addTree(nHeight, outer);
            }
            forest.endRow();
        }
        return forest;
    }

    string answer_a(const strings& input_data)
    {
        // read input data into grid (set height, set is visible if on edge)
        Forest forest = generateForest(input_data);
        forest.determineRelationships();
        forest.determineScenicScores();
        return to_string(forest.numberVisible());
    }

    string answer_b(const vector<string>& input_data)
    {
        Forest forest = generateForest(input_data);
        forest.determineRelationships();
        forest.determineScenicScores();
        return to_string(forest.highestScenicScore());
    }
} // namespace day_08
