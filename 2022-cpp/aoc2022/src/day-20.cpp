#include <days.h>
#include <algorithm>
#include <iostream>

namespace day_20
{
    using std::string;
    using std::vector;
    using std::ostream;

    struct Value
    {
        int64_t actualValue = 0;
        int originalIndex = 0;
    };

    ostream& operator<<(ostream& os, vector<Value> const& values)
    {
        int first = true;
        for (Value const& value : values) {
            if (!first) os << ",";
            os << value.actualValue;
            first = false;
        }
        return os;
    }

    class Mixer
    {
    public:
        Mixer(vector<Value>&& sequence) : sequence_(std::move(sequence)) {}
        void mix()
        {
            int64_t size = sequence_.size();

            vector<Value> newSequence(sequence_);

            for (int i = 0; i < sequence_.size(); i++) {
                auto itItemToMove = std::find_if(
                    newSequence.begin(), newSequence.end(),
                    [i](auto& value) { return value.originalIndex == i; }
                );
                int64_t currentIndex = itItemToMove - newSequence.begin();
                int64_t distanceToMove = itItemToMove->actualValue;

                int64_t destinationIndex = (currentIndex + distanceToMove) % (size - 1);
                if (destinationIndex < 0)
                    destinationIndex = (size - 1) + destinationIndex;
                if (destinationIndex == 0 && currentIndex!=destinationIndex)
                    destinationIndex = (size - 1) + destinationIndex;

                if (destinationIndex == currentIndex) {
                    continue;
                }

                vector<Value>::iterator first, middle, last;
                if (destinationIndex < currentIndex) {
                    first = newSequence.begin() + destinationIndex;
                    middle = newSequence.begin() + currentIndex;
                    last = middle + 1;
                } else {
                    first = newSequence.begin() + currentIndex;
                    middle = first + 1;
                    last = newSequence.begin() + destinationIndex + 1;
                }
                std::rotate(first, middle, last);
            }

            sequence_ = std::move(newSequence);
            //std::cout << sequence_ << std::endl;
        }

        int64_t indexOfFirstZero() const
        {
            auto itZero = std::find_if(
                sequence_.begin(), sequence_.end(),
                [](auto& value) { return value.actualValue == 0; }
            );
            return itZero - sequence_.begin();
        }

        int64_t at(int location) const
        {
            int realIndex = location % sequence_.size();
            return sequence_[realIndex].actualValue;
        }
    private:
        vector<Value> sequence_;
    };

    string answer_a(vector<string>const& inputData)
    {
        vector<Value> sequence(inputData.size());

        // the numbers in the sequence are NOT unique! 
        // so we have to keep track of their original locations
        for (int i = 0; i < inputData.size(); i++) {
            sequence[i] = Value{ std::stoi(inputData[i]), i };
        }

        Mixer mixer(std::move(sequence));
        mixer.mix();

        int zeroIndex = mixer.indexOfFirstZero();

        return std::to_string(
            mixer.at(zeroIndex + 1000) +
            mixer.at(zeroIndex + 2000) +
            mixer.at(zeroIndex + 3000)
        );
    }

    string answer_b(vector<string>const& inputData)
    {
        const auto decryption_key = 811589153;
        // Multiply each number by the decryption key before you begin

        vector<Value> sequence(inputData.size());

        // the numbers in the sequence are NOT unique! 
        // so we have to keep track of their original locations
        for (int i = 0; i < inputData.size(); i++) {
            sequence[i] = Value{ decryption_key * std::stoll(inputData[i]), i };
        }

        Mixer mixer(std::move(sequence));
        for(int i=0; i<10; ++i) mixer.mix();

        int zeroIndex = mixer.indexOfFirstZero();

        return std::to_string(
            mixer.at(zeroIndex + 1000) +
            mixer.at(zeroIndex + 2000) +
            mixer.at(zeroIndex + 3000)
        );
    }

}
