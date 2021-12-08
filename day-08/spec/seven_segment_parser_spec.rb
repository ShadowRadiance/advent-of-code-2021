require './seven_segment_parser'

RSpec.describe "Seven Segment Parser" do
  subject { SevenSegmentParser.new }

  describe "#parse" do
    let(:data) {
      <<~DATA
      be cfbegad cbdgef fgaecd cgeb fdcge agebfd fecdb fabcd edb | fdgacbe cefdb cefbgd gcbe
      edbfga begcd cbg gc gcadebf fbgde acbgfd abcde gfcbed gfec | fcgedb cgb dgebacf gc
      fgaebd cg bdaec gdafb agbcfd gdcbef bgcad gfac gcb cdgabef | cg cg fdcagb cbg
      fbegcd cbd adcefb dageb afcb bc aefdc ecdab fgdeca fcdbega | efabcd cedba gadfec cb
      aecbfdg fbg gf bafeg dbefa fcge gcbea fcaegb dgceab fcbdga | gecf egdcabf bgf bfgea
      fgeab ca afcebg bdacfeg cfaedg gcfdb baec bfadeg bafgc acf | gebdcfa ecba ca fadegcb
      dbcfg fgd bdegcaf fgec aegbdf ecdfab fbedc dacgb gdcebf gf | cefg dcbef fcge gbcadfe
      bdfegc cbegaf gecbf dfcage bdacg ed bedf ced adcbefg gebcd | ed bcgafe cdgba cbgef
      egadfb cdbfeg cegd fecab cgb gbdefca cg fgcdab egfdb bfceg | gbdfcae bgc cg cgb
      gcafb gcf dcaebfg ecagb gf abcdeg gaef cafbge fdbac fegbdc | fgae cfgab fg bagce      
      DATA
    }

    it "parses all the lines" do
      parsed = subject.parse(data)
      expect(parsed).to have_attributes(size: 10)
    end

    it "splits each line into two sets of letter-sorted words" do
      parsed = subject.parse(data)
      expect(parsed).to all( be_a(SevenSegmentParser::Line))
      expect(parsed.first.signal_patterns).to eq(
        %w[be abcdefg bcdefg acdefg bceg cdefg abdefg bcdef abcdf bde]
      )
      expect(parsed.first.output_value).to eq(
        %w[abcdefg bcdef bcdefg bceg]
      )
    end
  end
end