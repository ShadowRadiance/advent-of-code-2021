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
        int actualValue = 0;
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
            //std::cout << sequence_ << std::endl;

            int size = sequence_.size();

            vector<Value> newSequence(sequence_);

            for (int i = 0; i < sequence_.size(); i++) {
                auto itItemToMove = std::find_if(
                    newSequence.begin(), newSequence.end(),
                    [i](auto& value) { return value.originalIndex == i; }
                );
                int currentIndex = itItemToMove - newSequence.begin();
                int distanceToMove = itItemToMove->actualValue;

                int destinationIndex = (currentIndex + distanceToMove) % (size - 1);
                if (destinationIndex <= 0) destinationIndex = (size - 1) + destinationIndex;
                //std::cout << "Moving the " << distanceToMove << " at index " << currentIndex << " to index " << destinationIndex << "...\n";

                if (destinationIndex == currentIndex) {
                    //std::cout << newSequence << std::endl;
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
                //std::cout << newSequence << std::endl;
            }

            sequence_ = std::move(newSequence);
        }

        int indexOfFirstZero() const
        {
            auto itZero = std::find_if(
                sequence_.begin(), sequence_.end(),
                [](auto& value) { return value.actualValue == 0; }
            );
            return itZero - sequence_.begin();
        }

        int at(int location) const
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

    string answer_b(vector<string>const& input_data)
    {
        return "PENDING";
    }

}
