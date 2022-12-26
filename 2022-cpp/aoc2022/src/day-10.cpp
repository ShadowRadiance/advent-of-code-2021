#include <days.h>

#include <algorithm>
#include <iterator>
#include <numeric>
#include <array>
#include <sstream>

#include <cpu_10.h>

namespace day_10
{
    using std::string;
    using std::vector;
    using std::array;
    using std::stringstream;
    using std::to_string;

    using strings = vector<string>;

    string answer_a(const strings& input_data)
    {
        CPU_10 cpu{ input_data };

        vector<int> signalStrengths;
        for (size_t idx = 0; idx < cpu.cycles_required(); idx++) {
            cpu.begin_cycle();
            if (cpu.current_cycle() % 40 == 20) {
                signalStrengths.push_back(cpu.signal_strength());
            }
            cpu.end_cycle();
        }

        return to_string(accumulate(signalStrengths.begin(), signalStrengths.end(), 0));
    }

    string answer_b(const strings& input_data)
    {
        // 40x6 CRT
        //  CRT draws pixel N (1..240) into x-location (0..239)%40 (0-39, 6 times)
        // sprite location where cpu's X register is middle of 3-char sprite
        //  if CRT draw-location is within sprite - draw LIT (#) else draw DARK (.)

        CPU_10 cpu{ input_data };
        stringstream crt;
        int spriteCentre{ 0 };
        for (int currentPixel = 0; currentPixel < 240; currentPixel++) {
            int xPixel = currentPixel % 40;
            cpu.begin_cycle();
            spriteCentre = cpu.x();
            if (xPixel == 0) {
                crt << '\n';
            }
            bool overlaps = (xPixel >= (spriteCentre - 1) && xPixel <= (spriteCentre + 1));
            crt << (overlaps ? '#' : '.');
            cpu.end_cycle();
        }
        crt << '\n';
        crt << '\n';

        return crt.str();
    }
}
