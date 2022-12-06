# frozen_string_literal: true

require "./snailfish_number"

RSpec.describe SnailfishNumber do

  [
    { sum: [[1,2],
            [[3,4],5]],
      expect: { to_s: "[[1,2],[[3,4],5]]",
                magnitude: 143 }},
    { sum: [[[[[4,3],4],4],[7,[[8,4],9]]],
            [1,1]],
      expect: { to_s: "[[[[0,7],4],[[7,8],[6,0]]],[8,1]]",
                magnitude: 1384 } },
    { sum: [[1,1],
            [2,2],
            [3,3],
            [4,4]],
      expect: { to_s: "[[[[1,1],[2,2]],[3,3]],[4,4]]",
                magnitude: 445 } },
    { sum: [[1,1],
            [2,2],
            [3,3],
            [4,4],
            [5,5]],
      expect: { to_s: "[[[[3,0],[5,3]],[4,4]],[5,5]]",
                magnitude: 791 } },
    { sum: [[1,1],
            [2,2],
            [3,3],
            [4,4],
            [5,5],
            [6,6]],
      expect: { to_s: "[[[[5,0],[7,4]],[5,5]],[6,6]]",
                magnitude: 1137 } },
    { sum: [
        [[[0,[4,5]],[0,0]],[[[4,5],[2,6]],[9,5]]],
        [7,[[[3,7],[4,3]],[[6,3],[8,8]]]],
        [[2,[[0,8],[3,4]]],[[[6,7],1],[7,[1,6]]]],
        [[[[2,4],7],[6,[0,5]]],[[[6,8],[2,8]],[[2,1],[4,5]]]],
        [7,[5,[[3,8],[1,4]]]],
        [[2,[2,2]],[8,[8,1]]],
        [2,9],
        [1,[[[9,3],9],[[9,0],[0,7]]]],
        [[[5,[7,4]],7],1],
        [[[[4,2],2],6],[8,7]],
      ],
      expect: { to_s: "[[[[8,7],[7,7]],[[8,6],[7,7]]],[[[0,7],[6,6]],[8,7]]]",
                magnitude: 3488 } },
  ].each_with_index do |test_data, index|
    context "with test data ##{index+1} summing #{test_data[:sum].size} elements" do
      subject do
        test_data[:sum].map { |arr| SnailfishNumber.new(arr) }.reduce { |memo, addend| memo+addend }
      end

      test_data[:expect].each do |method, expectation|
        specify "calling the method #{method} yields the expectation #{expectation}" do
          expect(subject.send(method)).to eq(expectation)
        end
      end
    end
  end


end
