#include <iostream>
#include <chrono>

#include <days.h>
#include <utility.h>

#if defined(__clang__) && defined(__apple_build_version__)
    #if __apple_build_version__ <= 14000029
        #include <iomanip>
        #include <sstream>
        std::ostream& operator<<(std::ostream& os, const std::chrono::microseconds ms) {
            std::stringstream ss;
            ss.flags(os.flags());
            ss.imbue(os.getloc());
            ss.precision(os.precision());
            ss << ms.count() << "µs";
            return os << ss.str();
        }
    #endif
#endif

std::tuple<std::string, std::chrono::microseconds> time(auto fn, const std::vector<std::string>& data) {
    auto start = std::chrono::high_resolution_clock::now();
    auto result = fn(data);
    auto stop = std::chrono::high_resolution_clock::now();
    auto duration = duration_cast<std::chrono::microseconds>(stop - start);
    return { result, duration };
}

std::string path(int n) {
    if (n < 10) return std::string{ "./data/day-0" + std::to_string(n) + ".txt" };
    return std::string{ "./data/day-" + std::to_string(n) + ".txt" };
}

using AnswerFn = std::string(*)(const std::vector<std::string>&);
auto methods = std::vector<AnswerFn>{
    day_01::answer_a,
    day_01::answer_b,
    day_02::answer_a,
    day_02::answer_b,
    day_03::answer_a,
    day_03::answer_b,
    day_04::answer_a,
    day_04::answer_b,
    day_05::answer_a,
    day_05::answer_b,
    day_06::answer_a,
    day_06::answer_b,
    day_07::answer_a,
    day_07::answer_b,
    day_08::answer_a,
    day_08::answer_b,
    day_09::answer_a,
    day_09::answer_b,
    day_10::answer_a,
    day_10::answer_b,
    day_11::answer_a,
    day_11::answer_b,
    day_12::answer_a,
    day_12::answer_b,
    day_13::answer_a,
    day_13::answer_b,
    day_14::answer_a,
    day_14::answer_b,
    day_15::answer_a,
    day_15::answer_b,
    day_16::answer_a,
    day_16::answer_b,
    day_17::answer_a,
    day_17::answer_b,
    day_18::answer_a,
    day_18::answer_b,
    day_19::answer_a,
    day_19::answer_b,
    day_20::answer_a,
    day_20::answer_b,
    day_21::answer_a,
    day_21::answer_b,
    day_22::answer_a,
    day_22::answer_b,
    day_23::answer_a,
    day_23::answer_b,
    day_24::answer_a,
    day_24::answer_b,
    day_25::answer_a,
    day_25::answer_b,
};

int main(int argc, char** argv)
{
    using std::cout;

    for (int i = 0; i <= 49; i++) {
        auto day = 1 + (i / 2);
        auto data = load_data(path(day));
        auto fn = methods[i];

        if (fn != day_18::answer_a && fn != day_18::answer_b) continue;

        auto [result, duration] = time(fn, data);

        cout << "Day " << ((day < 10) ? "0" : "") << day << " Answer " << ((i % 2 == 0) ? "A" : "B") << ": "
             << "(" << std::setw(10) << duration << ") " << result << "\n";
    }
}
