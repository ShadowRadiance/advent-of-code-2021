#include <gtest/gtest.h>

#include <utility.h>
#include <filesystem>
#include <fstream>

namespace fs = std::filesystem;

using std::ofstream;
using std::string;
using std::vector;
using std::cerr;

class UtilityFileTest : public testing::Test
{
protected:
    static void SetUpTestSuite()
    {
        if (filename.empty()) {
            filename = fs::temp_directory_path().string() + "test.txt";
            ofstream out{ filename };
            out << "1234" << "\n"
                << "2345" << "\n"
                << "" << "\n"
                << "3456" << "\n"
                << "" << "\n"
                << "This is a long string" << "\n"
                << "" << "\n"
                << "" << "\n"
                ;
        }
    }

    static string filename;

};

string UtilityFileTest::filename;

TEST_F(UtilityFileTest, LoadData)
{
    cerr << "[          ] current_path = " << fs::current_path() << "\n";
    vector<string> expected{ "1234", "2345", "", "3456", "", "This is a long string", "" , "" };
    EXPECT_EQ(expected, load_data(filename));
    // executing it twice should return the same list each time
    EXPECT_EQ(expected, load_data(filename));
}

