require "./origami"

RSpec.describe Origami do
  let(:origami) { Origami.new(dots) }
  let(:parser) { OrigamiParser.new(data)}
  let(:dots) { parser.dots }
  let(:instructions) { parser.instructions }
  let(:data) {
    <<~DATA
      6,10
      0,14
      9,10
      0,3
      10,4
      4,11
      6,0
      6,12
      4,1
      0,13
      10,12
      3,4
      3,0
      8,4
      1,10
      2,14
      8,10
      9,0

      fold along y=7
      fold along x=5
    DATA
  }

  context "initial state" do
    it "has the expected state" do
      expected = <<~EXPECTED
        ...#..#..#.
        ....#......
        ...........
        #..........
        ...#....#.#
        ...........
        ...........
        ...........
        ...........
        ...........
        .#....#.##.
        ....#......
        ......#...#
        #..........
        #.#........
      EXPECTED
      expect(origami.display).to eq(expected.chomp)
    end
  end

  context "after processing a single instruction" do
    before do
      origami.process_one(instructions[0])
    end

    it "has the expected state" do
      expected = <<~EXPECTED
        #.##..#..#.
        #...#......
        ......#...#
        #...#......
        .#.#..#.###
        ...........
        ...........
      EXPECTED
      expect(origami.display).to eq(expected.chomp)
      expect(origami.display.count("#")).to eq(17)
    end
  end

  context "after processing all instructions" do
    before do
      origami.process_all(instructions)
    end

    it "has the expected state" do
      expected = <<~EXPECTED
        #####
        #...#
        #...#
        #...#
        #####
        .....
        .....
      EXPECTED
      expect(origami.display).to eq(expected.chomp)
      expect(origami.display.count("#")).to eq(16)
    end
  end
end