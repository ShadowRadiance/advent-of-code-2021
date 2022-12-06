# frozen_string_literal: true

require "./solver"
require "./pods/room"

RSpec.describe Solver do
  subject { Solver.new(initial, solved) }

  context "with the sample data" do
    let(:initial) { ".|.|BA|.|CD|.|BC|.|DA|.|." }
    let(:solved)  { ".|.|AA|.|BB|.|CC|.|DD|.|." }

    it "is all cool and shit" do
      expect(subject.solve).to eq(12521)
    end
  end

  context "close to done" do
    #############
    #.B...A.....#
    ###.#.#C#D###
    ###A#B#C#D###
    #############
    let(:initial) { ".|B|A|.|B|A|CC|.|DD|.|." }
    let(:solved)  { ".|.|AA|.|BB|.|CC|.|DD|.|." }

    it "has the right steps" do
      expect(Pods::Map.new(initial).all_possible_moves.map(&:to_s)).to eq([
        "[hall1->roomB]:40",
        "[hall5->roomA]:4",
      ])
    end

    it "is all cool and shit" do
      expect(subject.solve).to eq(44)
    end
  end

  context "with the FULL sample data" do
    let(:initial) { ".|.|BDDA|.|CCBD|.|BBAC|.|DACA|.|." }
    let(:solved)  { ".|.|AAAA|.|BBBB|.|CCCC|.|DDDD|.|." }

    it "is all cool and shit" do
      solver = Solver.new(initial, solved, room_size: 4)
      expect(solver.solve).to eq(44169)
    end
  end

end
