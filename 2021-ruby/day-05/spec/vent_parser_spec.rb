require './vent/parser'
require './vent/point'

RSpec.describe "Vent Data Parser" do
  let(:parser)  { Vent::Parser.new }

  describe "#parse" do
    subject { parser.parse(data.lines(chomp: true)) }

    let(:data) {
      <<~DATA
      0,9 -> 5,9
      8,0 -> 0,8
      9,4 -> 3,4
      2,2 -> 2,1
      7,0 -> 7,4
      6,4 -> 2,0
      0,9 -> 2,9
      3,4 -> 1,4
      0,0 -> 8,8
      5,5 -> 8,2
      DATA
    }
  
    it { is_expected.to be_an Array }
    it { is_expected.to all( be_a Vent::Line ) }
    it { is_expected.to have_attributes(size:10) }

    it { is_expected.to satisfy {|arr| arr[0].start_point == Vent::Point.new(0,9)  } }
    it { is_expected.to satisfy {|arr| arr[0].end_point == Vent::Point.new(5,9)  } }
        
    it { is_expected.to satisfy {|arr| arr[2].start_point == Vent::Point.new(9,4) } }
    it { is_expected.to satisfy {|arr| arr[2].end_point == Vent::Point.new(3,4) } }
  
    it { is_expected.to satisfy {|arr| arr.last.start_point == Vent::Point.new(5,5) } }
    it { is_expected.to satisfy {|arr| arr.last.end_point == Vent::Point.new(8,2) } }
  end

end